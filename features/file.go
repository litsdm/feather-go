package file

import (
	"context"
	"log"
	"net/http"
	"time"

	"feather.com/internal"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type File struct {
	ID        primitive.ObjectID   `bson:"_id,omitempty" json:"_id,omitempty"`
	Name      string               `json:"name" bson:"name"`
	Size      float64              `json:"size" bson:"size"`
	S3Url     string               `json:"s3Url" bson:"s3Url"`
	Type      string               `json:"type" bson:"type"`
	IsGroup   bool                 `json:"isGroup" bson:"isGroup"`
	To        []primitive.ObjectID `json:"to" bson:"to"`
	From      primitive.ObjectID   `json:"from" bson:"from"`
	CreatedAt time.Time            `json:"createdAt" bson:"createdAt"`
	ExpiresAt time.Time            `json:"expiresAt" bson:"expiresAt"`
}

func Routes(configuration *config.Config) *chi.Mux {
	router := chi.NewRouter()
	// router.Post("/", CreateFile(configuration))
	// router.Delete("/{userId}/files/{fileId}", DeleteFile(configuration))
	// router.Get("/", GetAllFiles(configuration))
	router.Get("/{userId}", GetUserFiles(configuration)) // MOVE TO USER
	return router
}

func GetUserFiles(configuration *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := chi.URLParam(r, "userId")
		objectID, _ := primitive.ObjectIDFromHex(userID)
		filter := bson.D{{Key: "to", Value: objectID}}

		collection := configuration.Database.Collection("files")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		cursor, err := collection.Find(ctx, filter)
		if err != nil {
			log.Fatal(err)
		}
		defer cursor.Close(ctx)

		var results []*File
		for cursor.Next(ctx) {
			file := &File{}
			er := cursor.Decode(file)
			if er != nil {
				log.Fatal(er)
			}
			results = append(results, file)
		}

		render.JSON(w, r, results)
	}
}
