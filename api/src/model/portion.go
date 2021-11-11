package model

import (
	"errors"
	"strings"
)

type Portion struct {
	ID       uint64 `json:"id,omitempty"`
	Quantity string `json:"quantity,omitempty"`
}

func (portion *Portion) Prepare(stage string) error {
	if erro := portion.validate(stage); erro != nil {
		return erro
	}

	if erro := portion.format(stage); erro != nil {
		return erro
	}

	return nil
}

func (portion *Portion) validate(stage string) error {
	if portion.Quantity == "" {
		return errors.New("the Portion quantity field cannot be blank")
	}

	return nil
}

func (portion *Portion) format(stage string) error {

	portion.Quantity = strings.TrimSpace(portion.Quantity)

	return nil

}
