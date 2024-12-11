package netopia

import "errors"

var (
	ErrMissingAPIKey       = errors.New("missing apiKey")
	ErrMissingPosSignature = errors.New("missing posSignature")
	ErrInvalidOrder        = errors.New("invalid order data")
	ErrMissingVerification = errors.New("missing Verification-token header")
	ErrInvalidPublicKey    = errors.New("invalid public key")
	ErrInvalidIssuer       = errors.New("invalid issuer")
	ErrEmptyAudience       = errors.New("empty audience in token")
	ErrInvalidAudience     = errors.New("invalid audience")
	ErrAudienceNotInSet    = errors.New("audience not in posSignatureSet")
	ErrPayloadHashMismatch = errors.New("payload hash mismatch")
	ErrInvalidToken        = errors.New("invalid token")
	ErrMissingHashMethod   = errors.New("hash method missing")
)
