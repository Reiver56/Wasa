package api

import (
	"Wasa-photo-1905917/service/api/reqcontext"
	"image/jpeg"
	"io"
)
//  Function that check format of photo

func getUserFolder(id string) error{
	_, errJpg := jpeg.Decode(b)
	if errJpg != nil {

		b = newReader
}