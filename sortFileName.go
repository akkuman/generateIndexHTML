package main

import (
	"path"
	"sort"
)

// 自定义排序类型
type MyFileNameList []string

func (filenames MyFileNameList) Len() int { return len(filenames) }
func (filenames MyFileNameList) Swap(i, j int) {
	filenames[i], filenames[j] = filenames[j], filenames[i]
}
func (filenames MyFileNameList) Less(i, j int) bool {
	iSuffixName := path.Ext(filenames[i]) + filenames[i]
	jSuffixName := path.Ext(filenames[j]) + filenames[j]
	return iSuffixName > jSuffixName
}

func sortFileName(filenameList []string) {
	sort.Sort(MyFileNameList(filenameList))
}
