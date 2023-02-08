package dortsolver

import "errors"

var (
	ErrNoApiKey          = errors.New("no api key specified")
	ErrNoPublicKey       = errors.New("no public/site key")
	ErrNoSiteUrl         = errors.New("no site url")
	ErrUnsolvableCaptcha = errors.New("unsolvable captcha")
)
