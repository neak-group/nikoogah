package entity


type Credibility struct{
	AverageRallyRating float32 `bson:"average_rally_rating"`
	AverageRallyPerYear float32 `bson:"average_rally_per_year"`
}