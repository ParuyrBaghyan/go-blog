package routes

import (
	"go-blog/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getAllPosts(context *gin.Context) {
	posts, err := models.GetAllPosts()

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
	var post models.Post
	err := context.ShouldBindJSON(&post)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	userId := context.GetInt64("userId")
	post.AuthorId = userId

	err = post.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create post.Try again later."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Post created", "post": post})
}
