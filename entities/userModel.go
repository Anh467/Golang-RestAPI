package entities

const UserModelTable = "Users"

type UserModel struct {
	UserID   int    `json:"userid" gorm:"column:UserID;primaryKey"`
	FullName string `json:"fullname" gorm:"column:FullName;"`
	Email    string `json:"email" gorm:"column:Email;"`
	Password string `json:"password" gorm:"column:Password;"`
	Role     string `json:"role" gorm:"column:Role;"`
}
