package signers

import (
	"encoding/json"
	"fmt"

	"github.com/zrbecker/signing/signing"
)

const SimpleSignerType = "simple_signer"

type SimpleSignerConfig struct {
	PrivateKey string `json:"private_key"`
}

type SimpleSigner struct {
	privateKey string
}

func SimpleSignerFactory(cfg signing.SignerConfig) (signing.Signer, error) {
	bz, err := json.Marshal(cfg.Config)
	if err != nil {
		return nil, err
	}

	var signerCfg SimpleSignerConfig
	if err := json.Unmarshal(bz, &signerCfg); err != nil {
		return nil, err
	}

	return &SimpleSigner{
		privateKey: signerCfg.PrivateKey,
	}, nil
}

func (s *SimpleSigner) Sign(params interface{}) (interface{}, error) {
	// Implementation irrelevant to exercise
	return nil, nil
}

func (s *SimpleSigner) DebugMessage() string {
	return fmt.Sprintf("SimpleSigner(PrivateKey=%s)", s.privateKey)
}
