package examples

import (
	"fmt"

	"github.com/thuongtruong1009/gouse/io"
	"github.com/thuongtruong1009/gouse/shared"
)

func ioCheckDir() {
	isExist, err1 := io.IsExistDir("tmp")
	if err1 != nil {
		println(err1.Error())
	}
	if isExist {
		println("dir exist")
	} else {
		println("dir not exist")
	}
}

func ioCreateDir() {
	err2 := io.CreateDir("tmp")
	if err2 != nil {
		println(err2.Error())
	}
	println("dir created")
}

func ioRemoveDir() {
	err3 := io.RemoveDir("tmp")
	if err3 != nil {
		println(err3.Error())
	}
	println("dir removed")
}

func ioLsDir() {
	data, err := io.LsDir(".")
	if err != nil {
		println(err.Error())
		return
	}

	for _, v := range data {
		println(v)
	}
}

func ioHierarchyDir() {
	data, err := io.HierarchyDir(".")
	if err != nil {
		println(err.Error())
		return
	}

	for _, v := range data {
		println(v)
	}
}

func ioCurrentDir() {
	data, err := io.CurrentDir()
	if err != nil {
		println(err.Error())
		return
	}

	println(data)
}

func ioCheckFile() {
	isExist, err := io.IsExistFile("data.json")
	if err != nil {
		println(err.Error())
	}
	if isExist {
		println("file exist")
	} else {
		println("file not exist")
	}
}

func ioCreateFile() {
	err := io.CreateFile("data.json")
	if err != nil {
		println(err.Error())
	}
	println("file created")
}

func ioRemoveFile() {
	err := io.RemoveFile("data.json")
	if err != nil {
		println(err.Error())
	}
	println("file removed")
}

func ioReadFileByLine() {
	data, err := io.ReadFileByLine("main.go")
	if err != nil {
		println(err.Error())
	}
	for _, v := range data {
		println(v)
	}
}

func ioFileInfo() {
	data, err := io.FileInfo("main.go")
	if err != nil {
		println(err.Error())
	}
	fmt.Printf("File info: %+v\n", data.All)
	fmt.Println("File info (with name):", data.Name)
	fmt.Printf("File info (with size): %d bytes\n", data.Size)
	fmt.Println("File info (with permissions):", data.Mode)
	fmt.Println("File info (with last modified):", data.ModTime)
	fmt.Println("File info (with directory check): ", data.IsDir)
	fmt.Printf("File info (with system process): %+v\n", data.Sys)
}

func ioRenameFile() {
	err := io.RenameFile("data.json", "data2.json")
	if err != nil {
		println(err.Error())
	}
	println("file renamed")
}

func ioCopyFile() {
	err := io.CopyFile("data.json", "data2.json")
	if err != nil {
		println(err.Error())
	}
	println("file copied")
}

func ioTruncateFile() {
	err := io.TruncateFile("data.json", 10)
	if err != nil {
		println(err.Error())
	}
	println("file truncated to 10 bytes")
}

func ioCleanFile() {
	err := io.CleanFile("data.json")
	if err != nil {
		println(err.Error())
	}

	// or using truncate with 0 bytes
	// err := helper.TruncateFile("data.json", 0)
	// if err != nil {
	// 	println(err.Error())
	// }

	println("file cleaned")
}

func ioWriteToFile() {
	err := io.WriteFile("data.json", []string{"this is data 1", "this is data 2"})
	if err != nil {
		println(err.Error())
	}
	println("file written")
}

func ioAppendToFile() {
	err := io.AppendFile("data.json", []string{"this is data 3", "this is data 4"})
	if err != nil {
		println(err.Error())
	}
	println("file appended")
}

func ioFileObj() {
	type User struct {
		Name string
		Age  int
	}

	exampleFile := "data.json"

	// read file
	data, err := io.ReadFileObj[User](exampleFile)
	if err != nil {
		println(err.Error())
	}
	fmt.Printf("data: %+v\n", data)

	// write file
	updateData := append(data, User{
		Name: fmt.Sprintf("name %d", len(data)+1),
		Age:  len(data) + 1,
	})

	err2 := io.WriteFileObj(exampleFile, updateData)
	if err2 != nil {
		println(err2.Error())
	}
	println("data written")
}

func ioClearConsole() {
	io.ClearConsole()
	println("console cleared")
}

func ioOutputColor() {
	io.OutputColor(shared.DEFAULT_FG, "this is default")
	io.OutputColor(shared.WHITE_FG, "this is white")
	io.OutputColor(shared.RED_FG, "this is red")
	io.OutputColor(shared.GREEN_FG, "this is green")
	io.OutputColor(shared.YELLOW_FG, "this is yellow")
	io.OutputColor(shared.BLUE_FG, "this is blue")
	io.OutputColor(shared.MAGENTA_FG, "this is magenta")
	io.OutputColor(shared.CYAN_FG, "this is cyan")
}

func IOExample() {
	ioCheckDir()
	ioCreateDir()
	ioRemoveDir()
	ioLsDir()
	ioHierarchyDir()
	ioCurrentDir()

	ioCheckFile()
	ioCreateFile()
	ioRemoveFile()
	ioReadFileByLine()
	ioFileInfo()
	ioRenameFile()
	ioCopyFile()
	ioTruncateFile()
	ioCleanFile()
	ioWriteToFile()
	ioAppendToFile()
	ioFileObj()

	ioClearConsole()
	ioOutputColor()
}
