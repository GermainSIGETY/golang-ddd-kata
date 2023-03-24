package e2e

import (
	"flag"
	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
	"github.com/gin-gonic/gin"
	"os"
	"testing"
)

var opts = godog.Options{
	Output:      colors.Colored(os.Stdout),
	Format:      "pretty",
	Paths:       []string{"features"},
	Concurrency: 1,
	Tags:        "~buggy",
}

func init() {
	godog.BindCommandLineFlags("godog.", &opts)
}

func TestFeatures(t *testing.T) {

	flag.Parse()
	// Set Gin to Test Mode
	gin.SetMode(gin.TestMode)
	opts.TestingT = t

	testSuite := godog.TestSuite{
		Name:                 "Todos API integration tests suite",
		TestSuiteInitializer: InitializeTestSuite,
		ScenarioInitializer:  InitializeScenario,
		Options:              &opts,
	}

	if testSuite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}
