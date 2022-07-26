package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

var (
	clientID     = ""
	clientSecret = ""
)

func main() {
	http.HandleFunc("/login/mercedes/", mercedesLoginHandler)
	http.HandleFunc("/login/mercedes/callback", mercedesCallbackHandler)

	log.Panic(
		http.ListenAndServe(":3000", nil),
	)
}

func mercedesLoginHandler(w http.ResponseWriter, r *http.Request) {
	redirectURL := fmt.Sprintf(
		"https://id.mercedes-benz.com/as/authorization.oauth2?response_type=code&client_id=%v&redirect_uri=%v&scope=%v&state=%v",
		clientID,
		"http://localhost:3000/login/mercedes/callback",
		"mb:vehicle:mbdata:vehiclestatus",
		"login",
	)

	http.Redirect(w, r, redirectURL, http.StatusTemporaryRedirect)
}

func mercedesCallbackHandler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")

	getMercedesAccessToken(code)
}

func getMercedesAccessToken(code string) string {
	authToken := base64.StdEncoding.EncodeToString([]byte(clientID + ":" + clientSecret))

	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("redirect_uri", "http://localhost:3000/login/mercedes/callback")
	data.Set("code", code)

	req, err := http.NewRequest(
		http.MethodPost,
		"https://id.mercedes-benz.com/as/token.oauth2",
		strings.NewReader(data.Encode()),
	)
	if err != nil {
		log.Panic("Request creation failed")
	}
	req.Header.Set("Authorization", "Basic "+authToken)
	req.Header.Set("content-type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Panic("Request failed")
	}

	body, _ := ioutil.ReadAll(resp.Body)

	type mercedesAccessTokenResponse struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		Scope       string `json:"scope"`
	}

	var res mercedesAccessTokenResponse
	json.Unmarshal(body, &res)

	return res.AccessToken
}
