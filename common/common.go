package common

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"time"
	"fmt"
	"yp/yunpan"
)



//设置模板方法
func TempLateCommon(r *gin.Engine){
	r.SetFuncMap(template.FuncMap{
		"GetPreviewImg"		: GetPreviewImg,
		"Date"				: Date,
		"FileSize"			: FileSize,
	})
}


//计算文件大小
func FileSize(fileSize int) string{
	if fileSize == 0{
		return "-"
	}
	if fileSize < 1024 {
		//return strconv.FormatInt(fileSize, 10) + "B"
		return fmt.Sprintf("%.2fB", float64(fileSize)/float64(1))
	} else if fileSize < (1024 * 1024) {
		return fmt.Sprintf("%.2fKB", float64(fileSize)/float64(1024))
	} else if fileSize < (1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fMB", float64(fileSize)/float64(1024*1024))
	} else if fileSize < (1024 * 1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fGB", float64(fileSize)/float64(1024*1024*1024))
	} else if fileSize < (1024 * 1024 * 1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fTB", float64(fileSize)/float64(1024*1024*1024*1024))
	} else { //if fileSize < (1024 * 1024 * 1024 * 1024 * 1024 * 1024)
		return fmt.Sprintf("%.2fEB", float64(fileSize)/float64(1024*1024*1024*1024*1024))
	}

}

//时间格式化
func Date(timeString string,templateNum int) string{
	templateT := [5]string{"2006-01-02 15:04:05","2006/01/02 15:04:05","2006-01-02","20060102","15:04:05"}
	to, _ := time.Parse("2006-01-02T15:04:05Z", timeString)
	stamp := to.Format(templateT[templateNum])
	return stamp
}

//获取对应的预览图
func GetPreviewImg(info yunpan.DataItem) string {
	if info.Type == "folder"{
		return "https://img.alicdn.com/imgextra/i3/O1CN01qSxjg71RMTCxOfTdi_!!6000000002097-2-tps-80-80.png"
	}
	if info.Type == "file"{
		if info.Thumbnail !=""{
			//存在预览图直接返回预览
			return info.Thumbnail
		}

		//压缩包图片
		if info.FileExtension == "zip" || info.FileExtension == "7z" || info.FileExtension == "rar" || info.FileExtension == "cab" || info.FileExtension == "arj" || info.FileExtension == "lzh" || info.FileExtension == "ace" || info.FileExtension == "7-zip" || info.FileExtension == "tar" || info.FileExtension == "gzip" || info.FileExtension == "uue" || info.FileExtension == "bz2" || info.FileExtension == "jar" || info.FileExtension == "iso" || info.FileExtension == "z"{
			return "https://img.alicdn.com/imgextra/i4/O1CN01nkRtEq1UWx7RA6wAg_!!6000000002526-2-tps-80-80.png"
		}
		if  info.FileExtension == "jpg"  || info.FileExtension == "bmp"  || info.FileExtension == "jpeg"  || info.FileExtension == "png"  || info.FileExtension == "ico"  || info.FileExtension == "gif"  || info.FileExtension == "psd"  || info.FileExtension == "psb"  || info.FileExtension == "tiff"  || info.FileExtension == "ai"  || info.FileExtension == "eps"{
			//图片
			return "https://img.alicdn.com/imgextra/i2/O1CN01nTpqfa1r3Oix8W5gQ_!!6000000005575-2-tps-80-80.png";
		}
		if  info.FileExtension == "mp3"  || info.FileExtension == "flac"  || info.FileExtension == "wav"  || info.FileExtension == "cda"  || info.FileExtension == "ape"  {
			//音频文件
			return "https://img.alicdn.com/imgextra/i2/O1CN016MMvV325VhpDSUyrK_!!6000000007532-2-tps-80-80.png"
		}
		if info.FileExtension == "mp4"  || info.FileExtension == "wmv"  || info.FileExtension == "rmk"  || info.FileExtension == "mkv" {
			//视频文件
			return "https://img.alicdn.com/imgextra/i2/O1CN01H7FCkb1P6mPJxDEFa_!!6000000001792-2-tps-80-80.png"
		}
		if info.FileExtension == "txt"{
			return "https://img.alicdn.com/imgextra/i2/O1CN01kHskgT2ACzipXL4Ra_!!6000000008168-2-tps-80-80.png"
		}
		if info.FileExtension == "apk"{
			return "https://img.alicdn.com/imgextra/i1/O1CN01c7Eyle1yNHE5orvXh_!!6000000006566-2-tps-80-80.png"
		}
		if info.FileExtension == "doc"{
			return "https://img.alicdn.com/imgextra/i3/O1CN01lsQ3Re1dD6UgfjcSf_!!6000000003701-2-tps-80-80.png"
		}
		if info.FileExtension == "pdf"{
			return "https://img.alicdn.com/imgextra/i2/O1CN01RFXLvR1z5FFSCtDy9_!!6000000006662-2-tps-80-80.png"
		}
		if info.FileExtension == "eot" ||  info.FileExtension == "otf" ||  info.FileExtension == "fon" ||  info.FileExtension == "font" ||  info.FileExtension == "ttf" ||  info.FileExtension == "woff"{
			return "https://img.alicdn.com/imgextra/i3/O1CN018X1UHQ29FudCDSfDu_!!6000000008039-2-tps-80-80.png"
		}
		if info.FileExtension == "json"{
			return "https://img.alicdn.com/imgextra/i4/O1CN010YFhZp1s1r0KikhUQ_!!6000000005707-2-tps-80-80.png"
		}
		if info.FileExtension == "js"{
			return "https://img.alicdn.com/imgextra/i4/O1CN01bHoRs41okxY3ZdYcf_!!6000000005264-2-tps-80-80.png"
		}
		if info.FileExtension == "css"{
			return "https://img.alicdn.com/imgextra/i3/O1CN01voAknR1i1GKvaV1kk_!!6000000004352-2-tps-80-80.png"
		}
		if info.FileExtension == "html"{
			return "https://img.alicdn.com/imgextra/i1/O1CN012j3pSY1Kl17oTVZTa_!!6000000001203-2-tps-80-80.png"
		}
	}
	return "https://img.alicdn.com/imgextra/i1/O1CN01mhaPJ21R0UC8s9oik_!!6000000002049-2-tps-80-80.png"
}
