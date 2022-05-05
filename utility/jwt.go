package utility

import (
	"errors"
	"identification_email/config"
	"identification_email/constants"
	"identification_email/utility/logger"
	"net/http"
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
	// getting expire time for token from configs
	logger.I("CreateToken invoked")
	defer logger.I("CreateToken returned")

	var expiry string
	var ok bool

	if expiry, ok = config.Configs[constants.ConstToken][constants.ConstExpiry]; !ok {
		logger.E(constants.ConstToken, constants.ConstExpiry, constants.ConstTokenExpiryNotfound)
		err = errors.New(constants.ConstTokenExpiryNotfound)
		return
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

	envMode := config.GetEnvMode()
	if envMode == "" {
		logger.E(constants.ConstEnvModeNotfound)
		err = errors.New(constants.ConstEnvModeNotfound)
		return
	}

	// getting jwt sceret from env
	sceret := config.GetValueByEnvMode(constants.ConstToken, envMode, constants.ConstSceret)
	if sceret == "" {
		logger.E(constants.ConstTokenSceretNotFound)
		err = errors.New(constants.ConstTokenSceretNotFound)
		return
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	if userToken, err = token.SignedString([]byte(sceret)); err != nil {
		logger.E("SignedString failed", err)
		return
	}

	return
}

// ValidateToken ...
// function will receive token as input and it will check token valid or not with respective
// then signature and expairy time check also if all condition passed it will return user id
func ValidateToken(input string) (userID int, isValid bool, err error) {
	// getting jwt sceret from env
	logger.I("ValidateToken invoked")
	defer logger.I("ValidateToken returned")

	envMode := config.GetEnvMode()
	if envMode == "" {
		logger.E(constants.ConstEnvModeNotfound)
		err = errors.New(constants.ConstEnvModeNotfound)
		return
	}

	sceret := config.GetValueByEnvMode(constants.ConstToken, envMode, constants.ConstSceret)
	if sceret == "" {
		logger.E(constants.ConstTokenSceretNotFound)
		err = errors.New(constants.ConstTokenSceretNotFound)
		return
	}

	// Initialize a new instance of `Claims`
	var token *jwt.Token
	tokenData := TokenClaims{}
	token, err = jwt.ParseWithClaims(input, &tokenData, func(token *jwt.Token) (interface{}, error) {
		logger.I(" func(token *jwt.Token) invoked and return")
		return []byte(sceret), nil
	})
	if err != nil {
		logger.E("ParseWithClaims failed", err)
		return
	}
	// check for token expired or not
	if !token.Valid {
		logger.E("token.Valid not valid")
		return
	}
	// if token passed above all condition finally setting valid true
	isValid = true
	userID = cast.ToInt(tokenData.UserID)

	return
}

func GetUserFromToken(r *http.Request) (userID int, err error) {
	logger.I("GetUserFromToken invoked")
	defer logger.I("GetUserFromToken returned")

	tokenStr := r.Header.Get("Authorization")

	envMode := config.GetEnvMode()
	if envMode == "" {
		logger.E(constants.ConstEnvModeNotfound)
		err = errors.New(constants.ConstEnvModeNotfound)
		return
	}

	sceret := config.GetValueByEnvMode(constants.ConstToken, envMode, constants.ConstSceret)
	if sceret == "" {
		logger.E(constants.ConstTokenSceretNotFound)
		err = errors.New(constants.ConstTokenSceretNotFound)
		return
	}

	tokenData := TokenClaims{}
	_, err = jwt.ParseWithClaims(tokenStr, &tokenData, func(token *jwt.Token) (interface{}, error) {
		logger.I(" func(token *jwt.Token) invoked and return")
		return []byte(sceret), nil
	})

	if err != nil {
		logger.E("ParseWithClaims failed", err)
		return
	}

	userID = cast.ToInt(tokenData.UserID)

	return
}
