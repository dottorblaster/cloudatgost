# cloudatGOst
CloudAtCost API wrapper written in Go.

API Version: `v1`

## References
https://github.com/cloudatcost/api

## Example
To interact with the CaC API you simply need to initialize a new client, passing to it your API key and your login.

```go
package main

import (
  "fmt"
  "cloudatgost"
)

func main() {
	// Initializes a new CloudAtCost client
    client := cloudatgost.NewClient("dottorblaster@gmail.com", "myApiKey", nil)

    // d is defined as a TemplateList, a type that maps the JSON
    // response of the listtemplates.php endpoint
    d := client.ListTemplates()

    // c is defined as a ServerList, a type that maps the JSON
    // response of the listservers.php endpoint
    c := client.ListServers()

    // e is defined as a CacConsole, a type that maps the JSON
    // response of the console.php endpoint. You need to provide
    // to it a string containing the server ID.
    e := client.Console("254484472")
  }
```

## License
cloudatGOst is licensed under MIT License. (See LICENSE)

## TODO
- Power operations
- Tests
- Refactoring to make all of this stuff more readable
