package string

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

//GetCurrentDirectory get current exec program directory.
func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}
