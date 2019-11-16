package database

import (
	"database/sql"

	"github.com/Linchpins-team/ilumi-backend/pkg/config"
	"github.com/Linchpins-team/ilumi-backend/pkg/model"
	"github.com/jinzhu/gorm"

	// driver
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type DB interface {
	Gorm() *gorm.DB
	SQL() *sql.DB
}

func New(conf config.Config) DB {
	db := initDB(conf)
	return dbImpl{
		db: db,
	}
}

func initDB(conf config.Config) *gorm.DB {
	db, err := gorm.Open("sqlite3", conf.DBPath)
	if err != nil {
		panic(err)
	}
	if err = db.AutoMigrate(&model.User{}, &model.Mission{}, &model.Answer{}, &model.AnswerReply{}, &model.Tag{}).Error; err != nil {
		panic(err)
	}
	return db
}

type dbImpl struct {
	db *gorm.DB
}

func (d dbImpl) Gorm() *gorm.DB {
	return d.db
}

func (d dbImpl) SQL() *sql.DB {
	return d.db.DB()
}
