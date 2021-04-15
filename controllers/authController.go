package controllers

import (
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/yuki0920/company_board/database"
	"github.com/yuki0920/company_board/models"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string

	// リクエストボディをdataにパースする
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	if data["password"] != data["password_confirm"] {
		return c.JSON(fiber.Map{
			"message": "password do not match",
		})
	}

	// パスワードのハッシュ化
	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user := models.User{
		FirstName: data["first_name"],
		LastName:  data["last_name"],
		Email:     data["email"],
		Password:  password,
	}

	database.DB.Create(&user)

	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User

	database.DB.Where("email = ?", data["email"]).First(&user)

	if user.Id == 0 {
		c.Status(404)
		return c.JSON(fiber.Map{
			"message": "not found",
		})
	}

	if err := bcrypt.CompareHashAndPassword((user.Password), []byte(data["password"])); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "incorrect password",
		})
	}
	clams := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		// strconv.FormatInt((int(user.id))) と同義
		Issuer:    strconv.Itoa(int(user.Id)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // 1day
	})

	token, err := clams.SignedString([]byte("secret"))

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24), // 1day
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

type Claims struct {
	jwt.StandardClaims
}

func User(c *fiber.Ctx) error {
	// クッキー取得
	cookie := c.Cookies("jwt")

	// デコード
	token, err := jwt.ParseWithClaims(cookie, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	if err != nil || !token.Valid {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	// structから要素Issuerを取り出すための書き方
	claims := token.Claims.(*Claims)

	var user models.User

	database.DB.Where("id = ?", claims.Issuer).First(&user)

	return c.JSON(user)
}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}
