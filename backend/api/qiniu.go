package api

import (
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"my-blog/backend/common"
	"my-blog/backend/config"
	"net/http"
)

/*
@author: shg
@since: 2022/10/21
@desc: //TODO
*/

func (*Api) GetQiniuToken(w http.ResponseWriter, r *http.Request) {
	// 自定义凭证有效期（示例2小时，Expires 单位为秒，为上传凭证的有效时间）
	bucket := "qiniu-oss-blog-img"
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	putPolicy.Expires = 7200 //示例2小时有效期
	mac := qbox.NewMac(config.Config.System.QiniuAccessKey, config.Config.System.QiniuSecretKey)
	upToken := putPolicy.UploadToken(mac)
	common.Success(w, upToken)
}
