package externalsigner

import (
	"encoding/json"
	"fmt"

	"github.com/zrbecker/signing/signing"
)

const ExternalSignerType = "external_signer"

type ExternalSignerConfig struct {
	Endpoint string `json:"endpoint"`
	APIKey   string `json:"api_key"`
}

type ExternalSigner struct {
	endpoint string
	apiKey   string
}

func ExternalSignerFactory(cfg signing.SignerConfig) (signing.Signer, error) {
	bz, err := json.Marshal(cfg.Config)
	if err != nil {
		return nil, err
	}

	var signerCfg ExternalSignerConfig
	if err := json.Unmarshal(bz, &signerCfg); err != nil {
		return nil, err
	}

	return &ExternalSigner{
		endpoint: signerCfg.Endpoint,
		apiKey:   signerCfg.APIKey,
	}, nil
}

func (s *ExternalSigner) Sign(params interface{}) (interface{}, error) {
	// Implementation irrelevant to exercise
	return nil, nil
}

func (s *ExternalSigner) DebugMessage() string {
	return fmt.Sprintf("ExternalSigner(Endpoint=%s, APIKey=%s)", s.endpoint, s.apiKey)
}
