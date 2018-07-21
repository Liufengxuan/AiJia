package models

import (
	"github.com/weilaihui/fdfs_client"
	"github.com/astaxie/beego"
)

func UploadByFilename(filename string)(groupName,fileId string,err error){
	fdfsCLient,errClient:=fdfs_client.NewFdfsClient("conf/client.conf")//go文件的文件夹才是代码的目录起始位置
	if errClient!=nil {
		beego.Error("fdfs_client.NewFdfsClient err=", errClient.Error())
		return "","",errClient
	}
	uploadResponse,errUpload:=fdfsCLient.UploadByFilename(filename)
	if errUpload!=nil{
		beego.Info("fdfsCLient.UploadByFilename err= ",errUpload.Error())
		return "","",errUpload
	}

	beego.Info("groupName and FileId=",uploadResponse.RemoteFileId)
	return uploadResponse.GroupName,uploadResponse.RemoteFileId,err

}





