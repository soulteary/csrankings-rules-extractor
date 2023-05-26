package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/dop251/goja"
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

	buf, _ := os.ReadFile("node-src/dist.js")
	vm := goja.New()

	csrankingsTS, _ := os.ReadFile("data/emeryberger/CSrankings/csrankings.ts")
	vm.Set("code", string(csrankingsTS))

	v, err := vm.RunString(strings.Join([]string{
		string(buf),
		"GetCSRankingsConfig(code)",
	}, "\n"))
	if err != nil {
		panic(err)
	}

	fmt.Println(v.Export().(string))
}
