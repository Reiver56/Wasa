package api

import (
	"Wasa-photo-1905917/service/api/reqcontext"
	"fmt"
	"os"
)

func createFolder(id_user string, ctx reqcontext.RequestContext) error {
	path := "./storage/" + fmt.Sprintf(id_user) + "/photos"
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		ctx.Logger.WithError(err).Error("can't create folder")
		return err
	}
	return nil
}
