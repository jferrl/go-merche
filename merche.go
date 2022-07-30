package merche

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	defaultBaseURL = "https://api.mercedes-benz.com/"
	userAgent      = "go-merche"
)

// A Client manages communication with the Mercedes API.
type Client struct {
	client *http.Client

	// Base URL for API requests. Defaults to the Mercedes API.
	//BaseURL should always be specified with a trailing slash.
	BaseURL *url.URL

	// User agent used when communicating with the Mercedes API.
	UserAgent string

	common service // Reuse a single struct instead of allocating one for each service on the heap.

	Resources     *ResourcesService
	VehicleStatus *VehicleStatusService
	FuelStatus    *FuelStatusService
}

type service struct {
	client *Client
}

// NewClient returns a new Mercedes API client. If a nil httpClient is
// provided, a new http.Client will be used. To use API methods which require
// authentication, provide an http.Client that will perform the authentication
// like golang.org/x/oauth2 library.
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	baseURL, _ := url.Parse(defaultBaseURL)

	c := &Client{
		client:    httpClient,
		BaseURL:   baseURL,
		UserAgent: userAgent,
	}
	c.common.client = c

	c.Resources = (*ResourcesService)(&c.common)
	c.VehicleStatus = (*VehicleStatusService)(&c.common)
	c.FuelStatus = (*FuelStatusService)(&c.common)

	return c
}

func (c *Client) newRequest(ctx context.Context, method, path string, body io.Reader) (*http.Request, error) {
	if !strings.HasSuffix(c.BaseURL.Path, "/") {
		return nil, fmt.Errorf("BaseURL must have a trailing slash, but %q does not", c.BaseURL)
	}

	req, err := http.NewRequestWithContext(ctx, method, c.BaseURL.String()+path, body)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}
	return req, nil
}

func (c *Client) do(req *http.Request, v interface{}) (*Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return &Response{resp}, err
	}
	defer resp.Body.Close()

	err = checkResponse(resp)
	if err != nil {
		return &Response{resp}, err
	}

	switch v := v.(type) {
	case nil:
	case io.Writer:
		_, err = io.Copy(v, resp.Body)
	default:
		decErr := json.NewDecoder(resp.Body).Decode(v)
		if decErr == io.EOF {
			decErr = nil // ignore EOF errors caused by empty response body
		}
		if decErr != nil {
			err = decErr
		}
	}

	return &Response{resp}, err
}

func checkResponse(r *http.Response) error {
	if code := r.StatusCode; http.StatusOK <= code && code <= 299 {
		return nil
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return errors.New("check response: error reading response body")
	}

	if isExVeError(r.StatusCode) {
		var exVeError ExVeError
		json.Unmarshal(body, &exVeError)
		return &exVeError
	}
	if r.StatusCode == http.StatusUnauthorized {
		var authErr UnauthorizedError
		json.Unmarshal(body, &authErr)
		return &authErr
	}

	return &MercedesAPIError{r.StatusCode}
}

// Bool is a helper routine that allocates a new bool value
// to store v and returns a pointer to it.
func Bool(v bool) *bool { return &v }

// Int is a helper routine that allocates a new int value
// to store v and returns a pointer to it.
func Int(v int) *int { return &v }

// Int64 is a helper routine that allocates a new int64 value
// to store v and returns a pointer to it.
func Int64(v int64) *int64 { return &v }

// String is a helper routine that allocates a new string value
// to store v and returns a pointer to it.
func String(v string) *string { return &v }
