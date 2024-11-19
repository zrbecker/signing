package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/zrbecker/signing/cmd/main/externalsigner"
	"github.com/zrbecker/signing/signing"
	"github.com/zrbecker/signing/signing/signers"
)

func main() {
	// This registry `r` can now be passed to any service that needs to initialize a signer
	r := signing.NewRegistry()
	signers.RegisterStandardSigners(r)
	r.RegisterSigner(externalsigner.ExternalSignerType, externalsigner.ExternalSignerFactory)

	// at some point later in the program, we will obtain a config.
	jsonCfg := []byte(`
		{
			"type": "external_signer",
			"config": {
				"endpoint": "localhost:8080",
				"api_key": "secret"
			}
		}
	`)

	// alternative config
	// jsonCfg := []byte(`
	// 	{
	// 		"type": "simple_signer",
	// 		"config": {
	// 			"private_key": "mysecretprivatekey"
	// 		}
	// 	}
	// `)

	var cfg signing.SignerConfig
	if err := json.Unmarshal(jsonCfg, &cfg); err != nil {
		log.Panic(err)
	}

	// now any service that has access to this cfg and the registry can construct a signer
	signer, err := r.CreateSigner(cfg)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(signer.DebugMessage())
}
