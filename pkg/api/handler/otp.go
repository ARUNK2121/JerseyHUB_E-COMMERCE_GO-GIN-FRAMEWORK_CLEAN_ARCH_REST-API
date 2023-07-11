package handler

import (
	"net/http"

	services "jerseyhub/pkg/usecase/interface"
	models "jerseyhub/pkg/utils/models"
	"jerseyhub/pkg/utils/response"

	"github.com/gin-gonic/gin"
)

type OtpHandler struct {
	otpUseCase services.OtpUseCase
}

func NewOtpHandler(useCase services.OtpUseCase) *OtpHandler {
	return &OtpHandler{
		otpUseCase: useCase,
	}
}

// @Summary		Send OTP
// @Description	OTP login send otp
// @Tags			User
// @Accept			json
// @Produce		    json
// @Param			otp  body  models.OTPData true	"otp-data"
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/users/otplogin [post]
func (ot *OtpHandler) SendOTP(c *gin.Context) {

	var phone models.OTPData
	if err := c.BindJSON(&phone); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
	}

	err := ot.otpUseCase.SendOTP(phone.PhoneNumber)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Could not send OTP", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "OTP sent successfully", nil, nil)
	c.JSON(http.StatusOK, successRes)

}

// @Summary		Verify OTP
// @Description	OTP login verify otp
// @Tags			User
// @Accept			json
// @Produce		    json
// @Param			otp  body  models.VerifyData  true	"otp-verify"
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/users/verifyotp [post]
func (ot *OtpHandler) VerifyOTP(c *gin.Context) {

	var code models.VerifyData
	if err := c.BindJSON(&code); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	users, err := ot.otpUseCase.VerifyOTP(code)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Could not verify OTP", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully verified OTP", users, nil)
	c.JSON(http.StatusOK, successRes)

}
