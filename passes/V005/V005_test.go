package V005_test

import (
	"testing"

	"github.com/bflad/tfproviderlint/passes/V005"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestV005(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, V005.Analyzer, "a")
}
