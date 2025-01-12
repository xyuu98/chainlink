package v2

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	evmclient "github.com/smartcontractkit/chainlink/core/chains/evm/client"
	evmcfg "github.com/smartcontractkit/chainlink/core/chains/evm/config/v2"
	"github.com/smartcontractkit/chainlink/core/config"
	"github.com/smartcontractkit/chainlink/core/logger"
	"github.com/smartcontractkit/chainlink/core/services/chainlink"
	"github.com/smartcontractkit/chainlink/core/store/dialects"
	"github.com/smartcontractkit/chainlink/core/store/models"
	"github.com/smartcontractkit/chainlink/core/utils"
)

// NewTestGeneralConfig returns a new config.GeneralConfig with default test overrides.
func NewTestGeneralConfig(t testing.TB) config.GeneralConfig { return NewGeneralConfig(t, nil) }

// NewGeneralConfig returns a new config.GeneralConfig with overrides.
// The default test overrides are applied before overrideFn.
func NewGeneralConfig(t testing.TB, overrideFn func(*chainlink.Config, *chainlink.Secrets)) config.GeneralConfig {
	tempDir := t.TempDir()
	g, err := chainlink.GeneralConfigOpts{
		OverrideFn: func(c *chainlink.Config, s *chainlink.Secrets) {
			overrides(c, s)
			c.RootDir = &tempDir
			if fn := overrideFn; fn != nil {
				fn(c, s)
			}
		},
	}.New(logger.TestLogger(t))
	require.NoError(t, err)
	return g
}

func overrides(c *chainlink.Config, s *chainlink.Secrets) {
	c.DevMode = true
	c.InsecureFastScrypt = ptr(true)

	c.Database.Dialect = dialects.TransactionWrappedPostgres
	c.Database.Lock.Mode = "none"
	c.Database.MaxIdleConns = ptr[int64](20)
	c.Database.MaxOpenConns = ptr[int64](20)
	c.Database.MigrateOnStartup = ptr(false)

	c.WebServer.SessionTimeout = models.MustNewDuration(2 * time.Minute)
	c.WebServer.BridgeResponseURL = models.MustParseURL("http://localhost:6688")

	chainID := utils.NewBigI(evmclient.NullClientChainID)
	chain, _ := evmcfg.Defaults(chainID)
	enabled := true
	c.EVM = append(c.EVM, &evmcfg.EVMConfig{
		ChainID: chainID,
		Chain:   chain,
		Enabled: &enabled,
		Nodes:   evmcfg.EVMNodes{{}},
	})
}

func ptr[T any](v T) *T { return &v }
