package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"jerseyhub/pkg/domain"
	"jerseyhub/pkg/mock/mockusecase"
	"jerseyhub/pkg/utils/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUserSignup(t *testing.T) {

	testCase := map[string]struct {
		input         interface{}
		buildStub     func(useCaseMock *mockusecase.MockUserUseCase, signup interface{})
		checkResponse func(t *testing.T, responseRecorder *httptest.ResponseRecorder)
	}{
		"Valid Signup": {
			input: models.UserDetails{
				Name:            "Arun K",
				Phone:           "6282246077",
				Email:           "arthurbishop120@gmail.com",
				Password:        "132456",
				ConfirmPassword: "123456",
			},
			buildStub: func(useCaseMock *mockusecase.MockUserUseCase, signupData interface{}) {
				// copying signupData to domain.user for pass to Mock usecase
				err := validator.New().Struct(signupData)
				if err != nil {
					fmt.Println("validation failed")
				}

				useCaseMock.EXPECT().UserSignUp(signupData, "").Times(1).Return(models.TokenUsers{
					Users: models.UserDetailsResponse{
						Id:    1,
						Name:  "Arun K",
						Email: "arthurbishop120@gmail.com",
						Phone: "6282246077",
					},
					Token: "aksjgnal.fiugliagbldfgbldf.gdbladfjnb",
				}, nil)
			},
			checkResponse: func(t *testing.T, responseRecorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusCreated, responseRecorder.Code)

			},
		},

		"fields provided in wrong format": {
			input: "",
			buildStub: func(useCaseMock *mockusecase.MockUserUseCase, signupData interface{}) {
				// copying signupData to domain.user for pass to Mock usecase
				err := validator.New().Struct(signupData)
				if err != nil {
					fmt.Println("validation failed")
				}
			},
			checkResponse: func(t *testing.T, responseRecorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)

			},
		},

		"struct validation fails": {
			input: models.SetNewName{},
			buildStub: func(useCaseMock *mockusecase.MockUserUseCase, signupData interface{}) {
				// copying signupData to domain.user for pass to Mock usecase
				err := validator.New().Struct(signupData)
				if err != nil {
					fmt.Println("validation failed")
				}
			},
			checkResponse: func(t *testing.T, responseRecorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)

			},
		},

		"user couldnot sign up": {
			input: models.UserDetails{
				Name:            "Arun K",
				Phone:           "6282246077",
				Email:           "arthurbishop120@gmail.com",
				Password:        "132456",
				ConfirmPassword: "123456",
			},
			buildStub: func(useCaseMock *mockusecase.MockUserUseCase, signupData interface{}) {
				// copying signupData to domain.user for pass to Mock usecase
				err := validator.New().Struct(signupData)
				if err != nil {
					fmt.Println("validation failed")
				}

				useCaseMock.EXPECT().UserSignUp(signupData, "").Times(1).Return(models.TokenUsers{}, errors.New("cannot sign up"))
			},
			checkResponse: func(t *testing.T, responseRecorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)

			},
		},
	}

	for testName, test := range testCase {
		test := test
		t.Run(testName, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			mockUseCase := mockusecase.NewMockUserUseCase(ctrl)
			test.buildStub(mockUseCase, test.input)

			userHandler := NewUserHandler(mockUseCase)

			server := gin.Default()
			server.POST("/signup", userHandler.UserSignUp)

			jsonData, err := json.Marshal(test.input)
			assert.NoError(t, err)
			body := bytes.NewBuffer(jsonData)

			mockRequest, err := http.NewRequest(http.MethodPost, "/signup", body)
			assert.NoError(t, err)
			responseRecorder := httptest.NewRecorder()

			server.ServeHTTP(responseRecorder, mockRequest)

			test.checkResponse(t, responseRecorder)

		})

	}
}

