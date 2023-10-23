package controller

import (
	"fmt"
	"net/http"
	"regexp"
	"time"

	"github.com/GatorsTigers/ConcurrentBookingSystem/database"
	"github.com/GatorsTigers/ConcurrentBookingSystem/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// this should come from ENV or a configuration file
var jwtKey = []byte("cbs-jwt-token-key")
var jwtCookieKey = "cbs-jwt-token"

// Create a struct to read the username and password from the request body
type credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Create a struct that will be encoded to a JWT.
// We add jwt.StandardClaims as an embedded type, to provide fields like expiry time
type claims struct {
	EmailId string
	jwt.StandardClaims
}

// LoginUser method to login user by generating JWT Token and setting cookie
func LoginUser(context *gin.Context) {
	var creds credentials
	if err := context.BindJSON(&creds); err != nil {
		context.JSON(http.StatusBadRequest, "Structure of the request body is invalid")
		return
	}
	result, err := database.ValidateUserCredentials(creds.Email, creds.Password)
	if err != nil || !result {
		context.JSON(http.StatusUnauthorized, "Credentials invalid or some error occured")
		return
	}

	if result {
		// Declare the expiration time of the token
		// here, we have kept it as 15 minutes
		expirationTime := time.Now().Add(15 * time.Minute)
		// Create the JWT claims, which includes the username and expiry time
		claims := &claims{
			EmailId: creds.Email,
			StandardClaims: jwt.StandardClaims{
				// In JWT, the expiry time is expressed as unix milliseconds
				ExpiresAt: expirationTime.Unix(),
			},
		}

		// Declare the token with the algorithm used for signing, and the claims
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		// Create the JWT string
		tokenString, err := token.SignedString(jwtKey)
		if err != nil {
			// If there is an error in creating the JWT return an internal server error
			context.JSON(http.StatusInternalServerError, "Some error occured in generating a JWT token")
			return
		}

		// Finally, we set the client cookie for "referralboard-jwt-token" as the JWT we just generated
		// we also set an expiry time which is the same as the token itself
		fmt.Println("Here", jwtCookieKey)
		context.SetCookie(jwtCookieKey, tokenString, expirationTime.Second(), "/", "localhost", false, true)

		context.JSON(http.StatusOK, tokenString)
	}
}

// ValidateLogin is the middleware.
func ValidateLogin(handlerFunc gin.HandlerFunc) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		// We can obtain the session token from the requests cookies, which come with every request
		cookie, err := ctx.Cookie(jwtCookieKey)
		if err != nil {
			if err == http.ErrNoCookie {
				// If the cookie is not set, return an unauthorized status
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, "No cookie in cookies found")
				return
			}
			// For any other type of error, return a bad request status
			ctx.AbortWithStatusJSON(http.StatusBadRequest, "Some error occured in getting token from cookie")
			return
		}

		// Initialize a new instance of `Claims`
		claims := &claims{}

		// Parse the JWT string and store the result in `claims`.
		// Note that we are passing the key in this method as well. This method will return an error
		// if the token is invalid (if it has expired according to the expiry time we set on sign in),
		// or if the signature does not match
		tkn, err := jwt.ParseWithClaims(cookie, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				ctx.JSON(http.StatusUnauthorized, "Invalid signature in token cookie")
				return
			}
			ctx.JSON(http.StatusBadRequest, "Some error occured while verifying token cookie signature")
			return
		}
		if !tkn.Valid {
			ctx.JSON(http.StatusUnauthorized, "The token is not valid")
			return
		}
		handlerFunc(ctx)
	}
}

func getTokenBody(context *gin.Context) *claims {
	tknStr, _ := context.Cookie(jwtCookieKey)
	claims := &claims{}
	jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	return claims
}

// LogoutUser method logs out the user by deleting the cookie
func LogoutUser(context *gin.Context) {
	claims := getTokenBody(context)
	expirationTime := time.Now()
	claims.ExpiresAt = expirationTime.Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		context.JSON(http.StatusInternalServerError, "Some error occured while signing the token for 0 expiration time")
		return
	}

	// Set the new token as the users `token` cookie
	context.SetCookie(jwtCookieKey, tokenString, expirationTime.Second(), "/", "localhost", false, true)
	context.JSON(http.StatusOK, "Logged out the user")
}

func RegisterUser(context *gin.Context) {
	var user models.User
	var err error

	if err = context.BindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "could not parse user response",
		})
		return
	}
	// Check if email is valid
	if !isEmailValid(user.EmailId) {
		context.JSON(http.StatusBadRequest, "Error occured while trying to add user - email is not valid")
	} else {
		// Extract company domain from email
		// Add user
		err := database.AddUser(&user)
		if err != nil {
			context.JSON(http.StatusBadRequest, fmt.Sprintf("Could not add user %v", err.Error()))
		} else {
			context.JSON(http.StatusOK, "Registration Successfull")
		}
	}
}

// isEmailValid is for Email validation
func isEmailValid(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}
