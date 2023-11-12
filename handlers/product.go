package handlers

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yaswanthsaivendra/prod_mang/helper"
	"github.com/yaswanthsaivendra/prod_mang/model"
)

func HandleProductUpload(c *gin.Context) {

	form, err := c.MultipartForm()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	imageFiles := form.File["images"]

	//getting user id
	user, err := helper.CurrentUser(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	productPrice, err := strconv.ParseFloat(form.Value["product_price"][0], 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newProduct := model.Product{
		ProductName:        form.Value["product_name"][0],
		ProductDescription: form.Value["product_description"][0],
		ProductPrice:       productPrice,
	}

	newProduct.UserID = user.ID

	uploadPath := "../static/product_images/"

	//uploading all images
	for i, image := range imageFiles {

		filename := filepath.Join(uploadPath, image.Filename)

		if err := c.SaveUploadedFile(image, filename); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error saving file %d", i)})
			return
		}

		newImage := model.Image{
			ProductImage: image.Filename,
			ProductID:    newProduct.ID,
		}
		newProduct.Images = append(newProduct.Images, newImage)
	}

	// creating the product

	savedProduct, err := newProduct.Save()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": savedProduct})
}

func GetAllProductsOfUser(c *gin.Context) {
	user, err := helper.CurrentUser(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user.Products})
}
