package bid_controller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joaolima7/leilao-goexpert/configuration/rest_err"
	"github.com/joaolima7/leilao-goexpert/internal/infra/api/web/validation"
	"github.com/joaolima7/leilao-goexpert/internal/infra/usecase/bid_usecase"
)

type BidController struct {
	bidUseCase bid_usecase.BidUseCaseInterface
}

func NewBidController(bidUseCase bid_usecase.BidUseCaseInterface) *BidController {
	return &BidController{
		bidUseCase: bidUseCase,
	}
}

func (u *BidController) CreateBid(c *gin.Context) {
	var bidInputDTO bid_usecase.BidInputDTO

	if err := c.ShouldBindJSON(&bidInputDTO); err != nil {
		restErr := validation.ValidateErr(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	err := u.bidUseCase.CreateBid(context.Background(), bidInputDTO)
	if err != nil {
		restErr := rest_err.ConvertError(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	c.Status(http.StatusCreated)
}
