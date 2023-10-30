package entities

const ProductModelTable = "Products"

type Product struct {
	ProductID   int            `json:"productid" gorm:"column:ProductID;primaryKey"`
	Name        string         `json:"name" gorm:"column:Name"`
	Description string         `json:"description" gorm:"column:Description"`
	Price       float64        `json:"price" gorm:"column:Price"`
	Category    CategorieModel `json:"Category" gorm:"foreignKey:CategoryID"`
	CategoryID  int            `json:"categoryid" gorm:"column:CategoryID"`
}

type ProductCreation struct {
	ProductID   int    `json:"productid" gorm:"column:ProductID;primaryKey"`
	Name        string `json:"name" gorm:"column:Name;primaryKey"`
	Description int    `json:"description" gorm:"column:Description;primaryKey"`
	Price       int    `json:"price" gorm:"column:Price;primaryKey"`
	CategoryID  int    `json:"categoryid" gorm:"column:CategoryID;primaryKey"`
}

func (Product) TableName() string {
	return ProductModelTable
}

func (ProductCreation) TableName() string {
	return ProductModelTable
}