func TestLoginHandler(t *testing.T) {
	testCase := map[string]struct {
		input         interface{}
		buildStub     func(useCaseMock *mockusecase.MockUserUseCase, signup interface{})
		checkResponse func(t *testing.T, responseRecorder *httptest.ResponseRecorder)
	}{
		"successfull login": {
			input: models.UserLogin{
				Email:    "arthurbishop120@gmail.com",
				Password: "12345",
			},
			buildStub: func(useCaseMock *mockusecase.MockUserUseCase, loginData interface{}) {
				// copying signupData to domain.user for pass to Mock usecase
				err := validator.New().Struct(loginData)
				if err != nil {
					fmt.Println("validation failed")
				}

				useCaseMock.EXPECT().LoginHandler(loginData).Times(1).Return(models.TokenUsers{
					Users: models.UserDetailsResponse{
						Id:    1,
						Name:  "Arun K",
						Email: "arthurbishop120@gmail.com",
						Phone: "6282246077",
					},
					Token: "aksjgnal.fiugliagbldfgbldf.gdbladfjnb",
				}, nil)
			},
			checkResponse: func(t *testing.T, responseRecorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusOK, responseRecorder.Code)

			},
		},
		"fields provided in wrong format": {
			input: "",
			buildStub: func(useCaseMock *mockusecase.MockUserUseCase, loginData interface{}) {
				// copying signupData to domain.user for pass to Mock usecase
				err := validator.New().Struct(loginData)
				if err != nil {
					fmt.Println("validation failed")
				}
			},
			checkResponse: func(t *testing.T, responseRecorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)

			},
		},
		"struct validation fails": {
			input: models.SetNewName{},
			buildStub: func(useCaseMock *mockusecase.MockUserUseCase, loginData interface{}) {
				// copying signupData to domain.user for pass to Mock usecase
				err := validator.New().Struct(loginData)
				if err != nil {
					fmt.Println("validation failed")
				}
			},
			checkResponse: func(t *testing.T, responseRecorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)

			},
		},
		"user couldnot login": {
			input: models.UserLogin{
				Email:    "anonymous@some.com",
				Password: "no password",
			},
			buildStub: func(useCaseMock *mockusecase.MockUserUseCase, loginData interface{}) {
				// copying signupData to domain.user for pass to Mock usecase
				err := validator.New().Struct(loginData)
				if err != nil {
					fmt.Println("validation failed")
				}

				useCaseMock.EXPECT().LoginHandler(loginData).Times(1).Return(models.TokenUsers{}, errors.New("cannot login up"))
			},
			checkResponse: func(t *testing.T, responseRecorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)

			},
		},
	}

	for testName, test := range testCase {
		test := test
		t.Run(testName, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			mockUseCase := mockusecase.NewMockUserUseCase(ctrl)
			test.buildStub(mockUseCase, test.input)

			userHandler := NewUserHandler(mockUseCase)

			server := gin.Default()
			server.POST("/login", userHandler.LoginHandler)

			jsonData, err := json.Marshal(test.input)
			assert.NoError(t, err)
			body := bytes.NewBuffer(jsonData)

			mockRequest, err := http.NewRequest(http.MethodPost, "/login", body)
			assert.NoError(t, err)
			responseRecorder := httptest.NewRecorder()

			server.ServeHTTP(responseRecorder, mockRequest)

			test.checkResponse(t, responseRecorder)

		})

	}
}

func TestAddAddress(t *testing.T) {
	testCase := map[string]struct {
		input         interface{}
		buildStub     func(useCaseMock *mockusecase.MockUserUseCase, Data interface{})
		checkResponse func(t *testing.T, responseRecorder *httptest.ResponseRecorder)
	}{
		"successfull": {
			input: models.AddAddress{
				Name:      "Arun K",
				HouseName: "nellikkal arun bhavan",
				Street:    "pallippuram",
				City:      "cherthala",
				State:     "kerala",
				Pin:       "688541",
			},
			buildStub: func(useCaseMock *mockusecase.MockUserUseCase, Data interface{}) {

				useCaseMock.EXPECT().AddAddress(1, Data).Times(1).Return(nil)
			},
			checkResponse: func(t *testing.T, responseRecorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusOK, responseRecorder.Code)

			},
		},
		"parameter problem": {
			input: models.AddAddress{
				Name:      "Arun K",
				HouseName: "nellikkal arun bhavan",
				Street:    "pallippuram",
				City:      "cherthala",
				State:     "kerala",
				Pin:       "688541",
			},
			buildStub: func(useCaseMock *mockusecase.MockUserUseCase, Data interface{}) {

			},
			checkResponse: func(t *testing.T, responseRecorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)

			},
		},
		"fields provided in wrong format": {
			input: "",
			buildStub: func(useCaseMock *mockusecase.MockUserUseCase, Data interface{}) {

			},
			checkResponse: func(t *testing.T, responseRecorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)

			},
		},
		"could not add the address": {
			input: models.AddAddress{
				Name:      "Arun K",
				HouseName: "nellikkal arun bhavan",
				Street:    "pallippuram",
				City:      "cherthala",
				State:     "kerala",
				Pin:       "688541",
			},
			buildStub: func(useCaseMock *mockusecase.MockUserUseCase, Data interface{}) {
				useCaseMock.EXPECT().AddAddress(1, Data).Times(1).Return(errors.New("could not add the address"))
			},
			checkResponse: func(t *testing.T, responseRecorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
			},
		},
	}

	for testName, test := range testCase {
		testName := testName
		test := test
		t.Run(testName, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			mockUseCase := mockusecase.NewMockUserUseCase(ctrl)
			test.buildStub(mockUseCase, test.input)

			userHandler := NewUserHandler(mockUseCase)

			server := gin.Default()
			server.POST("/add_address", userHandler.AddAddress)

			jsonData, err := json.Marshal(test.input)
			assert.NoError(t, err)
			body := bytes.NewBuffer(jsonData)

			mockRequest, err := http.NewRequest(http.MethodPost, "/add_address?id=1", body)
			assert.NoError(t, err)
			if testName == "parameter problem" {
				mockRequest, err = http.NewRequest(http.MethodPost, "/add_address?id=invalid", body)
				assert.NoError(t, err)
			}
			responseRecorder := httptest.NewRecorder()

			server.ServeHTTP(responseRecorder, mockRequest)

			test.checkResponse(t, responseRecorder)

		})

	}

}

