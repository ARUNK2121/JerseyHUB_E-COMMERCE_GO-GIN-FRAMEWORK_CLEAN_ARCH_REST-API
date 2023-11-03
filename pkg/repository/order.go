package repository

import (
	"errors"
	"jerseyhub/pkg/domain"
	"jerseyhub/pkg/utils/models"

	"gorm.io/gorm"
)

type orderRepository struct {
	DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *orderRepository {
	return &orderRepository{
		DB: db,
	}
}

func (or *orderRepository) GetOrders(id int) ([]domain.Order, error) {

	var orders []domain.Order

	if err := or.DB.Raw("select * from orders where user_id=?", id).Scan(&orders).Error; err != nil {
		return []domain.Order{}, err
	}

	return orders, nil

}

func (ad *orderRepository) GetCart(id int) ([]models.GetCart, error) {

	var cart []models.GetCart

	if err := ad.DB.Raw("SELECT inventories.product_name,cart_products.quantity,cart_products.total_price AS Total FROM cart_products JOIN inventories ON cart_products.inventory_id=inventories.id WHERE user_id=$1", id).Scan(&cart).Error; err != nil {
		return []models.GetCart{}, err
	}
	return cart, nil

}

func (i *orderRepository) OrderItems(userid, addressid, paymentid int, total float64, coupon string) (int, error) {

	var id int
	query := `
    INSERT INTO orders (user_id,address_id, payment_method_id, final_price,coupon_used)
    VALUES (?, ?, ?, ?, ?)
    RETURNING id
    `
	i.DB.Raw(query, userid, addressid, paymentid, total, coupon).Scan(&id)

	return id, nil

}

func (i *orderRepository) AddOrderProducts(order_id int, cart []models.GetCart) error {

	query := `
    INSERT INTO order_items (order_id,inventory_id,quantity,total_price)
    VALUES (?, ?, ?, ?)
    `

	for _, v := range cart {
		var inv int
		if err := i.DB.Raw("select id from inventories where product_name=$1", v.ProductName).Scan(&inv).Error; err != nil {
			return err
		}

		if err := i.DB.Exec(query, order_id, inv, v.Quantity, v.Total).Error; err != nil {
			return err
		}
	}

	return nil

}

func (i *orderRepository) CancelOrder(id int) error {

	if err := i.DB.Exec("update orders set order_status='CANCELED' where id=$1", id).Error; err != nil {
		return err
	}

	return nil

}

func (i *orderRepository) EditOrderStatus(status string, id int) error {

	if err := i.DB.Exec("update orders set order_status=$1 where id=$2", status, id).Error; err != nil {
		return err
	}

	return nil

}

func (or *orderRepository) AdminOrders(status string) ([]domain.OrderDetails, error) {

	var orders []domain.OrderDetails
	if err := or.DB.Raw("SELECT orders.id AS id, users.name AS username, CONCAT('House Name:',addresses.house_name, ',', 'Street:', addresses.street, ',', 'City:', addresses.city, ',', 'State', addresses.state, ',', 'Phone:', addresses.phone) AS address, payment_methods.payment_name AS payment_method, orders.final_price As total FROM orders JOIN users ON users.id = orders.user_id JOIN payment_methods ON payment_methods.id = orders.payment_method_id JOIN addresses ON orders.address_id = addresses.id WHERE order_status = $1", status).Scan(&orders).Error; err != nil {
		return []domain.OrderDetails{}, err
	}

	return orders, nil

}

func (o *orderRepository) CheckOrder(orderID string, userID int) error {

	var count int
	err := o.DB.Raw("select count(*) from orders where order_id = ?", orderID).Scan(&count).Error
	if err != nil {
		return err
	}
	if count < 0 {
		return errors.New("no such order exist")
	}
	var checkUser int
	err = o.DB.Raw("select user_id from orders where order_id = ?", orderID).Scan(&checkUser).Error
	if err != nil {
		return err
	}

	if userID != checkUser {
		return errors.New("the order is not did by this user")
	}

	return nil
}

func (o *orderRepository) GetOrderDetail(orderID string) (domain.Order, error) {

	var orderDetails domain.Order
	err := o.DB.Raw("select * from orders where order_id = ?", orderID).Scan(&orderDetails).Error
	if err != nil {
		return domain.Order{}, err
	}

	return orderDetails, nil

}

func (i *orderRepository) ReturnOrder(id int) error {

	if err := i.DB.Exec("update orders set order_status='RETURNED' where id=$1", id).Error; err != nil {
		return err
	}

	return nil

}

func (o *orderRepository) CheckOrderStatusByID(id int) (string, error) {

	var status string
	err := o.DB.Raw("select order_status from orders where id = ?", id).Scan(&status).Error
	if err != nil {
		return "", err
	}

	return status, nil
}

func (o *orderRepository) FindAmountFromOrderID(id int) (float64, error) {

	var amount float64
	err := o.DB.Raw("select final_price from orders where id = ?", id).Scan(&amount).Error
	if err != nil {
		return 0, err
	}

	return amount, nil
}

func (i *orderRepository) CreditToUserWallet(amount float64, walletId int) error {

	if err := i.DB.Exec("update wallets set amount=$1 where id=$2", amount, walletId).Error; err != nil {
		return err
	}

	return nil

}

func (o *orderRepository) FindUserIdFromOrderID(id int) (int, error) {

	var userID int
	err := o.DB.Raw("select user_id from orders where id = ?", id).Scan(&userID).Error
	if err != nil {
		return 0, err
	}

	return userID, nil
}

func (o *orderRepository) FindWalletIdFromUserID(userId int) (int, error) {

	var count int
	err := o.DB.Raw("select count(*) from wallets where user_id = ?", userId).Scan(&count).Error
	if err != nil {
		return 0, err
	}

	var walletID int
	if count > 0 {
		err := o.DB.Raw("select id from wallets where user_id = ?", userId).Scan(&walletID).Error
		if err != nil {
			return 0, err
		}
	}

	return walletID, nil

}

func (o *orderRepository) CreateNewWallet(userID int) (int, error) {

	var walletID int
	err := o.DB.Exec("Insert into wallets(user_id,amount) values($1,$2)", userID, 0).Error
	if err != nil {
		return 0, err
	}

	if err := o.DB.Raw("select id from wallets where user_id=$1", userID).Scan(&walletID).Error; err != nil {
		return 0, err
	}

	return walletID, nil
}

func (o *orderRepository) MakePaymentStatusAsPaid(id int) error {

	err := o.DB.Exec("UPDATE orders SET payment_status = 'PAID' WHERE id = $1", id).Error
	if err != nil {
		return err
	}

	return nil
}

func (o *orderRepository) GetProductImagesInAOrder(id int) ([]string, error) {

	var images []string
	err := o.DB.Raw(`SELECT inventories.image
	FROM order_items 
	JOIN inventories ON inventories.id = order_items.inventory_id
	JOIN orders ON orders.id = order_items.order_id 
	WHERE orders.id = $1`, id).Scan(&images).Error
	if err != nil {
		return []string{}, err
	}

	return images, nil
}

func (o *orderRepository) GetIndividualOrderDetails(id int) (models.IndividualOrderDetails, error) {

	var details models.IndividualOrderDetails
	err := o.DB.Raw(`SELECT orders.id AS order_id,
	CONCAT('House Name:',addresses.house_name, ' ', 'Street:', addresses.street, ' ', 'City:', addresses.city, ' ', 'State', addresses.state) AS address,
	addresses.phone AS phone, 
	orders.coupon_used,
	payment_methods.payment_name AS payment_method, 
	orders.final_price As total_amount ,
	orders.order_status,
	orders.payment_status
	FROM orders 
	 JOIN payment_methods ON payment_methods.id = orders.payment_method_id 
	JOIN addresses ON orders.address_id = addresses.id 
	WHERE orders.id = $1`, id).Scan(&details).Error
	if err != nil {
		return models.IndividualOrderDetails{}, err
	}

	return details, nil
}

func (o *orderRepository) GetProductDetailsInOrder(id int) ([]models.ProductDetails, error) {

	var products []models.ProductDetails
	err := o.DB.Raw(`SELECT  inventories.product_name,
	inventories.image,
	order_items.quantity,
	order_items.total_price 
	FROM order_items 
	JOIN inventories ON inventories.id = order_items.inventory_id 
	JOIN orders ON order_items.order_id = orders.id 
	WHERE orders.id = $1`, id).Scan(&products).Error
	if err != nil {
		return []models.ProductDetails{}, err
	}

	return products, nil
}

func (o *orderRepository) FindPaymentMethodOfOrder(id int) (string, error) {

	var payment string

	if err := o.DB.Raw(`select payment_methods.payment_name
	 from payment_methods
	  join orders on orders.payment_method_id = payment_methods.id
	   where orders.id = $1`, id).Scan(&payment).Error; err != nil {
		return "", err
	}
	return payment, nil
}
