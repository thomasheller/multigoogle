package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"
	"sync"

	"github.com/skratchdot/open-golang/open"
	be "github.com/thomasheller/braceexpansion"
)

const (
	googleSearch = "https://www.google.com/search?q="
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	input := strings.Join(os.Args[1:], " ")
	tree := parse(input)
	queries := tree.Expand()
	openAll(queries)
}

func parse(input string) *be.Tree {
	opts := be.ParseOpts{
		OpenBrace:             "(",
		CloseBrace:            ")",
		Separator:             ",",
		TreatRootAsList:       true,
		TreatSingleAsOptional: true,
	}

	tree, err := be.New().ParseCustom(input, opts)
	if err != nil {
		log.Fatalf("Parse error: %v", err)
	}

	return tree
}

func openAll(queries []string) {
	var wg sync.WaitGroup

	for _, query := range queries {
		wg.Add(1)
		go openBrowser(query, &wg)
	}

	wg.Wait()
}

func openBrowser(query string, wg *sync.WaitGroup) {
	t := strings.TrimSpace(query)
	e := url.QueryEscape(t)
	openURL := fmt.Sprintf("%s%s", googleSearch, e)
	fmt.Println(openURL)
	open.Run(openURL)
	wg.Done()
}

func printUsage() {
	fmt.Println(`usage: multigoogle [brace expression]
Examples:
  multigoogle San Francisco,Berlin
    => "San Francisco"
    => "Berlin"
  multigoogle Go (San Francisco,Berlin)
    => "Go San Francisco"
    => "Go Berlin"
  multigoogle (Go,Golang) (San Francisco,Berlin)
    => "Go San Francisco"
    => "Go Berlin"
    => "Golang San Francisco"
    => "Golang Berlin"
  multigoogle San Francisco (Golang)
    => "San Francisco"
    => "San Francisco (Golang)"
  multigoogle (San Francisco,Berlin) (Golang)
    => "San Francisco"
    => "San Francisco (Golang)"
    => "Berlin"
    => "Berlin (Golang)"`)
}
