package storage

import (
	"fmt"
	"net/url"

	"github.com/getsentry/sentry-go"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Storage struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	DB       *gorm.DB
}

var DBClient *Storage

func (s *Storage) NewSession() {
	dsn := url.URL{
		User:   url.UserPassword(s.User, s.Password),
		Scheme: "postgres",
		Host:   fmt.Sprintf("%s:%s", s.Host, s.Port),
		Path:   s.DBName,
	}
	db, err := gorm.Open(postgres.Open(dsn.String()), &gorm.Config{})
	if err != nil {
		sentry.CaptureMessage(fmt.Sprintf("Error initializing database: %v", err))
		panic("failed to connect database")
	}
	s.DB = db
	DBClient = s
}
