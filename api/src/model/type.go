package model

import (
	"errors"
	"strings"
)

type Type struct {
	ID   uint64 `json:"id,omitempty"`
	Type string `json:"typ,omitempty"`
}

func (typ *Type) Prepare(stage string) error {
	if erro := typ.validate(stage); erro != nil {
		return erro
	}

	if erro := typ.format(stage); erro != nil {
		return erro
	}

	return nil
}

func (typ *Type) validate(stage string) error {
	if typ.Type == "" {
		return errors.New("the Type field cannot be blank")
	}

	return nil
}

func (typ *Type) format(stage string) error {

	typ.Type = strings.TrimSpace(typ.Type)

	return nil

}