func TestGetAddresses(t *testing.T) {
	testCase := map[string]struct {
		buildStub     func(useCaseMock *mockusecase.MockUserUseCase)
		checkResponse func(t *testing.T, responseRecorder *httptest.ResponseRecorder)
		expected      []domain.Address
	}{
		"successfull": {
			buildStub: func(useCaseMock *mockusecase.MockUserUseCase) {

				useCaseMock.EXPECT().GetAddresses(1).Times(1).Return([]domain.Address{}, nil)
			},
			checkResponse: func(t *testing.T, responseRecorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusOK, responseRecorder.Code)

			},
		},
		"parameter problem": {
			buildStub: func(useCaseMock *mockusecase.MockUserUseCase) {
			},
			checkResponse: func(t *testing.T, responseRecorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)

			},
		},
		"error retrieving records": {
			buildStub: func(useCaseMock *mockusecase.MockUserUseCase) {

				useCaseMock.EXPECT().GetAddresses(1).Times(1).Return(nil, errors.New("error retrieving records"))
			},
			checkResponse: func(t *testing.T, responseRecorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)

			},
		},
	}

	for testName, test := range testCase {
		testName := testName
		test := test
		t.Run(testName, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			mockUseCase := mockusecase.NewMockUserUseCase(ctrl)
			test.buildStub(mockUseCase)

			userHandler := NewUserHandler(mockUseCase)

			server := gin.Default()
			server.POST("/getAddresses", userHandler.GetAddresses)

			mockRequest, err := http.NewRequest(http.MethodPost, "/getAddresses?id=1", nil)
			assert.NoError(t, err)
			if testName == "parameter problem" {
				mockRequest, err = http.NewRequest(http.MethodPost, "/getAddresses?id=invalid", nil)
				assert.NoError(t, err)
			}
			responseRecorder := httptest.NewRecorder()

			server.ServeHTTP(responseRecorder, mockRequest)

			test.checkResponse(t, responseRecorder)

		})

	}
}

func TestGetUserDetails(t *testing.T) {
	testCase := map[string]struct {
		buildStub     func(useCaseMock *mockusecase.MockUserUseCase)
		checkResponse func(t *testing.T, responseRecorder *httptest.ResponseRecorder)
		expected      []models.UserDetailsResponse
	}{
		"successfull": {
			buildStub: func(useCaseMock *mockusecase.MockUserUseCase) {

				useCaseMock.EXPECT().GetUserDetails(1).Times(1).Return(models.UserDetailsResponse{}, nil)
			},
			checkResponse: func(t *testing.T, responseRecorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusOK, responseRecorder.Code)

			},
		},
		"parameter problem": {
			buildStub: func(useCaseMock *mockusecase.MockUserUseCase) {
			},
			checkResponse: func(t *testing.T, responseRecorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)

			},
		},
		"error retrieving records": {
			buildStub: func(useCaseMock *mockusecase.MockUserUseCase) {

				useCaseMock.EXPECT().GetUserDetails(1).Times(1).Return(models.UserDetailsResponse{}, errors.New("error retrieving records"))
			},
			checkResponse: func(t *testing.T, responseRecorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)

			},
		},
	}

	for testName, test := range testCase {
		testName := testName
		test := test
		t.Run(testName, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			mockUseCase := mockusecase.NewMockUserUseCase(ctrl)
			test.buildStub(mockUseCase)

			userHandler := NewUserHandler(mockUseCase)

			server := gin.Default()
			server.POST("/get", userHandler.GetUserDetails)

			mockRequest, err := http.NewRequest(http.MethodPost, "/get?id=1", nil)
			assert.NoError(t, err)
			if testName == "parameter problem" {
				mockRequest, err = http.NewRequest(http.MethodPost, "/get?id=invalid", nil)
				assert.NoError(t, err)
			}
			responseRecorder := httptest.NewRecorder()

			server.ServeHTTP(responseRecorder, mockRequest)

			test.checkResponse(t, responseRecorder)

		})

	}
}

