package main

import "apex-oled/internal/oled"

func main() {
	//DisplayBinaryData(oled.GetTestData())
	bytes := oled.GetImageBytes()
	//DisplayHexData(bytes)
	DisplayBinaryData(bytes)
	WriteDataToOled(bytes)
}
