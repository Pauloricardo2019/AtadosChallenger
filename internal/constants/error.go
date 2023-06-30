package constants

import "errors"

var (
	ActionNotFound    = errors.New("action not found")
	VoluntaryNotFound = errors.New("voluntary not found")
)
