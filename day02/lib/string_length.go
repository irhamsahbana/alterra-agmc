package lib

import "errors"

func MaxStringLength(str string, max int) error {
	if len(str) > max {
		return errors.New("string length is too long")
	}
	return nil
}

func MinStringLength(str string, min int) error {
	if len(str) < min {
		return errors.New("string length is too short")
	}
	return nil
}
