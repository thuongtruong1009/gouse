package dir

import "github.com/thuongtruong1009/gouse/io"

func SampleIoCurrentDir() {
	data, err := io.CurrentDir()
	if err != nil {
		println(err.Error())
		return
	}

	println(data)
}
