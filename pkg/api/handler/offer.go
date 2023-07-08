package handler

import (
	services "jerseyhub/pkg/usecase/interface"
	"jerseyhub/pkg/utils/models"
	"jerseyhub/pkg/utils/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OfferHandler struct {
	usecase services.OfferUseCase
}

func NewOfferHandler(usecase services.OfferUseCase) *OfferHandler {
	return &OfferHandler{
		usecase: usecase,
	}
}

// @Summary		Add Offer
// @Description	Admin can add new offers forspecified categories
// @Tags			Admin
// @Accept			json
// @Produce		    json
// @Param			offer	body	models.OfferMaking	true	"offer"
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/admin/offers/add [post]
func (off *OfferHandler) AddNewOffer(c *gin.Context) {
	var model models.OfferMaking
	if err := c.BindJSON(&model); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	err := off.usecase.AddNewOffer(model)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Could not add the Offer", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully added Offer", nil, nil)
	c.JSON(http.StatusOK, successRes)

}

func (o *OfferHandler) MakeOfferExpire(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	if err := o.usecase.MakeOfferExpire(id); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Coupon cannot be made invalid", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully made Coupon as invaid", nil, nil)
	c.JSON(http.StatusOK, successRes)

}
