package storage

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"rxdrag.com/entify/consts"
)

type File struct {
	File     multipart.File
	Filename string
	Size     int64
}

type FileInfo struct {
	Path     string `json:"path"`
	Filename string `json:"fileName"`
	Size     int64  `json:"size"`
	MimeType string `json:"mimeType"`
	ExtName  string `json:"extName"`
}

var mimeTypes = map[string]string{
	".css":  "text/css; charset=utf-8",
	".gif":  "image/gif",
	".htm":  "text/html; charset=utf-8",
	".html": "text/html; charset=utf-8",
	".jpg":  "image/jpeg",
	".js":   "application/x-javascript",
	".pdf":  "application/pdf",
	".png":  "image/png",
	".xml":  "text/xml; charset=utf-8",
}

func (f *File) extName() string {
	return filepath.Ext(f.Filename)
}

func (f *File) mimeType() string {
	//mtype, err := mimetype.DetectReader(f.File)

	return mimeTypes[f.extName()]
}

func (f *File) Save() FileInfo {
	name := fmt.Sprintf("%s%s", uuid.New().String(), f.extName())
	localPath := fmt.Sprintf("%s/%s", consts.UPLOAD_PATH, name)
	file, err := os.OpenFile(
		localPath,
		os.O_WRONLY|os.O_CREATE,
		0666,
	)
	defer file.Close()
	if err != nil {
		panic(err.Error())
	}
	io.Copy(file, f.File)
	return FileInfo{Path: name, Filename: f.Filename, Size: f.Size, MimeType: f.mimeType(), ExtName: f.extName()}
}
