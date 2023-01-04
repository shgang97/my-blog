package qiniu

import (
	"backend/common"
	"backend/config"
	"github.com/gin-gonic/gin"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

func GetToken(ctx *gin.Context) {
	// 自定义凭证有效期（示例2小时，Expires 单位为秒，为上传凭证的有效时间）
	bucket := "blog-img-p1"
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	putPolicy.Expires = 7200 //示例2小时有效期
	mac := qbox.NewMac(config.Config.Qiniu.QiniuAccessKey, config.Config.Qiniu.QiniuSecretKey)
	upToken := putPolicy.UploadToken(mac)
	common.Success(ctx, upToken)
}
