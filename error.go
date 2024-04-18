package gomockmatcher

import "errors"

var (
	ErrIgnoreMethodAlreadyUsed = errors.New("the Ignore() method has already been used; cannot use both Check() and Ignore() on the same instance")
	ErrCheckMethodAlreadyUsed  = errors.New("the Check() method has already been used; cannot use both Check() and Ignore() on the same instance")
)
