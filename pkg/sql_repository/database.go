package sql_repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/gilperopiola/go-rest-example/pkg/common"
	"github.com/gilperopiola/go-rest-example/pkg/common/config"
	"github.com/gilperopiola/go-rest-example/pkg/common/models"

	mysqlLogger "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/*------------------
//  MySQL Database
/------------------*/

type Database struct {
	*gorm.DB
}

func NewDatabase() *Database {
	var database Database

	// Create connection. It's closed automatically by gorm
	// Retry connection if it fails due to Docker's orchestration
	if err := database.connect(); err != nil {
		log.Fatalf("error connecting to database: %v", err)
	}

	// Set connection pool limits
	// Log queries if debug = true
	// Destroy or clean tables
	// AutoMigrate fields
	// Create admin
	database.configure()

	return &database
}

/*---------------------------
//  Connect to DB & Ping it
//-------------------------*/

func (database *Database) connect() error {
	sqlConfig := common.Cfg.Database.SQL
	gormConfig := &gorm.Config{Logger: common.Logger}
	retries := 0
	var err error

	// Retry connection if it fails due to Docker's orchestration
	for retries < sqlConfig.MaxRetries {
		if database.DB, err = gorm.Open(mysql.Open(sqlConfig.GetMySQLConnectionString()), gormConfig); err == nil {
			break
		}

		if retries < sqlConfig.MaxRetries-1 {
			common.Logger.Info(context.Background(), "error connecting to mysql database, retrying... ", nil)
			time.Sleep(time.Duration(sqlConfig.RetryDelay) * time.Second)
		}

		retries++
	}

	// Return last db connection error
	if err != nil {
		return err
	}

	// Ping database to check if it's alive
	return database.GetSQLDB().Ping()
}

/*--------------------------
//    DB Configuration
//------------------------*/

func (database *Database) configure() {
	dbConfig := common.Cfg.Database
	sqlConfig := dbConfig.SQL

	// Set logger
	mysqlLogger.SetLogger(common.Logger)

	// Set connection pool limits
	setConnectionPoolLimits(database.GetSQLDB(), sqlConfig)

	// Destroy / Clean / AutoMigrate database
	prepareSchema(database.DB, dbConfig)

	// Insert admin user
	if dbConfig.AdminInsert {
		insertAdmin(database.DB, dbConfig)
	}

	// Just for formatting the logs :)
	if common.Cfg.LogInfo {
		fmt.Println("")
	}
}

func setConnectionPoolLimits(mySQLDB *sql.DB, sqlConfig config.SQL) {
	mySQLDB.SetMaxIdleConns(sqlConfig.MaxIdleConns)
	mySQLDB.SetMaxOpenConns(sqlConfig.MaxOpenConns)
	mySQLDB.SetConnMaxLifetime(time.Hour)
}

func prepareSchema(db *gorm.DB, dbConfig config.Database) {

	// Destroy or clean tables
	if dbConfig.Destroy {
		for _, model := range models.AllModels {
			db.Migrator().DropTable(model)
		}
	} else if dbConfig.Clean {
		for _, model := range models.AllModels {
			db.Delete(model)
		}
	}

	// AutoMigrate fields
	db.AutoMigrate(models.AllModels...)
}

func insertAdmin(db *gorm.DB, config config.Database) {
	adminPassword := common.Hash(common.Cfg.Database.AdminPassword, common.Cfg.HashSalt)
	admin := &models.User{Username: "admin", Email: "ferra.main@gmail.com", Password: adminPassword, IsAdmin: true}
	if err := db.Create(admin).Error; err != nil {
		fmt.Println(err.Error())
	}
}

/*----------------
//    Helpers
//---------------*/

func (database *Database) GetDB() *gorm.DB {
	if database == nil {
		return nil
	}
	return database.DB
}

func (database *Database) GetSQLDB() *sql.DB {
	if database == nil || database.DB == nil {
		return nil
	}
	sqlDB, _ := database.DB.DB()
	return sqlDB
}
