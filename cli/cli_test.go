package cli

import (
	"errors"
	"testing"

	"github.com/monitoror/monitoror/cli/version"

	coreModels "github.com/monitoror/monitoror/models"
)

func TestPrintBanner(t *testing.T) {
	PrintBanner()
}

func TestPrintDevMode(t *testing.T) {
	PrintDevMode()
}

func TestPrintMonitorableHeader(t *testing.T) {
	PrintMonitorableHeader()
}

func TestPrintMonitorable(t *testing.T) {
	PrintMonitorable("TEST", nil, nil)
	PrintMonitorable("TEST", []coreModels.Variant{coreModels.DefaultVariant}, nil)
	PrintMonitorable("TEST", []coreModels.Variant{coreModels.DefaultVariant, "variant1"}, map[coreModels.Variant]error{"variant2": errors.New("boom")})
	PrintMonitorable("TEST", nil, map[coreModels.Variant]error{"default": errors.New("boom")})
}

func TestPrintMonitorableFooter(t *testing.T) {
	version.Version = "0.0.0"
	PrintMonitorableFooter(true)
}

func TestPrintServerStartup(t *testing.T) {
	PrintServerStartup(nil, 8080)
}
