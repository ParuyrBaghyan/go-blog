package models

import (
	"fmt"
	"go-blog/db"
	"go-blog/utils"
	"time"

)

type PostMedia struct {
	Id        int64
	PostID    int    `json:"post_id" binding:"required"`
	MediaURL  string `json:"media_url" binding:"required"`
	MediaType string `json:"media_type" binding:"required"`
	Order     int    `json:"order"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (pm *PostMedia) SavePostMedia() error {
	query := "INSERT INTO post_media(post_id, media_url, media_type, `order`) VALUES (?, ?, ?, ?)"

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	filenameWithoutExt, err := utils.GetImageName(pm.MediaURL)
	if err != nil {
		return err
	}

	fmt.Println("pm.MediaURL----------------------", filenameWithoutExt)

	//--------------------------webp-------------------------------------------------
	

	//--------------------------webp-------------------------------------------------

	defer stmt.Close()
	result, err := stmt.Exec(pm.PostID, pm.MediaURL, pm.MediaType, pm.Order)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	pm.Id = id

	return err
}

func GetMediaByPostId(postId int64) ([]string, error) {
	query := "SELECT * FROM post_media WHERE post_id = ?"
	rows, err := db.DB.Query(query, postId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var mediaArray []string
	for rows.Next() {
		var pm PostMedia
		err := rows.Scan(&pm.Id, &pm.PostID, &pm.MediaURL, &pm.MediaType, &pm.Order, &pm.CreatedAt, &pm.UpdatedAt)
		if err != nil {
			return nil, err
		}

		mediaArray = append(mediaArray, pm.MediaURL)
	}

	return mediaArray, nil
}

func GetPostMedia(postId int64) (string, error) {

	mediaList, err := GetMediaByPostId(postId)
	if err != nil {
		return "", err
	}

	imageURL := fmt.Sprintf("%s/media/%s", utils.BaseURL, mediaList[0])
	return imageURL, nil
}
