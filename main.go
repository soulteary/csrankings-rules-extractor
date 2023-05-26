package main

import (
	"fmt"
	"os"
	"strings"

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

	// fmt.Println("Start parsing filter.xq")
	// inproceedings, articles, err := extractor.GetFilters("data/emeryberger/CSrankings/filter.xq")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("inproceedings:", len(inproceedings))
	// fmt.Println("articles:", len(articles))

	// tsConfig, err := extractor.GetCSRankingsTS("node-src/dist.js")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(tsConfig)

	buf, err := os.ReadFile("data/emeryberger/CSrankings/util/csrankings.py")
	if err != nil {
		panic(err)
	}

	// docker run --rm -it -p 8181:8081 soulteary/go-python-ast:alpine
	// trick: Comment out the code that is irrelevant to the parsing configuration to avoid errors in AST parsing (lower versions of Python)
	code := strings.ReplaceAll(string(buf), `if (pvmatcher := TECSCounterColon.match(pages)):`, "#")
	python, err := extractor.GetPythonAST("localhost:8181", code)
	if err != nil {
		panic(err)
	}
	fmt.Println(python)
}
