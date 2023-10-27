package entities

const CategorieModelTable = "Categories"

type CategorieModel struct {
	CategoryID int    `json:"categoryid" gorm:"column:CategoryID;primaryKey"`
	Name       string `json:"name" gorm:"column:Name;"`
}
