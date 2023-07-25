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

	db.AutoMigrate(&domain.Inventories{})
	db.AutoMigrate(&domain.Category{})
	db.AutoMigrate(&domain.Users{})
	db.AutoMigrate(&domain.Admin{})
	db.AutoMigrate(domain.Cart{})
	db.AutoMigrate(domain.Address{})
	db.AutoMigrate(domain.Order{})
	db.AutoMigrate(domain.OrderItem{})
	db.AutoMigrate(domain.PaymentMethod{})
	db.AutoMigrate(domain.Coupons{})
	db.AutoMigrate(domain.Wallet{})
	db.AutoMigrate(domain.Offer{})
	db.AutoMigrate(domain.LineItems{})
	db.AutoMigrate(domain.Wishlist{})
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
