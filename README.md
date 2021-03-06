[![GoDoc](https://godoc.org/github.com/rays/go-pagerduty?status.svg)](http://godoc.org/github.com/rays/go-pagerduty) [![Go Report Card](https://goreportcard.com/badge/github.com/rays/go-pagerduty)](https://goreportcard.com/report/github.com/rays/go-pagerduty) [![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/gojp/goreportcard/blob/master/LICENSE)
# go-pagerduty

go-pagerduty is a CLI and [go](https://golang.org/) client library for the [PagerDuty v2 API](https://v2.developer.pagerduty.com/v2/page/api-reference).

## Installation

First, download the source code
```cli
go get github.com/rays/go-pagerduty
```

Next build the application.
```cli
cd $GOPATH/src/github.com/rays/go-pagerduty
make install
```

## Usage

### CLI

The CLI requires an [authentication token](https://v2.developer.pagerduty.com/docs/authentication), which can be specified in `.pd.yml`
file in the home directory of the user, or passed as a command-line argument.
Example of config file:

```yaml
---
authtoken: fooBar
```

#### Commands
`pd` command provides a single entrypoint for all the API endpoints, with individual
API represented by their own sub commands. For an exhaustive list of sub-commands, try:
```
pd --help
```

An example of the `service` sub-command

```
pd service list
```


### Client Library

```go
package main

import (
	"fmt"
	"github.com/rays/go-pagerduty"
)

var	authtoken = "" // Set your auth token here

func main() {
	var opts pagerduty.ListEscalationPoliciesOptions
	client := pagerduty.NewClient(authtoken)
	eps, err := client.ListEscalationPolicies(opts)
	if err != nil {
		panic(err)
	}
	for _, p := range eps.EscalationPolicies {
		fmt.Println(p.Name)
	}
}
```

The PagerDuty API client also exposes its HTTP client as the `HTTPClient` field.
If you need to use your own HTTP client, for doing things like defining your own
transport settings, you can replace the default HTTP client with your own by
simply by setting a new value in the `HTTPClient` field.

## Contributing

1. Fork it ( https://github.com/rays/go-pagerduty/fork )
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create a new Pull Request

## License
[Apache 2](http://www.apache.org/licenses/LICENSE-2.0)
