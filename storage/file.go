package storage

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/gabriel-vasile/mimetype"
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
}

func (f *File) extName() string {
	return filepath.Ext(f.Filename)
}

func (f *File) mimeType() string {
	mtype, err := mimetype.DetectReader(f.File)

	if err != nil {
		panic(err.Error())
	}

	return mtype.String()
}

func (f *File) Save() FileInfo {
	path := fmt.Sprintf("./%s/%s%s", consts.UPLOAD_PATH, uuid.New().String(), f.extName())
	file, err := os.OpenFile(
		path,
		os.O_WRONLY|os.O_CREATE,
		0666,
	)
	defer file.Close()
	if err != nil {
		panic(err.Error())
	}
	io.Copy(file, f.File)
	return FileInfo{Path: path, Filename: f.Filename, Size: f.Size, MimeType: f.mimeType()}
}
