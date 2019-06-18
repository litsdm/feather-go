package feather

import "time"

type File struct {
	Name      string    `json:"name" bson:"name"`
	Size      float64   `json:"size" bson:"size"`
	S3Url     string    `json:"s3Url" bson:"s3Url"`
	Type      string    `json:"type" bson:"type"`
	IsGroup   bool      `json:"isGroup" bson:"isGroup"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	ExpiresAt time.Time `json:"expiresAt" bson:"expiresAt"`
}
