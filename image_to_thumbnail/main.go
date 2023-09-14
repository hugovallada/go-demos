package main

import "github.com/hugovallada/image-to-thumbnail/pkg/image"

func main() {
	e := image.GenerateThumbnail("./lambda-tracing.png")
	if e != nil {
		panic(e)
	}
}
