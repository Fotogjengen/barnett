package handlers

import (
	"barnett/api/database"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

const (
	MIN_PASSWORD_LENGTH                     = 8
	MIN_USERNAME_LENGTH                     = 6
	REQUIRED_PASSWORD_SYMBOLS_NUMBERS       = "1234567890"
	REQUIRED_PASSWORD_SYMBOLS_SPECIAL_SIGNS = "!#$%&/()=?|[]{}*@^-_.:;‚<>"
	REQUIRED_PASSWORD_LOWERCASE             = "qwertyuiopåasdfghjkløæzxcvbnm"
	REQUIRED_PASSWORD_UPPERCASE             = "QWERTYUIOPÅASDFGHJKLØÆZXCVBNM"

	SESSION_KEY_LENGTH = 20
	SESSION_KEY_BYTES = REQUIRED_PASSWORD_SYMBOLS_SPECIAL_SIGNS + REQUIRED_PASSWORD_SYMBOLS_NUMBERS + REQUIRED_PASSWORD_UPPERCASE + REQUIRED_PASSWORD_LOWERCASE
)

func hashPassword(pw string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pw), 14)
	return string(bytes), err
}

func checkPassword(pw, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pw))
	return err == nil
}

func signup(usr, pw string) (bool, string) {
	// Create query string
	q := `INSERT INTO bar_user(username, password_hash) VALUES('%s', '%s')`
	q = fmt.Sprintf(q, usr, pw)

	// Insert into db
	err := database.Insert(q)
	if err != nil {
		pqErr := err.(*pq.Error) // Convert to *pq.Error to be able to get error code

		errDesc := "Error: " + err.Error()
		fmt.Println(pqErr.Code.Name())
		if pqErr.Code == pq.ErrorCode(23505) { // User already exists
			errDesc = "Error: User already exists"
		}
		return false, errDesc
	}
	return true, "User created"
}

func Signup(ctx *gin.Context) {
	// Read body
	buf := make([]byte, 1024)
	num, _ := ctx.Request.Body.Read(buf) // Ignore error, will most likely get EOF (buffer too big)
	reqBody := string(buf[0:num])

	// Get username and password from body
	s := strings.Split(reqBody, "&")
	usr, pw := strings.Split(s[0], "=")[1], strings.Split(s[1], "=")[1]

	// Fulfill requirements
	reqStr := ""
	if len(usr) < MIN_USERNAME_LENGTH {
		reqStr += fmt.Sprintf("Username must be at least %d characters long \n", MIN_USERNAME_LENGTH)
	}
	if len(pw) < MIN_PASSWORD_LENGTH {
		reqStr += fmt.Sprintf("Password must be at least %d characters long \n", MIN_PASSWORD_LENGTH)
	}
	if !strings.ContainsAny(pw, REQUIRED_PASSWORD_LOWERCASE) {
		reqStr += "Password must contain at least one lowercase letter \n"
	}
	if !strings.ContainsAny(pw, REQUIRED_PASSWORD_UPPERCASE) {
		reqStr += "Password must contain at least one uppercase letter \n"
	}
	if !strings.ContainsAny(pw, REQUIRED_PASSWORD_SYMBOLS_NUMBERS) {
		reqStr += "Password must contain at least one number \n"
	}
	if !strings.ContainsAny(pw, REQUIRED_PASSWORD_SYMBOLS_SPECIAL_SIGNS) {
		reqStr += fmt.Sprintf("Password must contain at least one of following characters: '%s' \n",
			REQUIRED_PASSWORD_SYMBOLS_SPECIAL_SIGNS)
	}

	// If any requirements are not fulfilled, don't create user
	if len(reqStr) > 0 {
		ctx.JSON(http.StatusNotAcceptable, reqStr)
		return
	}

	// hash password
	pw, err := hashPassword(pw)

	if err != nil {
		log.Println("Error: ", err.Error())
	}

	if signedUp, str := signup(usr, pw); signedUp == false {
		ctx.JSON(http.StatusInternalServerError, str)
	} else {
		ctx.JSON(http.StatusOK, str)
	}
}

func generateSessionKey() string {
	/*
	Creates string of length SESSION_KEY_LENGTH
	Bytes in the string is from string constant SESSION_KEY_BYTES
	*/
	var key string
	for i := 0; i < SESSION_KEY_LENGTH; i++ {
		source := rand.NewSource(time.Now().UnixNano())
		random := rand.New(source)
		key += string(SESSION_KEY_BYTES[random.Intn(len(SESSION_KEY_BYTES))])
	}
	fmt.Println(key)
	return key
}

func createSession(id int) string {
	sessionKey := generateSessionKey()
	q := fmt.Sprintf("INSERT INTO user_session(session_key, user_id) VALUES('%s', %d);", sessionKey, id)
	err := database.Insert(q)
	if err != nil {
		fmt.Println("Error creating session: ", err.Error())
		return ""
	}
	return sessionKey
}

func Login(ctx *gin.Context) {
	// Read body
	buf := make([]byte, 1024)
	num, _ := ctx.Request.Body.Read(buf) // Ignore error, will most likely get EOF (buffer too big)
	reqBody := string(buf[0:num])

	// Get username and password from body
	s := strings.Split(reqBody, "&")
	usr, pw := strings.Split(s[0], "=")[1], strings.Split(s[1], "=")[1]

	// Get stored password hash from DB
	q := fmt.Sprintf(`SELECT password_hash, id FROM bar_user WHERE username='%s'`, usr)
	row := database.QueryOne(q)

	var password_hash string
	var id int
	err := row.Scan(&password_hash, &id)
	if err != nil {
		fmt.Println("Error: ", err.Error())
		if err.Error() == "sql: no rows in result set" {
			ctx.JSON(http.StatusNotFound, fmt.Sprintf("Could not find user %s", usr))
		} else {
			ctx.JSON(http.StatusInternalServerError, fmt.Sprintf("Something went wrong. Error: %s", err.Error()))
		}
		return
	}

	// Compare password hash from db and password from request body
	err = bcrypt.CompareHashAndPassword([]byte(password_hash), []byte(pw))
	if err != nil { // Passwords does not match
		ctx.JSON(http.StatusBadRequest, "Incorrect Password")
		return
	}

	// Create a new session
	sessionKey := createSession(id)
	if len(sessionKey) < SESSION_KEY_LENGTH {
		ctx.JSON(http.StatusBadRequest, "Failed to create session")
		return
	}

	// Set cookie
	ctx.SetCookie("sessionKey", sessionKey, 10, "/", "localhost", true, true)

	ctx.JSON(http.StatusOK, "Login successful")
}

func Logout(ctx *gin.Context) {
	cookie, err := ctx.Request.Cookie("sessionKey") // get sessionKey cookie
	if err != nil {
		// TODO: No cookie found, what to do?
		fmt.Println("No cookie set")
	}
	fmt.Println(cookie.Name)
}
