package api

import "path/filepath"

func getUserFolder(id_user string ) (string, error){
	return filepath.Join(photoFolder,id_user,"photos"), nil

}