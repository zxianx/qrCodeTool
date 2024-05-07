package main

import (
    "fmt"
    "time"

    qrcode "github.com/skip2/go-qrcode"
)

func main() {
    // 输入网址
    var url string
    fmt.Print("请输入一段文本或网址：\n")
    fmt.Scan(&url)

    // 生成二维码
    qr, err := qrcode.New(url, qrcode.Medium)
    if err != nil {
        fmt.Println("无法生成二维码：", err)
        return
    }

    // 保存二维码图像到文件
    t := time.Now().Unix()
    pos := fmt.Sprintf("./qrCodeRes_%d.png", t)
    err = qr.WriteFile(480, pos)
    if err != nil {
        fmt.Println("无法保存二维码图像：", err)
        return
    }

    fmt.Println("二维码生成完成，图片位置：\n", pos)
}
