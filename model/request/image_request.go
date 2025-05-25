package request

import "mime/multipart"

type ImageRequest struct {
	File multipart.File
	Ext  string
}
