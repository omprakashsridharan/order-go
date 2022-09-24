package db

import (
	"gorm.io/gorm"
)

var dbConn *gorm.DB

func init() {
	// envErr := godotenv.Load()
	// if envErr != nil {
	// 	log.Fatal("Error loading .env file")
	// }
	// dsn := os.Getenv("DATABASE_URL")
	// dsn := ""
	// fmt.Printf("%s", dsn)
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	panic(err)
	// }
	// sqlDB, err := db.DB()
	// if err != nil {
	// 	panic(err)
	// }
	// sqlDB.SetConnMaxIdleTime(time.Minute * 5)
	// sqlDB.SetMaxIdleConns(10)
	// sqlDB.SetMaxOpenConns(100)
	// sqlDB.SetConnMaxLifetime(time.Hour)
	// dbConn = db
}

func NewDb() (*DB, error) {
	sqlDB, err := dbConn.DB()
	db := &DB{
		connection: dbConn,
	}
	if err != nil {
		return db, err
	}
	if err := sqlDB.Ping(); err != nil {
		return db, err
	}
	return db, nil
}

type DB struct {
	connection *gorm.DB
}

type Database interface {
	Create(model interface{}) int64
}

func (db *DB) Create(model interface{}) int64 {
	return db.connection.Create(model).RowsAffected
}
