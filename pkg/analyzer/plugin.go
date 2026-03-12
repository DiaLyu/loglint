package analyzer

import "golang.org/x/tools/go/analysis"

type AnalyzerPlugin struct{}

func (AnalyzerPlugin) GetAnalyzers() []*analysis.Analyzer {
	return []*analysis.Analyzer{
		Analyzer,
	}
}

var Plugin AnalyzerPlugin