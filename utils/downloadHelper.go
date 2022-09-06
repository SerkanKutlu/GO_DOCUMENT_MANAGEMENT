package utils

import (
	"archive/zip"
	"io"
	"os"
)

func CreateZip() {
	archieve, err := os.Create("downloading.zip")
	if err != nil {
		panic(err)
	}
	defer archieve.Close()
	zipWriter := zip.NewWriter(archieve)
	file1, err := os.Open("documents/1662383062_arka.jpg.pdf")
	if err != nil {
		panic(err)
	}
	defer file1.Close()
	w1, err := zipWriter.Create("deneme/1662383062_arka.jpg.pdf")
	if err != nil {
		panic(err)
	}
	if _, err := io.Copy(w1, file1); err != nil {
		panic(err)
	}
	zipWriter.Close()
}
