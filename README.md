# Tranco Go Client

This module provides a Go client for
the [Tranco API](https://tranco-list.eu/api_documentation). [Tranco](https://tranco-list.eu/) is a research-oriented top
sites ranking hardened against
manipulation.

## Supported Endpoints

Currently, the following API endpoints are supported:

* /ranks/domain/{domain} (GET)
* /lists/id/{list_id} (GET)
* /lists/date/{date} (GET)

Note: The authenticated endpoint /lists/create (PUT) is not currently supported by this client.

## Usage

For examples on how to use this client, see the _example directory.

```go
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/shigaichi/tranco"
)

func main() {
	cli := tranco.New()

	ranks, err := cli.GetRanks(context.Background(), "tranco-list.eu")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ranks)

	lists, err := cli.GetListMetadataByDate(context.Background(), time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(lists)
}
```


