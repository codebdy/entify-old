package storage

import "mime/multipart"

type File struct {
	File     multipart.File
	Filename string
	Size     int64
}
