package entities

const UserModelTable = "Users"

// ERORR TAG
const FULL_NAME_BLANK = "Full Name can't be blank"
const EMAILL_BLANK = "Email can't be blank"
const PASS_WORD_BLANK = "Password can't be blank"
const EMAILL_WRONG_REGEX = "Email wrong format"
const EMAILL_DUPLICATE = "Email duplicate"

type UserModel struct {
	UserID   int    `json:"userid" gorm:"column:UserID;primaryKey"`
	FullName string `json:"fullname" gorm:"column:FullName;"`
	Email    string `json:"email" gorm:"column:Email;"`
	Password string `json:"password" gorm:"column:Password;"`
	Role     string `json:"role" gorm:"column:Role;"`
}

type UserCreateModel struct {
	UserID   int    `json:"userid" gorm:"column:UserID;primaryKey"`
	FullName string `json:"fullname" gorm:"column:FullName;"`
	Email    string `json:"email" gorm:"column:Email;"`
	Password string `json:"password" gorm:"column:Password;"`
	Role     string `json:"role" gorm:"column:Role;"`
}

type UserJWTModel struct {
	UserID   int    `json:"userid" gorm:"column:UserID;primaryKey"`
	FullName string `json:"fullname" gorm:"column:FullName;"`
	Email    string `json:"email" gorm:"column:Email;"`
	Role     string `json:"role" gorm:"column:Role;"`
	Token    string `json:"token" gorm:"column:Token;"`
}
