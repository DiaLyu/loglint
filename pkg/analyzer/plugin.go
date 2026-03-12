package analyzer 

import "golang.org/x/tools/go/analysis" 

// Для golangci-lint плагина 
type analyzerPlugin struct{} 

func (s *analyzerPlugin) GetAnalyzers() []*analysis.Analyzer { 
	return []*analysis.Analyzer{ Analyzer, } 
} 

var AnalyzerPlugin analyzerPlugin