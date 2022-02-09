package component

import (
	"food-delivery/component/uploadprovider"
	"gorm.io/gorm"
)

type AppContext interface {
	GetMainDBConnection() *gorm.DB
	GetUploadProvider() uploadprovider.UploadProvider
}

type appCtx struct {
	db         *gorm.DB
	upProvider uploadprovider.UploadProvider
}

func NewAppContext(db *gorm.DB, upProvider uploadprovider.UploadProvider) *appCtx {
	return &appCtx{db: db, upProvider: upProvider}
}

func (ctx appCtx) GetMainDBConnection() *gorm.DB {
	return ctx.db
}

func (ctx appCtx) GetUploadProvider() uploadprovider.UploadProvider {
	return ctx.upProvider
}
