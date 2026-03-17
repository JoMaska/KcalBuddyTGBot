package vo

import "errors"

type Gender string

const (
	Male   Gender = "male"
	Female Gender = "female"
)

func ParseGender(s string) (Gender, error) {
	switch Gender(s) {
	case Male, Female:
		return Gender(s), nil
	default:
		return "", errors.New("Invalid gender")
	}

}
