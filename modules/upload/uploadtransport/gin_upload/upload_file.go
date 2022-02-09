package gin_upload

import (
	"food-delivery/common"
	"food-delivery/component"
	"food-delivery/modules/upload/uploadbusiness"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Upload(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		fileHeader, err := c.FormFile("file")

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		folder := c.DefaultPostForm("folder", "img")

		file, err := fileHeader.Open()

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		defer file.Close()

		dataBytes := make([]byte, fileHeader.Size)

		_, err = file.Read(dataBytes)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		biz := uploadbusiness.NewUploadBiz(appCtx.GetUploadProvider(), nil)
		biz.Upload(c.Request.Context(), dataBytes, folder, fileHeader.Filename)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
