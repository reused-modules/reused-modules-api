package validation

import (
	"errors"

	"reused-modules-api/adapter/validator"
)

var (
	errInvalidValidatorInstance = errors.New("invalid validator instance")
)

const (
	InstanceGoPlayground int = iota
)

func NewValidatorFactory(instance int) (validator.Validator, error) {
	switch instance {
	case InstanceGoPlayground:
		return NewGoPlayground()
	default:
		return nil, errInvalidValidatorInstance
	}
}
