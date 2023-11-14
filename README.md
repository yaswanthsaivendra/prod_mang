# prod_mang

## tech stack
framework : gin
database : sqlite (file)

## run and test
```
go run main.go
go test ./...

```


## Implementations
- Models created 
  - User
  - Product
  - Image

- Image Analysis 
  - Goroutines make the image compression process isolated from the user interaction with the APIs
  - images are compressed and new file names are stored in <i>compressed_product_images</i> field


