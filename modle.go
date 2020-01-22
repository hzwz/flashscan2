package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

type Poc struct {
	Name   string  `yaml:"name"`
	Rule   Rules   `yaml:"rules"`
	Detail Details `yaml:"detail"`
}

type Rules struct {
	Method   string `yaml:"method"`
	Path     string `yaml:"path"`
	Headers  Header `yaml:"headers"`
	Body     string `yaml:"body"`
	Status   int    `yaml:"status"`
	Contains string `yaml:"contains"`
}
type Header struct {
	ContentType string `yaml:"contenttype`
	Cookie      string `yaml:cookie`
	UserAgent   string `yaml:useragent`
}
type Details struct {
	Link   string `yaml:link`
	Author string `yaml:author`
}

func (c *Poc) getConf(path string) Poc {
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	bs := Poc{}

	err = yaml.Unmarshal(yamlFile, &bs)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return bs
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
