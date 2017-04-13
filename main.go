package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

//删除[]string切片中的值
func remove(slice []string, elems string) []string {
	for i, strText := range slice {
		if strText == elems {
			slice = append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

//读取每一行，handler是对每一行进行自定义操作的函数类型参数
func ReadPerLine(fileName string, handler func(string)) error {
	f, err := os.Open(fileName)
	defer f.Close()
	if err != nil {
		return err
	}
	buf := bufio.NewScanner(f)
	for buf.Scan() {
		line := buf.Text()
		handler(line)
	}
	return nil
}

//递归遍历返回所有dir文件夹
func getDirList(path string) []string {
	var dirNames []string
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if info == nil {
			return err
		}
		if info.IsDir() {
			dirNames = append(dirNames, path)
			return nil
		}
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() return %v\n", err)
	}
	return dirNames
}

//表层遍历dirname目录文件（夹），不是递归遍历，并且去掉指定文件（夹）名
func walkDir(dirname string) []string {
	var filenames []string
	//遍历dir目录下的所有文件（注意不是递归遍历，这里只是一次遍历）
	//并且生成的文件名排除.sync和generateHTML.exe和index.html
	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		os.Exit(1)
	}
	for _, file := range files {
		filenames = append(filenames, file.Name())
	}
	// 读取noView.txt每一行，如果文件名在noView中就从filenames中删除这个元素
	if err := ReadPerLine("noView.txt", func(line string) {
		if line != "" {
			filenames = remove(filenames, line)
		}
	}); err != nil {
		fmt.Printf("read noView.txt Error, return %v\n", err)
	}
	return filenames
}

//读取解析配置文件，返回一个值键对对应的map
func getTemplateConfig(filename string) (map[string]string, error) {
	templateMap := map[string]string{}
	if err := ReadPerLine(filename, func(line string) {
		if line != "" {
			config := strings.SplitN(line, ": ", 2)
			templateMap[config[0]] = config[1]
		}
	}); err != nil {
		fmt.Printf("parse config.config Error,Please Check its validity,return %s\n", err)
		return templateMap, err
	}
	return templateMap, nil

}

//打开并保存文件filename，如果文件不存在就创建，如果文件存在就打开并且清空，然后写入content
func saveFile(filename string, content *string) error {
	TxtHTML, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0660)
	if err != nil {
		fmt.Println("OpenFile Error:" + err.Error())
		return err
	}
	HTMLWriter := bufio.NewWriter(TxtHTML)
	defer TxtHTML.Close()
	HTMLWriter.WriteString(*content)
	HTMLWriter.Flush()
	return nil
}

func main() {
	start := time.Now().Nanosecond()
	for _, dirname := range getDirList(".") {
		var liText string
		filenames := walkDir(dirname)

		templateHTML, err := ioutil.ReadFile("Template.html")
		if err != nil {
			fmt.Println("open Template.html Error,Please check if the Template.html exists")
			os.Exit(1)
		}

		//文件名信息拼接成li列表,然后将解析好的内容放入AllHTML
		for _, name := range filenames {
			liText += "			<li><a href=\"./" + name + "\" target=\"_Blank\">" + name + "</a></li>\n"
		}
		AllHTML := strings.Replace(string(templateHTML), "{{content_li}}", liText, -1)
		configHTMLMap, err := getTemplateConfig("config.config")
		if err != nil {
			fmt.Printf("Config ReadLine Error %v\n", err)
		}
		for k, v := range configHTMLMap {
			AllHTML = strings.Replace(AllHTML, "{{"+k+"}}", v, -1)
		}
		//保存index.html
		if err := saveFile(dirname+"/index.html", &AllHTML); err != nil {
			fmt.Printf("saveFile Error,return %v\n", err)
		}
	}
	end := time.Now().Nanosecond()
	fmt.Printf("程序用时:%fms", float32(end-start)/1000000)
}
