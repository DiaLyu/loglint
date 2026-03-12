package main 

import ( 
	"github.com/youruser/loglint/pkg/analyzer" 
	"golang.org/x/tools/go/analysis/singlechecker" 
) 

func main() { 
	// singlechecker адаптирует наш Analyzer под стандартный CLI-инструмент 
	singlechecker.Main(analyzer.Analyzer) 
}