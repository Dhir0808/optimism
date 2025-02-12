package opcm

import (
	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum-optimism/optimism/op-chain-ops/script"
)

type DeployProxyInput struct {
	Owner common.Address
}

func (input *DeployProxyInput) InputSet() bool {
	return true
}

type DeployProxyOutput struct {
	Proxy common.Address
}

type DeployProxyScript struct {
	Run func(input, output common.Address) error
}

func DeployProxy(
	host *script.Host,
	input DeployProxyInput,
) (DeployProxyOutput, error) {
	return RunBasicScript[DeployProxyInput, DeployProxyOutput](host, input, "DeployProxy.s.sol", "DeployProxy")
}