func TestChangePassword(t *testing.T) {
	testCase := map[string]struct {
		input         interface{}
		buildStub     func(useCaseMock *mockusecase.MockUserUseCase)
		checkResponse func(t *testing.T, responseRecorder *httptest.ResponseRecorder)
	}{
		"successfull": {
			input: models.ChangePassword{
				Oldpassword: "1234",
				Password:    "4321",
				Repassword:  "4321",
			},
			buildStub: func(useCaseMock *mockusecase.MockUserUseCase) {

				useCaseMock.EXPECT().ChangePassword(1, "1234", "4321", "4321").Times(1).Return(nil)
			},
			checkResponse: func(t *testing.T, responseRecorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusOK, responseRecorder.Code)

			},
		},
		"parameter problem": {
			input: models.ChangePassword{
				Oldpassword: "1234",
				Password:    "4321",
				Repassword:  "4321",
			},
			buildStub: func(useCaseMock *mockusecase.MockUserUseCase) {
			},
			checkResponse: func(t *testing.T, responseRecorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)

			},
		},
		"fields provided in wrong format": {
			input: "",
			buildStub: func(useCaseMock *mockusecase.MockUserUseCase) {

			},
			checkResponse: func(t *testing.T, responseRecorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)

			},
		},
		"could not change the password": {
			input: models.ChangePassword{
				Oldpassword: "1234",
				Password:    "4321",
				Repassword:  "4321",
			},
			buildStub: func(useCaseMock *mockusecase.MockUserUseCase) {

				useCaseMock.EXPECT().ChangePassword(1, "1234", "4321", "4321").Times(1).Return(errors.New("could not change the password"))
			},
			checkResponse: func(t *testing.T, responseRecorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)

			},
		},
	}
	for testName, test := range testCase {
		testName := testName
		test := test
		t.Run(testName, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			mockUseCase := mockusecase.NewMockUserUseCase(ctrl)
			test.buildStub(mockUseCase)

			userHandler := NewUserHandler(mockUseCase)

			server := gin.Default()
			server.POST("/change_password", userHandler.ChangePassword)

			jsonData, err := json.Marshal(test.input)
			assert.NoError(t, err)
			body := bytes.NewBuffer(jsonData)

			mockRequest, err := http.NewRequest(http.MethodPost, "/change_password?id=1", body)
			assert.NoError(t, err)
			if testName == "parameter problem" {
				mockRequest, err = http.NewRequest(http.MethodPost, "/change_password?id=invalid", body)
				assert.NoError(t, err)
			}
			responseRecorder := httptest.NewRecorder()

			server.ServeHTTP(responseRecorder, mockRequest)

			test.checkResponse(t, responseRecorder)

		})

	}
}

func TestForgotPasswordSend(t *testing.T) {
	testCase := map[string]struct {
		input         interface{}
		buildStub     func(useCaseMock *mockusecase.MockUserUseCase)
		checkResponse func(t *testing.T, responseRecorder *httptest.ResponseRecorder)
	}{
		"successfull": {
			input: models.ForgotPasswordSend{
				Phone: "6282246077",
			},
			buildStub: func(useCaseMock *mockusecase.MockUserUseCase) {

				useCaseMock.EXPECT().ForgotPasswordSend("6282246077").Times(1).Return(nil)
			},
			checkResponse: func(t *testing.T, responseRecorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusOK, responseRecorder.Code)

			},
		},
		"fields provided in wrong format": {
			input: "",
			buildStub: func(useCaseMock *mockusecase.MockUserUseCase) {
			},
			checkResponse: func(t *testing.T, responseRecorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)

			},
		},
		"cannot send otp": {
			input: models.ForgotPasswordSend{
				Phone: "6282246077",
			},
			buildStub: func(useCaseMock *mockusecase.MockUserUseCase) {

				useCaseMock.EXPECT().ForgotPasswordSend("6282246077").Times(1).Return(errors.New("cannot send otp"))
			},
			checkResponse: func(t *testing.T, responseRecorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)

			},
		},
	}

	for testName, test := range testCase {
		testName := testName
		test := test
		t.Run(testName, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			mockUseCase := mockusecase.NewMockUserUseCase(ctrl)
			test.buildStub(mockUseCase)

			userHandler := NewUserHandler(mockUseCase)

			server := gin.Default()
			server.POST("/forgot_password_send", userHandler.ForgotPasswordSend)

			jsonData, err := json.Marshal(test.input)
			assert.NoError(t, err)
			body := bytes.NewBuffer(jsonData)

			mockRequest, err := http.NewRequest(http.MethodPost, "/forgot_password_send?id=1", body)
			assert.NoError(t, err)
			if testName == "parameter problem" {
				mockRequest, err = http.NewRequest(http.MethodPost, "/change_password?id=invalid", body)
				assert.NoError(t, err)
			}
			responseRecorder := httptest.NewRecorder()

			server.ServeHTTP(responseRecorder, mockRequest)

			test.checkResponse(t, responseRecorder)

		})

	}
}

