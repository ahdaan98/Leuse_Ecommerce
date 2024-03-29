package handler

import (
	interfaces "github.com/ahdaan98/pkg/usecase/interface"
	"github.com/ahdaan98/pkg/utils/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CartHandler struct {
	usecase interfaces.CartUseCase
}

func NewCartHandler(usecase interfaces.CartUseCase) *CartHandler {
	return &CartHandler{
		usecase: usecase,
	}
}

func (i *CartHandler) AddToCart(c *gin.Context) {
	idString, _ := c.Get("id")
	UserID, _ := idString.(int)

	InventoryID, err := strconv.Atoi(c.Query("inventory_id"))
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Inventory Id not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	qty, err := strconv.Atoi(c.Query("quantity"))
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "check parameters properly", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	if err := i.usecase.AddToCart(UserID, InventoryID, qty); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Could not add to the cart", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully added to cart", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

func (i *CartHandler) CheckOut(c *gin.Context) {
	idString, _ := c.Get("id")
	id, _ := idString.(int)

	products, err := i.usecase.CheckOut(id)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not open checkout", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "Successfully got all records", products, nil)
	c.JSON(http.StatusOK, successRes)
}