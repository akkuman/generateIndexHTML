//用来作为自己的file.codecat.one小型index目录使用，配合resilio sync同步然后生成index.html
package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	var htmltext string
	var filenames []string

	//遍历.（本）目录下的所有文件（注意不是递归遍历，这里只是一次遍历）
	//并且生成的文件名排除.sync和generateHTML.exe和index.html
	files, err := ioutil.ReadDir(".")
	if err != nil {
		os.Exit(1)
	}
	for _, file := range files {
		if file.Name() != ".sync" && file.Name() != "generateHTML.exe" && file.Name() != "index.html" {
			filenames = append(filenames, file.Name())
		}
	}

	//文件名信息拼接成所需的html代码
	for _, name := range filenames {
		htmltext += "			<li><a href=\"./" + name + "\">" + name + "</a></li>\n"
	}
	AllHTML := "<html>\n\t<head>\n\t\t<meta charset=\"utf-8\">\n\t\t<title>Akkuman's Content</title>\n\t</head>\n\t<body>\n\t\t<h1>Akkuman's Content:</h1>\n\t\t<ul type=\"circle\">\n" +
		htmltext +
		"\t\t<img src=\"http://ip.ntrqq.net/images/titan.png?wd=b25pY3lhbiBkYWlzdWtp\" border=\"0\" />\n\t\t</ul>\n\t</body>\n</html>"

	//打开文件，如果文件不存在就创建，如果文件存在就打开并且清空，然后写入html代码
	TxtHTML, err := os.OpenFile("index.html", os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0660)
	if err != nil {
		fmt.Println("OpenFile Error:" + err.Error())
	}
	HTMLWriter := bufio.NewWriter(TxtHTML)
	defer TxtHTML.Close()
	HTMLWriter.WriteString(AllHTML)
	HTMLWriter.Flush()
}
