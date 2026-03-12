package analyzer 

import ( 
	"golang.org/x/tools/go/analysis/analysistest" 
	"testing" 
) 

func TestAll(t *testing.T) { 
	testdata := analysistest.TestData() 
	analysistest.Run(t, testdata, Analyzer, "a") 
}