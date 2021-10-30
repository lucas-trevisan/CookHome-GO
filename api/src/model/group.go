package model

import (
	"errors"
	"strings"
)

type Group struct {
	ID   uint64 `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

func (group *Group) Prepare(stage string) error {
	if erro := group.validate(stage); erro != nil {
		return erro
	}

	if erro := group.format(stage); erro != nil {
		return erro
	}

	return nil
}

func (group *Group) validate(stage string) error {
	if group.Name == "" {
		return errors.New("the Group name field cannot be blank")
	}

	return nil
}

func (group *Group) format(stage string) error {

	group.Name = strings.TrimSpace(group.Name)

	return nil

}
