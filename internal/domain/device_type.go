package domain

import (
	"time"

	"github.com/google/uuid"
)

type GoogleHome struct {
	Type            string   `bson:"type"`
	Traits          []string `bson:"traits"`
	WillReportState bool     `bson:"will_report_state"`
}

type DeviceType struct {
	Id          uuid.UUID  `bson:"_id"`
	Name        string     `bson:"name"`
	Description string     `bson:"description"`
	GoogleHome  GoogleHome `bson:"google_home"`
	CreatedAt   time.Time  `bson:"created_at"`
	UpdatedAt   time.Time  `bson:"updated_at"`
	DeletedAt   *time.Time `bson:"deleted_at"`
}
