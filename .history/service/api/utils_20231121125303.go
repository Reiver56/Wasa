package api

import (
	"Wasa-photo-1905917/service/api/reqcontext"
	"image/jpeg"
	"io"
)
//  Function that check format of photo

func getUserFolder(b io.ReadCloser, n io.ReadCloser, ctx reqcontext.RequestContext) error{
	_, errJpg := jpeg.Decode(b)
	if errJpg != nil {

		b = newReader
}