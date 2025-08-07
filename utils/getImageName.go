package utils

import (
	"errors"

	"strings"
)

func GetImageName(imageUrl string) (string, error) {
	dotIndex := strings.LastIndex(imageUrl, ".")
	if dotIndex == -1 {
		return "", errors.New("Could not find extention")
	}
	return imageUrl[:dotIndex], nil
}
