package main

import (
	"github.com/nobishino/live20210624"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(live20210624.Analyzer) }

