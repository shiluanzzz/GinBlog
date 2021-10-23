package model

import (
	"GinBlog/utils"
	"GinBlog/utils/errmsg"
	"context"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"log"
	"mime/multipart"
)

var accessKey = utils.ACCESSKEY
var secretKey = utils.SK
var ImgServer = utils.OSSServer
var BucketName = utils.BucketName

func UploadFile(file multipart.File, fileSize int64) (string, int) {
	putPolicy := storage.PutPolicy{
		Scope: BucketName,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{
		Zone:          &storage.ZoneHuanan,
		UseCdnDomains: false,
		UseHTTPS:      false,
	}
	// 额外参数
	putExtra := storage.PutExtra{}
	formUploader := storage.NewFormUploader(&cfg)
	res := storage.PutRet{}

	// upload
	err := formUploader.PutWithoutKey(context.Background(), &res, upToken, file, fileSize, &putExtra)
	if err != nil {
		log.Println("oss upload error:", err)
		return "", errmsg.ERROR
	}
	url := ImgServer + res.Key
	return url, errmsg.SUCCESS

}
