package main

import (
	"image"
	"log"
	"io/fs"
	"fmt"
	"image/png"
	_ "image/jpeg"
   	"path/filepath"
	"os"
	"github.com/disintegration/gift"
)

func loadImage(filename string) image.Image {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("os.Open failed: %v", err)
	}
	defer f.Close()
	fmt.Println(filename)
	img, _, err := image.Decode(f)
	if err != nil {
		log.Fatalf("image.Decode failed: %v", err)
	}
	return img
}

func saveImage(filename string, img image.Image) {
	f, err := os.Create(filename)
	if err != nil {
		log.Fatalf("os.Create failed: %v", err)
	}
	defer f.Close()
	err = png.Encode(f, img)
	if err != nil {
		log.Fatalf("png.Encode failed: %v", err)
	}
}

func walk(path string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}

	if !d.IsDir() {
		src := loadImage(path)
		g := gift.New(gift.Colorize(263, 30, 70))
		dst := image.NewNRGBA(g.Bounds(src.Bounds()))
		g.Draw(dst, src)

		os.Remove(path)
		saveImage(path, dst)
	}

	return nil
}

func Colorize() {
	filepath.WalkDir("./images", walk)
}