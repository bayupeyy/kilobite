package auth

import "github.com/dgrijalva/jwt-go"

type Service interface {
	GenerateToken(userID int) (string, error)
}

type jwtService struct {
}

//Membuat Secret Key
var secretKey = []byte("koclok_clok")

func (s *jwtService) GenerateToken(userID int) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = userID

	//Membuat Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	//Signature untuk secret key
	signedToken, err := token.SignedString(secretKey)

	//Melakukan pengecekan error
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}
