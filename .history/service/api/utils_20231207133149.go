package api

import (
	"Wasa-photo-1905917/service/api/reqcontext"
	"fmt"
	"os"
	"path/filepath"
)

func createFolder(id_user string, ctx reqcontext.RequestContext) error {
	path := "./storage/" + fmt.Sprintf(id_user) + "/photos"
	
}
func getUserFolder(id_user string) (string, error) {
	return filepath.Join(photoFolder, id_user, "users"), nil
}
