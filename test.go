package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func main()  {
	ListAll("C:\\Users\\29673\\go\\src\\awesomeProject")
}
func ListAll(dir string)  {
	dlist,flist, err := list(dir,"2")
	if err != nil {
		fmt.Println(err)
		return
	}
	//fmt.Println(dlist,flist)
	for _, v := range flist {
		fmt.Println(v)
	}
	for _, w := range dlist {
		fmt.Println(w)
	}
}

func list(dirpath ,query string) ([]string,[]string, error) {
	var dirList []string
	var fileList []string
	dirErr := filepath.Walk(dirpath,
		func(ppath string, f os.FileInfo, err error) error {//ppath=path,区分path包
			if f == nil {
				return err
			}else {
				var name =path.Base(f.Name())
				//fmt.Println(name)
				if strings.Contains(name,query) {
					if f.IsDir() {
						dirList = append(dirList, ppath)
						return nil
					}else
					if !f.IsDir() {
						fileList = append(fileList,ppath)
					}
				}
				return nil
			}
		})
	return dirList,fileList, dirErr
}