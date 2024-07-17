package constants

import "errors"

var (
	Config     = "config"
	ErrMissVar = errors.New("error reading config")
)
