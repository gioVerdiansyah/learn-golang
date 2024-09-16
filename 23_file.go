package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func makeFile() {
	filePath, err := os.Getwd()
	if err != nil {
		log.Fatal("Failed to get working directory.\n", err)
	}

	var fileName string = "test.txt"
	filePath = fmt.Sprintf("%s\\%s", filePath, fileName)

	var _, statErr = os.Stat(filePath)

	if os.IsNotExist(statErr) {
		var file, err = os.Create(filePath)
		if err != nil {
			fmt.Printf("Failed to create file: %s\n", fileName)
		} else {
			fmt.Printf("Successfully create %s file!\n", fileName)
		}
		defer file.Close()
	}

	if os.IsExist(statErr) {
		fmt.Printf("%s already exists! Skip\n", fileName)
	}
}

func writeFile() {
	filePath, err := os.Getwd()
	if err != nil {
		log.Fatal("Failed to get working directory.\n", err)
	}

	filePath = fmt.Sprintf("%s\\test.txt", filePath)

	file, err := os.OpenFile("test.txt", os.O_WRONLY, 0666)
	if err != nil {
		fmt.Printf("Failed to open file: %s\n", err)
	}

	_, err = file.WriteString("Hello my name is \"Verdi\" \n")
	if err != nil {
		fmt.Printf("Failed to write file: %s\n", err)
	}
	_, err = file.WriteString("I'm is from Indonesia. \n")
	if err != nil {
		fmt.Printf("Failed to write file: %s\n", err)
	}
}

func readFile() {
	var filePath string = "D:\\Document\\Golang\\learn_golang\\test.txt"
	var fileName string = filepath.Base(filePath)

	file, err := os.OpenFile(filePath, os.O_RDONLY, 0644)
	if err != nil {
		fmt.Printf("Failed to open file: %s\n", fileName)
		fmt.Println(err)
	}

	defer file.Close()

	var text = make([]byte, 1024)
	for {
		n, err := file.Read(text)
		if err != io.EOF {
			if err != nil {
				break
			}
		}

		if n == 0 {
			break
		}
	}

	fmt.Printf("-----%s-----\n", fileName)
	fmt.Println(string(text))
}

func deleteFile() {
	filePath, err := os.Getwd()
	if err != nil {
		log.Fatal("Failed to get working directory.\n", err)
	}

	filePath = fmt.Sprintf("%s\\", "test.txt")
	fileName := filepath.Base(filePath)

	err = os.Remove(filePath)
	if err != nil {
		fmt.Printf("Failed to delete \"%s\" file\n", fileName)
		fmt.Println(err)
		return
	}

	fmt.Printf("Successfully deleted \"%s\" file!\n", fileName)
}

func writeAndReadFileWithFullPath() {
	var fullPath string = "D:\\Document\\Golang\\learn_golang\\results\\test.txt"
	var dirPath string = filepath.Dir(fullPath)
	var fileName string = filepath.Base(fullPath)
	var _, alreadyExist = os.Stat(fullPath)

	if os.IsNotExist(alreadyExist) {
		err := os.MkdirAll(dirPath, os.ModePerm)
		if err != nil {
			fmt.Println("Failed to create directory.\n", err)
			return
		}

		file, err := os.Create(fullPath)
		if err != nil {
			fmt.Printf("Failed to create \"%s\" file.\n", fileName)
			fmt.Println(err)
			return
		}

		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				fmt.Println("Failed to close file.")
				return
			}
		}(file)
	} else {
		fmt.Printf("%s already exists! Skip\n", fileName)
	}

	fmt.Println("Writing file...")
	file, err := os.OpenFile(fullPath, os.O_WRONLY, 0666)
	if err != nil {
		fmt.Printf("Failed to open \"%s\" file.\n", fileName)
		fmt.Println(err)
		return
	}

	_, err = file.WriteString("Hello my name is Verdi \n")
	if err != nil {
		fmt.Printf("Failed to write to \"%s\" file.\n", fileName)
		return
	}
	_, err = file.WriteString("I'm is a BackEnd Developer \n")
	if err != nil {
		fmt.Printf("Failed to write to \"%s\" file.\n", fileName)
		return
	}
	_, err = file.WriteString("Nice to meet you!")
	if err != nil {
		fmt.Printf("Failed to write to \"%s\" file.\n", fileName)
		return
	}

	fmt.Println("Successfully create and write file!")

	//	Read
	fmt.Println()

	file, err = os.OpenFile(fullPath, os.O_RDONLY, 0644)
	if err != nil {
		fmt.Printf("Failed to open \"%s\" file.\n", fileName)
		fmt.Println(err)
		return
	}

	var text = make([]byte, 2048)
	for {
		n, err := file.Read(text)
		if err != io.EOF {
			if err != nil {
				fmt.Printf("Failed to read from \"%s\" file.\n", fileName)
				fmt.Println(err)

				break
			}
		}

		if n == 0 {
			break
		}
	}

	fmt.Printf("----%s----\n", fileName)
	fmt.Println(string(text))

	err = file.Close()
	if err != nil {
		fmt.Printf("Failed to close \"%s\" file.\n", fileName)
	}
}

func main() {
	//makeFile()
	//writeFile()
	//readFile()
	deleteFile()
	//writeAndReadFileWithFullPath()
}
