package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	strings, err := recognizeText("./test.jpeg")
	if err != nil {
		fmt.Println(err)
	}
	for _, value := range strings {
		fmt.Println(value)
	}
}

func recognizeText(imagePath string) ([]string, error) {
	cmd := exec.Command("tesseract", imagePath, "-")
	text, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	return strings.Split(string(text), " "), nil
}
