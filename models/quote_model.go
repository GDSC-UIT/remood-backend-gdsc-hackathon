package models

import (
	"context"
	"log"

	"remood/pkg/const/collections"
	"remood/pkg/database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Quote struct {
	BaseModel `json:",inline" bson:",inline"`

	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id"`
	Text   string             `json:"text"`
	Author string             `json:"author"`
}


func (q *Quote) CreateMany(quotes []Quote) ([]Quote, error) {
	// Create on database
	insert := make([]interface{}, 0)
	for i := range quotes {
		quotes[i].ID = primitive.NewObjectID()
		insert = append(insert, quotes[i])
	}

	log.Println(quotes)

	collection := database.GetMongoInstance().Db.Collection(string(collections.Quote))
	_, err := collection.InsertMany(context.Background(), insert)
	if err != nil {
		return quotes, err
	}

	return quotes, nil
}

func (q *Quote) GetRandom(number int) ([]Quote, error) {
	collection := database.GetMongoInstance().Db.Collection("quotes")

	var quotes []Quote
	pipeline := []interface{}{
		bson.M{
			"$sample": bson.M{
				"size": number,
			},
		},
	}
	cursor, err := collection.Aggregate(context.Background(), pipeline)
	if err != nil {
		log.Println(err)
		return quotes, err
	}

	if err = cursor.All(context.Background(), &quotes); err != nil {
		log.Println(err)
		return quotes, err
	}

	return quotes, nil
}