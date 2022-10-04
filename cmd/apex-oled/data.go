package main

import (
	"apex-oled/internal/oled"
	"encoding/hex"
	"fmt"
	"github.com/aleshkashell/usb"
	"log"
	"strings"
)

func WriteDataToOled(payload []byte) {
	hids, err := usb.Enumerate(0x1038, 0x1610)
	if err != nil {
		log.Fatal(err)
	}
	for i, hid := range hids {
		fmt.Println(strings.Repeat("-", 128))
		fmt.Printf("HID #%d\n", i)
		fmt.Printf("  OS Path:      %s\n", hid.Path)
		fmt.Printf("  Vendor ID:    %#04x\n", hid.VendorID)
		fmt.Printf("  Product ID:   %#04x\n", hid.ProductID)
		fmt.Printf("  Release:      %d\n", hid.Release)
		fmt.Printf("  Serial:       %s\n", hid.Serial)
		fmt.Printf("  Manufacturer: %s\n", hid.Manufacturer)
		fmt.Printf("  Product:      %s\n", hid.Product)
		fmt.Printf("  Usage Page:   %d\n", hid.UsagePage)
		fmt.Printf("  Usage:        %d\n", hid.Usage)
		fmt.Printf("  Interface:    %d\n", hid.Interface)
	}
	if len(hids) == 0 {
		log.Fatalf("No found device\n")
	}
	log.Printf("Opening device...")
	device, err := hids[1].Open()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Device is opened")
	defer device.Close()
	log.Printf("Sending data....")
	data := []byte{0x00, 0x65}
	data = append(data, payload...)
	result, err := device.Write(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Result %d\n", result)
	log.Printf("Data is sent")
	fmt.Println(strings.Repeat("=", 128))
}

func DisplayHexData(payload []byte) {
	for n, i := range hex.EncodeToString(payload) {
		if n%2 == 0 {
			fmt.Printf(" ")
		}
		if n%(oled.DisplayWidth/4) == 0 {
			fmt.Println()
		}
		fmt.Printf(string(i))
	}
	fmt.Println()
}

func DisplayBinaryData(payload []byte) {
	offset := 0
	for i := 0; i < offset; i++ {
		if i%(oled.DisplayWidth/8) == 0 {
			fmt.Println()
		}
		fmt.Printf("%08b", 0x00)
	}
	for n, i := range payload {
		//if n < offset {
		//	continue
		//}
		//if n%2 == 0 {
		//	fmt.Printf(" ")
		//}
		if (n+offset)%(oled.DisplayWidth/8) == 0 {
			fmt.Println()
		}
		fmt.Printf("%08b", i)
	}
	fmt.Println()
}