func TestForgotPasswordVerifyAndChange(t *testing.T) {
	testCase := map[string]struct {
		input         interface{}
		buildStub     func(useCaseMock *mockusecase.MockUserUseCase, data interface{})
		checkResponse func(t *testing.T, responseRecorder *httptest.ResponseRecorder)
	}{
		"successfull": {
			input: models.ForgotVerify{
				Phone:       "6282246077",
				Otp:         "1234",
				NewPassword: "4321",
			},
			buildStub: func(useCaseMock *mockusecase.MockUserUseCase, data interface{}) {

				useCaseMock.EXPECT().ForgotPasswordVerifyAndChange(data).Times(1).Return(nil)
			},
			checkResponse: func(t *testing.T, responseRecorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusOK, responseRecorder.Code)

			},
		},
		"fields provided in wrong format": {
			input: "",
			buildStub: func(useCaseMock *mockusecase.MockUserUseCase, data interface{}) {
			},
			checkResponse: func(t *testing.T, responseRecorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)

			},
		},
		"failure": {
			input: models.ForgotVerify{
				Phone:       "6282246077",
				Otp:         "1234",
				NewPassword: "4321",
			},
			buildStub: func(useCaseMock *mockusecase.MockUserUseCase, data interface{}) {

				useCaseMock.EXPECT().ForgotPasswordVerifyAndChange(data).Times(1).Return(errors.New("could not change password"))
			},
			checkResponse: func(t *testing.T, responseRecorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)

			},
		},
	}

	for testName, test := range testCase {
		testName := testName
		test := test
		t.Run(testName, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			mockUseCase := mockusecase.NewMockUserUseCase(ctrl)
			test.buildStub(mockUseCase, test.input)

			userHandler := NewUserHandler(mockUseCase)

			server := gin.Default()
			server.POST("/forgot_password_verify", userHandler.ForgotPasswordVerifyAndChange)

			jsonData, err := json.Marshal(test.input)
			assert.NoError(t, err)
			body := bytes.NewBuffer(jsonData)

			mockRequest, err := http.NewRequest(http.MethodPost, "/forgot_password_verify", body)
			assert.NoError(t, err)

			responseRecorder := httptest.NewRecorder()

			server.ServeHTTP(responseRecorder, mockRequest)

			test.checkResponse(t, responseRecorder)

		})

	}

}

func TestEditName(t *testing.T) {
	testCase := map[string]struct {
		input         interface{}
		buildStub     func(useCaseMock *mockusecase.MockUserUseCase)
		checkResponse func(t *testing.T, responseRecorder *httptest.ResponseRecorder)
	}{
		"successfull": {
			input: models.EditName{
				Name: "Arun K",
			},
			buildStub: func(useCaseMock *mockusecase.MockUserUseCase) {

				useCaseMock.EXPECT().EditName(1, "Arun K").Times(1).Return(nil)
			},
			checkResponse: func(t *testing.T, responseRecorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusOK, responseRecorder.Code)

			},
		},
		"parameter problem": {
			input: models.EditName{
				Name: "Arun K",
			},
			buildStub: func(useCaseMock *mockusecase.MockUserUseCase) {
			},
			checkResponse: func(t *testing.T, responseRecorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)

			},
		},
		"fields provided in wrong format": {
			input: "",
			buildStub: func(useCaseMock *mockusecase.MockUserUseCase) {

			},
			checkResponse: func(t *testing.T, responseRecorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)

			},
		},
		"could not change the Name": {
			input: models.EditName{
				Name: "Arun K",
			},
			buildStub: func(useCaseMock *mockusecase.MockUserUseCase) {

				useCaseMock.EXPECT().EditName(1, "Arun K").Times(1).Return(errors.New("could not change the name"))
			},
			checkResponse: func(t *testing.T, responseRecorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)

			},
		},
	}
	for testName, test := range testCase {
		testName := testName
		test := test
		t.Run(testName, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			mockUseCase := mockusecase.NewMockUserUseCase(ctrl)
			test.buildStub(mockUseCase)

			userHandler := NewUserHandler(mockUseCase)

			server := gin.Default()
			server.POST("/edit_name", userHandler.EditName)

			jsonData, err := json.Marshal(test.input)
			assert.NoError(t, err)
			body := bytes.NewBuffer(jsonData)

			mockRequest, err := http.NewRequest(http.MethodPost, "/edit_name?id=1", body)
			assert.NoError(t, err)
			if testName == "parameter problem" {
				mockRequest, err = http.NewRequest(http.MethodPost, "/edit_name?id=invalid", body)
				assert.NoError(t, err)
			}
			responseRecorder := httptest.NewRecorder()

			server.ServeHTTP(responseRecorder, mockRequest)

			test.checkResponse(t, responseRecorder)

		})

	}
}

