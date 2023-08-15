package gorm_client

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type IGormDB interface {
	Create(value interface{}) *gorm.DB
	Where(query interface{}, args ...interface{}) *gorm.DB
	Find(out interface{}, where ...interface{}) *gorm.DB
	Model(value interface{}) *gorm.DB
	Select(query interface{}, args ...interface{}) *gorm.DB
	AutoMigrate(dst ...interface{}) error
	Connect() IGormDB
}

type GormDB struct {
	client *gorm.DB
}

func NewGormDB() IGormDB {
	gormClient := &GormDB{}
	return gormClient.Connect()
}

func (db *GormDB) Connect() IGormDB {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      false,       // Disable color
		},
	)
	sqliteClient, err := gorm.Open(sqlite.Open("movieApp.db"), &gorm.Config{Logger: newLogger})
	if err != nil {
		log.Fatal("Failed to initialize database connection.", err.Error())
	}

	return &GormDB{
		client: sqliteClient,
	}
}

func (db *GormDB) AutoMigrate(dst ...interface{}) error {
	return db.client.AutoMigrate(dst...)
}

func (db *GormDB) Create(value interface{}) *gorm.DB {
	return db.client.Create(value)
}

func (db *GormDB) Where(query interface{}, args ...interface{}) *gorm.DB {
	return db.client.Where(query, args...)
}

func (db *GormDB) Find(out interface{}, where ...interface{}) *gorm.DB {
	return db.client.Find(out, where...)
}

func (db *GormDB) Model(value interface{}) *gorm.DB {
	return db.client.Model(value)
}

func (db *GormDB) Select(query interface{}, args ...interface{}) *gorm.DB {
	return db.client.Select(query, args...)
}
