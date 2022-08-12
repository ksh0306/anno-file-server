package handler

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/nicewook/authjwt/db"
	"github.com/nicewook/authjwt/helper"
	"github.com/nicewook/authjwt/models"

	"github.com/labstack/echo/v4"
)

// start - util
func CreateUser(user *models.User) (int64, error) {
	db := db.Connect()
	result, err := db.Exec("INSERT INTO users VALUES(?,?);", user.Username, user.Password)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func FindUser(user *models.User) (models.User, error) {
	log.Println("findUser from db: ", user.Username)
	foundUser := models.User{}
	var username, password string
	db := db.Connect()
	row := db.QueryRow("SELECT username, password FROM users WHERE username=?", user.Username)

	if err := row.Scan(&username, &password); err != nil {
		if err == sql.ErrNoRows {
			return foundUser, errors.New("fail to find user")
		}
		return foundUser, err
	}
	log.Println("username, password: ", username, password)
	foundUser.Username = username
	foundUser.Password = password
	return foundUser, nil
}

// end - util

func SignUp(c echo.Context) error {
	user := new(models.User)

	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "bad request",
		})
	}

	// 이미 이메일이 존재할 경우의 처리
	// if _, err := FindUser(user); err == nil { // no error, already exist
	// 	return c.JSON(http.StatusBadRequest, map[string]string{
	// 		"message": "existing user",
	// 	})
	// }

	// 비밀번호를 bycrypt 라이브러리로 해싱 처리
	log.Println("sign up attempt:", user.Username, user.Password)

	hashpw, err := helper.HashPassword(user.Password)
	if err != nil {
		log.Println("password hashing failed")
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	user.Password = hashpw

	// 위의 두단계에서 err가 nil일 경우 DB에 유저를 생성
	var userID int64

	if userID, err = CreateUser(user); err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed SignUp",
		})
	}

	// 모든 처리가 끝난 후 200, Success 메시지를 반환
	return c.JSON(http.StatusOK, map[string]string{
		"message": fmt.Sprintf("create user %s as ID %v successfully", user.Username, userID),
	})
}

func SignIn(c echo.Context) error {
	user := new(models.User)

	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "bad request",
		})
	}
	log.Println("log in attempt:", user.Username, user.Password)
	inputpw := user.Password

	// 존재하지않는 아이디일 경우
	var (
		foundUser models.User
		err       error
	)

	if foundUser, err = FindUser(user); err != nil { // not found
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "no user found",
		})
	}

	log.Println("user exist:", foundUser.Username, foundUser.Password)
	res := helper.CheckPasswordHash(foundUser.Password, inputpw)

	// 비밀번호 검증에 실패한 경우
	if !res {
		log.Println("wrong password:", foundUser.Password, inputpw)
		return echo.ErrUnauthorized
	}
	// 검증완료 클레임 생성
	accessToken, err := helper.CreateJWT(user.Username)
	if err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Login Success",
		"token":   accessToken,
	})
}

// 테스트용 API
func MockData() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Mock Data를 생성한다.
		list := map[string]string{
			"1": "고양이",
			"2": "사자",
			"3": "호랑이",
		}
		return c.JSON(http.StatusOK, list)
	}
}
