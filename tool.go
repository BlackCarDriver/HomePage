package main
import (
	"fmt"
	"io/ioutil"
	// "strings"
	"os"
	"log"
	"strconv"
)
func RenameJPG() {
	path := "/home/blackcardriver/Documents/date/images/"
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	i:=18
	for _, file := range files {
		oldpath := path + file.Name()
		newpath := path + strconv.Itoa(i) + ".jpg"
		fmt.Println(oldpath)
		fmt.Println(newpath)
		os.Rename(oldpath,newpath)
		i++
	}
return
}