package helpers

import (
	"log"

	"github.com/gin-gonic/gin"
)

func UploadFile(c *gin.Context) (string, error) {
	// single file
	file, _ := c.FormFile("Avatar")
	log.Println(file.Filename)
	// Upload the file to specific dst.
	dst := "./views/assets/image/" + file.Filename
	err := c.SaveUploadedFile(file, dst)
	if err != nil {
		return "", err
	}
	log.Printf("'%s' uploaded!", file.Filename)
	return file.Filename, nil
}
