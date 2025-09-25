package main

import (
	"fmt"
	"os"
	"path/filepath"

	textprocessor "textprocessor/internal/transformers"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Ошибка: необходимо указать входной и выходной файлы.")
		fmt.Println("Использование: go run . <input_file> <output_file>")
		os.Exit(1)
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	if inputFile == outputFile {
		fmt.Println("Ошибка: входной и выходной файлы должны быть разными.")
		os.Exit(1)
	}

	if filepath.Ext(outputFile) != ".txt" {
		fmt.Println("Ошибка: выходной файл должен иметь расширение .txt.")
		os.Exit(1)
	}

	if _, err := os.Stat(inputFile); os.IsNotExist(err) {
		fmt.Printf("Ошибка: входной файл %s не существует.\n", inputFile)
		os.Exit(1)
	}

	content, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Ошибка при чтении файла %s: %v\n", inputFile, err)
		os.Exit(1)
	}

	processedContent := textprocessor.ProcessText(string(content))

	err = os.WriteFile(outputFile, []byte(processedContent), 0644)
	if err != nil {
		fmt.Printf("Ошибка при записи в файл %s: %v\n", outputFile, err)
		os.Exit(1)
	}
}
