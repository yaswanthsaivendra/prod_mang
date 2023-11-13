package helper

import (
	"fmt"
	"os"

	"image"
	"image/png"

	"github.com/google/uuid"
	compression "github.com/nurlantulemisov/imagecompression"
	"github.com/yaswanthsaivendra/prod_mang/database"
	"github.com/yaswanthsaivendra/prod_mang/model"
)

func TriggerImagesCompression(imageFiles *[]model.Image) {

	for _, imageFile := range *imageFiles {
		ImageProcessing(imageFile)
	}
	return
}

func ImageProcessing(imageFile model.Image) {

	filePath := fmt.Sprintf("./static/product_images/"+"%s", imageFile.ProductImage)

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err.Error())
	}

	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Println(err.Error())
	}

	compressing, _ := compression.New(50)
	compressedImage := compressing.Compress(img)

	filename := uuid.New().String() + ".png"
	fmt.Println(filename)

	outputPath := fmt.Sprintf("./static/compressed_product_images/"+"%s", filename)

	outputFile, err := os.Create(outputPath)
	if err != nil {
		fmt.Printf("error creating file: %s", err)
	}

	defer outputFile.Close()

	// Encode the compressed image to PNG format
	err = png.Encode(outputFile, compressedImage)

	if err != nil {
		fmt.Println(err.Error())
	}

	//save the image path
	imageFile.CompressedProductImage = filename
	err = database.Database.Save(imageFile).Error
	if err != nil {
		fmt.Println(err.Error())
	}
	return
}
