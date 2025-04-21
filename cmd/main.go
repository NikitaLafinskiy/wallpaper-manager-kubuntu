package main

import (
	"os"
)

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "changeWallpaper":
			ChangeWallpaper()
		case "colorize":
			Colorize()
		}
	} else {
		ChangeWallpaper()
	}
}