func TestEditEmail(t *testing.T) {
	testCase := map[string]struct {
		input         interface{}
		buildStub     func(useCaseMock *mockusecase.MockUserUseCase)
		checkResponse func(t *testing.T, responseRecorder *httptest.ResponseRecorder)
	}{
		"successfull": {
			input: models.EditEmail{
				Email: "arthurbishop120@gmail.com",
			},
			buildStub: func(useCaseMock *mockusecase.MockUserUseCase) {

				useCaseMock.EXPECT().EditEmail(1, "arthurbishop120@gmail.com").Times(1).Return(nil)
			},
			checkResponse: func(t *testing.T, responseRecorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusOK, responseRecorder.Code)

			},
		},
		"parameter problem": {
			input: models.EditEmail{
				Email: "arthurbishop120@gmail.com",
			},
			buildStub: func(useCaseMock *mockusecase.MockUserUseCase) {
			},
			checkResponse: func(t *testing.T, responseRecorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)

			},
		},
		"fields provided in wrong format": {
			input: "",
			buildStub: func(useCaseMock *mockusecase.MockUserUseCase) {

			},
			checkResponse: func(t *testing.T, responseRecorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)

			},
		},
		"could not change the Email": {
			input: models.EditEmail{
				Email: "arthurbishop120@gmail.com",
			},
			buildStub: func(useCaseMock *mockusecase.MockUserUseCase) {

				useCaseMock.EXPECT().EditEmail(1, "arthurbishop120@gmail.com").Times(1).Return(errors.New("could not change the email"))
			},
			checkResponse: func(t *testing.T, responseRecorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)

			},
		},
	}
	for testName, test := range testCase {
		testName := testName
		test := test
		t.Run(testName, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			mockUseCase := mockusecase.NewMockUserUseCase(ctrl)
			test.buildStub(mockUseCase)

			userHandler := NewUserHandler(mockUseCase)

			server := gin.Default()
			server.POST("/edit_email", userHandler.EditEmail)

			jsonData, err := json.Marshal(test.input)
			assert.NoError(t, err)
			body := bytes.NewBuffer(jsonData)

			mockRequest, err := http.NewRequest(http.MethodPost, "/edit_email?id=1", body)
			assert.NoError(t, err)
			if testName == "parameter problem" {
				mockRequest, err = http.NewRequest(http.MethodPost, "/edit_email?id=invalid", body)
				assert.NoError(t, err)
			}
			responseRecorder := httptest.NewRecorder()

			server.ServeHTTP(responseRecorder, mockRequest)

			test.checkResponse(t, responseRecorder)

		})

	}
}

func TestEditPhone(t *testing.T) {
	testCase := map[string]struct {
		input         interface{}
		buildStub     func(useCaseMock *mockusecase.MockUserUseCase)
		checkResponse func(t *testing.T, responseRecorder *httptest.ResponseRecorder)
	}{
		"successfull": {
			input: models.EditPhone{
				Phone: "6282246077",
			},
			buildStub: func(useCaseMock *mockusecase.MockUserUseCase) {

				useCaseMock.EXPECT().EditPhone(1, "6282246077").Times(1).Return(nil)
			},
			checkResponse: func(t *testing.T, responseRecorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusOK, responseRecorder.Code)

			},
		},
		"parameter problem": {
			input: models.EditPhone{
				Phone: "6282246077",
			},
			buildStub: func(useCaseMock *mockusecase.MockUserUseCase) {
			},
			checkResponse: func(t *testing.T, responseRecorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)

			},
		},
		"fields provided in wrong format": {
			input: "",
			buildStub: func(useCaseMock *mockusecase.MockUserUseCase) {

			},
			checkResponse: func(t *testing.T, responseRecorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)

			},
		},
		"could not change the phone": {
			input: models.EditPhone{
				Phone: "6282246077",
			},
			buildStub: func(useCaseMock *mockusecase.MockUserUseCase) {

				useCaseMock.EXPECT().EditPhone(1, "6282246077").Times(1).Return(errors.New("could not change the phone"))
			},
			checkResponse: func(t *testing.T, responseRecorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
			},
		},
	}
	for testName, test := range testCase {
		testName := testName
		test := test
		t.Run(testName, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			mockUseCase := mockusecase.NewMockUserUseCase(ctrl)
			test.buildStub(mockUseCase)

			userHandler := NewUserHandler(mockUseCase)

			server := gin.Default()
			server.POST("/edit_phone", userHandler.EditPhone)

			jsonData, err := json.Marshal(test.input)
			assert.NoError(t, err)
			body := bytes.NewBuffer(jsonData)

			mockRequest, err := http.NewRequest(http.MethodPost, "/edit_phone?id=1", body)
			assert.NoError(t, err)
			if testName == "parameter problem" {
				mockRequest, err = http.NewRequest(http.MethodPost, "/edit_phone?id=invalid", body)
				assert.NoError(t, err)
			}
			responseRecorder := httptest.NewRecorder()

			server.ServeHTTP(responseRecorder, mockRequest)

			test.checkResponse(t, responseRecorder)

		})

	}
}

