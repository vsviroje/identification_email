package utility

import (
	"errors"
	"fmt"
	"identification_email/config"
	"identification_email/constants"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/cast"
)

// TokenClaims ...
// Create a struct that will be encoded to a JWT.
// We add jwt.StandardClaims as an embedded type, to provide fields like expiry time
type TokenClaims struct {
	UserID interface{} `json:"sub"`
	jwt.StandardClaims
}

// CreateToken ...
// function will receive http request info as input
// based on user info it will create user token and it will return user token as response
func CreateToken(input interface{}) (userToken string, err error) {
	for {
		// getting expire time for token from configs
		var expiry string
		var ok bool
		if expiry, ok = config.Configs[constants.ConstToken][constants.ConstExpiry]; !ok {
			err = errors.New(constants.ConstTokenExpiryNotfound)
			break
		}
		expirationTime := time.Now().Add(time.Duration(cast.ToInt(expiry)) * time.Minute)
		// Create the JWT claims, which includes the username and expiry time
		claims := &TokenClaims{UserID: input,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expirationTime.Unix(),
				IssuedAt:  time.Now().Local().Unix(),
				NotBefore: time.Now().Local().Unix(),
			},
		}
		// getting jwt sceret from env
		var sceret string
		if sceret, ok = config.Configs[constants.ConstToken][constants.ConstSceret]; !ok {
			err = errors.New(constants.ConstTokenSceretNotFound)
			break
		}
		// Declare the token with the algorithm used for signing, and the claims
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		// Create the JWT string
		if userToken, err = token.SignedString([]byte(sceret)); err != nil {
			fmt.Println("error occured while genrating token", err)
			break
		}
		break
	}
	return
}

// ValidateToken ...
// function will receive token as input and it will check token valid or not with respective
// then signature and expairy time check also if all condition passed it will return user id
func ValidateToken(input string) (userID int, isValid bool, err error) {
	for {
		// getting jwt sceret from env
		var sceret string
		var ok bool
		if sceret, ok = config.Configs[constants.ConstToken][constants.ConstSceret]; !ok {
			err = errors.New(constants.ConstTokenSceretNotFound)
			break
		}
		// Initialize a new instance of `Claims`
		var token *jwt.Token
		tokenData := TokenClaims{}
		if token, err = jwt.ParseWithClaims(input, &tokenData, func(token *jwt.Token) (interface{}, error) {
			return []byte(sceret), nil
		}); err != nil {
			fmt.Println("Error occured while parsing token", err)
			break
		}
		// check for token expired or not
		if !token.Valid {
			fmt.Println("Error token already expired", token.Valid)
			err = nil
			break
		}
		// if token passed above all condition finally setting valid true
		isValid = true
		userID = cast.ToInt(tokenData.UserID)
		break
	}
	return
}
