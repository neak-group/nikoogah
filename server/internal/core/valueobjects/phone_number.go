package valueobjects


type PhoneNumber struct{
	Region string
	PhoneNumber string `bson:"phone_number"`
}