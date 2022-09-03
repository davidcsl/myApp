package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/goDavid/myapp/api"
)

// bga --name "ticket to ride" --clientid abc123 --skip 10 --limit 5
// define the cmd line arguments
func main() {
	query := flag.String("query", "", "Boardgame name to search")
	clientId := flag.String("clientId", "", "Boardgame Atlas client_id")
	limit := flag.Uint("limit", 10, "Limit the number of results returned")
	skip := flag.Uint("skip", 0, "Skip the number of results provided")
	timeout := flag.Uint("timeout", 0, "defer duration")

	//Parse the command line

	flag.Parse()

	fmt.Printf("query=%s, clientId=%s, limit=%d, skip=%d\n", *query, *clientId, *limit, *skip)

	if isNull(*query) {
		log.Fatalln("plesae use --query to set the boardgame name to search")
	}
	if isNull(*clientId) {
		log.Fatalln("plesae use --cleintId to set your Boardgame Atlas client id")
	}

	bga := api.New(*clientId)

	// create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(*timeout*uint(time.Second)))
	defer cancel()

	//make the invocation
	result, err := bga.Search(context.Background(), *query, *limit, *skip)

	if nil != err {
		log.Fatalf("Cannot search for boardgame: %v", err)
	}

	boldGreen = color.New(color.Bold).Add(color.FgHiGreen).FprintFunc()

	for _, g := range result.Games {
		fmt.Printf("%s: %s\n", boldGreen("Name"), g.Name)
		fmt.Printf("%s: %s\n", boldGreen("Description"), g.Description)
		fmt.Printf("%s: %s\n", boldGreen("Url"), g.Url)
	}

}

func isNull(s string) bool {
	return len(strings.TrimSpace(s)) <= 0
}
