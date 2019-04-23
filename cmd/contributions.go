package cmd

import (
	"errors"
	"flag"

	contributions "github.com/ajardin/contributions/internal"
)

func init() {
	flag.StringVar(&contributions.Path, "path", ".", "Path of the project to be analysed.")
	flag.StringVar(&contributions.Start, "start", "01 Jan 2000", "Date from which commits are counted.")
	flag.IntVar(&contributions.Threshold, "threshold", 10, "Minimum threshold for contributions.")
}

func Run() {
	flag.Parse()

	if len(contributions.Path) == 0 {
		panic(errors.New("a valid project path must be provided"))
	}

	if len(contributions.Start) == 0 {
		panic(errors.New("a valid start date must be provided"))
	}

	// Start the analysis process
	contributions.Analyze()
}
