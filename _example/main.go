package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Shigaichi/tranco"
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
