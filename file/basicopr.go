package file

import (
	"fmt"
	"log"
	"os"
)

func ReadFile(name string) {
	dirs, err := os.ReadDir(name)
	if err != nil {
		log.Fatal(err)
		return
	}
	for _, dir := range dirs {
		fmt.Println(dir.IsDir(), dir.Name(), dir.Type())
	}
}
