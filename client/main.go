package main

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var cacheItems = map[string]Item{}
var cacheUsers = map[string]Result{}

func init() {
	for i:=0;i<=10;i++{
		result, _ := getUserData(strconv.Itoa(i))
		cacheUsers[strconv.Itoa(i)] = result
	}
	fmt.Printf("[api:client] =========================== START CACHE OK ===========================\n",)
}

func main() {

	r := gin.Default()

	r.GET("/client/:id", func(c *gin.Context) {
		user_id := c.Param("id")
		if user_id == "" {
			c.JSON(http.StatusBadRequest, gin.H{"message": "users id not valid"})
			return
		}

		result, err := getUserData(user_id)
		if err != nil{
			c.AbortWithError(http.StatusInternalServerError, errors.Wrap(err, "unexpected response"))
			return
		}

		cacheUsers[user_id] = result

		c.JSON(http.StatusOK, result)
		return
	})

	r.Run()
}

func getUserData(user_id string) (Result, error) {
	var result Result
	var exists bool
	if result, exists = cacheUsers[user_id]; !exists {
		fmt.Printf("[api:client] NOT CACHE : user : %s exists \n",user_id)
		var err error
		userResponse, err := get("http://front/users/" + user_id)
		if err != nil {
			return result, err
		}

		fmt.Printf("[api:client] Response from users:\n %s\n", userResponse)
		soldItemsResponse, err := get("http://front/sold_items/" + user_id)
		if err != nil {
			return result, err
		}
		fmt.Printf("[api:client] Response from sold_items:\n %s\n", soldItemsResponse)

		soldItems := []SoldItem{}
		if err := json.Unmarshal(soldItemsResponse, &soldItems); err != nil {
			return result, err
		}

		items := []Item{}
		for _, soldItem := range soldItems {
			if item, exists := cacheItems[soldItem.ID]; !exists {
				fmt.Printf("[api:client] CACHE : items : %s doesnt exists \n",soldItem.ID)

				itemResponse, err := get("http://front/items/" + soldItem.ID)
				if err != nil {
					return result, err
				}
				item := Item{}
				if err := json.Unmarshal(itemResponse, &item); err != nil {
					return result, err
				}
				items = append(items, item)
				cacheItems[soldItem.ID] = item
			} else {
				fmt.Printf("[api:client] CACHE : items : %s exists \n",user_id)
				items = append(items, item)
			}
		}

		var sum float64 = 0
		var sumUSD float64 = 0
		for _, item := range items {
			// currencyResponse, err := get("http://front/currency_conversions/" + item.Currency)
			// if err != nil {
			// 	c.AbortWithError(http.StatusInternalServerError, errors.Wrap(err, "fail currency_conversions api"))
			// 	return
			// }
			// currencyValue, err := strconv.ParseFloat(fmt.Sprintf("%s", currencyResponse), 64)
			// if err != nil {
			// 	c.AbortWithError(http.StatusInternalServerError, errors.Wrap(err, "unexpectd value from currency_conversions api"))
			// 	return
			// }

			sumUSD = sumUSD + item.Price
			sum = sum + (item.Price * 155.0)
		}

		result = Result{
			User_id:   user_id,
			Total:     sum,
			Total_usd: sumUSD,
		}

		return result, nil
	}else {
		fmt.Printf("[api:client] CACHE : user : %s exists \n",user_id)
		return result, nil
	}
}