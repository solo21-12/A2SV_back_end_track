package infrastructure

import "golang.org/x/crypto/bcrypt"

func EncryptPassword(password string) (string, error) {
	cur_pass := []byte(password)
	encryptedPassword, err := bcrypt.GenerateFromPassword(cur_pass, bcrypt.DefaultCost)

	return string(encryptedPassword), err

}

func ValidatePassword(password string, hashedPassword string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) == nil
}
