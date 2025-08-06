package routes

import (
	"fmt"
	"go-blog/models"
	"go-blog/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func getAllPosts(context *gin.Context) {
	page, err := strconv.Atoi(context.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(context.DefaultQuery("limit", "10"))
	if err != nil || limit < 1 {
		limit = 10
	}

	offset := (page - 1) * limit

	posts, err := models.GetPaginatedPosts(limit, offset)


	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch posts."})
		return
	}

	context.JSON(http.StatusOK, posts)
}

func getPost(context *gin.Context) {
	postId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse post id"})
	}

	post, err := models.GetPostById(postId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch post" + err.Error()})
		return
	}

	context.JSON(http.StatusOK, post)
}

func createPost(context *gin.Context) {
	title := context.PostForm("title")
	content := context.PostForm("content")
	tags := context.PostForm("tags")
	dateTimeStr := context.PostForm("dateTime")

	parsedTime, err := time.Parse(time.RFC3339, dateTimeStr)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid dateTime format"})
		return
	}

	userId := context.GetInt64("userId")

	post := models.Post{
		Title:    title,
		Content:  content,
		Tags:     tags,
		DateTime: parsedTime,
		AuthorId: int(userId),
	}

	err = post.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create post. Try again later.", "error": err.Error()})
		return
	}

	err = utils.AddMedia(context, post.Id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save post media. Try again later.", "error": err.Error()})
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Post created", "post": post})
}

func updatePost(context *gin.Context) {
	postId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse post id"})
		return
	}

	userId := context.GetInt64("userId")
	post, err := models.GetPostById(postId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the post."})
		return
	}

	if post.AuthorId != int(userId) {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to update post."})
		return
	}

	var updatePost models.Post
	err = context.ShouldBindJSON(&updatePost)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data." + err.Error()})
		return
	}

	updatePost.Id = postId

	err = updatePost.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update the event."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event updated successfully!"})

}

func deletePost(context *gin.Context) {
	postId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse post id"})
		return
	}

	userId := context.GetInt64("userId")
	post, err := models.GetPostById(postId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the post."})
		return
	}
	fmt.Println("authorId", post.AuthorId)
	fmt.Println("userId", userId)
	if post.AuthorId != int(userId) {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to delete post."})
		return
	}

	err = post.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete the post."})
		return
	}

	err = utils.RemoveMedia(context, post.Id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not remove post media. Try again later.", "error": err.Error()})
	}

	context.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully!"})
}
