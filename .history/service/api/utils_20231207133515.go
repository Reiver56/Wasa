package api

import (
	"Wasa-photo-1905917/service/api/reqcontext"
	"fmt"
	"os"
	"path/filepath"
)

func createFolder(id_user string, ctx reqcontext.RequestContext) error {
	path := "./storage/" + fmt.Sprintf(id_user) + "/photos"
	fmt-
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		ctx.Logger.WithError(err).Error("can't create folder")
		return err
	}
	return nil
}
func getUserFolder(id_user string) (string, error) {
	return filepath.Join(photoFolder, id_user, "users"), nil
}