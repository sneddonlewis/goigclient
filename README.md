![Go CI](https://github.com/sneddonlewis/goigclient/actions/workflows/run-tests.yml/badge.svg) 
![Go Report Card](https://goreportcard.com/badge/github.com/sneddonlewis/goigclient) 
![License](https://img.shields.io/github/license/sneddonlewis/goigclient)

# Go IG Client
Golang Client Library for IG REST trading API documented at [IG Labs](https://labs.ig.com/rest-trading-api-guide.html)

## Status

This package is in pre-release. All the available endpoints from the IG REST API are provided. There may be bugs and the API is subject to change until a stable release from V1.  
The streaming API has not yet been added.

## Installation

```
go get github.com/sneddonlewis/goigclient
```

## How to use

```go
import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/sneddonlewis/goigclient/client"
)

func main() {

    // Create a client with your API key, username and password either for
    // a demo account or for a live account. In this case we're using a demo
    // account.
	client := client.NewIGClient(
        "IgDemoAPIKey",
        "IgDemoUsername",
        "IgDemoPassword",
        false, // Whether to connect to a Live account
    )

    // Loggin in returns OAuth token and refresh token in the response
    // but these are also tracked in the client so the response can be
    // ignored and doesn't need to be provided with each request
	authResponse, err = client.Login()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

    indented, _ := json.MarshalIndent(authResponse, "", " ")
    fmt.Println(string(indented))

    // Retrieve all the open positions on the active account
    // ( The active account is the default account for Live or Demo )
    positionsResponse, err := client.Positions()

    if err != nil {
        fmt.Println(err)
	    os.Exit(1)
    }

    for _, pos := range positionsResponse.Positions {
        indented, _ = json.MarshalIndent(pos, "", " ")
	    fmt.Println(string(indented))
    }
}
```
