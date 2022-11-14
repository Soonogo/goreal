package data

import (
	"goreal/internal/conf"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewDB, NewGreeterRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	db *gorm.DB
}

type Article struct {
	gorm.Model
	Slug        string
	Title       string
	Description string
	Body        string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, db *gorm.DB) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{db: db}, cleanup, nil
}

func NewDB(c *conf.Data) *gorm.DB {
	db, err := gorm.Open(mysql.Open(c.Database.Dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	if err := db.AutoMigrate(&Article{}); err != nil {
		panic(err)
	}
	return db
}
