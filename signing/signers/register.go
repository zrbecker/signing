package signers

import "github.com/zrbecker/signing/signing"

func RegisterStandardSigners(r *signing.Registry) {
	r.RegisterSigner(SimpleSignerType, SimpleSignerFactory)
}
