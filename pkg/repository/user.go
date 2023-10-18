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

func (c *userDatabase) UserSignUp(user models.UserDetails, referral string) (models.UserDetailsResponse, error) {

	var userDetails models.UserDetailsResponse
	err := c.DB.Raw("INSERT INTO users (name, email, password, phone,referral_code) VALUES (?, ?, ?, ?,?) RETURNING id, name, email, phone", user.Name, user.Email, user.Password, user.Phone, referral).Scan(&userDetails).Error

	if err != nil {
		return models.UserDetailsResponse{}, err
	}

	return userDetails, nil
}

func (cr *userDatabase) UserBlockStatus(email string) (bool, error) {
	var isBlocked bool
	err := cr.DB.Raw("select blocked from users where email = ?", email).Scan(&isBlocked).Error
	if err != nil {
		return false, err
	}
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

func (i *userDatabase) AddAddress(id int, address models.AddAddress, result bool) error {
	err := i.DB.Exec(`
		INSERT INTO addresses (user_id, name, house_name, street, city, state, phone, pin,"default")
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9 )`,
		id, address.Name, address.HouseName, address.Street, address.City, address.State, address.Phone, address.Pin, result).Error
	if err != nil {
		return errors.New("could not add address")
	}

	return nil
}

func (c *userDatabase) CheckIfFirstAddress(id int) bool {

	var count int
	// query := fmt.Sprintf("select count(*) from addresses where user_id='%s'", id)
	if err := c.DB.Raw("select count(*) from addresses where user_id=$1", id).Scan(&count).Error; err != nil {
		return false
	}
	// if count is greater than 0 that means the user already exist
	return count > 0

}

func (ad *userDatabase) GetAddresses(id int) ([]domain.Address, error) {

	var addresses []domain.Address

	if err := ad.DB.Raw("select * from addresses where user_id=?", id).Scan(&addresses).Error; err != nil {
		return []domain.Address{}, errors.New("error in getting addresses")
	}

	return addresses, nil

}

func (ad *userDatabase) GetUserDetails(id int) (models.UserDetailsResponse, error) {

	var details models.UserDetailsResponse

	if err := ad.DB.Raw("select id,name,email,phone from users where id=?", id).Scan(&details).Error; err != nil {
		return models.UserDetailsResponse{}, errors.New("could not get user details")
	}

	return details, nil

}

func (i *userDatabase) ChangePassword(id int, password string) error {

	err := i.DB.Exec("UPDATE users SET password=$1 WHERE id=$2", password, id).Error
	if err != nil {
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

func (ad *userDatabase) RemoveFromCart(cart, inventory int) error {

	if err := ad.DB.Exec(`DELETE FROM line_items WHERE cart_id = $1 AND inventory_id = $2`, cart, inventory).Error; err != nil {
		return err
	}

	return nil

}

func (ad *userDatabase) UpdateQuantityAdd(id, inv_id int) error {

	query := `
		UPDATE line_items
		SET quantity = quantity + 1
		WHERE cart_id=$1 AND inventory_id=$2
	`

	result := ad.DB.Exec(query, id, inv_id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (ad *userDatabase) UpdateQuantityLess(id, inv_id int) error {

	if err := ad.DB.Exec(`UPDATE line_items
	SET quantity = quantity - 1
	WHERE cart_id = $1 AND inventory_id=$2;
	`, id, inv_id).Error; err != nil {
		return err
	}

	return nil

}

func (ad *userDatabase) GetCartID(id int) (int, error) {

	var cart_id int

	if err := ad.DB.Raw("select id from carts where user_id=?", id).Scan(&cart_id).Error; err != nil {
		return 0, err
	}

	return cart_id, nil

}

func (ad *userDatabase) GetProductsInCart(cart_id int) ([]int, error) {

	var cart_products []int

	if err := ad.DB.Raw("select inventory_id from line_items where cart_id=?", cart_id).Scan(&cart_products).Error; err != nil {
		return []int{}, err
	}

	return cart_products, nil

}

func (ad *userDatabase) FindProductNames(inventory_id int) (string, error) {

	var product_name string

	if err := ad.DB.Raw("select product_name from inventories where id=?", inventory_id).Scan(&product_name).Error; err != nil {
		return "", err
	}

	return product_name, nil

}

func (ad *userDatabase) FindCartQuantity(cart_id, inventory_id int) (int, error) {

	var quantity int

	if err := ad.DB.Raw("select quantity from line_items where cart_id=$1 and inventory_id=$2", cart_id, inventory_id).Scan(&quantity).Error; err != nil {
		return 0, err
	}

	return quantity, nil

}

func (ad *userDatabase) FindPrice(inventory_id int) (float64, error) {

	var price float64

	if err := ad.DB.Raw("select price from inventories where id=?", inventory_id).Scan(&price).Error; err != nil {
		return 0, err
	}

	return price, nil

}

func (ad *userDatabase) FindCategory(inventory_id int) (int, error) {

	var category int

	if err := ad.DB.Raw("select category_id from inventories where id=?", inventory_id).Scan(&category).Error; err != nil {
		return 0, err
	}

	return category, nil

}

func (ad *userDatabase) FindofferPercentage(category_id int) (int, error) {
	var percentage int
	err := ad.DB.Raw("select discount_rate from offers where category_id=$1 and valid=true", category_id).Scan(&percentage).Error
	if err != nil {
		return 0, err
	}

	return percentage, nil
}

func (ad *userDatabase) FindUserFromReference(ref string) (int, error) {
	var user int

	if err := ad.DB.Raw("SELECT id FROM users WHERE referral_code = ?", ref).Find(&user).Error; err != nil {
		return 0, err
	}

	return user, nil
}

func (i *userDatabase) CreditReferencePointsToWallet(user_id int) error {
	err := i.DB.Exec("Update wallets set amount=amount+20 where user_id=$1", user_id).Error
	if err != nil {
		return err
	}

	return nil
}

func (i *userDatabase) GetReferralCodeFromID(id int) (string, error) {
	var referral string
	err := i.DB.Raw("SELECT referral_code FROM users WHERE id=?", id).Scan(&referral).Error
	if err != nil {
		return "", err
	}

	return referral, nil
}

func (i *userDatabase) FindProductImage(id int) (string, error) {
	var image string
	err := i.DB.Raw("SELECT image FROM inventories WHERE id = ?", id).Scan(&image).Error
	if err != nil {
		return "", err
	}

	return image, nil
}

func (i *userDatabase) FindStock(id int) (int, error) {
	var stock int
	err := i.DB.Raw("SELECT stock FROM inventories WHERE id = ?", id).Scan(&stock).Error
	if err != nil {
		return 0, err
	}

	return stock, nil
}
