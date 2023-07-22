package usecase

import (
	"testing"

	"jerseyhub/pkg/config"
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

	cfg := config.Config{}
	userUseCase := NewUserUseCase(userRepo, cfg, otpRepo, inventoryRepo, orderRepo)

	testData := []struct {
		name           string
		input          models.UserDetails
		buildStub      func(mockrepo.MockUserRepository, mockrepo.MockOrderRepository, models.UserDetails)
		expectedOutput models.TokenUsers
		expectedError  error
	}{
		{
			name: "success",
			input: models.UserDetails{
				Name:            "Arun K",
				Email:           "arthurbishop120@gmail.com",
				Phone:           "6282246077",
				Password:        gomock.Any().String(),
				ConfirmPassword: gomock.Any().String(),
			},
			buildStub: func(userRepo mockrepo.MockUserRepository, orderRepo mockrepo.MockOrderRepository, signupData models.UserDetails) {
				userRepo.EXPECT().CheckUserAvailability("arthurbishop120@gmail.com").Times(1).Return(false)

				userRepo.EXPECT().FindUserFromReference("12345").Times(1).Return(1, nil)

				userRepo.EXPECT().
					UserSignUp(models.UserDetails{
						Name:            "Arun K",
						Email:           "arthurbishop120@gmail.com",
						Phone:           "6282246077",
						Password:        gomock.Any().String(),
						ConfirmPassword: gomock.Any().String(),
					}, "12345").
					Times(1).
					Return(
						models.UserDetailsResponse{
							Id:    1,
							Name:  "Arun K",
							Email: "arthurbishop120@gmail.com",
							Phone: "6282246077",
						}, nil)

				userRepo.EXPECT().CreditReferencePointsToWallet(1).Times(1).Return(nil)

				orderRepo.EXPECT().CreateNewWallet(1).Times(1).Return(1, nil)
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
		// Add more test cases here if needed
	}

	for _, tt := range testData {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {

			gomock.InOrder(
				userRepo.EXPECT().CheckUserAvailability(tt.input.Email).Times(1).Return(false),
				userRepo.EXPECT().FindUserFromReference("12345").Times(1).Return(1, nil),
				userRepo.EXPECT().UserSignUp(models.UserDetails{
					Name:            "Arun K",
					Email:           "arthurbishop120@gmail.com",
					Phone:           "6282246077",
					Password:        gomock.Any().String(),
					ConfirmPassword: gomock.Any().String(),
				}, gomock.Any().String()).Times(1).Return(
					models.UserDetailsResponse{
						Id:    1,
						Name:  "Arun K",
						Email: "arthurbishop120@gmail.com",
						Phone: "6282246077",
					}, nil,
				),
				userRepo.EXPECT().CreditReferencePointsToWallet(1).Times(1).Return(nil),
				orderRepo.EXPECT().CreateNewWallet(1).Times(1).Return(1, nil),
			)

			// tt.buildStub(*userRepo, *orderRepo, tt.input)

			tokenusers, err := userUseCase.UserSignUp(tt.input, "12345")

			assert.Equal(t, tt.expectedOutput.Users.Id, tokenusers.Users.Id)
			assert.Equal(t, tt.expectedOutput.Users.Name, tokenusers.Users.Name)
			assert.Equal(t, tt.expectedOutput.Users.Email, tokenusers.Users.Email)
			assert.Equal(t, tt.expectedError, err)
		})
	}
}
