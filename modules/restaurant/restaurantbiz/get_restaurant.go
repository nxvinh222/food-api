package restaurantbiz

import (
	"context"
	"food-delivery/modules/restaurant/restaurantmodel"
)

type GetRestaurantStorage interface {
	FindDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*restaurantmodel.Restaurant, error)
}

type GetRestaurantBiz struct {
	store GetRestaurantStorage
}

func NewGetRestaurantBiz(store GetRestaurantStorage) *GetRestaurantBiz {
	return &GetRestaurantBiz{store: store}
}

func (biz *GetRestaurantBiz) GetRestaurant(ctx context.Context, id int) (*restaurantmodel.Restaurant, error) {
	data, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"id": id})

	return data, err
}