func TestGetCart(t *testing.T) {
	testCase := map[string]struct {
		buildStub     func(useCaseMock *mockusecase.MockUserUseCase)
		checkResponse func(t *testing.T, responseRecorder *httptest.ResponseRecorder)
		expected      []models.GetCart
	}{
		"successfull": {
			buildStub: func(useCaseMock *mockusecase.MockUserUseCase) {

				useCaseMock.EXPECT().GetCart(1).Times(1).Return([]models.GetCart{}, nil)
			},
			checkResponse: func(t *testing.T, responseRecorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusOK, responseRecorder.Code)

			},
		},
		"parameter problem": {
			buildStub: func(useCaseMock *mockusecase.MockUserUseCase) {
			},
			checkResponse: func(t *testing.T, responseRecorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)

			},
		},
		"error retrieving records": {
			buildStub: func(useCaseMock *mockusecase.MockUserUseCase) {

				useCaseMock.EXPECT().GetCart(1).Times(1).Return(nil, errors.New("error retrieving records"))
			},
			checkResponse: func(t *testing.T, responseRecorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)

			},
		},
	}

	for testName, test := range testCase {
		testName := testName
		test := test
		t.Run(testName, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			mockUseCase := mockusecase.NewMockUserUseCase(ctrl)
			test.buildStub(mockUseCase)

			userHandler := NewUserHandler(mockUseCase)

			server := gin.Default()
			server.POST("/get_cart", userHandler.GetCart)

			mockRequest, err := http.NewRequest(http.MethodPost, "/get_cart?id=1", nil)
			assert.NoError(t, err)
			if testName == "parameter problem" {
				mockRequest, err = http.NewRequest(http.MethodPost, "/get_cart?id=invalid", nil)
				assert.NoError(t, err)
			}
			responseRecorder := httptest.NewRecorder()

			server.ServeHTTP(responseRecorder, mockRequest)

			test.checkResponse(t, responseRecorder)

		})

	}
}

func TestRemoveFromCart(t *testing.T) {
	testCase := map[string]struct {
		buildStub     func(useCaseMock *mockusecase.MockUserUseCase)
		checkResponse func(t *testing.T, responseRecorder *httptest.ResponseRecorder)
	}{
		"successfull": {
			buildStub: func(useCaseMock *mockusecase.MockUserUseCase) {

				useCaseMock.EXPECT().RemoveFromCart(1).Times(1).Return(nil)
			},
			checkResponse: func(t *testing.T, responseRecorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusOK, responseRecorder.Code)

			},
		},
		"parameter problem": {
			buildStub: func(useCaseMock *mockusecase.MockUserUseCase) {

			},
			checkResponse: func(t *testing.T, responseRecorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)

			},
		},
		"could not remove from cart": {
			buildStub: func(useCaseMock *mockusecase.MockUserUseCase) {

				useCaseMock.EXPECT().RemoveFromCart(1).Times(1).Return(errors.New("could not remove from cart"))
			},
			checkResponse: func(t *testing.T, responseRecorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)

			},
		},
	}
	for testName, test := range testCase {
		testName := testName
		test := test
		t.Run(testName, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			mockUseCase := mockusecase.NewMockUserUseCase(ctrl)
			test.buildStub(mockUseCase)

			userHandler := NewUserHandler(mockUseCase)

			server := gin.Default()
			server.POST("/remove_from_cart", userHandler.RemoveFromCart)

			mockRequest, err := http.NewRequest(http.MethodPost, "/remove_from_cart?id=1", nil)
			assert.NoError(t, err)
			if testName == "parameter problem" {
				mockRequest, err = http.NewRequest(http.MethodPost, "/remove_from_cart?id=invalid", nil)
				assert.NoError(t, err)
			}
			responseRecorder := httptest.NewRecorder()

			server.ServeHTTP(responseRecorder, mockRequest)

			test.checkResponse(t, responseRecorder)

		})

	}
}

func TestUpdateQuantityAdd(t *testing.T) {
	testCase := map[string]struct {
		buildStub     func(useCaseMock *mockusecase.MockUserUseCase)
		checkResponse func(t *testing.T, responseRecorder *httptest.ResponseRecorder)
	}{
		"successfull": {
			buildStub: func(useCaseMock *mockusecase.MockUserUseCase) {

				useCaseMock.EXPECT().UpdateQuantityAdd(1, 1).Times(1).Return(nil)
			},
			checkResponse: func(t *testing.T, responseRecorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusOK, responseRecorder.Code)

			},
		},
		"parameter problem id": {
			buildStub: func(useCaseMock *mockusecase.MockUserUseCase) {

			},
			checkResponse: func(t *testing.T, responseRecorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)

			},
		},
		"parameter problem inventory": {
			buildStub: func(useCaseMock *mockusecase.MockUserUseCase) {

			},
			checkResponse: func(t *testing.T, responseRecorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)

			},
		},
		"could not add the quantity": {
			buildStub: func(useCaseMock *mockusecase.MockUserUseCase) {

				useCaseMock.EXPECT().UpdateQuantityAdd(1, 1).Times(1).Return(errors.New("could not add the quantity"))
			},
			checkResponse: func(t *testing.T, responseRecorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)

			},
		},
	}

	for testName, test := range testCase {
		testName := testName
		test := test
		t.Run(testName, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			mockUseCase := mockusecase.NewMockUserUseCase(ctrl)
			test.buildStub(mockUseCase)

			userHandler := NewUserHandler(mockUseCase)

			server := gin.Default()
			server.POST("/remove_from_cart", userHandler.UpdateQuantityAdd)

			mockRequest, err := http.NewRequest(http.MethodPost, "/remove_from_cart?id=1&inventory=1", nil)
			assert.NoError(t, err)
			if testName == "parameter problem id" {
				mockRequest, err = http.NewRequest(http.MethodPost, "/remove_from_cart?id=invalid&inventory=1", nil)
				assert.NoError(t, err)
			}
			if testName == "parameter problem inventory" {
				mockRequest, err = http.NewRequest(http.MethodPost, "/remove_from_cart?id=1&inventory=invalid", nil)
				assert.NoError(t, err)
			}
			responseRecorder := httptest.NewRecorder()

			server.ServeHTTP(responseRecorder, mockRequest)

			test.checkResponse(t, responseRecorder)

		})

	}
}

