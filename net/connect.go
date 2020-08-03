package net

import (
	m "awesomeProject/lcffilemanager/manager"
	"fmt"
	"github.com/axgle/mahonia"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/url"
	"os"
	"path"
	"strings"
)


func InitConnection()  *gin.Engine{

	r := gin.Default()

	r.GET("/favicon.ico", func(c *gin.Context) {
		c.String(403,"!")
	})

	r.GET("/view/", func(c *gin.Context) {
		q:=c.Query("h")
		c.String(200,q)
		kai, _ := url.QueryUnescape(q)
		fmt.Print(kai)
	})

	r.GET("/index/", func(c *gin.Context) {
		h:=c.Query("h")
		//h,_=url.QueryUnescape(c.Request.RequestURI[10:])
		//fmt.Println(c.ClientIP())
		//fmt.Println(c.Request.RemoteAddr)//192.168.18.104:54783
		//fmt.Println(c.Request.Host)
		host:=c.Request.Host
		types,exist:=m.IsExistFileOrDir(h)
		//fmt.Println(m.ImgInSameDir(h))
		//fmt.Println(exist)
		if exist {
			if types {//dir
				list,err:=m.ListSingle(h)
				if err!=nil {
					fmt.Println(err)
				}
				pos:=pathPos(h)
				strs:=strSlicer(h,pos)
				c.HTML(200,"index.tmpl",gin.H{
					"href":h,
					"strs":strs,
					"host":host,
					"list":list,
				})
			}else {//file
				//c.String(200,"file desu.")
				ext:=path.Ext(strings.ToLower(h))
				//fmt.Println(ext)
				pos:=pathPos(h)
				strs:=strSlicer(h,pos)
				if ext==".txt"||ext==".html"{
					fmt.Println("txt")
					c.HTML(200,"txt.tmpl",gin.H{
						"strs":strs,
						"file":readTxt(h),
						"host":host,
					})
				}else if ext==".mp4"  {
					modTime,size,mode:=m.FileInfo(h)
					c.HTML(200,"video.tmpl",gin.H{
						"host":host,
						"strs":strs,
						"filename":path.Base(h),
						"modTime":modTime,
						"mode":mode,
						"size":size,
					})
				}else if ext==".jpg"||ext==".gif"||ext==".png" {
					modTime,size,mode:=m.FileInfo(h)
					c.HTML(200,"image.tmpl",gin.H{
						"host":host,
						"strs":strs,
						"filename":path.Base(h),
						"modTime":modTime,
						"mode":mode,
						"size":size,
						"lastDirImgs":m.ImgInSameDir(lastDirPath(h)),
					})
				}else {
					modTime,size,mode:=m.FileInfo(h)
					c.HTML(200,"download.tmpl",gin.H{
						"host":host,
						"strs":strs,
						"filename":path.Base(h),
						"modTime":modTime,
						"mode":mode,
						"size":size,
					})
				}
			}
		}else {//not found
			c.String(200,"not found")
		}
	})

	//r.GET("/index/", func(c *gin.Context) {
	//	c.Redirect(http.StatusMovedPermanently, "/index/?h=G://")
	//})

	r.GET("/search/", func(c *gin.Context) {
		p:=c.Query("p")
		q:=c.Query("q")
		host:=c.Request.Host
		list, _ :=m.List(p,q)
		c.HTML(200,"search.tmpl",gin.H{
			"query":q,
			"host":host,
			"list":list,
			"href":p,//搜索文件夹
		})
	})

	r.GET("/download/", func(c *gin.Context) {
		h:=c.Query("h")
		types,exist:=m.IsExistFileOrDir(h)
		if exist {
			if types {//dir
				c.String(403,"leave me plz.")
			}else {//file
				filename:=path.Base(h)
				c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
				c.Writer.Header().Add("Content-Type", "application/octet-stream")
				c.File(h)
			}

		}else {
			c.String(404,"not found.")
		}

	})

	return r
}

func pathPos(h string)[]int{
	length:=strings.Count(h,"/")
	var pos []int
	for i:=0;i<=length;i++ {
		if len(pos)==0{
			pos = append(pos, strings.Index(h[0:],"/"))
		}else if strings.Index(h[pos[i-1]+1:],"/")!=-1 {
			//fmt.Println(strings.Index(h[pos[i-1]+1:], "/"))
			pos = append(pos, strings.Index(h[pos[i-1]+1:], "/")+pos[i-1]+1)
		}
	}
	return pos
}

func lastDirPath(h string)string  {
	length:=strings.Count(h,"/")
	var pos []int
	for i:=0;i<=length;i++ {
		if len(pos)==0{
			pos = append(pos, strings.Index(h[0:],"/"))
		}else if strings.Index(h[pos[i-1]+1:],"/")!=-1 {
			pos = append(pos, strings.Index(h[pos[i-1]+1:], "/")+pos[i-1]+1)
		}
	}

	lastDir:=h[0:pos[len(pos)-1]]
	return lastDir
}

func strSlicer(str string,pos []int)[]string  {
	if pos[0]==-1 {
		return nil
	}
	var strs []string
	strs = append(strs,str[0:pos[0]])
	for i:=0;i<=len(pos)-2;i++ {
		strs = append(strs,str[pos[i]+1:pos[i+1]])
	}
	strs = append(strs,str[pos[len(pos)-1]+1:])
	return strs
}

func readTxt(h string)string{
	fi, err := os.Open(h)
	if err != nil {
		//fmt.Println("read file fail", err)
		return "fail1"
	}
	defer fi.Close()
	decoder := mahonia.NewDecoder("gbk") // 把原来ANSI格式的文本文件里的字符，用gbk进行解码。
	fd, err := ioutil.ReadAll(decoder.NewReader(fi))
	if err != nil {
		//fmt.Println("read to fd fail", err)
		return "fail2"
	}
	return string(fd)
}
