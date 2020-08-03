package manager

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
)
// 全部，包括子路径下文件



type Name struct {
	Num int
	Name string
	IsDir bool
}
type Lists struct {
	Num int
	Href string
	Name string
	IsDir bool
}
type Img struct{
	Num int
	Name string
}

func ListSingle(pwd string)(result []Name,errs error) {
	var Names []Name
	//获取文件或目录相关信息
	fileInfoList,_ := ioutil.ReadDir(pwd)
	//fmt.Println(len(fileInfoList))

	for i := range fileInfoList {
		Names = append(Names,Name{
			Num: i,
			Name: fileInfoList[i].Name(),
			IsDir: fileInfoList[i].IsDir(),
		})//打印当前文件或目录下的文件或目录名
	}
	//fmt.Println(fileInfoList[0].Name())
	return Names,nil
}

func FileInfo(pwd string)(modTime string ,size int64,mode os.FileMode)  {
	fileInfo,err:=os.Stat(pwd)
	if err!=nil {
		return
	}
	modTime = fileInfo.ModTime().Format("2006-01-02 15:04:05")
	size = fileInfo.Size()
	mode = fileInfo.Mode() //权限
	return modTime,size,mode
}

//判断文件(夹)是否存在
func IsExistFileOrDir (path string)(types bool,exist bool)  {
	s, err := os.Stat(path)
	if err != nil {
		return false,false
	}else {
		return s.IsDir(),true
	}
}

func ImgInSameDir(pwd string)[]Img{
	var Imgs []Img
	fileInfoList,err := ioutil.ReadDir(pwd)
	if err!=nil {
		fmt.Println(err)
	}
	for i := range fileInfoList {
		h:=pwd+"/"+fileInfoList[i].Name()
		//fmt.Println(h)
		types,exist:=IsExistFileOrDir(h)
		//fmt.Println(exist,types)
		if exist&&!types {
			var a =strings.ToLower(path.Ext(h))
			//fmt.Println(a)
			if a==".jpg"||a==".png"||a==".gif" {
				Imgs = append(Imgs,Img{
					Num: i,
					Name: fileInfoList[i].Name(),
				})
			}
		}

	}
	fmt.Println(Imgs)
	return Imgs
}

func List(dirpath ,query string) ([]Lists, error) {
	var list []Lists
	err := filepath.Walk(dirpath,
		func(ppath string, f os.FileInfo, err error) error { //ppath=path,区分path包
			if f == nil {
				return err
			} else {
				var name = path.Base(f.Name())
				if strings.Contains(name, query) {
					list =append(list,Lists{
						Num: len(list),
						Name: name,
						Href: ppath,
						IsDir: f.IsDir(),
					})
					return nil
				}
				return nil
			}
		})
	return list,err
}