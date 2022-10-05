package main

import "apex-oled/internal/oled"

func main() {
	//DisplayBinaryData(oled.GetTestData())
	labels := []string{
		"     ",
		"        ^ ^",
		"      \\(^w^)/",
	}
	bytes := oled.SimpleTextToImage(labels)
	//DisplayHexData(bytes)
	DisplayBinaryData(bytes)
	WriteDataToOled(bytes)
}
