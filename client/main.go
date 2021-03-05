package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func main() {

	r := gin.Default()

	r.GET("/client/:id", func(c *gin.Context) {
		var err error
		userResponse, err := get("http://front/users/" + c.Param("id"))
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, errors.Wrap(err, "users api fail"))
			return
		}

		fmt.Printf("[api:client] Response from users:\n %s\n", userResponse)
		soldItemsResponse, err := get("http://front/sold_items/" + c.Param("id"))
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, errors.Wrap(err, "solds_items api fail"))
			return
		}
		fmt.Printf("[api:client] Response from users:\n %s\n", soldItemsResponse)

		soldItems := []SoldItem{}
		if err := json.Unmarshal(soldItemsResponse, &soldItems); err != nil {
			c.AbortWithError(http.StatusInternalServerError, errors.Wrap(err, "unexpected sold_items response"))
			return
		}

		items := []Item{}
		for _, soldItem := range soldItems {
			itemResponse, err := get("http://front/items/" + soldItem.ID)
			if err != nil {
				c.AbortWithError(http.StatusInternalServerError, errors.Wrap(err, "fail item api"))
				return
			}
			item := Item{}
			if err := json.Unmarshal(itemResponse, &item); err != nil {
				c.AbortWithError(http.StatusInternalServerError, errors.Wrap(err, "unexpected item response"))
				return
			}

			items = append(items, item)
		}

		var sum float64 = 0
		var sumUSD float64 = 0
		for _, item := range items {
			currencyResponse, err := get("http://front/currency_conversions/" + item.Currency)
			if err != nil {
				c.AbortWithError(http.StatusInternalServerError, errors.Wrap(err, "fail currency_conversions api"))
				return
			}
			currencyValue, err := strconv.ParseFloat(fmt.Sprintf("%s", currencyResponse), 64)
			if err != nil {
				c.AbortWithError(http.StatusInternalServerError, errors.Wrap(err, "unexpectd value from currency_conversions api"))
				return
			}

			sumUSD = sumUSD + item.Price
			sum = sum + (item.Price * currencyValue)
		}

		c.JSON(http.StatusOK, gin.H{
			"user":      c.Param("id"),
			"total":     sum,
			"total_usd": sumUSD,
		})
	})

	r.Run()
}
