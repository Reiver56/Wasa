package api

import (
	"Wasa-photo-1905917/service/api/reqcontext"
	"os"
	"path/filepath"
)

func createFolder(id_user string, ctx reqcontext.RequestContext) error {
	path := ""

	err := os.MkdirAll(filepath.Join(path, "users"), os.ModePerm)
	if err != nil {
		ctx.Logger.WithError(err).Error("can't create folder")
		return err
	}
	return nil
}
func getUserFolder(id_user string) (string, error) {
	return filepath.Join(photoFolder, id_user, "users"), nil
}
