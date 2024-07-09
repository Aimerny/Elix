package render

import (
	"fmt"
	"github.com/nfnt/resize"
	"image"
	"image/draw"
	"image/jpeg"
	"os"
	"testing"
)

func TestImg(t *testing.T) {
	// 打开背景图片
	backgroundFile, err := os.Open("plain.jpg")
	if err != nil {
		fmt.Println("Error: Unable to open background image.")
		return
	}
	defer backgroundFile.Close()

	// 解码背景图片
	backgroundImg, err := jpeg.Decode(backgroundFile)
	if err != nil {
		fmt.Println("Error: Unable to decode background image.")
		return
	}

	// 打开素材图片
	overlayFile, err := os.Open("cat.jpg")
	if err != nil {
		fmt.Println("Error: Unable to open overlay image.")
		return
	}
	defer overlayFile.Close()

	// 解码素材图片
	overlayImg, err := jpeg.Decode(overlayFile)
	if err != nil {
		fmt.Println("Error: Unable to decode overlay image.")
		return
	}

	// 获取背景图的大小
	bgBounds := backgroundImg.Bounds()

	// 创建一个新的 RGBA 图像用于合成
	resultImg := image.NewRGBA(bgBounds)

	// 将背景图绘制到新图像上
	draw.Draw(resultImg, bgBounds, backgroundImg, image.Point{0, 0}, draw.Src)

	// 素材图的位置
	offset := image.Pt(100, 100) // 可以根据需要调整位置

	// 获取素材图的大小
	scaledWidth := 1000
	scaledHeight := 1000

	// 缩放
	m := resize.Resize(uint(scaledWidth), uint(scaledHeight), overlayImg, resize.Bilinear)

	// 将素材图绘制到新图像的指定位置上
	draw.Draw(resultImg, image.Rect(offset.X, offset.Y, offset.X+scaledWidth, offset.Y+scaledHeight), m, image.Point{0, 0}, draw.Over)

	// 保存结果图像到文件
	outputFile, err := os.Create("result.jpg")
	if err != nil {
		fmt.Println("Error: Unable to create result image.")
		return
	}
	defer outputFile.Close()

	// 将合成的图像编码为 JPEG 并写入文件
	if err := jpeg.Encode(outputFile, resultImg, nil); err != nil {
		fmt.Println("Error: Unable to encode result image.")
		return
	}

	fmt.Println("Image composition complete.")
}
