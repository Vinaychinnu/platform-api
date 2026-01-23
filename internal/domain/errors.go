package domain

import "errors"

var (
	ErrInvalidName        = errors.New("invalid name")
	ErrInvalidEnvironment = errors.New("invalid environment")
	ErrInvalidVersion     = errors.New("invalid version")
)
