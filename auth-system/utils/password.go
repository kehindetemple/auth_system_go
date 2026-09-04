package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(Password string)(string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(Password), bcrypt.DefaultCost )

	if err != nil{
		return "", err
	}

	return string(hashed), nil

}
func CheckPassword(user_password string, dp_password string)  error{

	compare := bcrypt.CompareHashAndPassword(
		[]byte(dp_password),
		[]byte(user_password) )

	return compare
}
