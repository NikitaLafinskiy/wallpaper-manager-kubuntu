package main

import (
	"os"
	"log"
	"fmt"
	"time"
	"math/rand/v2"
	"os/exec"
)

func ChangeWallpaper() {
	BASE_PATH := "/home/tainticulus/Documents/wallpapers/"
	BASE_DIR := BASE_PATH + "images"
	FILE_SEPARATOR := "/"
	HORIZONTAL_DIR := "horizontal"
	VERTICAL_DIR := "vertical"
	HORIZONTAL_COPY_LOCATION := BASE_PATH + "image-horizontal.jpg"
	VERTICAL_COPY_LOCATION := BASE_PATH + "image-vertical.jpg"
	REFRESH_SCRIPT_PATH := BASE_PATH + "change-wallpaper.sh"

	// Find the amount of folders in a dir
	files, err := os.ReadDir(BASE_DIR)
	if err != nil {
		log.Fatal(err)
	}

	// Find the current day of the month and the dir to choose
	hour := time.Now().Hour()
	currDir := files[hour % len(files)]

	// Select a random vertical and horizontal image out of that dir
	dirName := currDir.Name()
	horizontalPath := BASE_DIR + FILE_SEPARATOR + dirName + FILE_SEPARATOR + HORIZONTAL_DIR
	verticalPath := BASE_DIR + FILE_SEPARATOR + dirName + FILE_SEPARATOR + VERTICAL_DIR

	horizontalDir, err := os.ReadDir(horizontalPath)
	verticalDir, err := os.ReadDir(verticalPath)

	if (err != nil) {
		log.Fatal(err)
	}

	horizontalIndex := rand.IntN(len(horizontalDir))
	verticalIndex := rand.IntN(len(verticalDir))

	fmt.Println(verticalIndex)

	// Copy the image to the desired lcoation
	horizontalCopyPath := horizontalPath + FILE_SEPARATOR + horizontalDir[horizontalIndex].Name()
	verticalCopyPath := verticalPath + FILE_SEPARATOR + verticalDir[verticalIndex].Name()
	fmt.Println(verticalCopyPath)

	os.Remove(HORIZONTAL_COPY_LOCATION)
	os.Link(horizontalCopyPath, HORIZONTAL_COPY_LOCATION)

	os.Remove(VERTICAL_COPY_LOCATION)
	os.Link(verticalCopyPath, VERTICAL_COPY_LOCATION)

	// Refresh the current wallpaper
	cmd := exec.Command("sh", REFRESH_SCRIPT_PATH)
	cmd.Run()
}