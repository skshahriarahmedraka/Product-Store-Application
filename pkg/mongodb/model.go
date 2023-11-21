package mongodatabase

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LoginModel struct {
	Email    string `json:"email" bson:"email" gorm:"type:varchar(100)" validate:"min=2,max=100"`
	Password string `json:"password" bson:"password" gorm:"type:varchar(100)" validate:"min=2,max=100"`
}

type RegModel struct {
	Firstname string `json:"firstname" bson:"firstname" gorm:"type:varchar(100)" validate:"min=2,max=100"`
	Lastname  string `json:"lastname" bson:"lastname" gorm:"type:varchar(100)" validate:"min=2,max=100"`
	Email     string `json:"email" bson:"email" gorm:"type:varchar(100)" validate:"min=2,max=100"`
	Password  string `json:"password" bson:"password" gorm:"type:varchar(100)" validate:"min=2,max=100"`
}

type UserData struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	Firstname   string             `json:"firstname" bson:"firstname" gorm:"type:varchar(100)" validate:"min=2,max=100"`
	Lastname    string             `json:"lastname" bson:"lastname" gorm:"type:varchar(100)" validate:"min=2,max=100"`
	Email       string             `json:"email" bson:"email" gorm:"unique;not null;type:varchar(100)" validate:"required,min=3,max=100"`
	Password    string             `json:"password" bson:"password" gorm:"type:varchar(100)" validate:"required,min=3,max=50"`
	Telephone   string             `json:"telephone" bson:"telephone" gorm:"type:varchar(100)" validate:"min=2,max=100"`
	Address     string             `json:"address" bson:"address" gorm:"type:varchar(100)" validate:"min=2,max=100"`
	AccountType string             `json:"accounttype" bson:"accounttype" gorm:"type:varchar(100)" validate:"min=2,max=100"`
}

type Brands struct {
	Id         primitive.ObjectID `json:"_id" bson:"_id"`
	Name       string             `json:"name" bson:"name" gorm:"type:varchar(100)" validate:"min=2,max=100"`
	Status_id  bool               `json:"status_id" bson:"status_id"`
	Created_at time.Time          `json:"created_at" bson:"created_at"`
}

type Catagories struct {
	Id         primitive.ObjectID `json:"_id" bson:"_id"`
	Name       string             `json:"name" bson:"name" gorm:"type:varchar(100)" validate:"min=2,max=100"`
	Parent_id  primitive.ObjectID `json:"parent_id" bson:"parent_id" gorm:"type:varchar(100)" validate:"min=2,max=100"`
	Sequence   string             `json:"sequence" bson:"sequence" gorm:"type:varchar(100)" validate:"min=2,max=100"`
	Status_id  bool               `json:"status_id" bson:"status_id" `
	Created_at time.Time          `json:"created_at" bson:"created_at" `
}

type Suppliers struct {
	Id                   primitive.ObjectID `json:"_id" bson:"_id"`
	Email                string             `json:"email" bson:"email" gorm:"unique;not null;type:varchar(100)" validate:"required,min=3,max=100"`
	Phone                string             `json:"phone" bson:"phone" gorm:"type:varchar(100)" validate:"min=2,max=100"`
	Status_id            bool               `json:"status_id" bson:"status_id"`
	Is_verified_supplier bool               `json:"is_verified_supplier" bson:"is_verified_supplier"`
	Created_at           time.Time          `json:"created_at" bson:"created_at"`
}

type Products struct {
	Id             primitive.ObjectID `json:"_id" bson:"_id"`
	Name           string             `json:"name" bson:"name" gorm:"type:varchar(100)" validate:"min=2,max=100"`
	Description    string             `json:"description" bson:"description" gorm:"type:varchar(100)" validate:"min=2,max=100"`
	Specifications string             `json:"specifications" bson:"specifications" gorm:"type:varchar(100)" validate:"min=2,max=100"`
	Brand_id       primitive.ObjectID `json:"brand_id" bson:"brand_id"`
	Catagory_id    primitive.ObjectID `json:"catagory_id" bson:"catagory_id" `
	Supplier_id    primitive.ObjectID `json:"supplier_id" bson:"supplier_id"`
	Unit_price     int                `json:"unit_price" bson:"unit_price" `
	Discount_price int                `json:"discount_price" bson:"discount_price"`
	Tags           []string           `json:"tags" bson:"tags"`
	Status_id      bool               `json:"status_id" bson:"status_id"`
}

type Product_Stocks struct {
	Id             primitive.ObjectID `json:"_id" bson:"_id"`
	Product_id     primitive.ObjectID `json:"product_id" bson:"product_id"`
	Stock_quantity int             `json:"stock_quantity" bson:"stock_quantity" `
	Updated_at     time.Time          `json:"updated_at" bson:"updated_at"`
}

type Brand_List struct {
	Brand_id []primitive.ObjectID `json:"brand_id" bson:"brand_id"`
}
type Category_List struct {
	Category_id []primitive.ObjectID `json:"category_id" bson:"category_id"`
}