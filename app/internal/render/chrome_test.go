package render

import (
	"context"
	"github.com/chromedp/chromedp"
	"io/ioutil"
	"log"
	"testing"
)

func Test_chrome(t *testing.T) {
	options := []chromedp.ExecAllocatorOption{
		chromedp.WindowSize(2560, 1440),
		chromedp.Flag("headless", false),
	}
	options = append(chromedp.DefaultExecAllocatorOptions[:], options...)
	ctx, _ := chromedp.NewExecAllocator(context.Background(), options...)

	ctx, _ = chromedp.NewContext(ctx, chromedp.WithLogf(log.Printf)) // 会打开浏览器并且新建一个标签页进行操作

	var buf []byte
	if err := chromedp.Run(ctx, ScreenshotTasks("http://aimerny.top", &buf)); err != nil {
		log.Fatal(err)
	}
	// 将截图保存到文件
	err := ioutil.WriteFile("screenshot.png", buf, 0644)
	if err != nil {
		log.Fatal(err)
	}

}

func elementScreenshot(urlstr, sel string, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.Screenshot(sel, res, chromedp.NodeVisible),
	}
}
func ScreenshotTasks(url string, imageBuf *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.FullScreenshot(imageBuf, 100),
	}
}
