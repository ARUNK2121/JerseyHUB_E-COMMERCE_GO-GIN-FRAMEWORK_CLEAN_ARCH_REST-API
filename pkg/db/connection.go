package db

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	config "jerseyhub/pkg/config"
	domain "jerseyhub/pkg/domain"
)

func ConnectDatabase(cfg config.Config) (*gorm.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", cfg.DBHost, cfg.DBUser, cfg.DBName, cfg.DBPort, cfg.DBPassword)
	db, dbErr := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{SkipDefaultTransaction: true})

	if err := db.AutoMigrate(&domain.Inventories{}); err != nil {
		return db, err
	}
	if err := db.AutoMigrate(&domain.Category{}); err != nil {
		return db, err
	}
	if err := db.AutoMigrate(&domain.Users{}); err != nil {
		return db, err
	}
	if err := db.AutoMigrate(&domain.Admin{}); err != nil {
		return db, err
	}
	if err := db.AutoMigrate(domain.Cart{}); err != nil {
		return db, err
	}
	if err := db.AutoMigrate(domain.Address{}); err != nil {
		return db, err
	}
	if err := db.AutoMigrate(domain.Order{}); err != nil {
		return db, err
	}
	if err := db.AutoMigrate(domain.OrderItem{}); err != nil {
		return db, err
	}
	if err := db.AutoMigrate(domain.PaymentMethod{}); err != nil {
		return db, err
	}
	if err := db.AutoMigrate(domain.Coupons{}); err != nil {
		return db, err
	}
	if err := db.AutoMigrate(domain.Wallet{}); err != nil {
		return db, err
	}
	if err := db.AutoMigrate(domain.Offer{}); err != nil {
		return db, err
	}
	if err := db.AutoMigrate(domain.LineItems{}); err != nil {
		return db, err
	}
	if err := db.AutoMigrate(domain.Wishlist{}); err != nil {
		return db, err
	}
	CheckAndCreateAdmin(db)

	return db, dbErr
}

func CheckAndCreateAdmin(db *gorm.DB) {
	var count int64
	db.Model(&domain.Admin{}).Count(&count)
	if count == 0 {
		password := "comebuyjersey"
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return
		}

		admin := domain.Admin{
			ID:       1,
			Name:     "jerseyhub",
			Username: "jerseyhub@gmail.com",
			Password: string(hashedPassword),
		}
		db.Create(&admin)
	}
}
