package secure

import "golang.org/x/crypto/bcrypt"

func HashString(str string) (string, error) {

	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(str), 12)
	if err != nil {
		return "", err
	}

	return string(hashedBytes), nil
}

func CompareHashStrWithStr(hashStr string, str string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashStr), []byte(str))
	return err == nil
}
