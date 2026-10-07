package main

import (
	"fmt"
	"io/fs"
	"log"

	"github.com/ShauryaBisht/GoDrive/internal/filemanager"
)

func printPath(path string,d fs.DirEntry,err error) error{
	if err != nil{
		return err
	}
	info:=d.Info()
    if(d.IsDir()){
	   fmt.Println("DIR ",path)
	}else {
		fmt.Println("FILE ",path)
	}
	return nil
}

func main(){
	err:=filemanager.WalkDirectory(".", printPath)
	if err != nil{
		log.Fatal(err)
	}
}