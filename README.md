# go-merche

[![GoDoc](https://img.shields.io/static/v1?label=godoc&message=reference&color=blue)](https://pkg.go.dev/github.com/jferrl/go-merche)
[![Test Status](https://github.com/jferrl/go-merche/workflows/tests/badge.svg)](https://github.com/jferrl/go-merche/actions?query=workflow%3Atests)
[![codecov](https://codecov.io/gh/jferrl/go-merche/branch/main/graph/badge.svg?token=68I4BZF235)](https://codecov.io/gh/jferrl/go-merche)

Go library for accessing the Mercedes-Benz vehicles API. :auto_rickshaw:

## Why go-merche?

With the launch of connected cars on the market, some brands such as Mercedes-Benz have created a set of APIs with which we can access in "real" time data that our vehicles are generating, in addition to other additional services that may not be intrinsically related to vehicles.

The idea of this package is to provide a set of methods to be able to access the Mercedes-Benz APIs using the Go programming language.

Last but not least, the name of the pkg :sunglasses: . I name it "go-merche" just because "Merche" the diminutive of "Mercedes" said affectively in Spanish.

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

 opts := &merche.Options{ VehicleID: "EXVETESTVIN000001" }
 resources, _, err := client.Resources.GetAvailableResources(ctx, opts)
}
```

Using the `context` package, you can easily pass cancelation signals and
deadlines to various services of the client for handling a request.

## How to enable mercedes APIs

1) Own a Mercedes Benz Car with Mercedes me installed and working.
2) Create an application in <https://developer.mercedes-benz.com/>
3) Register to the following APIs (all of them):
   - [Electric Vehicle Status](https://developer.mercedes-benz.com/products/electric_vehicle_status)
   - [Fuel Status](https://developer.mercedes-benz.com/products/fuel_status)
   - [Pay As You Drive Insurance](https://developer.mercedes-benz.com/products/pay_as_you_drive_insurance)
   - [Vehicle Lock Status](https://developer.mercedes-benz.com/products/vehicle_lock_status)
   - [Vehicle Status](https://developer.mercedes-benz.com/products/vehicle_status)

## Token Creation

To create a Mercedes API token it's necessary to perform the following steps:

1) Create the APP and register it to all APIs as describe in "How to enable mercedes APIs"
2) Logout in all browser from Mercedes Me (Developer site included)
3) Execute the "mercedes_api_oauth" example with your personal client_id and client_secret
4) Using your Mercedes Me credential log in: a new web-page will ask to authorize the APP created in the step 1 to access to the personal information associated to the APIs registered
5) Enable all information and press Allow

## Authentication

The `go-merche` library does not handle authentication. So that, when
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

When using an authenticated Client, all calls made by the client will
include the specified OAuth token.

Here you can find an example of how to authenticate with Mercedes
OAuth API <https://github.com/jferrl/go-merche/tree/main/example/mercedes_api_oauth>

## Mercedes API connections

The `go-merche` pkg implements services to connect against the following Mercedes APIs:

- [Fuel Status](https://developer.mercedes-benz.com/products/fuel_status) :white_check_mark:
- [Pay As You Drive Insurance](https://developer.mercedes-benz.com/products/pay_as_you_drive_insurance) :white_check_mark:
- [Vehicle Lock Status](https://developer.mercedes-benz.com/products/vehicle_lock_status) :white_check_mark:
- [Vehicle Status](https://developer.mercedes-benz.com/products/vehicle_status) :white_check_mark:
- [Electric Vehicle Status](https://developer.mercedes-benz.com/products/electric_vehicle_status) :white_check_mark:

Take into account that the pkg services are reaching API containers to get all the avalible resources
in the same API call. In future releases, `go-merche` will implement individual methods to get data from
a specific resource from the Mercedes API. :construction:

## Use cases

- Get real data from Mercedes-Benz vehicles.
- IOT proyects. Just take a look to <https://tinygo.org/>.
- ...and many more :stuck_out_tongue_winking_eye:

## License

This library is distributed under the BSD-style license found in the LICENSE file.
