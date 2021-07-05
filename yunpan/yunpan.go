package yunpan

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
	"yp/config"
)

type Yp struct {
	UserInfo 			UserInfo				`json:"user_info"`
	RefreshInfo			refreshInfo				`json:"refresh_info"`
	DataItems			DataItems				`json:"items"`
	DownloadUrl			DataItem				`json:"downloadUrl"`
}

//用户信息
type UserInfo struct {
	Avatar 				string 					`json:"avatar"`
	CreatedAt 			string					`json:"created_at"`
	DefaultDriveId 		string					`json:"default_drive_id"`
	Description 		string					`json:"description"`
	DomainId			string					`json:"domain_id"`
	Email				string					`json:"email"`
	NickName			string					`json:"nick_name"`
	UserId				string					`json:"user_id"`
	UserName			string					`json:"user_name"`
}

//刷新信息
type refreshInfo struct {
	AccessToken 		string					`json:"access_token"`
	Avatar 				string 					`json:"avatar"`
	DefaultDriveId 		string					`json:"default_drive_id"`
	DeviceId			string					`json:"device_id"`
	ExistLink			[]string				`json:"exist_link"`
	ExpireTime			string					`json:"expire_time"`
	ExpiresIn			int						`json:"expires_in"`
	IsFirstLogin		bool					`json:"is_first_login"`
	NeedLink			bool					`json:"need_link"`
	NeedRpVerify		bool					`json:"need_rp_verify"`
	NickName			string					`json:"nick_name"`
	PinSetup			bool					`json:"pin_setup"`
	RefreshToken		string					`json:"refresh_token"`
	Role				string					`json:"role"`
	State				string					`json:"state"`
	Status				string					`json:"status"`
	TokenType			string					`json:"token_type"`
	UserId				string					`json:"user_id"`
	UserName			string					`json:"user_name"`
}

//暂时不用
type RefreshUserData struct {

}

//列表
type DataItems struct {
	Item				[]DataItem				`json:"items"`
	NextMarker			string					`json:"next_marker"`
}

//列表详情
type DataItem struct {
	Category			string 					`json:"category"`
	ContentHash			string 					`json:"content_hash"`
	ContentHashName		string 					`json:"content_hash_name"`
	ContentType			string 					`json:"content_type"`
	Crc64Hash			string 					`json:"crc64_hash"`
	CreatedAt			string 					`json:"created_at"`
	DomainId			string 					`json:"domain_id"`
	DownloadUrl			string 					`json:"download_url"`
	DeviceId			string					`json:"device_id"`
	EncryptMode			string					`json:"encrypt_mode"`
	FileExtension		string					`json:"file_extension"`
	FileId				string					`json:"file_id"`
	Hidden				bool					`json:"hidden"`
	Name				string					`json:"name"`
	ParentFileId		string					`json:"parent_file_id"`
	PunishFlag			int						`json:"punish_flag"`
	Size				int						`json:"size"`
	Starred				bool					`json:"starred"`
	Status				string					`json:"status"`
	Type				string					`json:"type"`
	UpdatedAt			string					`json:"updated_at"`
	UploadId			string					`json:"upload_id"`
	Url					string					`json:"url"`
	Thumbnail			string					`json:"thumbnail"`
	VideoPreviewMetadata map[string]interface{} `json:"video_preview_metadata"`
}



var Yunpan = new(Yp)

func init(){
	Yunpan.Refresh()
	Yunpan.Heartbeat()
}

//身份信息刷新
func (y *Yp) Refresh()  {
	url := "https://auth.aliyundrive.com/v2/account/token"

	refresh_token := config.Conf.RefreshToken
	data := map[string]string{
		"grant_type"	: "refresh_token",
		"refresh_token" : refresh_token,
	}
	data_json,_ :=json.Marshal(data)
	header := map[string]string{
		"Content-Type"		:"application/json",
		"origin"			:"https://www.aliyundrive.com",
		"referer"			:"https://www.aliyundrive.com",
	}

	respByte, _ := y.curl(url,"POST",string(data_json),header)
	err := json.Unmarshal([]byte(respByte), &y.RefreshInfo)
	if err != nil{
		fmt.Println(err)
	}
	//写入新的刷新refresh_token
	config.Conf.RefreshToken = y.RefreshInfo.RefreshToken

}

//获取文件列表
func (y *Yp) GetList(data map[string]interface{}) (DataItems,error) {
	url := "https://api.aliyundrive.com/v2/file/list"
	data_json,_ :=json.Marshal(data)
	header := map[string]string{
		"Content-Type"		:"application/json;charset=UTF-8",
		"origin"			:"https://www.aliyundrive.com",
		"referer"			:"https://www.aliyundrive.com",
		"authorization"		:y.RefreshInfo.TokenType+" "+y.RefreshInfo.AccessToken,
	}

	respByte, _ := y.curl(url,"POST",string(data_json),header)
	list := y.DataItems
	err := json.Unmarshal([]byte(respByte), &list)
	return list,err
}

