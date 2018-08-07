package main

import (
	"github.com/skip2/go-qrcode"
	"log"
)

func main() {
	err := qrcode.WriteFile("https://kikna.io", qrcode.Medium, 256, "qr.png")

	if err != nil {
		log.Fatal(err)
	}

	println("Qr code has been generated successfully")
}
