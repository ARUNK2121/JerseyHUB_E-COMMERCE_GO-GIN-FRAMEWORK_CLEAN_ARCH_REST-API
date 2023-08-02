package usecase

import (
	"errors"
	"testing"

	"jerseyhub/pkg/config"
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
