package usecase

import (
	"errors"
	"fmt"

	"jerseyhub/pkg/config"
	"jerseyhub/pkg/domain"
	"jerseyhub/pkg/helper"
	interfaces "jerseyhub/pkg/repository/interface"
	"jerseyhub/pkg/utils/models"

	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
)

type userUseCase struct {
	userRepo            interfaces.UserRepository
	cfg                 config.Config
	otpRepository       interfaces.OtpRepository
	inventoryRepository interfaces.InventoryRepository
}

func NewUserUseCase(repo interfaces.UserRepository, cfg config.Config, otp interfaces.OtpRepository, inv interfaces.InventoryRepository) *userUseCase {
	return &userUseCase{
		userRepo:            repo,
		cfg:                 cfg,
		otpRepository:       otp,
		inventoryRepository: inv,
	}
}

func (u *userUseCase) UserSignUp(user models.UserDetails) (models.TokenUsers, error) {
	fmt.Println("add users")
	// Check whether the user already exist. If yes, show the error message, since this is signUp
	userExist := u.userRepo.CheckUserAvailability(user.Email)
	fmt.Println("user exists", userExist)
	if userExist {
		return models.TokenUsers{}, errors.New("user already exist, sign in")
	}
	fmt.Println(user)
	if user.Password != user.ConfirmPassword {
		return models.TokenUsers{}, errors.New("password does not match")
	}

	// Hash password since details are validated
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return models.TokenUsers{}, errors.New("internal server error")
	}
	user.Password = string(hashedPassword)

	// add user details to the database
	userData, err := u.userRepo.UserSignUp(user)
	if err != nil {
		return models.TokenUsers{}, err
	}

	// crete a JWT token string for the user
	tokenString, err := helper.GenerateTokenClients(userData)
	if err != nil {
		return models.TokenUsers{}, errors.New("could not create token due to some internal error")
	}

	// copies all the details except the password of the user
	var userDetails models.UserDetailsResponse
	err = copier.Copy(&userDetails, &userData)
	if err != nil {
		return models.TokenUsers{}, err
	}

	return models.TokenUsers{
		Users: userDetails,
		Token: tokenString,
	}, nil
}

func (u *userUseCase) LoginHandler(user models.UserLogin) (models.TokenUsers, error) {

	// checking if a username exist with this email address
	ok := u.userRepo.CheckUserAvailability(user.Email)
	if !ok {
		return models.TokenUsers{}, errors.New("the user does not exist")
	}

	isBlocked, err := u.userRepo.UserBlockStatus(user.Email)
	if err != nil {
		return models.TokenUsers{}, err
	}

	if isBlocked {
		return models.TokenUsers{}, errors.New("user is blocked by admin")
	}

	// Get the user details in order to check the password, in this case ( The same function can be reused in future )
	user_details, err := u.userRepo.FindUserByEmail(user)
	if err != nil {
		return models.TokenUsers{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user_details.Password), []byte(user.Password))
	if err != nil {
		return models.TokenUsers{}, errors.New("password incorrect")
	}

	var userDetails models.UserDetailsResponse
	err = copier.Copy(&userDetails, &user_details)
	if err != nil {
		return models.TokenUsers{}, err
	}

	tokenString, err := helper.GenerateTokenClients(userDetails)
	if err != nil {
		return models.TokenUsers{}, errors.New("could not create token")
	}

	return models.TokenUsers{
		Users: userDetails,
		Token: tokenString,
	}, nil

}

func (i *userUseCase) AddAddress(id int, address models.AddAddress) error {

	err := i.userRepo.AddAddress(id, address)
	if err != nil {
		return err
	}

	return nil

}

func (i *userUseCase) GetAddresses(id int) ([]domain.Address, error) {

	addresses, err := i.userRepo.GetAddresses(id)
	if err != nil {
		return []domain.Address{}, err
	}

	return addresses, nil

}

func (i *userUseCase) GetUserDetails(id int) (models.UserDetailsResponse, error) {

	details, err := i.userRepo.GetUserDetails(id)
	if err != nil {
		return models.UserDetailsResponse{}, err
	}

	return details, nil

}

func (i *userUseCase) ChangePassword(id int, old string, password string, repassword string) error {

	userPassword, err := i.userRepo.GetPassword(id)
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(old))
	if err != nil {
		return errors.New("password incorrect")
	}

	if password != repassword {
		return errors.New("passwords does not match")
	}

	newpassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return errors.New("internal server error")
	}

	return i.userRepo.ChangePassword(id, string(newpassword))

}

func (i *userUseCase) ForgotPasswordSend(phone string) error {

	ok := i.otpRepository.FindUserByMobileNumber(phone)
	if !ok {
		return errors.New("the user does not exist")
	}

	helper.TwilioSetup(i.cfg.ACCOUNTSID, i.cfg.AUTHTOKEN)
	fmt.Println("accsid:", i.cfg.ACCOUNTSID)
	fmt.Println("auth:", i.cfg.AUTHTOKEN)
	_, err := helper.TwilioSendOTP(phone, i.cfg.SERVICESID)
	if err != nil {
		return errors.New("error ocurred while generating OTP")
	}

	return nil

}

func (i *userUseCase) ForgotPasswordVerifyAndChange(model models.ForgotVerify) error {
	helper.TwilioSetup(i.cfg.ACCOUNTSID, i.cfg.AUTHTOKEN)
	err := helper.TwilioVerifyOTP(i.cfg.SERVICESID, model.Otp, model.Phone)
	if err != nil {
		return errors.New("error while verifying")
	}

	id, err := i.userRepo.FindIdFromPhone(model.Phone)
	if err != nil {
		return errors.New("cannot find user from mobile number")
	}

	newpassword, err := bcrypt.GenerateFromPassword([]byte(model.NewPassword), 10)
	if err != nil {
		return errors.New("hashing problem")
	}

	// if user is authenticated then change the password i the database
	if err := i.userRepo.ChangePassword(id, string(newpassword)); err != nil {
		return err
	}

	return nil
}

func (i *userUseCase) EditName(id int, name string) error {

	err := i.userRepo.EditName(id, name)
	if err != nil {
		return err
	}

	return nil

}

func (i *userUseCase) EditEmail(id int, email string) error {

	err := i.userRepo.EditName(id, email)
	if err != nil {
		return err
	}

	return nil

}

func (i *userUseCase) EditPhone(id int, phone string) error {

	err := i.userRepo.EditPhone(id, phone)
	if err != nil {
		return err
	}

	return nil

}

func (i *userUseCase) GetCart(id int) ([]models.GetCart, error) {

	cart, err := i.userRepo.GetCart(id)
	if err != nil {
		return []models.GetCart{}, err
	}

	return cart, nil

}

func (i *userUseCase) RemoveFromCart(id int) error {

	err := i.userRepo.RemoveFromCart(id)
	if err != nil {
		return err
	}

	return nil

}

func (i *userUseCase) UpdateQuantityAdd(id int) error {

	err := i.userRepo.UpdateQuantityAdd(id)
	if err != nil {
		return err
	}

	return nil

}

func (i *userUseCase) UpdateQuantityLess(id int) error {

	err := i.userRepo.UpdateQuantityLess(id)
	if err != nil {
		return err
	}

	return nil

}