package restaurantstorage

import (
	"context"
	"food-delivery/common"
	"food-delivery/modules/restaurant/restaurantmodel"
)

func (s *sqlStore) Create(ctx context.Context, data *restaurantmodel.RestaurantCreate) error {
	err := s.db.Create(data).Error
	if err != nil {
		return common.ErrDB(err)
	}

	return nil
}
