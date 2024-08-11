package infrastructure

import "golang.org/x/crypto/bcrypt"

func EncryptPassword(password string) (string, error) {
	cur_pass := []byte(password)
	encryptedPassword, err := bcrypt.GenerateFromPassword(cur_pass, bcrypt.DefaultCost)

	return string(encryptedPassword), err

}

func ValidatePassword(password string, userPassword string) bool {
	return bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(password)) == nil
}
