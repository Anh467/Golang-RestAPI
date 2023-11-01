package entities

const CategorieModelTable = "Categories"

type CategoryModel struct {
	CategoryID int    `json:"categoryid" gorm:"column:CategoryID;primaryKey"`
	Name       string `json:"name" gorm:"column:Name;"`
}

type CategoryUpdate struct {
	Name string `json:"name" gorm:"column:Name;"`
}
type CategoryCreate struct {
	Name string `json:"name" gorm:"column:Name;"`
}

func (CategoryModel) TableName() string {
	return CategorieModelTable
}
func (CategoryUpdate) TableName() string {
	return CategorieModelTable
}
func (CategoryCreate) TableName() string {
	return CategorieModelTable
}
