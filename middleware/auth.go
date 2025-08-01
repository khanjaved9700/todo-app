package middleware

import "golang.org/x/crypto/bcrypt"

func HashPassword(pass string) ([]byte, error) {

	hashed, err := bcrypt.GenerateFromPassword([]byte(pass), 14)
	if err != nil {
		return nil, err
	}
	return hashed, nil
}

func CheckHashPassword(hashPass []byte, password string) bool {
	err := bcrypt.CompareHashAndPassword(hashPass, []byte(password))
	return err == nil
}