func TestUpdateQuantityLess(t *testing.T) {
	testCase := map[string]struct {
		buildStub     func(useCaseMock *mockusecase.MockUserUseCase)
		checkResponse func(t *testing.T, responseRecorder *httptest.ResponseRecorder)
	}{
		"successfull": {
			buildStub: func(useCaseMock *mockusecase.MockUserUseCase) {

				useCaseMock.EXPECT().UpdateQuantityLess(1, 1).Times(1).Return(nil)
			},
			checkResponse: func(t *testing.T, responseRecorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusOK, responseRecorder.Code)

			},
		},
		"parameter problem id": {
			buildStub: func(useCaseMock *mockusecase.MockUserUseCase) {

			},
			checkResponse: func(t *testing.T, responseRecorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)

			},
		},
		"parameter problem inventory": {
			buildStub: func(useCaseMock *mockusecase.MockUserUseCase) {

			},
			checkResponse: func(t *testing.T, responseRecorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)

			},
		},
		"could not add the quantity": {
			buildStub: func(useCaseMock *mockusecase.MockUserUseCase) {

				useCaseMock.EXPECT().UpdateQuantityLess(1, 1).Times(1).Return(errors.New("could not add the quantity"))
			},
			checkResponse: func(t *testing.T, responseRecorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)

			},
		},
	}

	for testName, test := range testCase {
		testName := testName
		test := test
		t.Run(testName, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			mockUseCase := mockusecase.NewMockUserUseCase(ctrl)
			test.buildStub(mockUseCase)

			userHandler := NewUserHandler(mockUseCase)

			server := gin.Default()
			server.POST("/remove_from_cart", userHandler.UpdateQuantityLess)

			mockRequest, err := http.NewRequest(http.MethodPost, "/remove_from_cart?id=1&inventory=1", nil)
			assert.NoError(t, err)
			if testName == "parameter problem id" {
				mockRequest, err = http.NewRequest(http.MethodPost, "/remove_from_cart?id=invalid&inventory=1", nil)
				assert.NoError(t, err)
			}
			if testName == "parameter problem inventory" {
				mockRequest, err = http.NewRequest(http.MethodPost, "/remove_from_cart?id=1&inventory=invalid", nil)
				assert.NoError(t, err)
			}
			responseRecorder := httptest.NewRecorder()

			server.ServeHTTP(responseRecorder, mockRequest)

			test.checkResponse(t, responseRecorder)

		})

	}
}

func TestGetMyReferenceLink(t *testing.T) {
	testCase := map[string]struct {
		buildStub     func(useCaseMock *mockusecase.MockUserUseCase)
		checkResponse func(t *testing.T, responseRecorder *httptest.ResponseRecorder)
	}{
		"successfull": {
			buildStub: func(useCaseMock *mockusecase.MockUserUseCase) {

				useCaseMock.EXPECT().GetMyReferenceLink(1).Times(1).Return(gomock.Any().String(), nil)
			},
			checkResponse: func(t *testing.T, responseRecorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusOK, responseRecorder.Code)

			},
		},
		"parameter problem": {
			buildStub: func(useCaseMock *mockusecase.MockUserUseCase) {

			},
			checkResponse: func(t *testing.T, responseRecorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)

			},
		},
		"could not retrieve reference link": {
			buildStub: func(useCaseMock *mockusecase.MockUserUseCase) {

				useCaseMock.EXPECT().GetMyReferenceLink(1).Times(1).Return(gomock.Any().String(), errors.New("could not retrieve reference link"))
			},
			checkResponse: func(t *testing.T, responseRecorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)

			},
		},
	}
	for testName, test := range testCase {
		testName := testName
		test := test
		t.Run(testName, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			mockUseCase := mockusecase.NewMockUserUseCase(ctrl)
			test.buildStub(mockUseCase)

			userHandler := NewUserHandler(mockUseCase)

			server := gin.Default()
			server.POST("/remove_from_cart", userHandler.GetMyReferenceLink)

			mockRequest, err := http.NewRequest(http.MethodPost, "/remove_from_cart?id=1", nil)
			assert.NoError(t, err)
			if testName == "parameter problem" {
				mockRequest, err = http.NewRequest(http.MethodPost, "/remove_from_cart?id=invalid", nil)
				assert.NoError(t, err)
			}
			responseRecorder := httptest.NewRecorder()

			server.ServeHTTP(responseRecorder, mockRequest)

			test.checkResponse(t, responseRecorder)

		})
	}
}
