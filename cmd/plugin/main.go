package main

import (
	"golang.org/x/tools/go/analysis"
	"github.com/youruser/loglint/pkg/analyzer"
)

var AnalyzerPlugin = []*analysis.Analyzer{
	analyzer.Analyzer,
}