package main

import (
	"fmt"
	"path"
)

var SuffixToClass map[string]string = map[string]string{
	"mpg": "video_icon", "mp4": "video_icon", "avi": "video_icon", "mov": "video_icon", "rmvb": "video_icon", "flv": "video_icon", "rm": "video_icon", "wmv": "video_icon", "3gp": "video_icon", "mkv": "video_icon", "m4v": "video_icon",
	"flac": "music_icon", "mp3": "music_icon", "wma": "music_icon", "wav": "music_icon", "aac": "music_icon", "ape": "music_icon", "ogg": "music_icon", "m4a": "music_icon",
	"zip": "pack_icon", "rar": "pack_icon", "7z": "pack_icon", "cab": "pack_icon", "iso": "pack_icon", "tar": "pack_icon", "gz": "pack_icon", "gzip": "pack_icon", "bz2": "pack_icon", "xz": "pack_icon", "jar": "pack_icon", "kz": "pack_icon",
	"chm": "document_icon", "txt": "document_icon", "doc": "document_icon", "docx": "document_icon", "pdf": "document_icon", "xls": "document_icon", "xlsx": "document_icon", "html": "document_icon", "md": "document_icon", "ppt": "document_icon", "rtf": "document_icon", "wps": "document_icon",
	"png": "picture_icon", "jpg": "picture_icon", "jpeg": "picture_icon", "ico": "picture_icon", "psd": "picture_icon", "gif": "picture_icon", "bmp": "picture_icon", "svg": "picture_icon",
	"ttf": "fontfile_icon", "otf": "fontfile_icon",
	"exe": "program_icon", "msi": "program_icon",
	"py": "code_icon", "go": "code_icon", "c": "code_icon", "cpp": "code_icon", "js": "code_icon", "java": "code_icon", "cs": "code_icon", "php": "code_icon", "pl": "code_icon", "asp": "code_icon", "aspx": "code_icon", "jsp": "code_icon",
}

//取文件后缀名并返回类似与class="video_icon mp4"的字符串
func getFileClass(filename string) string {
	if filename == "" {
		return ""
	}
	suffixName := path.Ext(filename)
	if suffixName == "" {
		return ""
	}
	suffixNameWithoutDot := suffixName[1:]
	if className, ok := SuffixToClass[suffixNameWithoutDot]; ok {
		return fmt.Sprintf("class=li-style %s %s", className, suffixNameWithoutDot)
	}
	return ""
}
