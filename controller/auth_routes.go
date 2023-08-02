package controller

import (
	"encoding/json"
	"net/http"
	"regexp"
	"time"

	"github.com/GatorsTigers/ConcurrentBookingSystem/database"
	"github.com/GatorsTigers/ConcurrentBookingSystem/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// this should come from ENV or a configuration file
var jwtKey = []byte("referralboard-jwt-token-key")

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

// Adapter is an alias so I dont have to type so much.
type Adapter func(http.Handler) http.Handler

// LoginUser method to login user by generating JWT Token and setting cookie
func LoginUser(context *gin.Context) {
	var creds credentials
	// Get the JSON body and decode into credentials
	err := json.NewDecoder(context.Request.Body).Decode(&creds)
	if err != nil {
		// If the structure of the body is wrong, return an HTTP error
		context.JSON(http.StatusBadRequest, "Structure of the request body is invalid")
		return
	}
	result, err := database.ValidateUserCredentials(creds.Email, creds.Password)
	if err != nil || !result {
		context.JSON(http.StatusUnauthorized, "Credentials invalid or some error occured")
		return
	}

	user, err1 := database.GetUserByEmailID(creds.Email)
	if err1 != nil {
		context.JSON(http.StatusBadRequest, "Some error occured in searching UserId for this email")
		return
	}
	if result {
		// Declare the expiration time of the token
		// here, we have kept it as 5 minutes
		expirationTime := time.Now().Add(30 * time.Minute)
		// Create the JWT claims, which includes the username and expiry time
		claims := &claims{
			EmailId: user.EmailId,
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
		context.SetCookie("referralboard-jwt-token", tokenString, expirationTime.Second())

		context.JSON(http.StatusOK, tokenString)
	}
}

/*
// ValidateLogin is the middleware.
func ValidateLogin(next ...http.Handler) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// We can obtain the session token from the requests cookies, which come with every request
		c, err := r.Cookie("referralboard-jwt-token")
		if err != nil {
			if err == http.ErrNoCookie {
				// If the cookie is not set, return an unauthorized status
				services.RespondError(w, http.StatusUnauthorized, "No cookie in cookies found")
				return
			}
			// For any other type of error, return a bad request status
			services.RespondError(w, http.StatusBadRequest, "Some error occured in getting token from cookie")
			return
		}

		// Get the JWT string from the cookie
		tknStr := c.Value

		// Initialize a new instance of `Claims`
		claims := &claims{}

		// Parse the JWT string and store the result in `claims`.
		// Note that we are passing the key in this method as well. This method will return an error
		// if the token is invalid (if it has expired according to the expiry time we set on sign in),
		// or if the signature does not match
		tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				services.RespondError(w, http.StatusUnauthorized, "Invalid signature in token cookie")
				return
			}
			services.RespondError(w, http.StatusBadRequest, "Some error occured while verifying token cookie signature")
			return
		}
		if !tkn.Valid {
			services.RespondError(w, http.StatusUnauthorized, "The token is not valid")
			return
		}
		if len(next) == 1 {
			next[0].ServeHTTP(w, r)
		} else {
			services.RespondJSON(w, http.StatusOK, tknStr)
		}
	}
}

func getTokenBody(context *gin.Context) *claims {
	tknStr, _ := context.Cookie("cbs-jwt-token")
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
		services.RespondError(w, http.StatusInternalServerError, "Some error occured while signing the token for 0 expiration time")
		return
	}

	// Set the new token as the users `token` cookie
	http.SetCookie(w, &http.Cookie{
		Name:    "referralboard-jwt-token",
		Value:   tokenString,
		Expires: expirationTime,
	})
	services.RespondJSON(w, http.StatusOK, "Logged out the user")
}


// GetUserByID wraps the GET  User by Id method
func GetUserByEmailID(context *gin.Context) {
	claims := getTokenBody(context)
	user, er := database.GetUserByEmailID(claims.EmailId)
	if er != nil {
		context.JSON(http.StatusBadRequest, "Error occured while trying to fetch user")
	} else {
		context.JSON(http.StatusOK, user)
	}
}*/

// AddUser wraps the POST User method
func AddUser(context *gin.Context) {
	var user models.User
	// fmt.Println(json.NewDecoder(r.Body))
	var screenJson []models.Screen
	if err := context.BindJSON(&screenJson); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "could not parse screen response",
		})
		return
	}
	// Check if email is valid
	if !isEmailValid(user.EmailId) {
		context.JSON(http.StatusBadRequest, "Error occured while trying to add user - email is not valid")
	} else {
		// Extract company domain from email
		// Add user
		newUser, er2 := database.AddUser(&user)
		if er2 != nil {
			context.JSON(http.StatusBadRequest, er2.Error())
		} else {
			context.JSON(http.StatusOK, newUser)
		}
	}
}

// isEmailValid is for Email validation
func isEmailValid(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}
