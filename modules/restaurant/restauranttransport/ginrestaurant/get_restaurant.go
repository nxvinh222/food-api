package ginrestaurant

import (
	"food-delivery/common"
	"food-delivery/component"
	"food-delivery/modules/restaurant/restaurantbiz"
	"food-delivery/modules/restaurant/restaurantstorage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetRestaurant(ctx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid id",
			})

			return
		}

		store := restaurantstorage.NewSQLStore(ctx.GetMainDBConnection())
		biz := restaurantbiz.NewGetRestaurantBiz(store)

		result, err := biz.GetRestaurant(c.Request.Context(), id)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))

	}
}
