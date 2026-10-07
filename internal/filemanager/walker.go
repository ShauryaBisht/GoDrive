package filemanager

import (
	"io/fs"
	"path/filepath"
)

func WalkDirectory(root string, fn fs.WalkDirFunc) error{
	result:=filepath.WalkDir(root,fn)
	if(result!=nil){
		return result
	}
	return nil
}
