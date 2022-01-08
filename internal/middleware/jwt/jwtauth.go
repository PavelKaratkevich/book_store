package jwtAuth

import (
	// "errors"
	err "book_store/internal/response"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/twinj/uuid"
)

type User struct {
	// ID uint64            `json:"id"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

var user = User{
	Username: "username",
	Password: "password",
}

type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AccessUuid   string
	RefreshUuid  string
	AtExpires    int64
	RtExpires    int64
}

func Login() gin.HandlerFunc {
	var u User
	return func(c *gin.Context) {

		err := c.ShouldBindJSON(&u)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Please enter login and password"})
		} else {
			if u.Username != user.Username || u.Password != user.Password {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Please provide valid login details"})
				return
			}
			ts, err := CreateToken()
			if err != nil {
				c.JSON(http.StatusUnprocessableEntity, err.Error())
				return
			}
			tokens := map[string]string{
				"access_token":  ts.AccessToken,
				"refresh_token": ts.RefreshToken,
			}
			c.JSON(http.StatusOK, tokens)
		}
	}
}

func CreateToken() (*TokenDetails, error) {
	td := &TokenDetails{}
	td.AtExpires = time.Now().Add(time.Minute * 15).Unix()
	td.AccessUuid = uuid.NewV4().String()

	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	td.RefreshUuid = uuid.NewV4().String()

	var err error
	//Creating Access Token
	os.Getenv("ACCESS_SECRET") //this should be in an env file
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = td.AccessUuid
	// atClaims["user_id"] = userid
	atClaims["exp"] = td.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return nil, err
	}
	//Creating Refresh Token
	os.Setenv("REFRESH_SECRET", "mcmvmkmsdnfsdmfdsjf") //this should be in an env file
	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = td.RefreshUuid
	// rtClaims["user_id"] = userid
	rtClaims["exp"] = td.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(os.Getenv("REFRESH_SECRET")))
	if err != nil {
		return nil, err
	}
	return td, nil
}

func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	//normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func TokenValid(r *http.Request) error {
	token, err := VerifyToken(r)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

// CheckToken verifies if the token is provided, if it is verified and if it is valid 
func CheckToken(c *gin.Context) *err.AppError {
	var appErr err.AppError

	if token := ExtractToken(c.Request); token == "" {
		appErr.Code = http.StatusBadRequest
		appErr.Message = "Please provide a valid token"
		return &appErr
	}

	if _, err := VerifyToken(c.Request); err != nil {
		appErr.Code = http.StatusForbidden
		appErr.Message = "Unauthorized"
		return &appErr
	}

	if err := TokenValid(c.Request); err != nil {
		appErr.Code = http.StatusForbidden
		appErr.Message = "Unauthorized"
		return &appErr
	}
	return nil
}
