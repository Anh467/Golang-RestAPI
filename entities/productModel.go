package entities

const ProductModelTable = "Products"

type Product struct {
	ProductID   int           `json:"productid" gorm:"column:ProductID;primaryKey"`
	Name        string        `json:"name" gorm:"column:Name"`
	Description string        `json:"description" gorm:"column:Description"`
	Price       float64       `json:"price" gorm:"column:Price"`
	Category    CategoryModel `json:"Category" gorm:"foreignKey:CategoryID"`
	CategoryID  int           `json:"categoryid" gorm:"column:CategoryID"`
}

type ProductCreation struct {
	ProductID   int     `json:"productid" gorm:"column:ProductID;primaryKey"`
	Name        string  `json:"name" gorm:"column:Name;primaryKey"`
	Description string  `json:"description" gorm:"column:Description;primaryKey"`
	Price       float64 `json:"price" gorm:"column:Price;primaryKey"`
	CategoryID  int     `json:"categoryid" gorm:"column:CategoryID;primaryKey"`
}

type ProductUpdate struct {
	Name        string  `json:"name" gorm:"column:Name;primaryKey"`
	Description string  `json:"description" gorm:"column:Description;primaryKey"`
	Price       float64 `json:"price" gorm:"column:Price;primaryKey"`
	CategoryID  int     `json:"categoryid" gorm:"column:CategoryID;primaryKey"`
}

func (Product) TableName() string {
	return ProductModelTable
}

func (ProductCreation) TableName() string {
	return ProductModelTable
}

func (ProductUpdate) TableName() string {
	return ProductModelTable
}
