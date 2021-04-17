package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	Id        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" gorm:"unique"`
	Password  []byte `json:"-"`
}

func (user *User) SetPassword(password string) {
	// パスワードのハッシュ化
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)

	user.Password = hashedPassword
}

func (user *User) ComparePassword(password string) error {
	// ユーザーのハッシュ化されたパスワードと引数の文字列を比較している
	return bcrypt.CompareHashAndPassword((user.Password), []byte(password))
}
