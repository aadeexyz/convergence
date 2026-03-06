package lens

import (
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

type workflowFactoryTuple struct {
	Factory           common.Address
	Name              string
	Oracle            common.Address
	LatestMarket      common.Address
	CollateralBalance *big.Int
	RollingEMAWindow  []*big.Int
}

func TestCodecDecodeGetWorkflowDataMethodOutput(t *testing.T) {
	codec, err := NewCodec()
	require.NoError(t, err)

	parsed, err := abi.JSON(strings.NewReader(LensMetaData.ABI))
	require.NoError(t, err)

	collateralToken := common.HexToAddress("0x1000000000000000000000000000000000000001")
	expectedFactories := []workflowFactoryTuple{
		{
			Factory:           common.HexToAddress("0x2000000000000000000000000000000000000002"),
			Name:              "Taylor Swift",
			Oracle:            common.HexToAddress("0x3000000000000000000000000000000000000003"),
			LatestMarket:      common.HexToAddress("0x4000000000000000000000000000000000000004"),
			CollateralBalance: big.NewInt(123456789),
			RollingEMAWindow:  []*big.Int{big.NewInt(11), big.NewInt(22), big.NewInt(33)},
		},
	}

	encoded, err := parsed.Methods["getWorkflowData"].Outputs.Pack(collateralToken, expectedFactories)
	require.NoError(t, err)

	output, err := codec.DecodeGetWorkflowDataMethodOutput(encoded)
	require.NoError(t, err)
	require.Equal(t, collateralToken, output.CollateralToken)
	require.Len(t, output.Factories, 1)
	require.Equal(t, expectedFactories[0].Factory, output.Factories[0].Factory)
	require.Equal(t, expectedFactories[0].Name, output.Factories[0].Name)
	require.Equal(t, expectedFactories[0].Oracle, output.Factories[0].Oracle)
	require.Equal(t, expectedFactories[0].LatestMarket, output.Factories[0].LatestMarket)
	require.Zero(t, expectedFactories[0].CollateralBalance.Cmp(output.Factories[0].CollateralBalance))
	require.Len(t, output.Factories[0].RollingEMAWindow, 3)
	require.Zero(t, expectedFactories[0].RollingEMAWindow[0].Cmp(output.Factories[0].RollingEMAWindow[0]))
	require.Zero(t, expectedFactories[0].RollingEMAWindow[1].Cmp(output.Factories[0].RollingEMAWindow[1]))
	require.Zero(t, expectedFactories[0].RollingEMAWindow[2].Cmp(output.Factories[0].RollingEMAWindow[2]))
}
