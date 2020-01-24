package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func GetResult() {
	if len(result) > 0 {
		fmt.Printf("%d target detected Vulnerability\n", len(result))
		fmt.Println("----------------------------------------------------------------")
		for k, v := range result {
			fmt.Printf("%s    |    %s\n", k, v)
		}
		fmt.Println("----------------------------------------------------------------")
	} else {
		fmt.Println("----------------------------------------------------------------")
		fmt.Println("All target checked. none of them has Vulnerability")
		fmt.Println("----------------------------------------------------------------")
	}
}

func Getpoc() []Poc {
	pwd, _ := os.Getwd()
	//获取文件或目录相关信息
	fileInfoList, err := ioutil.ReadDir(pwd + "/poc")
	if err != nil {
		log.Fatal(err)
	}
	slice := []Poc{}
	for i := range fileInfoList {
		a := strings.ToUpper(PocPath)
		if a == "ALL" {
			var c Poc
			slice = append(slice, c.getConf(pwd+"/poc/"+fileInfoList[i].Name()))
		} else {
			if strings.Index(strings.ToUpper(fileInfoList[i].Name()), a) != -1 {
				var c Poc
				slice = append(slice, c.getConf(pwd+"/poc/"+fileInfoList[i].Name()))
			}

		}
	}
	if len(slice) == 0 {
		fmt.Println("-poc 参数设置错误")
		os.Exit(1)
	}
	return slice

}

func ReadAll(path string) []byte {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)

	}
	data, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)

	}
	return data
}

func Head() {
	fmt.Println(`
	.d888 888                   888                                          
	d88P"  888                   888                                          
	888    888                   888                                          
	888888 888  8888b.  .d8888b  88888b.  .d8888b   .d8888b  8888b.  88888b.  
	888    888     "88b 88K      888 "88b 88K      d88P"        "88b 888 "88b 
	888    888 .d888888 "Y8888b. 888  888 "Y8888b. 888      .d888888 888  888 
	888    888 888  888      X88 888  888      X88 Y88b.    888  888 888  888 
	888    888 "Y888888  88888P' 888  888  88888P'  "Y8888P "Y888888 888  888 																							  
	`)
}

func AppendToFile(fileName string, content string) error {
	// 以只写的模式，打开文件
	f, err := os.OpenFile(fileName, os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("cacheFileList.yml file create failed. err: " + err.Error())
	} else {
		// 查找文件末尾的偏移量
		n, _ := f.Seek(0, os.SEEK_END)
		// 从末尾的偏移量开始写入内容
		_, err = f.WriteAt([]byte(content), n)
	}
	defer f.Close()
	return err
}
