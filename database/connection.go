package database

import (
	"fmt"
	"main/config"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

var DB *gorm.DB

type DBConfig struct {
	RwDSN string // Read Write : create, update, delete
	RoDSN string // Read Only : select
}

func InitDB() (*gorm.DB, error) { // postgres 연동
	dns := generateDialector()
	db, err := gorm.Open(postgres.Open(dns.RwDSN), &gorm.Config{}) // err값이 null이면 RW 연동 성공
	if err != nil {
		return nil, err
	}

	err = db.Use(dbresolver.Register(dbresolver.Config{ // err값이 null이면 RO 연동 성공
		Replicas: []gorm.Dialector{postgres.Open(dns.RoDSN)}, // 레플리카 read only
		Policy:   dbresolver.RandomPolicy{},                  // 정책 설정
	}).SetMaxOpenConns(50).SetMaxIdleConns(10))
	if err != nil {
		return nil, err
	}

	DB = db

	return DB, nil
}

func generateDialector() DBConfig {
	cfg := DBConfig{
		// viper에서 dns 읽어오기
		RwDSN: fmt.Sprintf("host=%s:%d user=%s dbname=%s sslmode=disable password=%s",
			viper.GetString(config.DATABASE_HOST),
			viper.GetInt(config.DATABASE_PORT),
			viper.GetString(config.DATABASE_DATABASE),
			viper.GetString(config.DATABASE_RW_USER),
			viper.GetString(config.DATABASE_RW_PW),
		),
		RoDSN: fmt.Sprintf("host=%s:%d user=%s dbname=%s sslmode=disable password=%s",
			viper.GetString(config.DATABASE_HOST),
			viper.GetInt(config.DATABASE_PORT),
			viper.GetString(config.DATABASE_DATABASE),
			viper.GetString(config.DATABASE_RO_USER),
			viper.GetString(config.DATABASE_RO_PW),
		),
	}

	return cfg
}
