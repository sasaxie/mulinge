package main

import "github.com/fogleman/gg"

func main() {
	drawBackground()
}

func drawBackground() {
	// 加载图片
	img, err := gg.LoadJPG("./source/img/airdrop.jpg")
	if err != nil {
		panic(err)
	}

	// 获取图片尺寸
	size := img.Bounds().Size()

	width := size.X
	height := size.Y

	// 以加载图片的宽高作为新图片的大小
	dc := gg.NewContext(width, height)

	// 画图
	dc.DrawImage(img, 0, 0)

	// 画矩形
	dc.SetLineWidth(1) // 设置画线的宽度
	dc.SetRGB(1, 0, 0) // 设置画线的颜色，取值范围[0,1]，其实就是x/255，比如红色是255,0,0，除以255是1,0,0
	dc.DrawRectangle(0, 0, 100, 100)
	dc.Stroke()

	// 画圆
	dc.SetLineWidth(2)
	dc.SetRGB(0, 1, 0)
	dc.DrawCircle(50, 50, 50)
	dc.Stroke()

	// 画椭圆
	dc.SetLineWidth(3)
	dc.SetRGB(0, 1, 0)
	dc.DrawEllipse(50, 50, 50, 25)
	dc.Stroke()

	// 保存新图片，一般quality设置为75即可，最高可设置为100，值越高，质量越好，但是占空间大
	err = dc.SaveJPG("./source/img/airdrop-bg.jpg", 75)
	if err != nil {
		panic(err)
	}
}
