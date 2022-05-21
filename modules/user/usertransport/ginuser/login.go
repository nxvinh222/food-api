package ginuser

import (
	"food-delivery/common"
	"food-delivery/component"
	"food-delivery/component/hasher"
	"food-delivery/component/tokenprovider/jwt"
	"food-delivery/modules/user/userbiz"
	"food-delivery/modules/user/usermodel"
	"food-delivery/modules/user/userstorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(ctx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var userLoginData usermodel.UserLogin

		err := c.ShouldBind(&userLoginData)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := userstorage.NewSQLStore(ctx.GetMainDBConnection())
		tProvider := jwt.NewJwtProvider(ctx.GetSecretKey())
		md5 := hasher.NewMd5Hash()

		biz := userbiz.NewLoginBusiness(store, tProvider, md5, 60*60*24*30)
		account, err := biz.Login(c.Request.Context(), &userLoginData)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(account))
	}
}
