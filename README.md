# cloudatGOst [![GoDoc](https://godoc.org/github.com/dottorblaster/cloudatgost?status.svg)](https://godoc.org/github.com/dottorblaster/cloudatgost) [![Build Status](https://travis-ci.org/dottorblaster/cloudatgost.svg?branch=master)](https://travis-ci.org/dottorblaster/cloudatgost)

CloudAtCost API wrapper written in Go.

API Version: `v1`

## References
https://github.com/cloudatcost/api

## Installation
```bash
$ go get github.com/dottorblaster/cloudatgost
```

## Example
To interact with the CaC API you simply need to initialize a new client, passing to it your API key and your login.

```go
package main

import (
  "github.com/dottorblaster/cloudatgost"
)

func main() {
	// Initializes a new CloudAtCost client
    client := cloudatgost.NewClient("johndoe@example.com", "myApiKey", nil)

    // d is defined as a TemplateList, a type that maps the JSON
    // response of the listtemplates.php endpoint
    first := client.ListTemplates()

    // c is defined as a ServerList, a type that maps the JSON
    // response of the listservers.php endpoint
    second := client.ListServers()

    // e is defined as a CacConsole, a type that maps the JSON
    // response of the console.php endpoint. You need to provide
    // to it a string containing the server ID.
    third := client.Console("254484472")
  }
```

## License
cloudatGOst is licensed under MIT License. (See LICENSE)

## TODO
- Tests
- Refactoring to make all of this stuff more readable
