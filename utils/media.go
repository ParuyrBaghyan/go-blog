package utils

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func AddMedia(context *gin.Context, postId int64) error {
	file, err := context.FormFile("data")
	if err == nil {
		folderPath := filepath.Join("media", fmt.Sprintf("%d", postId))
		err = os.MkdirAll(folderPath, os.ModePerm)
		if err != nil {
			return err
		}

		filePath := filepath.Join(folderPath, file.Filename)
		if err := context.SaveUploadedFile(file, filePath); err != nil {
			return err
		}
	}

	return err
}

func RemoveMedia(context *gin.Context, postId int64) error {
	foldePath := filepath.Join("media", fmt.Sprintf("%d", postId))
	err := os.RemoveAll(foldePath)
	if err != nil {
		return err
	}

	return err
}
