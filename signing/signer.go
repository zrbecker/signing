package signing

type Signer interface {
	Sign(params interface{}) (interface{}, error)
	DebugMessage() string
}

type SignerConfig struct {
	Type   string      `json:"type"`
	Config interface{} `json:"config"`
}
