package domain

import "errors"

// Sentinel errors used across repository and service layers.
var (
	ErrNotFound    = errors.New("not found")
	ErrConflict    = errors.New("conflict")
	ErrValidation  = errors.New("validation error")
	ErrUnauthorized = errors.New("unauthorized")
)
