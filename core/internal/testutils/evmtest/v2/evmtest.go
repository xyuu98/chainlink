package v2

import (
	"testing"

	"github.com/smartcontractkit/chainlink/core/chains/evm/config"
	v2 "github.com/smartcontractkit/chainlink/core/chains/evm/config/v2"
	"github.com/smartcontractkit/chainlink/core/internal/testutils/configtest"
	"github.com/smartcontractkit/chainlink/core/logger"
	"github.com/smartcontractkit/chainlink/core/utils"
)

func ChainEthMainnet(t *testing.T) config.ChainScopedConfig      { return scopedConfig(t, 1) }
func ChainOptimismMainnet(t *testing.T) config.ChainScopedConfig { return scopedConfig(t, 10) }
func ChainOptimismKovan(t *testing.T) config.ChainScopedConfig   { return scopedConfig(t, 69) }
func ChainArbitrumMainnet(t *testing.T) config.ChainScopedConfig { return scopedConfig(t, 42161) }
func ChainArbitrumRinkeby(t *testing.T) config.ChainScopedConfig { return scopedConfig(t, 421611) }

func scopedConfig(t *testing.T, chainID int64) config.ChainScopedConfig {
	id := utils.NewBigI(chainID)
	def, _ := v2.Defaults(id)
	evmCfg := v2.EVMConfig{ChainID: id, Chain: def}
	return v2.NewTOMLChainScopedConfig(configtest.NewTestGeneralConfig(t), &evmCfg, logger.TestLogger(t))
}
