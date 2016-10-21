package bcrypt

import (
	"github.com/jsimonetti/pwscheme/ssha"
)

func GenerateFromPassword(pwd []byte, cost uint8) ([]byte, error) {
	hash, err := ssha.Generate(string(pwd), cost)
	return []byte(hash), err
}

func CompareHashAndPassword(hash, pwd []byte) error {
	_, err := ssha.Validate(string(pwd), string(hash))
	return err
}
