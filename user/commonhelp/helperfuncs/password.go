package helperfuncs

import "golang.org/x/crypto/bcrypt"

func PaaswordToHash(password string) string {
	passByte, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic("paasord not hashed")
	}

	return string(passByte)
}

// Check hashed pass and user entered is true or not
func CompareHashPassAndEnteredPass(hashed, entered string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(entered)); err != nil {
		return false
	} else {
		return true
	}
}
