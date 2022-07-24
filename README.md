# go-merche

[![GoDoc](https://img.shields.io/static/v1?label=godoc&message=reference&color=blue)](https://pkg.go.dev/github.com/jferrl/go-merche)
[![Test Status](https://github.com/jferrl/go-merche/workflows/tests/badge.svg)](https://github.com/jferrl/go-merche/actions?query=workflow%3Atests)
[![codecov](https://codecov.io/gh/jferrl/go-merche/branch/main/graph/badge.svg?token=68I4BZF235)](https://codecov.io/gh/jferrl/go-merche)

Go library for accessing the Mercedes-Benz vehicles API.

## Usage

go-merche is compatible with modern Go releases.

Build a new client, then you can use the services to reach different parts of the Mercedes API. For example:

```go
import (
    "context"

    "github.com/jferrl/go-merche"
    )

func main() {
 ctx := context.Background()
 
 client := merche.NewClient(nil)

 resources, _, err := client.VehicleStatus.GetAvailableResources(ctx, &merche.GetVehicleStatusOptions{ VehicleID: "EXVETESTVIN000001" })
}
```

Using the `context` package, you can easily pass cancelation signals and deadlines to various services of the client for handling a request.

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

When using an authenticated Client, all calls made by the client will include the specified OAuth token.

Here you can find an example of how to authenticate with Mercedes OAuth API <https://github.com/jferrl/go-merche/tree/main/example/mercedes_api_oauth>

## License

This library is distributed under the BSD-style license found in the LICENSE file.
