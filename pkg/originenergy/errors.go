package originenergy

import "errors"

var (
	ErrSessionFileNotFound = errors.New("session file not found")
	ErrAgreementIDNotSet   = errors.New("AGREEMENT_ID not set")
)