//获取下载地址
func (y *Yp) GetDownloadUrl(file_id string) (DataItem,error) {
	url := "https://api.aliyundrive.com/v2/file/get"
	data := map[string]interface{}{
		"drive_id"						: y.RefreshInfo.DefaultDriveId,
		"file_id"						: file_id,
		"image_thumbnail_process"		: "image/resize,w_400/format,jpeg",
		"fields"						: "*",
		"image_url_process"				: "image/resize,w_1920/format,jpeg",
		"order_by"						: "updated_at",
		"order_direction"				: "DESC",
		"video_thumbnail_process"		: "video/snapshot,t_0,f_jpg,ar_auto,w_300",
	}
	data_json,_ :=json.Marshal(data)
	header := map[string]string{
		"Content-Type"		:"application/json",
		"origin"			:"https://www.aliyundrive.com",
		"referer"			:"https://www.aliyundrive.com",
		"authorization"		:y.RefreshInfo.TokenType+" "+y.RefreshInfo.AccessToken,
	}
	respByte, _ := y.curl(url,"POST",string(data_json),header)
	res := DataItem{}
	err := json.Unmarshal([]byte(respByte),&res)
	return res,err
}

//获取音频详情
func (y *Yp)GetAudioPlayInfo(file_id string) (map[string]interface{},error){
	url := "https://api.aliyundrive.com/v2/databox/get_audio_play_info"
	data := map[string]interface{}{
		"drive_id"						: y.RefreshInfo.DefaultDriveId,
		"file_id"						: file_id,
	}
	data_json,_ :=json.Marshal(data)
	header := map[string]string{
		"Content-Type"		:"application/json",
		"origin"			:"https://www.aliyundrive.com",
		"referer"			:"https://www.aliyundrive.com",
		"authorization"		:y.RefreshInfo.TokenType+" "+y.RefreshInfo.AccessToken,
	}
	respByte, _ := y.curl(url,"POST",string(data_json),header)
	info := map[string]interface{}{}
	err := json.Unmarshal([]byte(respByte), &info)
	return info,err
}

//获取视频详情
func (y *Yp)GetVideoPlayInfo(file_id string) (map[string]interface{},error){
	url := "https://api.aliyundrive.com/v2/databox/get_video_play_info"
	data := map[string]interface{}{
		"drive_id"						: y.RefreshInfo.DefaultDriveId,
		"file_id"						: file_id,
	}
	data_json,_ :=json.Marshal(data)
	header := map[string]string{
		"Content-Type"		:"application/json",
		"origin"			:"https://www.aliyundrive.com",
		"referer"			:"https://www.aliyundrive.com",
		"authorization"		:y.RefreshInfo.TokenType+" "+y.RefreshInfo.AccessToken,
	}
	respByte, _ := y.curl(url,"POST",string(data_json),header)
	info := map[string]interface{}{}
	err := json.Unmarshal([]byte(respByte), &info)
	return info,err
}

//打包下载
func (y *Yp)MultiDownloadUrl(data map[string]interface{}) (map[string]interface{},error) {
	url := "https://api.aliyundrive.com/adrive/v1/file/multiDownloadUrl"
	data_json,_ :=json.Marshal(data)
	header := map[string]string{
		"Content-Type"		:"application/json",
		"origin"			:"https://www.aliyundrive.com",
		"referer"			:"https://www.aliyundrive.com",
		"authorization"		:y.RefreshInfo.TokenType+" "+y.RefreshInfo.AccessToken,
	}
	respByte, _ := y.curl(url,"POST",string(data_json),header)
	info := map[string]interface{}{}
	err := json.Unmarshal([]byte(respByte), &info)
	return info,err
}


//封装curl
func (y *Yp)curl(url string,options ...interface{}) ([]byte,error) {
	//options -》 method string,data string,hearder map[string]string
	//获取访问方法
	method := "GET"
	if options[0] != nil{
		method = options[0].(string)
	}
	//获取参数
	data := ""
	if options[1] != nil{
		data = options[1].(string)
	}
	//获取头
	header := map[string]string{}
	if options[2] != nil{
		header = options[2].(map[string]string)
	}
	req, _ := http.NewRequest(method,url,strings.NewReader(data))
	//设置请求头
	for key,value := range header{
		req.Header.Set(key,value)
	}
	resp, err := (&http.Client{}).Do(req)
	if(err != nil){
		return nil,err
	}
	defer resp.Body.Close()
	result ,err := ioutil.ReadAll(resp.Body)
	if(err != nil){
		return nil,err
	}
	return result,nil
}




//心跳包
func (y *Yp) Heartbeat()  {
	//var ch chan int
	//定时任务
	ticker := time.NewTicker(time.Second * 6500)
	go func() {
		for range ticker.C {
			fmt.Println("心跳启动")
			//执行
			y.Refresh()
		}

	}()

}
