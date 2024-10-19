package valueobjects


type PhoneNumber struct{
	Region string	`bson:"region"`
	PhoneNumber string `bson:"phone_number"`
}