package handler

import (
	services "jerseyhub/pkg/usecase/interface"
	"jerseyhub/pkg/utils/models"
	"jerseyhub/pkg/utils/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CouponHandler struct {
	usecase services.CouponUsecase
}

func NewCouponHandler(use services.CouponUsecase) *CouponHandler {
	return &CouponHandler{
		usecase: use,
	}
}

// @Summary		Add Coupon
// @Description	Admin can add new coupons
// @Tags			Admin
// @Accept			json
// @Produce		    json
// @Param			coupon	body	models.Coupons	true	"coupon"
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/admin/coupons/create [post]
func (coup *CouponHandler) CreateNewCoupon(c *gin.Context) {
	var coupon models.Coupons
	if err := c.BindJSON(&coupon); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	err := coup.usecase.AddCoupon(coupon)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Could not add the Coupon", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully added Coupon", nil, nil)
	c.JSON(http.StatusOK, successRes)

}

// @Summary		Make Coupon ad invalid
// @Description	Admin can make the coupons as invalid so that users cannot use that particular coupon
// @Tags			Admin
// @Accept			json
// @Produce		    json
// @Param			id	query	string	true	"id"
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/admin/coupons/delete [delete]
func (coup *CouponHandler) MakeCouponInvalid(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	if err := coup.usecase.MakeCouponInvalid(id); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Coupon cannot be made invalid", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully made Coupon as invaid", nil, nil)
	c.JSON(http.StatusOK, successRes)

}

func (co *CouponHandler) GetAllCoupons(c *gin.Context) {

	categories, err := co.usecase.GetAllCoupons()
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "error in getting coupons", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully got all categories", categories, nil)
	c.JSON(http.StatusOK, successRes)

}
