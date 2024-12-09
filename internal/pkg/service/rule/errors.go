package rule

import (
	"errors"
)

var (
	ErrRuleExists      = errors.New("rule account already exists")
	ErrRuleNotFound    = errors.New("rule account not found")
	ErrForeignKeyError = errors.New("foreign key error")
	ErrInvalidField    = errors.New("error invalid field")
)
