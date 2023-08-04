package usecase

import (
	"errors"
	"testing"

	"jerseyhub/pkg/config"
	"jerseyhub/pkg/domain"
	"jerseyhub/pkg/mock/mockhelper"
	"jerseyhub/pkg/mock/mockrepo"
	"jerseyhub/pkg/utils/models"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_UserSignUp(t *testing.T) {
	ctrl := gomock.NewController(t)

	// Create mock implementations for the repositories
	userRepo := mockrepo.NewMockUserRepository(ctrl)
	orderRepo := mockrepo.NewMockOrderRepository(ctrl)
	otpRepo := mockrepo.NewMockOtpRepository(ctrl)
	inventoryRepo := mockrepo.NewMockInventoryRepository(ctrl)
	helper := mockhelper.NewMockHelper(ctrl)
	cfg := config.Config{}

	userUseCase := NewUserUseCase(userRepo, cfg, otpRepo, inventoryRepo, orderRepo, helper)

	testData := map[string]struct {
		input          models.UserDetails
		StubDetails    func(mockrepo.MockUserRepository, mockrepo.MockOrderRepository, models.UserDetails, mockhelper.MockHelper)
		expectedOutput models.TokenUsers
		expectedError  error
	}{
		"success": {
			input: models.UserDetails{
				Name:            "Arun K",
				Email:           "arthurbishop120@gmail.com",
				Phone:           "6282246077",
				Password:        gomock.Any().String(),
				ConfirmPassword: gomock.Any().String(),
			},
			StubDetails: func(userRepo mockrepo.MockUserRepository, orderRepo mockrepo.MockOrderRepository, signupData models.UserDetails, helper mockhelper.MockHelper) {
				gomock.InOrder(
					userRepo.EXPECT().CheckUserAvailability(signupData.Email).Times(1).Return(false),
					userRepo.EXPECT().FindUserFromReference("12345").Times(1).Return(1, nil),
					helper.EXPECT().PasswordHashing(signupData.Password).Times(1).Return(gomock.Any().String(), nil),
					helper.EXPECT().GenerateRefferalCode().Times(1).Return(gomock.Any().String(), nil),
					userRepo.EXPECT().UserSignUp(signupData, gomock.Any().String()).Times(1).Return(
						models.UserDetailsResponse{
							Id:    1,
							Name:  signupData.Name,
							Email: signupData.Email,
							Phone: signupData.Phone,
						}, nil,
					),
					helper.EXPECT().GenerateTokenClients(models.UserDetailsResponse{
						Id:    1,
						Name:  "Arun K",
						Email: "arthurbishop120@gmail.com",
						Phone: "6282246077",
					}).Times(1).Return(gomock.Any().String(), nil),
					userRepo.EXPECT().CreditReferencePointsToWallet(1).Times(1).Return(nil),
					orderRepo.EXPECT().CreateNewWallet(1).Times(1).Return(1, nil),
				)
			},
			expectedOutput: models.TokenUsers{
				Users: models.UserDetailsResponse{
					Id:    1,
					Name:  "Arun K",
					Email: "arthurbishop120@gmail.com",
					Phone: "6282246077",
				},
				Token: "ajjsjsjsjsjs.sjsjsjsjsjs.sjsjsjsjs",
			},
			expectedError: nil,
		},

		"user already exists": {
			input: models.UserDetails{
				Name:            "Arun K",
				Email:           "arthurbishop120@gmail.com",
				Phone:           "6282246077",
				Password:        gomock.Any().String(),
				ConfirmPassword: gomock.Any().String(),
			},
			StubDetails: func(userRepo mockrepo.MockUserRepository, orderRepo mockrepo.MockOrderRepository, signupData models.UserDetails, helper mockhelper.MockHelper) {
				gomock.InOrder(
					userRepo.EXPECT().CheckUserAvailability(signupData.Email).Times(1).Return(true),
				)
			},
			expectedOutput: models.TokenUsers{},
			expectedError:  errors.New("user already exist, sign in"),
		},

		"password missmatch": {
			input: models.UserDetails{
				Name:            "Arun K",
				Email:           "arthurbishop120@gmail.com",
				Phone:           "6282246077",
				Password:        gomock.Any().String(),
				ConfirmPassword: "shshsh",
			},
			StubDetails: func(userRepo mockrepo.MockUserRepository, orderRepo mockrepo.MockOrderRepository, signupData models.UserDetails, helper mockhelper.MockHelper) {
				gomock.InOrder(
					userRepo.EXPECT().CheckUserAvailability(signupData.Email).Times(1).Return(false),
				)
			},
			expectedOutput: models.TokenUsers{},
			expectedError:  errors.New("password does not match"),
		},

		"could not find the owner of reference id": {
			input: models.UserDetails{
				Name:            "Arun K",
				Email:           "arthurbishop120@gmail.com",
				Phone:           "6282246077",
				Password:        gomock.Any().String(),
				ConfirmPassword: gomock.Any().String(),
			},
			StubDetails: func(userRepo mockrepo.MockUserRepository, orderRepo mockrepo.MockOrderRepository, signupData models.UserDetails, helper mockhelper.MockHelper) {
				gomock.InOrder(
					userRepo.EXPECT().CheckUserAvailability(signupData.Email).Times(1).Return(false),
					userRepo.EXPECT().FindUserFromReference("12345").Times(1).Return(0, errors.New("cannot find reference user")),
				)
			},
			expectedOutput: models.TokenUsers{},
			expectedError:  errors.New("cannot find reference user"),
		},

		"password hashing problem": {
			input: models.UserDetails{
				Name:            "Arun K",
				Email:           "arthurbishop120@gmail.com",
				Phone:           "6282246077",
				Password:        gomock.Any().String(),
				ConfirmPassword: gomock.Any().String(),
			},
			StubDetails: func(userRepo mockrepo.MockUserRepository, orderRepo mockrepo.MockOrderRepository, signupData models.UserDetails, helper mockhelper.MockHelper) {
				gomock.InOrder(
					userRepo.EXPECT().CheckUserAvailability(signupData.Email).Times(1).Return(false),
					userRepo.EXPECT().FindUserFromReference("12345").Times(1).Return(1, nil),
					helper.EXPECT().PasswordHashing(signupData.Password).Times(1).Return(gomock.Any().String(), errors.New("error in hashing password")),
				)
			},
			expectedOutput: models.TokenUsers{},
			expectedError:  errors.New("error in hashing password"),
		},

		"could not generate reference code": {
			input: models.UserDetails{
				Name:            "Arun K",
				Email:           "arthurbishop120@gmail.com",
				Phone:           "6282246077",
				Password:        gomock.Any().String(),
				ConfirmPassword: gomock.Any().String(),
			},
			StubDetails: func(userRepo mockrepo.MockUserRepository, orderRepo mockrepo.MockOrderRepository, signupData models.UserDetails, helper mockhelper.MockHelper) {
				gomock.InOrder(
					userRepo.EXPECT().CheckUserAvailability(signupData.Email).Times(1).Return(false),
					userRepo.EXPECT().FindUserFromReference("12345").Times(1).Return(1, nil),
					helper.EXPECT().PasswordHashing(signupData.Password).Times(1).Return(gomock.Any().String(), nil),
					helper.EXPECT().GenerateRefferalCode().Times(1).Return(gomock.Any().String(), errors.New("error in creating reference id")),
				)
			},
			expectedOutput: models.TokenUsers{},
			expectedError:  errors.New("internal server error"),
		},

		"could not add the user to database": {
			input: models.UserDetails{
				Name:            "Arun K",
				Email:           "arthurbishop120@gmail.com",
				Phone:           "6282246077",
				Password:        gomock.Any().String(),
				ConfirmPassword: gomock.Any().String(),
			},
			StubDetails: func(userRepo mockrepo.MockUserRepository, orderRepo mockrepo.MockOrderRepository, signupData models.UserDetails, helper mockhelper.MockHelper) {
				gomock.InOrder(
					userRepo.EXPECT().CheckUserAvailability(signupData.Email).Times(1).Return(false),
					userRepo.EXPECT().FindUserFromReference("12345").Times(1).Return(1, nil),
					helper.EXPECT().PasswordHashing(signupData.Password).Times(1).Return(gomock.Any().String(), nil),
					helper.EXPECT().GenerateRefferalCode().Times(1).Return(gomock.Any().String(), nil),
					userRepo.EXPECT().UserSignUp(signupData, gomock.Any().String()).Times(1).Return(
						models.UserDetailsResponse{}, errors.New("could not add the user"),
					),
				)
			},
			expectedOutput: models.TokenUsers{},
			expectedError:  errors.New("could not add the user"),
		},

		"could not generate the token": {
			input: models.UserDetails{
				Name:            "Arun K",
				Email:           "arthurbishop120@gmail.com",
				Phone:           "6282246077",
				Password:        gomock.Any().String(),
				ConfirmPassword: gomock.Any().String(),
			},
			StubDetails: func(userRepo mockrepo.MockUserRepository, orderRepo mockrepo.MockOrderRepository, signupData models.UserDetails, helper mockhelper.MockHelper) {
				gomock.InOrder(
					userRepo.EXPECT().CheckUserAvailability(signupData.Email).Times(1).Return(false),
					userRepo.EXPECT().FindUserFromReference("12345").Times(1).Return(1, nil),
					helper.EXPECT().PasswordHashing(signupData.Password).Times(1).Return(gomock.Any().String(), nil),
					helper.EXPECT().GenerateRefferalCode().Times(1).Return(gomock.Any().String(), nil),
					userRepo.EXPECT().UserSignUp(signupData, gomock.Any().String()).Times(1).Return(
						models.UserDetailsResponse{
							Id:    1,
							Name:  signupData.Name,
							Email: signupData.Email,
							Phone: signupData.Phone,
						}, nil,
					),
					helper.EXPECT().GenerateTokenClients(models.UserDetailsResponse{
						Id:    1,
						Name:  "Arun K",
						Email: "arthurbishop120@gmail.com",
						Phone: "6282246077",
					}).Times(1).Return(gomock.Any().String(), errors.New("could not generate the token")),
				)
			},
			expectedOutput: models.TokenUsers{},
			expectedError:  errors.New("could not create token due to some internal error"),
		},

		"could not credit the amount ": {
			input: models.UserDetails{
				Name:            "Arun K",
				Email:           "arthurbishop120@gmail.com",
				Phone:           "6282246077",
				Password:        gomock.Any().String(),
				ConfirmPassword: gomock.Any().String(),
			},
			StubDetails: func(userRepo mockrepo.MockUserRepository, orderRepo mockrepo.MockOrderRepository, signupData models.UserDetails, helper mockhelper.MockHelper) {
				gomock.InOrder(
					userRepo.EXPECT().CheckUserAvailability(signupData.Email).Times(1).Return(false),
					userRepo.EXPECT().FindUserFromReference("12345").Times(1).Return(1, nil),
					helper.EXPECT().PasswordHashing(signupData.Password).Times(1).Return(gomock.Any().String(), nil),
					helper.EXPECT().GenerateRefferalCode().Times(1).Return(gomock.Any().String(), nil),
					userRepo.EXPECT().UserSignUp(signupData, gomock.Any().String()).Times(1).Return(
						models.UserDetailsResponse{
							Id:    1,
							Name:  signupData.Name,
							Email: signupData.Email,
							Phone: signupData.Phone,
						}, nil,
					),
					helper.EXPECT().GenerateTokenClients(models.UserDetailsResponse{
						Id:    1,
						Name:  "Arun K",
						Email: "arthurbishop120@gmail.com",
						Phone: "6282246077",
					}).Times(1).Return(gomock.Any().String(), nil),
					userRepo.EXPECT().CreditReferencePointsToWallet(1).Times(1).Return(errors.New("error in crediting amount")),
				)
			},
			expectedOutput: models.TokenUsers{},
			expectedError:  errors.New("error in crediting gift"),
		},

		"could not create the wallet": {
			input: models.UserDetails{
				Name:            "Arun K",
				Email:           "arthurbishop120@gmail.com",
				Phone:           "6282246077",
				Password:        gomock.Any().String(),
				ConfirmPassword: gomock.Any().String(),
			},
			StubDetails: func(userRepo mockrepo.MockUserRepository, orderRepo mockrepo.MockOrderRepository, signupData models.UserDetails, helper mockhelper.MockHelper) {
				gomock.InOrder(
					userRepo.EXPECT().CheckUserAvailability(signupData.Email).Times(1).Return(false),
					userRepo.EXPECT().FindUserFromReference("12345").Times(1).Return(1, nil),
					helper.EXPECT().PasswordHashing(signupData.Password).Times(1).Return(gomock.Any().String(), nil),
					helper.EXPECT().GenerateRefferalCode().Times(1).Return(gomock.Any().String(), nil),
					userRepo.EXPECT().UserSignUp(signupData, gomock.Any().String()).Times(1).Return(
						models.UserDetailsResponse{
							Id:    1,
							Name:  signupData.Name,
							Email: signupData.Email,
							Phone: signupData.Phone,
						}, nil,
					),
					helper.EXPECT().GenerateTokenClients(models.UserDetailsResponse{
						Id:    1,
						Name:  "Arun K",
						Email: "arthurbishop120@gmail.com",
						Phone: "6282246077",
					}).Times(1).Return(gomock.Any().String(), nil),
					userRepo.EXPECT().CreditReferencePointsToWallet(1).Times(1).Return(nil),
					orderRepo.EXPECT().CreateNewWallet(1).Times(1).Return(1, errors.New("errror in creating new wallet")),
				)
			},
			expectedOutput: models.TokenUsers{},
			expectedError:  errors.New("errror in creating new wallet"),
		},
	}

	for _, test := range testData {

		test.StubDetails(*userRepo, *orderRepo, test.input, *helper)

		tokenusers, err := userUseCase.UserSignUp(test.input, "12345")

		assert.Equal(t, test.expectedOutput.Users.Id, tokenusers.Users.Id)
		assert.Equal(t, test.expectedOutput.Users.Name, tokenusers.Users.Name)
		assert.Equal(t, test.expectedOutput.Users.Email, tokenusers.Users.Email)
		assert.Equal(t, test.expectedError, err)

	}
}

func Test_LoginHandler(t *testing.T) {
	ctrl := gomock.NewController(t)

	// Create mock implementations for the repositories
	userRepo := mockrepo.NewMockUserRepository(ctrl)
	orderRepo := mockrepo.NewMockOrderRepository(ctrl)
	otpRepo := mockrepo.NewMockOtpRepository(ctrl)
	inventoryRepo := mockrepo.NewMockInventoryRepository(ctrl)
	helper := mockhelper.NewMockHelper(ctrl)
	cfg := config.Config{}

	userUseCase := NewUserUseCase(userRepo, cfg, otpRepo, inventoryRepo, orderRepo, helper)

	testData := map[string]struct {
		input          models.UserLogin
		StubDetails    func(mockrepo.MockUserRepository, mockrepo.MockOrderRepository, models.UserLogin, mockhelper.MockHelper)
		expectedOutput models.TokenUsers
		expectedError  error
	}{
		"success": {
			input: models.UserLogin{
				Email:    "arthurbishop120@gmail.com",
				Password: "4321",
			},
			StubDetails: func(userRepo mockrepo.MockUserRepository, orderRepo mockrepo.MockOrderRepository, Data models.UserLogin, helper mockhelper.MockHelper) {
				gomock.InOrder(
					userRepo.EXPECT().CheckUserAvailability(Data.Email).Times(1).Return(true),
					userRepo.EXPECT().UserBlockStatus(Data.Email).Times(1).Return(false, nil),
					userRepo.EXPECT().FindUserByEmail(Data).Times(1).Return(models.UserSignInResponse{
						Id:       1,
						UserID:   1,
						Name:     "Arun K",
						Email:    "arthurbishop120@gmail.com",
						Phone:    "6282246077",
						Password: "4321",
					}, nil),
					helper.EXPECT().CompareHashAndPassword("4321", "4321").Times(1).Return(nil),
					helper.EXPECT().GenerateTokenClients(models.UserDetailsResponse{
						Id:    1,
						Name:  "Arun K",
						Email: "arthurbishop120@gmail.com",
						Phone: "6282246077",
					}).Times(1).Return(gomock.Any().String(), nil),
				)
			},
			expectedOutput: models.TokenUsers{
				Users: models.UserDetailsResponse{
					Id:    1,
					Name:  "Arun K",
					Email: "arthurbishop120@gmail.com",
					Phone: "6282246077",
				},
				Token: "ajjsjsjsjsjs.sjsjsjsjsjs.sjsjsjsjs",
			},
			expectedError: nil,
		},
		"no user exists": {
			input: models.UserLogin{
				Email:    "arthurbishop120@gmail.com",
				Password: "4321",
			},
			StubDetails: func(userRepo mockrepo.MockUserRepository, orderRepo mockrepo.MockOrderRepository, Data models.UserLogin, helper mockhelper.MockHelper) {
				gomock.InOrder(
					userRepo.EXPECT().CheckUserAvailability(Data.Email).Times(1).Return(false),
				)
			},
			expectedOutput: models.TokenUsers{
				Users: models.UserDetailsResponse{},
				Token: "ajjsjsjsjsjs.sjsjsjsjsjs.sjsjsjsjs",
			},
			expectedError: errors.New("the user does not exist"),
		},
		"error in retrieving blocked status": {
			input: models.UserLogin{
				Email:    "arthurbishop120@gmail.com",
				Password: "4321",
			},
			StubDetails: func(userRepo mockrepo.MockUserRepository, orderRepo mockrepo.MockOrderRepository, Data models.UserLogin, helper mockhelper.MockHelper) {
				gomock.InOrder(
					userRepo.EXPECT().CheckUserAvailability(Data.Email).Times(1).Return(true),
					userRepo.EXPECT().UserBlockStatus(Data.Email).Times(1).Return(false, errors.New("internal error")),
				)
			},
			expectedOutput: models.TokenUsers{
				Users: models.UserDetailsResponse{},
				Token: "ajjsjsjsjsjs.sjsjsjsjsjs.sjsjsjsjs",
			},
			expectedError: errors.New("internal error"),
		},
		"blocked user is trying to login": {
			input: models.UserLogin{
				Email:    "arthurbishop120@gmail.com",
				Password: "4321",
			},
			StubDetails: func(userRepo mockrepo.MockUserRepository, orderRepo mockrepo.MockOrderRepository, Data models.UserLogin, helper mockhelper.MockHelper) {
				gomock.InOrder(
					userRepo.EXPECT().CheckUserAvailability(Data.Email).Times(1).Return(true),
					userRepo.EXPECT().UserBlockStatus(Data.Email).Times(1).Return(true, nil),
				)
			},
			expectedOutput: models.TokenUsers{
				Users: models.UserDetailsResponse{},
				Token: "ajjsjsjsjsjs.sjsjsjsjsjs.sjsjsjsjs",
			},
			expectedError: errors.New("user is blocked by admin"),
		},
		"Error in finding the user": {
			input: models.UserLogin{
				Email:    "arthurbishop120@gmail.com",
				Password: "4321",
			},
			StubDetails: func(userRepo mockrepo.MockUserRepository, orderRepo mockrepo.MockOrderRepository, Data models.UserLogin, helper mockhelper.MockHelper) {
				gomock.InOrder(
					userRepo.EXPECT().CheckUserAvailability(Data.Email).Times(1).Return(true),
					userRepo.EXPECT().UserBlockStatus(Data.Email).Times(1).Return(false, nil),
					userRepo.EXPECT().FindUserByEmail(Data).Times(1).Return(models.UserSignInResponse{}, errors.New("internal error")),
				)
			},
			expectedOutput: models.TokenUsers{
				Users: models.UserDetailsResponse{},
				Token: "ajjsjsjsjsjs.sjsjsjsjsjs.sjsjsjsjs",
			},
			expectedError: errors.New("internal error"),
		},
		"incorrect password": {
			input: models.UserLogin{
				Email:    "arthurbishop120@gmail.com",
				Password: "4321",
			},
			StubDetails: func(userRepo mockrepo.MockUserRepository, orderRepo mockrepo.MockOrderRepository, Data models.UserLogin, helper mockhelper.MockHelper) {
				gomock.InOrder(
					userRepo.EXPECT().CheckUserAvailability(Data.Email).Times(1).Return(true),
					userRepo.EXPECT().UserBlockStatus(Data.Email).Times(1).Return(false, nil),
					userRepo.EXPECT().FindUserByEmail(Data).Times(1).Return(models.UserSignInResponse{
						Id:       1,
						UserID:   1,
						Name:     "Arun K",
						Email:    "arthurbishop120@gmail.com",
						Phone:    "6282246077",
						Password: "4321",
					}, nil),
					helper.EXPECT().CompareHashAndPassword("4321", "4321").Times(1).Return(errors.New("error")),
				)
			},
			expectedOutput: models.TokenUsers{
				Users: models.UserDetailsResponse{},
				Token: "ajjsjsjsjsjs.sjsjsjsjsjs.sjsjsjsjs",
			},
			expectedError: errors.New("password incorrect"),
		},
		"error in generating token": {
			input: models.UserLogin{
				Email:    "arthurbishop120@gmail.com",
				Password: "4321",
			},
			StubDetails: func(userRepo mockrepo.MockUserRepository, orderRepo mockrepo.MockOrderRepository, Data models.UserLogin, helper mockhelper.MockHelper) {
				gomock.InOrder(
					userRepo.EXPECT().CheckUserAvailability(Data.Email).Times(1).Return(true),
					userRepo.EXPECT().UserBlockStatus(Data.Email).Times(1).Return(false, nil),
					userRepo.EXPECT().FindUserByEmail(Data).Times(1).Return(models.UserSignInResponse{
						Id:       1,
						UserID:   1,
						Name:     "Arun K",
						Email:    "arthurbishop120@gmail.com",
						Phone:    "6282246077",
						Password: "4321",
					}, nil),
					helper.EXPECT().CompareHashAndPassword("4321", "4321").Times(1).Return(nil),
					helper.EXPECT().GenerateTokenClients(models.UserDetailsResponse{
						Id:    1,
						Name:  "Arun K",
						Email: "arthurbishop120@gmail.com",
						Phone: "6282246077",
					}).Times(1).Return(gomock.Any().String(), errors.New("error")),
				)
			},
			expectedOutput: models.TokenUsers{
				Users: models.UserDetailsResponse{},
				Token: "ajjsjsjsjsjs.sjsjsjsjsjs.sjsjsjsjs",
			},
			expectedError: errors.New("could not create token"),
		},
	}

	for _, test := range testData {

		test.StubDetails(*userRepo, *orderRepo, test.input, *helper)

		tokenusers, err := userUseCase.LoginHandler(test.input)

		assert.Equal(t, test.expectedOutput.Users.Id, tokenusers.Users.Id)
		assert.Equal(t, test.expectedOutput.Users.Name, tokenusers.Users.Name)
		assert.Equal(t, test.expectedOutput.Users.Email, tokenusers.Users.Email)
		assert.Equal(t, test.expectedError, err)

	}
}

func Test_AddAddress(t *testing.T) {
	ctrl := gomock.NewController(t)

	// Create mock implementations for the repositories
	userRepo := mockrepo.NewMockUserRepository(ctrl)
	orderRepo := mockrepo.NewMockOrderRepository(ctrl)
	otpRepo := mockrepo.NewMockOtpRepository(ctrl)
	inventoryRepo := mockrepo.NewMockInventoryRepository(ctrl)
	helper := mockhelper.NewMockHelper(ctrl)
	cfg := config.Config{}

	userUseCase := NewUserUseCase(userRepo, cfg, otpRepo, inventoryRepo, orderRepo, helper)

	testData := map[string]struct {
		input          models.AddAddress
		StubDetails    func(mockrepo.MockUserRepository, mockrepo.MockOrderRepository, models.AddAddress, mockhelper.MockHelper)
		expectedOutput models.TokenUsers
		expectedError  error
	}{
		"success and as default": {
			input: models.AddAddress{
				Name:      "Arun K",
				HouseName: "nellikkal arun bhavan",
				Street:    "pallippuram",
				City:      "cherthala",
				State:     "kerala",
				Pin:       "688541",
			},
			StubDetails: func(userRepo mockrepo.MockUserRepository, orderRepo mockrepo.MockOrderRepository, Data models.AddAddress, helper mockhelper.MockHelper) {
				gomock.InOrder(
					userRepo.EXPECT().CheckIfFirstAddress(1).Times(1).Return(false),
					userRepo.EXPECT().AddAddress(1, Data, true).Times(1).Return(nil),
				)
			},
			expectedError: nil,
		},
		"success but not as default": {
			input: models.AddAddress{
				Name:      "Arun K",
				HouseName: "nellikkal arun bhavan",
				Street:    "pallippuram",
				City:      "cherthala",
				State:     "kerala",
				Pin:       "688541",
			},
			StubDetails: func(userRepo mockrepo.MockUserRepository, orderRepo mockrepo.MockOrderRepository, Data models.AddAddress, helper mockhelper.MockHelper) {
				gomock.InOrder(
					userRepo.EXPECT().CheckIfFirstAddress(1).Times(1).Return(true),
					userRepo.EXPECT().AddAddress(1, Data, false).Times(1).Return(nil),
				)
			},
			expectedError: nil,
		},
		"couldnt add the address": {
			input: models.AddAddress{
				Name:      "Arun K",
				HouseName: "nellikkal arun bhavan",
				Street:    "pallippuram",
				City:      "cherthala",
				State:     "kerala",
				Pin:       "688541",
			},
			StubDetails: func(userRepo mockrepo.MockUserRepository, orderRepo mockrepo.MockOrderRepository, Data models.AddAddress, helper mockhelper.MockHelper) {
				gomock.InOrder(
					userRepo.EXPECT().CheckIfFirstAddress(1).Times(1).Return(true),
					userRepo.EXPECT().AddAddress(1, Data, false).Times(1).Return(errors.New("could not add the address")),
				)
			},
			expectedError: errors.New("error in adding address"),
		},
	}
	for _, test := range testData {

		test.StubDetails(*userRepo, *orderRepo, test.input, *helper)

		err := userUseCase.AddAddress(1, test.input)
		assert.Equal(t, test.expectedError, err)

	}
}

func Test_GetAddresses(t *testing.T) {
	ctrl := gomock.NewController(t)

	// Create mock implementations for the repositories
	userRepo := mockrepo.NewMockUserRepository(ctrl)
	orderRepo := mockrepo.NewMockOrderRepository(ctrl)
	otpRepo := mockrepo.NewMockOtpRepository(ctrl)
	inventoryRepo := mockrepo.NewMockInventoryRepository(ctrl)
	helper := mockhelper.NewMockHelper(ctrl)
	cfg := config.Config{}

	userUseCase := NewUserUseCase(userRepo, cfg, otpRepo, inventoryRepo, orderRepo, helper)

	testData := map[string]struct {
		input          int
		StubDetails    func(mockrepo.MockUserRepository, mockrepo.MockOrderRepository, int, mockhelper.MockHelper)
		expectedOutput []domain.Address
		expectedError  error
	}{
		"success": {
			input: 1,
			StubDetails: func(userRepo mockrepo.MockUserRepository, orderRepo mockrepo.MockOrderRepository, id int, helper mockhelper.MockHelper) {
				gomock.InOrder(
					userRepo.EXPECT().GetAddresses(id).Times(1).Return([]domain.Address{}, nil),
				)
			},
			expectedOutput: []domain.Address{},
			expectedError:  nil,
		},
		"errorcase": {
			input: 1,
			StubDetails: func(userRepo mockrepo.MockUserRepository, orderRepo mockrepo.MockOrderRepository, id int, helper mockhelper.MockHelper) {
				gomock.InOrder(
					userRepo.EXPECT().GetAddresses(id).Times(1).Return([]domain.Address{}, errors.New("error")),
				)
			},
			expectedOutput: []domain.Address{},
			expectedError:  errors.New("error in getting addresses"),
		},
	}
	for _, test := range testData {

		test.StubDetails(*userRepo, *orderRepo, test.input, *helper)

		products, err := userUseCase.GetAddresses(test.input)
		assert.Equal(t, test.expectedOutput, products)
		assert.Equal(t, test.expectedError, err)

	}
}

func Test_GetUserDetails(t *testing.T) {
	ctrl := gomock.NewController(t)

	// Create mock implementations for the repositories
	userRepo := mockrepo.NewMockUserRepository(ctrl)
	orderRepo := mockrepo.NewMockOrderRepository(ctrl)
	otpRepo := mockrepo.NewMockOtpRepository(ctrl)
	inventoryRepo := mockrepo.NewMockInventoryRepository(ctrl)
	helper := mockhelper.NewMockHelper(ctrl)
	cfg := config.Config{}

	userUseCase := NewUserUseCase(userRepo, cfg, otpRepo, inventoryRepo, orderRepo, helper)

	testData := map[string]struct {
		input          int
		StubDetails    func(mockrepo.MockUserRepository, mockrepo.MockOrderRepository, int, mockhelper.MockHelper)
		expectedOutput models.UserDetailsResponse
		expectedError  error
	}{
		"success": {
			input: 1,
			StubDetails: func(userRepo mockrepo.MockUserRepository, orderRepo mockrepo.MockOrderRepository, id int, helper mockhelper.MockHelper) {
				gomock.InOrder(
					userRepo.EXPECT().GetUserDetails(id).Times(1).Return(models.UserDetailsResponse{
						Id:    1,
						Name:  "Arun K",
						Email: "arthurbishop120@gmail.com",
						Phone: "6282246077",
					}, nil),
				)
			},
			expectedOutput: models.UserDetailsResponse{
				Id:    1,
				Name:  "Arun K",
				Email: "arthurbishop120@gmail.com",
				Phone: "6282246077",
			},
			expectedError: nil,
		},
		"errorcase": {
			input: 1,
			StubDetails: func(userRepo mockrepo.MockUserRepository, orderRepo mockrepo.MockOrderRepository, id int, helper mockhelper.MockHelper) {
				gomock.InOrder(
					userRepo.EXPECT().GetUserDetails(id).Times(1).Return(models.UserDetailsResponse{}, errors.New("error")),
				)
			},
			expectedOutput: models.UserDetailsResponse{},
			expectedError:  errors.New("error in getting details"),
		},
	}
	for _, test := range testData {

		test.StubDetails(*userRepo, *orderRepo, test.input, *helper)

		products, err := userUseCase.GetUserDetails(test.input)
		assert.Equal(t, test.expectedOutput, products)
		assert.Equal(t, test.expectedError, err)

	}
}

func Test_ChangePassword(t *testing.T) {
	ctrl := gomock.NewController(t)

	// Create mock implementations for the repositories
	userRepo := mockrepo.NewMockUserRepository(ctrl)
	orderRepo := mockrepo.NewMockOrderRepository(ctrl)
	otpRepo := mockrepo.NewMockOtpRepository(ctrl)
	inventoryRepo := mockrepo.NewMockInventoryRepository(ctrl)
	helper := mockhelper.NewMockHelper(ctrl)
	cfg := config.Config{}

	userUseCase := NewUserUseCase(userRepo, cfg, otpRepo, inventoryRepo, orderRepo, helper)

	testData := map[string]struct {
		input struct {
			ID         int
			Old        string
			Password   string
			RePassword string
		}
		StubDetails func(mockrepo.MockUserRepository, mockrepo.MockOrderRepository, struct {
			ID         int
			Old        string
			Password   string
			RePassword string
		}, mockhelper.MockHelper)
		expectedOutput models.UserDetailsResponse
		expectedError  error
	}{
		"success": {
			input: struct {
				ID         int
				Old        string
				Password   string
				RePassword string
			}{
				ID:         1,
				Old:        "1234",
				Password:   "4321",
				RePassword: "4321",
			},
			StubDetails: func(userRepo mockrepo.MockUserRepository, orderRepo mockrepo.MockOrderRepository, data struct {
				ID         int
				Old        string
				Password   string
				RePassword string
			}, helper mockhelper.MockHelper) {
				gomock.InOrder(
					userRepo.EXPECT().GetPassword(data.ID).Times(1).Return("1234", nil),
					helper.EXPECT().CompareHashAndPassword(data.Old, "1234").Times(1).Return(nil),
					helper.EXPECT().PasswordHashing(data.Password).Times(1).Return(gomock.Any().String(), nil),
					userRepo.EXPECT().ChangePassword(data.ID, gomock.Any().String()).Times(1).Return(nil),
				)
			},
			expectedOutput: models.UserDetailsResponse{},
			expectedError:  nil,
		},
		"couldnt fetch password": {
			input: struct {
				ID         int
				Old        string
				Password   string
				RePassword string
			}{
				ID:         1,
				Old:        "1234",
				Password:   "4321",
				RePassword: "4321",
			},
			StubDetails: func(userRepo mockrepo.MockUserRepository, orderRepo mockrepo.MockOrderRepository, data struct {
				ID         int
				Old        string
				Password   string
				RePassword string
			}, helper mockhelper.MockHelper) {
				gomock.InOrder(
					userRepo.EXPECT().GetPassword(data.ID).Times(1).Return("", errors.New("could not fetch")),
				)
			},
			expectedOutput: models.UserDetailsResponse{},
			expectedError:  errors.New("internal error"),
		},
		"hashing problem": {
			input: struct {
				ID         int
				Old        string
				Password   string
				RePassword string
			}{
				ID:         1,
				Old:        "1234",
				Password:   "4321",
				RePassword: "4321",
			},
			StubDetails: func(userRepo mockrepo.MockUserRepository, orderRepo mockrepo.MockOrderRepository, data struct {
				ID         int
				Old        string
				Password   string
				RePassword string
			}, helper mockhelper.MockHelper) {
				gomock.InOrder(
					userRepo.EXPECT().GetPassword(data.ID).Times(1).Return("1234", nil),
					helper.EXPECT().CompareHashAndPassword(data.Old, "1234").Times(1).Return(errors.New("error")),
				)
			},
			expectedOutput: models.UserDetailsResponse{},
			expectedError:  errors.New("password incorrect"),
		},
		"password hashing problem": {
			input: struct {
				ID         int
				Old        string
				Password   string
				RePassword string
			}{
				ID:         1,
				Old:        "1234",
				Password:   "4321",
				RePassword: "4321",
			},
			StubDetails: func(userRepo mockrepo.MockUserRepository, orderRepo mockrepo.MockOrderRepository, data struct {
				ID         int
				Old        string
				Password   string
				RePassword string
			}, helper mockhelper.MockHelper) {
				gomock.InOrder(
					userRepo.EXPECT().GetPassword(data.ID).Times(1).Return("1234", nil),
					helper.EXPECT().CompareHashAndPassword(data.Old, "1234").Times(1).Return(nil),
					helper.EXPECT().PasswordHashing(data.Password).Times(1).Return(gomock.Any().String(), errors.New("error")),
				)
			},
			expectedOutput: models.UserDetailsResponse{},
			expectedError:  errors.New("error in hashing password"),
		},
		"password and repassword dont match": {
			input: struct {
				ID         int
				Old        string
				Password   string
				RePassword string
			}{
				ID:         1,
				Old:        "1234",
				Password:   "4321",
				RePassword: "5432",
			},
			StubDetails: func(userRepo mockrepo.MockUserRepository, orderRepo mockrepo.MockOrderRepository, data struct {
				ID         int
				Old        string
				Password   string
				RePassword string
			}, helper mockhelper.MockHelper) {
				gomock.InOrder(
					userRepo.EXPECT().GetPassword(data.ID).Times(1).Return("1234", nil),
					helper.EXPECT().CompareHashAndPassword(data.Old, "1234").Times(1).Return(nil),
				)
			},
			expectedOutput: models.UserDetailsResponse{},
			expectedError:  errors.New("passwords does not match"),
		},
	}
	for _, test := range testData {

		test.StubDetails(*userRepo, *orderRepo, test.input, *helper)

		err := userUseCase.ChangePassword(test.input.ID, test.input.Old, test.input.Password, test.input.RePassword)
		assert.Equal(t, test.expectedError, err)

	}
}

func Test_ForgotPasswordSend(t *testing.T) {
	ctrl := gomock.NewController(t)

	// Create mock implementations for the repositories
	userRepo := mockrepo.NewMockUserRepository(ctrl)
	orderRepo := mockrepo.NewMockOrderRepository(ctrl)
	otpRepo := mockrepo.NewMockOtpRepository(ctrl)
	inventoryRepo := mockrepo.NewMockInventoryRepository(ctrl)
	helper := mockhelper.NewMockHelper(ctrl)
	cfg := config.Config{}

	userUseCase := NewUserUseCase(userRepo, cfg, otpRepo, inventoryRepo, orderRepo, helper)

	testData := map[string]struct {
		input          string
		StubDetails    func(mockrepo.MockUserRepository, mockrepo.MockOrderRepository, mockhelper.MockHelper, string)
		expectedOutput models.UserDetailsResponse
		expectedError  error
	}{
		"success": {
			input: "6282246077",
			StubDetails: func(userRepo mockrepo.MockUserRepository, orderRepo mockrepo.MockOrderRepository, helper mockhelper.MockHelper, data string) {
				gomock.InOrder(
					otpRepo.EXPECT().FindUserByMobileNumber(data).Times(1).Return(true),
					helper.EXPECT().TwilioSetup(gomock.Any(), gomock.Any()).Times(1),
					helper.EXPECT().TwilioSendOTP("6282246077", cfg.SERVICESID).Times(1).Return(gomock.Any().String(), nil),
				)
			},
			expectedOutput: models.UserDetailsResponse{},
			expectedError:  nil,
		},
		"user not found": {
			input: "6282246077",
			StubDetails: func(userRepo mockrepo.MockUserRepository, orderRepo mockrepo.MockOrderRepository, helper mockhelper.MockHelper, data string) {
				gomock.InOrder(
					otpRepo.EXPECT().FindUserByMobileNumber(data).Times(1).Return(false),
				)
			},
			expectedOutput: models.UserDetailsResponse{},
			expectedError:  errors.New("the user does not exist"),
		},
		"error in generating otp": {
			input: "6282246077",
			StubDetails: func(userRepo mockrepo.MockUserRepository, orderRepo mockrepo.MockOrderRepository, helper mockhelper.MockHelper, data string) {
				gomock.InOrder(
					otpRepo.EXPECT().FindUserByMobileNumber(data).Times(1).Return(true),
					helper.EXPECT().TwilioSetup(gomock.Any(), gomock.Any()).Times(1),
					helper.EXPECT().TwilioSendOTP("6282246077", cfg.SERVICESID).Times(1).Return(gomock.Any().String(), errors.New("error")),
				)
			},
			expectedOutput: models.UserDetailsResponse{},
			expectedError:  errors.New("error ocurred while generating OTP"),
		},
	}
	for _, test := range testData {

		test.StubDetails(*userRepo, *orderRepo, *helper, test.input)

		err := userUseCase.ForgotPasswordSend(test.input)
		assert.Equal(t, test.expectedError, err)

	}
}

func Test_ForgotPasswordVerifyAndChange(t *testing.T) {
	ctrl := gomock.NewController(t)

	// Create mock implementations for the repositories
	userRepo := mockrepo.NewMockUserRepository(ctrl)
	orderRepo := mockrepo.NewMockOrderRepository(ctrl)
	otpRepo := mockrepo.NewMockOtpRepository(ctrl)
	inventoryRepo := mockrepo.NewMockInventoryRepository(ctrl)
	helper := mockhelper.NewMockHelper(ctrl)
	cfg := config.Config{}

	userUseCase := NewUserUseCase(userRepo, cfg, otpRepo, inventoryRepo, orderRepo, helper)

	testData := map[string]struct {
		input          models.ForgotVerify
		StubDetails    func(mockrepo.MockUserRepository, mockrepo.MockOrderRepository, mockhelper.MockHelper, models.ForgotVerify)
		expectedOutput models.UserDetailsResponse
		expectedError  error
	}{
		"success": {
			input: models.ForgotVerify{
				Phone:       "6282246077",
				Otp:         "1234",
				NewPassword: "4321",
			},
			StubDetails: func(userRepo mockrepo.MockUserRepository, orderRepo mockrepo.MockOrderRepository, helper mockhelper.MockHelper, data models.ForgotVerify) {
				gomock.InOrder(
					helper.EXPECT().TwilioSetup(gomock.Any(), gomock.Any()).Times(1),
					helper.EXPECT().TwilioVerifyOTP(cfg.SERVICESID, data.Otp, data.Phone).Times(1).Return(nil),
					userRepo.EXPECT().FindIdFromPhone("6282246077").Times(1).Return(1, nil),
					helper.EXPECT().PasswordHashing(data.NewPassword).Times(1).Return(gomock.Any().String(), nil),
					userRepo.EXPECT().ChangePassword(1, gomock.Any().String()).Times(1).Return(nil),
				)
			},
			expectedOutput: models.UserDetailsResponse{},
			expectedError:  nil,
		},
		"error in verifying otp": {
			input: models.ForgotVerify{
				Phone:       "6282246077",
				Otp:         "1234",
				NewPassword: "4321",
			},
			StubDetails: func(userRepo mockrepo.MockUserRepository, orderRepo mockrepo.MockOrderRepository, helper mockhelper.MockHelper, data models.ForgotVerify) {
				gomock.InOrder(
					helper.EXPECT().TwilioSetup(gomock.Any(), gomock.Any()).Times(1),
					helper.EXPECT().TwilioVerifyOTP(cfg.SERVICESID, data.Otp, data.Phone).Times(1).Return(errors.New("error")),
				)
			},
			expectedOutput: models.UserDetailsResponse{},
			expectedError:  errors.New("error while verifying"),
		},
		"cannot find user": {
			input: models.ForgotVerify{
				Phone:       "6282246077",
				Otp:         "1234",
				NewPassword: "4321",
			},
			StubDetails: func(userRepo mockrepo.MockUserRepository, orderRepo mockrepo.MockOrderRepository, helper mockhelper.MockHelper, data models.ForgotVerify) {
				gomock.InOrder(
					helper.EXPECT().TwilioSetup(gomock.Any(), gomock.Any()).Times(1),
					helper.EXPECT().TwilioVerifyOTP(cfg.SERVICESID, data.Otp, data.Phone).Times(1).Return(nil),
					userRepo.EXPECT().FindIdFromPhone("6282246077").Times(1).Return(0, errors.New("error")),
				)
			},
			expectedOutput: models.UserDetailsResponse{},
			expectedError:  errors.New("cannot find user from mobile number"),
		},
		"error in hashing": {
			input: models.ForgotVerify{
				Phone:       "6282246077",
				Otp:         "1234",
				NewPassword: "4321",
			},
			StubDetails: func(userRepo mockrepo.MockUserRepository, orderRepo mockrepo.MockOrderRepository, helper mockhelper.MockHelper, data models.ForgotVerify) {
				gomock.InOrder(
					helper.EXPECT().TwilioSetup(gomock.Any(), gomock.Any()).Times(1),
					helper.EXPECT().TwilioVerifyOTP(cfg.SERVICESID, data.Otp, data.Phone).Times(1).Return(nil),
					userRepo.EXPECT().FindIdFromPhone("6282246077").Times(1).Return(1, nil),
					helper.EXPECT().PasswordHashing(data.NewPassword).Times(1).Return(gomock.Any().String(), errors.New("error")),
				)
			},
			expectedOutput: models.UserDetailsResponse{},
			expectedError:  errors.New("error in hashing password"),
		},
		"could not change password": {
			input: models.ForgotVerify{
				Phone:       "6282246077",
				Otp:         "1234",
				NewPassword: "4321",
			},
			StubDetails: func(userRepo mockrepo.MockUserRepository, orderRepo mockrepo.MockOrderRepository, helper mockhelper.MockHelper, data models.ForgotVerify) {
				gomock.InOrder(
					helper.EXPECT().TwilioSetup(gomock.Any(), gomock.Any()).Times(1),
					helper.EXPECT().TwilioVerifyOTP(cfg.SERVICESID, data.Otp, data.Phone).Times(1).Return(nil),
					userRepo.EXPECT().FindIdFromPhone("6282246077").Times(1).Return(1, nil),
					helper.EXPECT().PasswordHashing(data.NewPassword).Times(1).Return(gomock.Any().String(), nil),
					userRepo.EXPECT().ChangePassword(1, gomock.Any().String()).Times(1).Return(errors.New("error")),
				)
			},
			expectedOutput: models.UserDetailsResponse{},
			expectedError:  errors.New("could not change password"),
		},
	}
	for _, test := range testData {

		test.StubDetails(*userRepo, *orderRepo, *helper, test.input)

		err := userUseCase.ForgotPasswordVerifyAndChange(test.input)
		assert.Equal(t, test.expectedError, err)

	}
}

func Test_EditName(t *testing.T) {
	ctrl := gomock.NewController(t)

	// Create mock implementations for the repositories
	userRepo := mockrepo.NewMockUserRepository(ctrl)
	orderRepo := mockrepo.NewMockOrderRepository(ctrl)
	otpRepo := mockrepo.NewMockOtpRepository(ctrl)
	inventoryRepo := mockrepo.NewMockInventoryRepository(ctrl)
	helper := mockhelper.NewMockHelper(ctrl)
	cfg := config.Config{}

	userUseCase := NewUserUseCase(userRepo, cfg, otpRepo, inventoryRepo, orderRepo, helper)

	testData := map[string]struct {
		input struct {
			id   int
			name string
		}
		StubDetails func(mockrepo.MockUserRepository, mockrepo.MockOrderRepository, mockhelper.MockHelper, struct {
			id   int
			name string
		})
		expectedOutput models.UserDetailsResponse
		expectedError  error
	}{
		"success": {
			input: struct {
				id   int
				name string
			}{
				id:   1,
				name: "Arun K",
			},
			StubDetails: func(userRepo mockrepo.MockUserRepository, orderRepo mockrepo.MockOrderRepository, helper mockhelper.MockHelper, data struct {
				id   int
				name string
			}) {
				gomock.InOrder(
					userRepo.EXPECT().EditName(data.id, data.name).Times(1).Return(nil),
				)
			},
			expectedOutput: models.UserDetailsResponse{},
			expectedError:  nil,
		},
		"error": {
			input: struct {
				id   int
				name string
			}{
				id:   1,
				name: "Arun K",
			},
			StubDetails: func(userRepo mockrepo.MockUserRepository, orderRepo mockrepo.MockOrderRepository, helper mockhelper.MockHelper, data struct {
				id   int
				name string
			}) {
				gomock.InOrder(
					userRepo.EXPECT().EditName(data.id, data.name).Times(1).Return(errors.New("error")),
				)
			},
			expectedOutput: models.UserDetailsResponse{},
			expectedError:  errors.New("could not change"),
		},
	}
	for _, test := range testData {

		test.StubDetails(*userRepo, *orderRepo, *helper, test.input)

		err := userUseCase.EditName(test.input.id, test.input.name)
		assert.Equal(t, test.expectedError, err)

	}
}

func Test_EditEmail(t *testing.T) {
	ctrl := gomock.NewController(t)

	// Create mock implementations for the repositories
	userRepo := mockrepo.NewMockUserRepository(ctrl)
	orderRepo := mockrepo.NewMockOrderRepository(ctrl)
	otpRepo := mockrepo.NewMockOtpRepository(ctrl)
	inventoryRepo := mockrepo.NewMockInventoryRepository(ctrl)
	helper := mockhelper.NewMockHelper(ctrl)
	cfg := config.Config{}

	userUseCase := NewUserUseCase(userRepo, cfg, otpRepo, inventoryRepo, orderRepo, helper)

	testData := map[string]struct {
		input struct {
			id    int
			email string
		}
		StubDetails func(mockrepo.MockUserRepository, mockrepo.MockOrderRepository, mockhelper.MockHelper, struct {
			id    int
			email string
		})
		expectedOutput models.UserDetailsResponse
		expectedError  error
	}{
		"success": {
			input: struct {
				id    int
				email string
			}{
				id:    1,
				email: "arthurbishop120@gmail.com",
			},
			StubDetails: func(userRepo mockrepo.MockUserRepository, orderRepo mockrepo.MockOrderRepository, helper mockhelper.MockHelper, data struct {
				id    int
				email string
			}) {
				gomock.InOrder(
					userRepo.EXPECT().EditEmail(data.id, data.email).Times(1).Return(nil),
				)
			},
			expectedOutput: models.UserDetailsResponse{},
			expectedError:  nil,
		},
		"error": {
			input: struct {
				id    int
				email string
			}{
				id:    1,
				email: "arthurbishop120@gmail.com",
			},
			StubDetails: func(userRepo mockrepo.MockUserRepository, orderRepo mockrepo.MockOrderRepository, helper mockhelper.MockHelper, data struct {
				id    int
				email string
			}) {
				gomock.InOrder(
					userRepo.EXPECT().EditEmail(data.id, data.email).Times(1).Return(errors.New("error")),
				)
			},
			expectedOutput: models.UserDetailsResponse{},
			expectedError:  errors.New("could not change"),
		},
	}
	for _, test := range testData {

		test.StubDetails(*userRepo, *orderRepo, *helper, test.input)

		err := userUseCase.EditEmail(test.input.id, test.input.email)
		assert.Equal(t, test.expectedError, err)

	}
}

func Test_EditPhone(t *testing.T) {
	ctrl := gomock.NewController(t)

	// Create mock implementations for the repositories
	userRepo := mockrepo.NewMockUserRepository(ctrl)
	orderRepo := mockrepo.NewMockOrderRepository(ctrl)
	otpRepo := mockrepo.NewMockOtpRepository(ctrl)
	inventoryRepo := mockrepo.NewMockInventoryRepository(ctrl)
	helper := mockhelper.NewMockHelper(ctrl)
	cfg := config.Config{}

	userUseCase := NewUserUseCase(userRepo, cfg, otpRepo, inventoryRepo, orderRepo, helper)

	testData := map[string]struct {
		input struct {
			id    int
			phone string
		}
		StubDetails func(mockrepo.MockUserRepository, mockrepo.MockOrderRepository, mockhelper.MockHelper, struct {
			id    int
			phone string
		})
		expectedOutput models.UserDetailsResponse
		expectedError  error
	}{
		"success": {
			input: struct {
				id    int
				phone string
			}{
				id:    1,
				phone: "6282246077",
			},
			StubDetails: func(userRepo mockrepo.MockUserRepository, orderRepo mockrepo.MockOrderRepository, helper mockhelper.MockHelper, data struct {
				id    int
				phone string
			}) {
				gomock.InOrder(
					userRepo.EXPECT().EditPhone(data.id, data.phone).Times(1).Return(nil),
				)
			},
			expectedOutput: models.UserDetailsResponse{},
			expectedError:  nil,
		},
		"error": {
			input: struct {
				id    int
				phone string
			}{
				id:    1,
				phone: "6282246077",
			},
			StubDetails: func(userRepo mockrepo.MockUserRepository, orderRepo mockrepo.MockOrderRepository, helper mockhelper.MockHelper, data struct {
				id    int
				phone string
			}) {
				gomock.InOrder(
					userRepo.EXPECT().EditPhone(data.id, data.phone).Times(1).Return(errors.New("error")),
				)
			},
			expectedOutput: models.UserDetailsResponse{},
			expectedError:  errors.New("could not change"),
		},
	}
	for _, test := range testData {

		test.StubDetails(*userRepo, *orderRepo, *helper, test.input)

		err := userUseCase.EditPhone(test.input.id, test.input.phone)
		assert.Equal(t, test.expectedError, err)

	}
}

func Test_UpdateQuantityAdd(t *testing.T) {
	ctrl := gomock.NewController(t)

	// Create mock implementations for the repositories
	userRepo := mockrepo.NewMockUserRepository(ctrl)
	orderRepo := mockrepo.NewMockOrderRepository(ctrl)
	otpRepo := mockrepo.NewMockOtpRepository(ctrl)
	inventoryRepo := mockrepo.NewMockInventoryRepository(ctrl)
	helper := mockhelper.NewMockHelper(ctrl)
	cfg := config.Config{}

	userUseCase := NewUserUseCase(userRepo, cfg, otpRepo, inventoryRepo, orderRepo, helper)

	testData := map[string]struct {
		input struct {
			id     int
			inv_id int
		}
		StubDetails func(mockrepo.MockUserRepository, mockrepo.MockOrderRepository, mockhelper.MockHelper, struct {
			id     int
			inv_id int
		})
		expectedOutput models.UserDetailsResponse
		expectedError  error
	}{
		"success": {
			input: struct {
				id     int
				inv_id int
			}{
				id:     1,
				inv_id: 1,
			},
			StubDetails: func(userRepo mockrepo.MockUserRepository, orderRepo mockrepo.MockOrderRepository, helper mockhelper.MockHelper, data struct {
				id     int
				inv_id int
			}) {
				gomock.InOrder(
					userRepo.EXPECT().UpdateQuantityAdd(data.id, data.inv_id).Times(1).Return(nil),
				)
			},
			expectedOutput: models.UserDetailsResponse{},
			expectedError:  nil,
		},
		"error": {
			input: struct {
				id     int
				inv_id int
			}{
				id:     1,
				inv_id: 1,
			},
			StubDetails: func(userRepo mockrepo.MockUserRepository, orderRepo mockrepo.MockOrderRepository, helper mockhelper.MockHelper, data struct {
				id     int
				inv_id int
			}) {
				gomock.InOrder(
					userRepo.EXPECT().UpdateQuantityAdd(data.id, data.inv_id).Times(1).Return(errors.New("error")),
				)
			},
			expectedOutput: models.UserDetailsResponse{},
			expectedError:  errors.New("error"),
		},
	}
	for _, test := range testData {

		test.StubDetails(*userRepo, *orderRepo, *helper, test.input)

		err := userUseCase.UpdateQuantityAdd(test.input.id, test.input.inv_id)
		assert.Equal(t, test.expectedError, err)

	}
}

func Test_UpdateQuantityLess(t *testing.T) {
	ctrl := gomock.NewController(t)

	// Create mock implementations for the repositories
	userRepo := mockrepo.NewMockUserRepository(ctrl)
	orderRepo := mockrepo.NewMockOrderRepository(ctrl)
	otpRepo := mockrepo.NewMockOtpRepository(ctrl)
	inventoryRepo := mockrepo.NewMockInventoryRepository(ctrl)
	helper := mockhelper.NewMockHelper(ctrl)
	cfg := config.Config{}

	userUseCase := NewUserUseCase(userRepo, cfg, otpRepo, inventoryRepo, orderRepo, helper)

	testData := map[string]struct {
		input struct {
			id     int
			inv_id int
		}
		StubDetails func(mockrepo.MockUserRepository, mockrepo.MockOrderRepository, mockhelper.MockHelper, struct {
			id     int
			inv_id int
		})
		expectedOutput models.UserDetailsResponse
		expectedError  error
	}{
		"success": {
			input: struct {
				id     int
				inv_id int
			}{
				id:     1,
				inv_id: 1,
			},
			StubDetails: func(userRepo mockrepo.MockUserRepository, orderRepo mockrepo.MockOrderRepository, helper mockhelper.MockHelper, data struct {
				id     int
				inv_id int
			}) {
				gomock.InOrder(
					userRepo.EXPECT().UpdateQuantityLess(data.id, data.inv_id).Times(1).Return(nil),
				)
			},
			expectedOutput: models.UserDetailsResponse{},
			expectedError:  nil,
		},
		"error": {
			input: struct {
				id     int
				inv_id int
			}{
				id:     1,
				inv_id: 1,
			},
			StubDetails: func(userRepo mockrepo.MockUserRepository, orderRepo mockrepo.MockOrderRepository, helper mockhelper.MockHelper, data struct {
				id     int
				inv_id int
			}) {
				gomock.InOrder(
					userRepo.EXPECT().UpdateQuantityLess(data.id, data.inv_id).Times(1).Return(errors.New("error")),
				)
			},
			expectedOutput: models.UserDetailsResponse{},
			expectedError:  errors.New("error"),
		},
	}
	for _, test := range testData {

		test.StubDetails(*userRepo, *orderRepo, *helper, test.input)

		err := userUseCase.UpdateQuantityLess(test.input.id, test.input.inv_id)
		assert.Equal(t, test.expectedError, err)

	}
}

func Test_RemoveFromCart(t *testing.T) {
	ctrl := gomock.NewController(t)

	// Create mock implementations for the repositories
	userRepo := mockrepo.NewMockUserRepository(ctrl)
	orderRepo := mockrepo.NewMockOrderRepository(ctrl)
	otpRepo := mockrepo.NewMockOtpRepository(ctrl)
	inventoryRepo := mockrepo.NewMockInventoryRepository(ctrl)
	helper := mockhelper.NewMockHelper(ctrl)
	cfg := config.Config{}

	userUseCase := NewUserUseCase(userRepo, cfg, otpRepo, inventoryRepo, orderRepo, helper)

	testData := map[string]struct {
		input          int
		StubDetails    func(mockrepo.MockUserRepository, mockrepo.MockOrderRepository, mockhelper.MockHelper, int)
		expectedOutput models.UserDetailsResponse
		expectedError  error
	}{
		"success": {
			input: 1,
			StubDetails: func(userRepo mockrepo.MockUserRepository, orderRepo mockrepo.MockOrderRepository, helper mockhelper.MockHelper, data int) {
				gomock.InOrder(
					userRepo.EXPECT().RemoveFromCart(data).Times(1).Return(nil),
				)
			},
			expectedOutput: models.UserDetailsResponse{},
			expectedError:  nil,
		},
		"error": {
			input: 1,
			StubDetails: func(userRepo mockrepo.MockUserRepository, orderRepo mockrepo.MockOrderRepository, helper mockhelper.MockHelper, data int) {
				gomock.InOrder(
					userRepo.EXPECT().RemoveFromCart(data).Times(1).Return(errors.New("error")),
				)
			},
			expectedOutput: models.UserDetailsResponse{},
			expectedError:  errors.New("error"),
		},
	}
	for _, test := range testData {

		test.StubDetails(*userRepo, *orderRepo, *helper, test.input)

		err := userUseCase.RemoveFromCart(test.input)
		assert.Equal(t, test.expectedError, err)

	}
}

func Test_GetMyReferenceLink(t *testing.T) {
	ctrl := gomock.NewController(t)

	// Create mock implementations for the repositories
	userRepo := mockrepo.NewMockUserRepository(ctrl)
	orderRepo := mockrepo.NewMockOrderRepository(ctrl)
	otpRepo := mockrepo.NewMockOtpRepository(ctrl)
	inventoryRepo := mockrepo.NewMockInventoryRepository(ctrl)
	helper := mockhelper.NewMockHelper(ctrl)
	cfg := config.Config{}

	userUseCase := NewUserUseCase(userRepo, cfg, otpRepo, inventoryRepo, orderRepo, helper)

	testData := map[string]struct {
		input          int
		StubDetails    func(mockrepo.MockUserRepository, mockrepo.MockOrderRepository, mockhelper.MockHelper, int)
		expectedOutput string
		expectedError  error
	}{
		"success": {
			input: 1,
			StubDetails: func(userRepo mockrepo.MockUserRepository, orderRepo mockrepo.MockOrderRepository, helper mockhelper.MockHelper, data int) {
				gomock.InOrder(
					userRepo.EXPECT().GetReferralCodeFromID(1).Times(1).Return("1234", nil),
				)
			},
			expectedOutput: "jerseyhub.com/users/signup?ref=1234",
			expectedError:  nil,
		},
		"error getting referal code": {
			input: 1,
			StubDetails: func(userRepo mockrepo.MockUserRepository, orderRepo mockrepo.MockOrderRepository, helper mockhelper.MockHelper, data int) {
				gomock.InOrder(
					userRepo.EXPECT().GetReferralCodeFromID(1).Times(1).Return("", errors.New("error")),
				)
			},
			expectedOutput: "",
			expectedError:  errors.New("error getting ref code"),
		},
	}
	for _, test := range testData {

		test.StubDetails(*userRepo, *orderRepo, *helper, test.input)

		url, err := userUseCase.GetMyReferenceLink(test.input)
		assert.Equal(t, test.expectedOutput, url)
		assert.Equal(t, test.expectedError, err)

	}
}

func Test_GetCart(t *testing.T) {
	ctrl := gomock.NewController(t)

	// Create mock implementations for the repositories
	userRepo := mockrepo.NewMockUserRepository(ctrl)
	orderRepo := mockrepo.NewMockOrderRepository(ctrl)
	otpRepo := mockrepo.NewMockOtpRepository(ctrl)
	inventoryRepo := mockrepo.NewMockInventoryRepository(ctrl)
	helper := mockhelper.NewMockHelper(ctrl)
	cfg := config.Config{}

	userUseCase := NewUserUseCase(userRepo, cfg, otpRepo, inventoryRepo, orderRepo, helper)

	testData := map[string]struct {
		input          int
		StubDetails    func(mockrepo.MockUserRepository, mockrepo.MockOrderRepository, mockhelper.MockHelper, int)
		expectedOutput []models.GetCart
		expectedError  error
	}{
		"success": {
			input: 1,
			StubDetails: func(userRepo mockrepo.MockUserRepository, orderRepo mockrepo.MockOrderRepository, helper mockhelper.MockHelper, data int) {
				gomock.InOrder(
					userRepo.EXPECT().GetCartID(data).Times(1).Return(1, nil),
					userRepo.EXPECT().GetProductsInCart(1).Times(1).Return([]int{1, 2, 3}, nil),
					userRepo.EXPECT().FindProductNames(1).Times(1).Return("a", nil),
					userRepo.EXPECT().FindProductNames(2).Times(1).Return("b", nil),
					userRepo.EXPECT().FindProductNames(3).Times(1).Return("c", nil),
					userRepo.EXPECT().FindCartQuantity(1, 1).Times(1).Return(5, nil),
					userRepo.EXPECT().FindCartQuantity(1, 2).Times(1).Return(6, nil),
					userRepo.EXPECT().FindCartQuantity(1, 3).Times(1).Return(7, nil),
					userRepo.EXPECT().FindPrice(1).Times(1).Return(500.00, nil),
					userRepo.EXPECT().FindPrice(2).Times(1).Return(520.00, nil),
					userRepo.EXPECT().FindPrice(3).Times(1).Return(600.00, nil),
					userRepo.EXPECT().FindCategory(1).Times(1).Return(1, nil),
					userRepo.EXPECT().FindCategory(2).Times(1).Return(2, nil),
					userRepo.EXPECT().FindCategory(3).Times(1).Return(3, nil),
					userRepo.EXPECT().FindofferPercentage(1).Times(1).Return(10, nil),
					userRepo.EXPECT().FindofferPercentage(2).Times(1).Return(10, nil),
					userRepo.EXPECT().FindofferPercentage(3).Times(1).Return(10, nil),
				)
			},
			expectedOutput: []models.GetCart([]models.GetCart{{ProductName: "a", Category_id: 1, Quantity: 5, Total: 500, DiscountedPrice: 450}, {ProductName: "b", Category_id: 2, Quantity: 6, Total: 520, DiscountedPrice: 468}, {ProductName: "c", Category_id: 3, Quantity: 7, Total: 600, DiscountedPrice: 540}}),
			expectedError:  nil,
		},
		"error from getting cart": {
			input: 1,
			StubDetails: func(userRepo mockrepo.MockUserRepository, orderRepo mockrepo.MockOrderRepository, helper mockhelper.MockHelper, data int) {
				gomock.InOrder(
					userRepo.EXPECT().GetCartID(data).Times(1).Return(1, errors.New("internal error")),
				)
			},
			expectedOutput: []models.GetCart([]models.GetCart{}),
			expectedError:  errors.New("internal error"),
		},
		"error getting products": {
			input: 1,
			StubDetails: func(userRepo mockrepo.MockUserRepository, orderRepo mockrepo.MockOrderRepository, helper mockhelper.MockHelper, data int) {
				gomock.InOrder(
					userRepo.EXPECT().GetCartID(data).Times(1).Return(1, nil),
					userRepo.EXPECT().GetProductsInCart(1).Times(1).Return([]int{}, errors.New("internal error")),
				)
			},
			expectedOutput: []models.GetCart([]models.GetCart{}),
			expectedError:  errors.New("internal error"),
		},
		"error in finding product names": {
			input: 1,
			StubDetails: func(userRepo mockrepo.MockUserRepository, orderRepo mockrepo.MockOrderRepository, helper mockhelper.MockHelper, data int) {
				gomock.InOrder(
					userRepo.EXPECT().GetCartID(data).Times(1).Return(1, nil),
					userRepo.EXPECT().GetProductsInCart(1).Times(1).Return([]int{1, 2, 3}, nil),
					userRepo.EXPECT().FindProductNames(1).Times(1).Return("", errors.New("internal error")),
				)
			},
			expectedOutput: []models.GetCart([]models.GetCart{}),
			expectedError:  errors.New("internal error"),
		},
		"cart quantity error": {
			input: 1,
			StubDetails: func(userRepo mockrepo.MockUserRepository, orderRepo mockrepo.MockOrderRepository, helper mockhelper.MockHelper, data int) {
				gomock.InOrder(
					userRepo.EXPECT().GetCartID(data).Times(1).Return(1, nil),
					userRepo.EXPECT().GetProductsInCart(1).Times(1).Return([]int{1, 2, 3}, nil),
					userRepo.EXPECT().FindProductNames(1).Times(1).Return("a", nil),
					userRepo.EXPECT().FindProductNames(2).Times(1).Return("b", nil),
					userRepo.EXPECT().FindProductNames(3).Times(1).Return("c", nil),
					userRepo.EXPECT().FindCartQuantity(1, 1).Times(1).Return(5, errors.New("internal error")),
				)
			},
			expectedOutput: []models.GetCart([]models.GetCart{}),
			expectedError:  errors.New("internal error"),
		},
		"price finding error": {
			input: 1,
			StubDetails: func(userRepo mockrepo.MockUserRepository, orderRepo mockrepo.MockOrderRepository, helper mockhelper.MockHelper, data int) {
				gomock.InOrder(
					userRepo.EXPECT().GetCartID(data).Times(1).Return(1, nil),
					userRepo.EXPECT().GetProductsInCart(1).Times(1).Return([]int{1, 2, 3}, nil),
					userRepo.EXPECT().FindProductNames(1).Times(1).Return("a", nil),
					userRepo.EXPECT().FindProductNames(2).Times(1).Return("b", nil),
					userRepo.EXPECT().FindProductNames(3).Times(1).Return("c", nil),
					userRepo.EXPECT().FindCartQuantity(1, 1).Times(1).Return(5, nil),
					userRepo.EXPECT().FindCartQuantity(1, 2).Times(1).Return(6, nil),
					userRepo.EXPECT().FindCartQuantity(1, 3).Times(1).Return(7, nil),
					userRepo.EXPECT().FindPrice(1).Times(1).Return(500.00, errors.New("internal error")),
				)
			},
			expectedOutput: []models.GetCart([]models.GetCart{}),
			expectedError:  errors.New("internal error"),
		},
		"find category error": {
			input: 1,
			StubDetails: func(userRepo mockrepo.MockUserRepository, orderRepo mockrepo.MockOrderRepository, helper mockhelper.MockHelper, data int) {
				gomock.InOrder(
					userRepo.EXPECT().GetCartID(data).Times(1).Return(1, nil),
					userRepo.EXPECT().GetProductsInCart(1).Times(1).Return([]int{1, 2, 3}, nil),
					userRepo.EXPECT().FindProductNames(1).Times(1).Return("a", nil),
					userRepo.EXPECT().FindProductNames(2).Times(1).Return("b", nil),
					userRepo.EXPECT().FindProductNames(3).Times(1).Return("c", nil),
					userRepo.EXPECT().FindCartQuantity(1, 1).Times(1).Return(5, nil),
					userRepo.EXPECT().FindCartQuantity(1, 2).Times(1).Return(6, nil),
					userRepo.EXPECT().FindCartQuantity(1, 3).Times(1).Return(7, nil),
					userRepo.EXPECT().FindPrice(1).Times(1).Return(500.00, nil),
					userRepo.EXPECT().FindPrice(2).Times(1).Return(520.00, nil),
					userRepo.EXPECT().FindPrice(3).Times(1).Return(600.00, nil),
					userRepo.EXPECT().FindCategory(1).Times(1).Return(1, errors.New("internal error")),
				)
			},
			expectedOutput: []models.GetCart([]models.GetCart{}),
			expectedError:  errors.New("internal error"),
		},
		"find offer percentage error": {
			input: 1,
			StubDetails: func(userRepo mockrepo.MockUserRepository, orderRepo mockrepo.MockOrderRepository, helper mockhelper.MockHelper, data int) {
				gomock.InOrder(
					userRepo.EXPECT().GetCartID(data).Times(1).Return(1, nil),
					userRepo.EXPECT().GetProductsInCart(1).Times(1).Return([]int{1, 2, 3}, nil),
					userRepo.EXPECT().FindProductNames(1).Times(1).Return("a", nil),
					userRepo.EXPECT().FindProductNames(2).Times(1).Return("b", nil),
					userRepo.EXPECT().FindProductNames(3).Times(1).Return("c", nil),
					userRepo.EXPECT().FindCartQuantity(1, 1).Times(1).Return(5, nil),
					userRepo.EXPECT().FindCartQuantity(1, 2).Times(1).Return(6, nil),
					userRepo.EXPECT().FindCartQuantity(1, 3).Times(1).Return(7, nil),
					userRepo.EXPECT().FindPrice(1).Times(1).Return(500.00, nil),
					userRepo.EXPECT().FindPrice(2).Times(1).Return(520.00, nil),
					userRepo.EXPECT().FindPrice(3).Times(1).Return(600.00, nil),
					userRepo.EXPECT().FindCategory(1).Times(1).Return(1, nil),
					userRepo.EXPECT().FindCategory(2).Times(1).Return(2, nil),
					userRepo.EXPECT().FindCategory(3).Times(1).Return(3, nil),
					userRepo.EXPECT().FindofferPercentage(1).Times(1).Return(10, errors.New("internal error")),
				)
			},
			expectedOutput: []models.GetCart([]models.GetCart{}),
			expectedError:  errors.New("internal error"),
		},
	}
	for _, test := range testData {

		test.StubDetails(*userRepo, *orderRepo, *helper, test.input)

		cart, err := userUseCase.GetCart(test.input)
		assert.Equal(t, test.expectedOutput, cart)
		assert.Equal(t, test.expectedError, err)

	}
}
