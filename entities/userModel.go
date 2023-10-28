package entities

import "github.com/dgrijalva/jwt-go"

const UserModelTable = "Users"

var JwtKey = []byte("my_secret_key")

// ERORR TAG
const FULL_NAME_BLANK = "Full Name can't be blank"
const EMAILL_BLANK = "Email can't be blank"
const PASS_WORD_BLANK = "Password can't be blank"
const EMAILL_WRONG_REGEX = "Email wrong format"
const EMAILL_DUPLICATE = "Email duplicate"

// ERORR TAG JWT
const TOKEN_NOT_APPROPRIATE = "Token not appropriate"

type UserModel struct {
	UserID   int    `json:"userid" gorm:"column:UserID;primaryKey"`
	FullName string `json:"fullname" gorm:"column:FullName;"`
	Email    string `json:"email" gorm:"column:Email;"`
	Password string `json:"password" gorm:"column:Password;"`
	Role     string `json:"role" gorm:"column:Role;"`
}

type UserCreateModel struct {
	UserID   int    `json:"userid" gorm:"column:UserID;primaryKey"`
	FullName string `json:"fullname" gorm:"column:FullName"`
	Email    string `json:"email" gorm:"column:Email"`
	Password string `json:"password" gorm:"column:Password"`
	Role     string `json:"role" gorm:"column:Role"`
}

type UserJWTModel struct {
	UserID   int    `json:"userid" gorm:"column:UserID"`
	FullName string `json:"fullname" gorm:"column:FullName;"`
	Email    string `json:"email" gorm:"column:Email;"`
	Role     string `json:"role" gorm:"column:Role;"`
	Token    string `json:"token" gorm:"column:Token;"`
}

type Claims struct {
	UserID   int    `json:"userid"`
	Email    string `json:"email"`
	Password string `json:"password"`
	jwt.StandardClaims
}
