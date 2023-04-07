package vo

import "strings"

type IValidate interface {
	ValidateError() (msg string, ok bool)
}

type ValidateMsg struct {
	arr []string
}

func NewValidateMsg() ValidateMsg {
	return ValidateMsg{
		arr: make([]string, 0),
	}
}

func (v *ValidateMsg) Add(msg string) {
	v.arr = append(v.arr, msg)
}

func (v *ValidateMsg) Message() string {
	return strings.Join(v.arr, "; ")
}

func (v *ValidateMsg) HasMessage() bool {
	return len(v.arr) > 0
}
