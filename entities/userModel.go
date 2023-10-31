package entities

import "github.com/dgrijalva/jwt-go"

const UserModelTable = "Users"
const ROLE_USER = "user"
const ROLE_ADMIN = "admin"
const ROLE_DELETED = "deleted"

type UserModel struct {
	UserID   int    `json:"userid" gorm:"column:UserID;primaryKey"`
	FullName string `json:"fullname" gorm:"column:FullName;"`
	Email    string `json:"email" gorm:"column:Email;"`
	Password string `json:"password" gorm:"column:Password;"`
	Role     string `json:"role" gorm:"column:Role;default:user"`
}

type UserCreateModel struct {
	UserID   int    `json:"userid" gorm:"column:UserID;primaryKey"`
	FullName string `json:"fullname" gorm:"column:FullName"`
	Email    string `json:"email" gorm:"column:Email"`
	Password string `json:"password" gorm:"column:Password"`
	Role     string `json:"role" gorm:"column:Role;default:user"`
}

type UserUpdate struct {
	FullName string `json:"fullname" gorm:"column:FullName"`
	Password string `json:"password" gorm:"column:Password"`
	Role     string `json:"role" gorm:"column:Role;default:user"`
}

type UserGetModel struct {
	UserID   int    `json:"userid" gorm:"column:UserID;primaryKey"`
	FullName string `json:"fullname" gorm:"column:FullName"`
	Email    string `json:"email" gorm:"column:Email"`
	Role     string `json:"role" gorm:"column:Role;default:user"`
}

type UserJWTModel struct {
	UserID   int    `json:"userid" gorm:"column:UserID;primaryKey"`
	FullName string `json:"fullname" gorm:"column:FullName;"`
	Email    string `json:"email" gorm:"column:Email;"`
	Role     string `json:"role" gorm:"column:Role;default:user"`
	Token    string `json:"token" gorm:"column:Token;"`
}

type Claims struct {
	UserID   int    `json:"userid"`
	Email    string `json:"email"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func (UserModel) TableName() string {
	return UserModelTable
}

func (UserCreateModel) TableName() string {
	return UserModelTable
}

func (UserGetModel) TableName() string {
	return UserModelTable
}

func (UserUpdate) TableName() string {
	return UserModelTable
}
