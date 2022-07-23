# go-merche

[![Test Status](https://github.com/jferrl/go-merche/workflows/tests/badge.svg)](https://github.com/jferrl/go-merche/actions?query=workflow%3Atests)
[![codecov](https://codecov.io/gh/jferrl/go-merche/branch/main/graph/badge.svg?token=68I4BZF235)](https://codecov.io/gh/jferrl/go-merche)

Go library for accessing the Mercedes-Benz vehicles API.

## Authentication

The go-merche library does not handle authentication. So that, when
creating a new client, pass an `http.Client` that can handle authentication.
The Recommended way to do this is using the oauth2 golang pkg.

```go
import "golang.org/x/oauth2"

func main() {
 ctx := context.Background()
 ts := oauth2.StaticTokenSource(
  &oauth2.Token{AccessToken: "... your access token ..."},
 )
 tc := oauth2.NewClient(ctx, ts)

 client := merche.NewClient(tc)
}
```
