/*
 * @Author: your name
 * @Date: 2020-01-21 10:46:49
 * @LastEditTime : 2020-01-24 22:02:14
 * @LastEditors  : Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /flashscan/modle.go
 */
package main

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Poc struct {
	Name   string  `yaml:"name"`
	Rule   Rules   `yaml:"rules"`
	Detail Details `yaml:"detail"`
}

type Rules struct {
	Method   string  `yaml:"method"`
	Path     string  `yaml:"path"`
	Header   Headers `yaml:"headers"`
	Body     string  `yaml:"body"`
	Status   int     `yaml:"status"`
	Contains string  `yaml:"contains"`
}
type Headers map[string]string

/*
type Header struct {
	ContentType string `yaml:"contenttype`
	Cookie      string `yaml:cookie`
	UserAgent   string `yaml:useragent`
}
*/
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
