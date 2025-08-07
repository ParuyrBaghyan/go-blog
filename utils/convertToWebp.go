// package utils

// import (
// 	"image/jpeg"
// 	"io/ioutil"

// 	"github.com/kolesa-team/go-webp/webp"

// 	"os"
// )

// func ConvertToWebp(inputPath, outputPath string) error {
// 	inFile, err := os.Open(inputPath)
// 	if err != nil {
// 		return err
// 	}
// 	defer inFile.Close()

// 	img, err := jpeg.Decode(inFile)
// 	if err != nil {
// 		return err
// 	}

// 	webpData, err := webp.
// 	if err != nil {
// 		return err
// 	}

// 	return ioutil.WriteFile(outputPath, webpData, 0644)
// }
