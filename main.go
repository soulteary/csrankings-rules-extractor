package main

import "github.com/soulteary/csrankings-rules-extractor/internal/network"

func FetchFiles() {
	// DBLP parsing rules.
	network.DownloadFile("https://github.com/emeryberger/CSrankings/raw/gh-pages/filter.xq")
	// Core Definitions.
	network.DownloadFile("https://github.com/emeryberger/CSrankings/raw/gh-pages/util/csrankings.py")
	// Interface Interaction Related Definitions.
	network.DownloadFile("https://github.com/emeryberger/CSrankings/raw/gh-pages/csrankings.ts")
}

const (
	FLAG_UPDATE_CSRANKINGS_FILES = false
)

func main() {
	if FLAG_UPDATE_CSRANKINGS_FILES {
		FetchFiles()
	}
}
