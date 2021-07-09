package main

import (
    "log"
    "os"
    "path/filepath"
)

func RemoveGlob(path string) (err error) {
    contents, err := filepath.Glob(path)
    if err != nil {
        return
    }
    for _, item := range contents {
        err = os.RemoveAll(item)
        if err != nil {
            return
        }
    }
    return
}    

func main() {
    err := RemoveGlob(os.Args[1])
    if err != nil {
        log.Fatalf("Error removing files: %+v", err)
    } else {
      log.Println("Removed files")
    }
}