package file

import "github.com/thuongtruong1009/gouse/io"

func SampleIoRemoveFile() {
	err := io.RemoveFile("data.json")
	if err != nil {
		println(err.Error())
	}
	println("file removed")
}
