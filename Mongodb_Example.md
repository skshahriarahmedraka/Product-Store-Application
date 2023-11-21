
## MongoDB Schema



### Collection: productcol
	Id             ObjectID         `json:"_id" bson:"_id"`
	Name           string             `json:"name" bson:"name" `
	Description    string             `json:"description" bson:"description" `
	Specifications string             `json:"specifications" bson:"specifications" `
	Brand_id       primitive.ObjectID `json:"brand_id" bson:"brand_id"`
	Catagory_id    primitive.ObjectID `json:"catagory_id" bson:"catagory_id" `
	Supplier_id    primitive.ObjectID `json:"supplier_id" bson:"supplier_id"`
	Unit_price     int                `json:"unit_price" bson:"unit_price" `
	Discount_price int                `json:"discount_price" bson:"discount_price"`
	Tags           []string           `json:"tags" bson:"tags"`
	Status_id      bool               `json:"status_id" bson:"status_id"`


### Collection: brandcol
	Id         primitive.ObjectID `json:"_id" bson:"_id"`
	Name       string             `json:"name" bson:"name"`
	Status_id  bool               `json:"status_id" bson:"status_id"`
	Created_at time.Time          `json:"created_at" bson:"created_at"`


### Collection: categorycol
	Id         primitive.ObjectID `json:"_id" bson:"_id"`
	Name       string             `json:"name" bson:"name" `
	Parent_id  primitive.ObjectID `json:"parent_id" bson:"parent_id" `
	Sequence   string             `json:"sequence" bson:"sequence" `
	Status_id  bool               `json:"status_id" bson:"status_id" `
	Created_at time.Time          `json:"created_at" bson:"created_at" `


### Collection: suppliercol
	Id                   primitive.ObjectID `json:"_id" bson:"_id"`
	Email                string             `json:"email" bson:"email" `
	Phone                string             `json:"phone" bson:"phone" `
	Status_id            bool               `json:"status_id" bson:"status_id"`
	Is_verified_supplier bool               `json:"is_verified_supplier" bson:"is_verified_supplier"`
	Created_at           time.Time          `json:"created_at" bson:"created_at"`

### Collection: productstockcol
	Id             primitive.ObjectID `json:"_id" bson:"_id"`
	Product_id     primitive.ObjectID `json:"product_id" bson:"product_id"`
	Stock_quantity int                `json:"stock_quantity" bson:"stock_quantity" `
	Updated_at     time.Time          `json:"updated_at" bson:"updated_at"`


