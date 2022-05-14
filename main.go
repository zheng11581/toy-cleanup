package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"time"
)

func init()  {

}

func main()  {
	dir := flag.String("directory", "/This/is/a/dir", "需要删除文件的目录")
	regexp := flag.String("regexp", "pdf", "需要删除文件的后缀")
	days := flag.Int64("before", 30, "需要保留多少天内的文件")
	flag.Parse()
	if *dir == "" || *dir == "/" {
		fmt.Printf("The directory is not allowed\n")
		os.Exit(0)
	}
	absDir, _ := filepath.Abs(*dir)
	err := filepath.WalkDir(absDir, func(path string, d fs.DirEntry, err error) error {
		result, _ := filepath.Match(absDir + "/*" + *regexp, path)
		if result {
			dInfo, _ := d.Info()
			modTime := dInfo.ModTime()
			before := time.Now().AddDate(0, 0, -int(*days))
			if modTime.Before(before) {
				fmt.Printf("The matched file is %s.\n", path)
				err := os.Remove(path)
				if err != nil {
					return err
				}

			}
		}
		if err == nil {
			return nil
		}
		return err
	})
	//err := filepath.Walk(absDir, func(path string, info fs.FileInfo, err error) error {
	//	result, _ := filepath.Match(absDir + "/*" + *regexp, path)
	//	if result {
	//		modTime := info.ModTime()
	//		before := time.Now().AddDate(0, 0, -int(*days))
	//		if modTime.Before(before) {
	//			fmt.Printf("The matched file is %s.\n", path)
	//			os.Remove(path)
	//		}
	//	}
	//	if err == nil {
	//		return nil
	//	}
	//	return err
	//})
	if err != nil {
		fmt.Println(err)
	}
	//os.Exit(0)

}
