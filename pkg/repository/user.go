package repository

import (
	"errors"
	"fmt"

	"jerseyhub/pkg/domain"
	interfaces "jerseyhub/pkg/repository/interface"
	"jerseyhub/pkg/utils/models"

	"gorm.io/gorm"
)

type userDatabase struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) interfaces.UserRepository {
	return &userDatabase{DB}
}

func (c *userDatabase) CheckUserAvailability(email string) bool {

	var count int
	query := fmt.Sprintf("select count(*) from users where email='%s'", email)
	if err := c.DB.Raw(query).Scan(&count).Error; err != nil {
		return false
	}
	// if count is greater than 0 that means the user already exist
	return count > 0

}

func (c *userDatabase) UserSignUp(user models.UserDetails) (models.UserDetailsResponse, error) {

	var userDetails models.UserDetailsResponse
	err := c.DB.Raw("INSERT INTO users (name, email, password, phone) VALUES (?, ?, ?, ?) RETURNING id, name, email, phone", user.Name, user.Email, user.Password, user.Phone).Scan(&userDetails).Error

	if err != nil {
		return models.UserDetailsResponse{}, err
	}

	return userDetails, nil
}

func (cr *userDatabase) UserBlockStatus(email string) (bool, error) {
	fmt.Println(email)
	var isBlocked bool
	err := cr.DB.Raw("select blocked from users where email = ?", email).Scan(&isBlocked).Error
	if err != nil {
		return false, err
	}
	fmt.Println(isBlocked)
	return isBlocked, nil
}

func (c *userDatabase) FindUserByEmail(user models.UserLogin) (models.UserSignInResponse, error) {

	var user_details models.UserSignInResponse

	err := c.DB.Raw(`
		SELECT *
		FROM users where email = ? and blocked = false
		`, user.Email).Scan(&user_details).Error

	if err != nil {
		return models.UserSignInResponse{}, errors.New("error checking user details")
	}

	return user_details, nil

}

// func (i *userDatabase) AddAddress(id int, address models.AddAddress) error {
// 	fmt.Println(id, address.Name, address.HouseName, address.Street, address.City, address.State, address.Pin)
// 	err := i.DB.Raw("insert into addresses(user_id,name,house_name,street,city,state,pin) values($1,$2,$3,$4,$5,$6,$7)returning id", id, address.Name, address.HouseName, address.Street, address.City, address.State, address.Pin).Error
// 	if err != nil {
// 		return err
// 	}

// 	return nil

// }

func (i *userDatabase) AddAddress(id int, address models.AddAddress) error {
	fmt.Println(id, address.Name, address.HouseName, address.Street, address.City, address.State, address.Pin)
	err := i.DB.Exec(`
		INSERT INTO addresses (user_id, name, house_name, street, city, state, pin)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id`,
		id, address.Name, address.HouseName, address.Street, address.City, address.State, address.Pin).Error
	if err != nil {
		return err
	}

	return nil
}

func (ad *userDatabase) GetAddresses(id int) ([]domain.Address, error) {

	var addresses []domain.Address

	if err := ad.DB.Raw("select * from addresses where user_id=?", id).Scan(&addresses).Error; err != nil {
		return []domain.Address{}, err
	}

	return addresses, nil

}

func (ad *userDatabase) GetUserDetails(id int) (models.UserDetailsResponse, error) {

	var details models.UserDetailsResponse

	if err := ad.DB.Raw("select id,name,email,phone from users where id=?", id).Scan(&details).Error; err != nil {
		return models.UserDetailsResponse{}, err
	}

	return details, nil

}

func (i *userDatabase) ChangePassword(id int, password string) error {

	err := i.DB.Exec("UPDATE users SET password=$1 WHERE id=$2", password, id).Error
	if err != nil {
		fmt.Println("Error updating password:", err)
		return err
	}

	return nil

}

func (i *userDatabase) GetPassword(id int) (string, error) {

	var userPassword string
	err := i.DB.Raw("select password from users where id = ?", id).Scan(&userPassword).Error
	if err != nil {
		return "", err
	}
	return userPassword, nil

}

func (ad *userDatabase) FindIdFromPhone(phone string) (int, error) {

	var id int

	if err := ad.DB.Raw("select id from users where phone=?", phone).Scan(&id).Error; err != nil {
		return id, err
	}

	return id, nil

}

func (i *userDatabase) EditName(id int, name string) error {
	err := i.DB.Exec(`update users set name=$1 where id=$2`, name, id).Error
	if err != nil {
		return err
	}

	return nil
}

func (i *userDatabase) EditEmail(id int, email string) error {
	err := i.DB.Exec(`update users set email=$1 where id=$2`, email, id).Error
	if err != nil {
		return err
	}

	return nil
}

func (i *userDatabase) EditPhone(id int, phone string) error {
	err := i.DB.Exec(`update users set phone=$1 where id=$2`, phone, id).Error
	if err != nil {
		return err
	}

	return nil
}

func (ad *userDatabase) GetCart(id int) ([]models.GetCart, error) {

	var cart []models.GetCart

	if err := ad.DB.Raw("select inventories.product_name,cart_products.quantity,cart_products.total_price from cart_products inner join inventories on cart_products.inventory_id=inventories.id where user_id=?", id).Scan(&cart).Error; err != nil {
		return []models.GetCart{}, err
	}

	return cart, nil

}

func (ad *userDatabase) RemoveFromCart(id int) error {

	if err := ad.DB.Exec(`delete from cart_products where id=$1`, id).Error; err != nil {
		return err
	}

	return nil

}

func (ad *userDatabase) UpdateQuantityAdd(id int) error {

	if err := ad.DB.Exec(`UPDATE cart_products
	SET quantity = quantity + 1,
	total_price = total_price + (
	SELECT price
	FROM inventories
	WHERE inventories.id = cart_products.inventory_id
	)
	WHERE id = $1;
	`, id).Error; err != nil {
		return err
	}

	return nil
}

func (ad *userDatabase) UpdateQuantityLess(id int) error {

	if err := ad.DB.Exec(`UPDATE cart_products
	SET quantity = quantity - 1,
	total_price = total_price - (
	SELECT price
	FROM inventories
	WHERE inventories.id = cart_products.inventory_id
	)
	WHERE id = $1;
	`, id).Error; err != nil {
		return err
	}

	return nil

}