package dir

import "github.com/thuongtruong1009/gouse/io"

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
