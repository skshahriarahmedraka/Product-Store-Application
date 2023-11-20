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
	Email    string `json:"email" bson:"email" gorm:"type:varchar(100)" validate:"min=2,max=100"`
	Password string `json:"password" bson:"password" gorm:"type:varchar(100)" validate:"min=2,max=100"`
}



type UserData struct {
	ID     primitive.ObjectID `json:"_id" bson:"_id"`
	Firstname string `json:"firstname" bson:"firstname" gorm:"type:varchar(100)" validate:"min=2,max=100"`
	Lastname  string `json:"lastname" bson:"lastname" gorm:"type:varchar(100)" validate:"min=2,max=100"`
	Email     string `json:"email" bson:"email" gorm:"unique;not null;type:varchar(100)" validate:"required,min=3,max=100"`
	Password  string `json:"password" bson:"password" gorm:"type:varchar(100)" validate:"required,min=3,max=50"`
	Telephone string `json:"telephone" bson:"telephone" gorm:"type:varchar(100)" validate:"min=2,max=100"`
	Address string `json:"address" bson:"address" gorm:"type:varchar(100)" validate:"min=2,max=100"`
	AccountType string `json:"accounttype" bson:"accounttype" gorm:"type:varchar(100)" validate:"min=2,max=100"`
}


type Brands struct {
	Id primitive.ObjectID
	Name string 
	Status_id string 
	Created_at time.Time 
}

type Catagories struct {
	Id primitive.ObjectID
	Name string
	Parent_id string
	Sequence string
	Status_id string
	Created_at time.Time
}

type Suppliers struct {
	Id primitive.ObjectID
	Email string
	Phone string
	Status_id string
	Is_verified_supplier string
	Created_at time.Time
}

type Products struct {
	Id primitive.ObjectID
	Name string
	Description string
	Specifications string
	Brand_id string
	Catagory_id string
	Supplier_id string
	Unit_price string
	Discount_price string
	Tags string
	Status_id string

}

type Product_Stocks struct {
	Id primitive.ObjectID
	Product_id string
	Stock_quantity string
	Updated_at time.Time
}