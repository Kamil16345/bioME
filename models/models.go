package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"

)
type User struct{
	ID				primitive.ObjectID
	FirstName		*string
	LastName		*string
	Password		*string
	Email			*string
	Phone			*string
	Token			*string
	RefreshToken	*string
	CreatedAt		time.Time
	UpdatedAt		time.Time
	UserID			string
	UserCart		[]ProductUser
	AddressDetails	[]Address
	OrderStatus		[]Order
}
type Product struct {
	ProductID		primitive.ObjectID
	ProductName		string
	Price			uint64
	Rating
	Image
}
type ProductUser struct{
	ProductID		primitive.ObjectID
	ProductName
	Price
	Rating
	Image
}
type Address struct{
	AddressID		primitive.ObjectID
	House
	Street
	City
	PinCode
}
type Order struct{
	OrderID			primitive.ObjectID
	OrderCart
	OrderedAt
	Price
	Discount
	PaymentMethod
}
type Payment struct{
	Digital bool
	COD bool	//cash on delivery
}