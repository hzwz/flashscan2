/*
 * @Author: wsm
 * @Date: 2020-01-19 15:26:26
 * @LastEditTime : 2020-02-12 11:08:22
 * @LastEditors  : Please set LastEditors
 * @Description:
 * @FilePath: /flashscan/main.go
 */
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"time"
)

var (
	h          bool
	HttpMethod string
	IpFile     string
	Port       int
	ThreadNum  int
	PocPath    string
	Output     string
)
var result map[string]string = map[string]string{}

func init() {
	flag.BoolVar(&h, "h", false, "Help")
	flag.StringVar(&HttpMethod, "m", "", "Http method,http or https")
	flag.StringVar(&IpFile, "f", "", "The file of the target")
	flag.IntVar(&Port, "p", 80, "The Port of target")
	flag.IntVar(&ThreadNum, "t", 100, "The num of threads")
	flag.StringVar(&PocPath, "poc", "", "The file of the target")
	flag.StringVar(&Output, "o", "", "The output file path of result")
	flag.Usage = usage
	flag.Parse()
	if h {
		flag.Usage()
		os.Exit(0)
	}
	if flag.NFlag() < 4 {
		usage()
		os.Exit(0)
	}
}

func main() {
	pocs := Getpoc()
	Head()
	p := NewPool(ThreadNum)

	file, err := os.Open(IpFile)
	if err != nil {
		fmt.Println("open file failed!, err:", err)
		return
	}
	rd := bufio.NewReader(file)
	go func() {
		for {
			line, err := rd.ReadString('\n')
			if err != nil {
				time.Sleep(10 * time.Second)

			}
			if io.EOF == err {
				p.EntryChannel <- NewTask(func() error {

					GetResult()
					os.Exit(0)
					return nil
				})

			}
			p.EntryChannel <- NewTask(func() error {
				Request(line[:len(line)-1], pocs)
				return nil
			})

		}

	}()

	//启动协程池p
	p.Run()

}
