package entities

const REVIEW_TABLE = "Reviews"

type ReviewModel struct {
	ReviewID  int          `json:"reviewid" gorm:"column:ReviewID ; primaryKey"`
	ProductID int          `json:"ProductID" gorm:"column:ProductID"`
	UserID    int          `json:"userid" gorm:"column:UserID"`
	Rating    int          `json:"rating" gorm:"column:Rating"`
	Comment   string       `json:"comment" gorm:"column:Comment"`
	Product   ProductModel `json:"product" gorm:"foreignKey:ProductID"`
	User      UserGetModel `json:"user" gorm:"foreignKey:UserID"`
}

type ReviewCreate struct {
	ProductID int    `json:"ProductID" gorm:"column:ProductID"`
	UserID    int    `json:"userid" gorm:"column:UserID"`
	Rating    int    `json:"rating" gorm:"column:Rating"`
	Comment   string `json:"comment" gorm:"column:Comment"`
}

type ReviewUpdate struct {
	ReviewID int    `json:"reviewid" gorm:"column:ReviewID ; primaryKey"`
	Rating   int    `json:"rating" gorm:"column:Rating"`
	Comment  string `json:"comment" gorm:"column:Comment"`
}

type ReviewGet struct {
	ReviewID  int          `json:"reviewid" gorm:"column:ReviewID ; primaryKey"`
	ProductID int          `json:"ProductID" gorm:"column:ProductID"`
	UserID    int          `json:"userid" gorm:"column:UserID"`
	Rating    int          `json:"rating" gorm:"column:Rating"`
	Comment   string       `json:"comment" gorm:"column:Comment"`
	User      UserGetModel `json:"user" gorm:"foreignKey:UserID"`
}

func (ReviewModel) TableName() string {
	return REVIEW_TABLE
}

func (ReviewCreate) TableName() string {
	return REVIEW_TABLE
}
func (ReviewUpdate) TableName() string {
	return REVIEW_TABLE
}
func (ReviewGet) TableName() string {
	return REVIEW_TABLE
}
