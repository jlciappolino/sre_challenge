package internal

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type currencyController struct{}

var conversions = map[string]float64{"usd": 155.00}

func NewCurrencyController() currencyController{
	return currencyController{}
}

func (c currencyController) Get(g *gin.Context){
	id := g.Param("id")
	value, err := getValue(	id)
	if err!=nil{
		g.JSON(http.StatusBadRequest,gin.H{"error": err.Error()})
		return
	}
	g.JSON(200, value)
	return
}

func getValue(simbol string) (float64, error){
	if value,exists:= conversions[simbol]; exists{
		return value, nil
	}
	return 0, errors.New("currency doesnt exists")
}
