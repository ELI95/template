package application

import "github.com/eli95/template/pkg/validator/domain/ports"

type Validator struct {
}

func NewValidator() *Validator {
	return &Validator{}
}

func (v *Validator) Struct(s ports.Evaluable) error {
	return s.Validate()
}
