package main

import (
	"fmt"
	"os"

	"github.com/soulteary/csrankings-rules-extractor/internal/define"
	"github.com/soulteary/csrankings-rules-extractor/internal/extractor"
	"github.com/soulteary/csrankings-rules-extractor/internal/network"
)

func FetchFiles() {
	// DBLP parsing rules.
	network.DownloadFile("https://github.com/emeryberger/CSrankings/raw/gh-pages/filter.xq")
	// Core Definitions.
	network.DownloadFile("https://github.com/emeryberger/CSrankings/raw/gh-pages/util/csrankings.py")
	// Interface Interaction Related Definitions.
	network.DownloadFile("https://github.com/emeryberger/CSrankings/raw/gh-pages/csrankings.ts")
}

func main() {
	if define.FLAG_UPDATE_CSRANKINGS_FILES {
		FetchFiles()
	}

	fmt.Println("Start parsing filter.xq")
	inproceedings, articles, err := extractor.GetFilters("data/emeryberger/CSrankings/filter.xq")
	if err != nil {
		panic(err)
	}
	fmt.Println("inproceedings:", len(inproceedings))
	fmt.Println("articles:", len(articles))
	fmt.Println()
	fmt.Println()

	fmt.Println("Start parsing csrankings.ts")
	tsConfig, err := extractor.GetCSRankingsTS("node-src/dist.js")
	if err != nil {
		panic(err)
	}
	fmt.Println(tsConfig)
	fmt.Println()
	fmt.Println()

	fmt.Println("Start parsing csrankings.py")
	// docker run --rm -it -p 8181:8081 soulteary/go-python-ast:alpine
	csrankingsPy, err := os.ReadFile("data/emeryberger/CSrankings/util/csrankings.py")
	if err != nil {
		panic(err)
	}
	jsConverter, err := os.ReadFile("internal/extractor/csrankings-py.js")
	if err != nil {
		panic(err)
	}
	GetPyConfig, err := extractor.GetPyConfig("localhost:8181", csrankingsPy, jsConverter)
	if err != nil {
		panic(err)
	}
	fmt.Println(GetPyConfig)
	fmt.Println()
	fmt.Println()

}
