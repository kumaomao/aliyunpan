package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"yp/yunpan"
)

func MainView(c *gin.Context) {
	yp := yunpan.Yunpan
	folder := c.DefaultQuery("folder","root")
	//获取文件列表
	data := map[string]interface{}{
		"all" 							: false,
		"drive_id"						: yp.RefreshInfo.DefaultDriveId,
		"fields"						: "*",
		"image_thumbnail_process"		: "image/resize,w_400/format,jpeg",
		"image_url_process"				: "image/resize,w_1920/format,jpeg",
		"limit"							: 100,
		"order_by"						: "updated_at",
		"order_direction"				: "DESC",
		"parent_file_id"				: folder,
		"video_thumbnail_process"		: "video/snapshot,t_0,f_jpg,ar_auto,w_300",
	}
	yp.GetList(data)
	count := len(yp.DataItems.Item)

	info,_ := yp.GetDownloadUrl(folder)
	parent := info.ParentFileId

	ref := map[string]string{}
	if parent == "" || folder == "root"{
		ref = map[string]string{
			"title"  : "刷新",
			"url"	 : "/main/index",
		}
	}else{
		ref = map[string]string{
			"title"  : "返回上级目录",
			"url"	 : "/main/index?folder="+parent,
		}
		if parent == "root"{
			ref["url"] =  "/main/index"
		}
	}

	c.HTML(http.StatusOK, "main/index.html", gin.H{
		"title"		: "阿里云盘分享",
		"userInfo" 	: yp.RefreshInfo,
		"count" 	: count,
		"list"		: yp.DataItems.Item,
		"ref"		: ref,
	})
}

//文件下载
func Download(c *gin.Context){
	yp := yunpan.Yunpan
	field_id := c.Query("file")
	info,err := yp.GetDownloadUrl(field_id)
	if err != nil{
		return
	}
	//d := yp.GetFile(downUrl.Url)
	c.Redirect(302,info.DownloadUrl)
}

//打包下载
func MultiDownload(c *gin.Context)  {
	yp := yunpan.Yunpan
	field_id := c.Query("file")
	info,_ := yp.GetDownloadUrl(field_id)
	data := map[string]interface{}{
		"archive_name"      : info.Name,
		"download_infos"	: []map[string]interface{}{
			{
				"drive_id" 	: yp.RefreshInfo.DefaultDriveId,
				"files"		: []map[string]string{
					{
						"file_id"	: field_id,
					},
				},
			},
		},
	}

	down,err := yp.MultiDownloadUrl(data)
	if err != nil{
		return
	}
	//d := yp.GetFile(downUrl.Url)
	c.Redirect(302,down["download_url"].(string))
}

//预览界面
func Preview(c *gin.Context)  {
	yp := yunpan.Yunpan
	file := c.Query("file")
	info,_ :=yp.GetDownloadUrl(file)

	cate := info.Category
	fileInfo := map[string]interface{}{}
	view := "preview/preview_other.html"
	vedio_json := ""
	switch cate {
		case "image":
			view = "preview/preview_image.html"
			break
		case "audio":
			view = "preview/preview_music.html"
			fileInfo,_ = yp.GetAudioPlayInfo(file)
			break
		case "video":
			view = "preview/preview_video.html"
			fileInfo,_ = yp.GetVideoPlayInfo(file)
			//清晰度
			quality := map[string]string{
				"LD"		: "流畅",
				"SD"		: "标清",
				"HD"		: "高清",
				"FHD"		: "超清",
			}
			//资源切片
			urlList := []map[string]string{}
			urlList = append(urlList,map[string]string{
				"name"		: "原画",
				"url"		: info.Url,
				"type"		: "auto",
			})
			for _,v := range fileInfo["template_list"].([]interface{}){
				urlList = append(urlList,map[string]string{
					"name"		: quality[v.(map[string]interface{})["template_id"].(string)],
					"url"		: v.(map[string]interface{})["url"].(string),
					"type"		: "auto",
				})
			}
			vedio,_ :=json.Marshal(urlList)
			vedio_json = string(vedio)
			break
	}
	parent := info.ParentFileId

	ref := map[string]string{
		"title"  : "返回上级目录",
		"url"	 : "/main/index?folder="+parent,
	}

	c.HTML(http.StatusOK, view, gin.H{
		"title"		: info.Name+"阿里云盘预览",
		"info"		: info,
		"ref"		: ref,
		"fileInfo"	: fileInfo,
		"vedioJson" : vedio_json,
	})

}