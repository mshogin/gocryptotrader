//+build mock_test_off

// This will build if build tag mock_test_off is parsed and will do live testing
// using all tests in (exchange)_test.go
package anx

import (
	"log"
	"os"
	"testing"

	"github.com/thrasher-corp/gocryptotrader/config"
	"github.com/thrasher-corp/gocryptotrader/exchanges/sharedtestvalues"
)

var mockTests = false

func TestMain(m *testing.M) {
	cfg := config.GetConfig()
	cfg.LoadConfig("../../testdata/configtest.json")
	anxConfig, err := cfg.GetExchangeConfig("ANX")
	if err != nil {
		log.Fatalf("Test Failed - ANX Setup() init error: %s", err)
	}
	anxConfig.API.AuthenticatedSupport = true
	anxConfig.API.Credentials.Key = apiKey
	anxConfig.API.Credentials.Secret = apiSecret
	a.SetDefaults()
	a.Setup(anxConfig)
	log.Printf(sharedtestvalues.LiveTesting, a.GetName(), a.API.Endpoints.URL)
	os.Exit(m.Run())
}