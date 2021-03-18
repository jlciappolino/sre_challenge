package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/pkg/errors"

	"github.com/gin-gonic/gin"
)

func suscribe() {
}

func main() {

	suscribe()
	r := gin.Default()

	r.POST("/sales/consumer", func(c *gin.Context) {
		d, _ := c.GetRawData()
		fmt.Printf("[consumer] recibido: %s\n", d)

		c.JSON(http.StatusOK, fmt.Sprintf("%s", d))
		return
	})

	r.GET("/sales/:id", func(c *gin.Context) {
		user_id := c.Param("id")
		if user_id == "" {
			c.JSON(http.StatusBadRequest, gin.H{"message": "users id not valid"})
			return
		}

		result, err := getUserData(user_id)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, errors.Wrap(err, "unexpected response"))
			return
		}

		c.JSON(http.StatusOK, result)
		return
	})

	r.Run()
}

func getUserData(user_id string) (Result, error) {
	var result Result

	var err error
	userResponse, err := get("http://front/users/" + user_id)
	if err != nil {
		return result, err
	}

	fmt.Printf("[api:sales] Response from users:\n %s\n", userResponse)
	soldItemsResponse, err := get("http://front/sold_items/" + user_id)
	if err != nil {
		return result, err
	}
	fmt.Printf("[api:sales] Response from sold_items:\n %s\n", soldItemsResponse)

	soldItems := []SoldItem{}
	if err := json.Unmarshal(soldItemsResponse, &soldItems); err != nil {
		return result, err
	}

	items := []Item{}
	for _, soldItem := range soldItems {
		itemResponse, err := get("http://front/items/" + soldItem.ID)
		if err != nil {
			return result, err
		}
		item := Item{}
		if err := json.Unmarshal(itemResponse, &item); err != nil {
			return result, err
		}
		items = append(items, item)
	}

	var sum float64 = 0
	var sumUSD float64 = 0
	for _, item := range items {
		currencyResponse, err := get("http://front/currency_conversions/" + item.Currency)
		if err != nil {
			return result, errors.Wrap(err, "fail currency_conversions api")
		}
		currencyValue, err := strconv.ParseFloat(fmt.Sprintf("%s", currencyResponse), 64)
		if err != nil {
			return result, errors.Wrap(err, "unexpectd value from currency_conversions api")
		}

		sumUSD = sumUSD + item.Price
		sum = sum + (item.Price * currencyValue)
	}

	result = Result{
		User_id:   user_id,
		Total:     sum,
		Total_usd: sumUSD,
	}

	return result, nil

}
