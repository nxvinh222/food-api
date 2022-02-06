package restaurantstorage

import (
	"context"
	"food-delivery/modules/restaurant/restaurantmodel"
)

func (s *sqlStore) FindDataByCondition(
	ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string,
) (*restaurantmodel.Restaurant, error) {
	var result restaurantmodel.Restaurant

	db := s.db

	for i := range moreKeys {
		db.Preload(moreKeys[i])
	}

	err := db.Where(conditions).First(&result).Error
	if err != nil {
		return nil, err
	}

	return &result, nil
}
