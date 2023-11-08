package main

import (
	"C"
	"log"
)
import (
	"bytes"
	"fmt"
	"image/png"
	"os"
	"unsafe"

	diff "github.com/olegfedoseev/image-diff"
)

func main() {}

//export helloWorld
func helloWorld() {
	log.Println("hello World")
}

//export compare
func compare(a_data *C.uchar, a_size C.int, b_data *C.uchar, b_size C.int) *C.char {

	a := (*[1 << 30]byte)(unsafe.Pointer(a_data))[:a_size:a_size]
	b := (*[1 << 30]byte)(unsafe.Pointer(b_data))[:b_size:b_size]

	a_img, err := png.Decode(bytes.NewReader(a))
	if err != nil {
		return C.CString(fmt.Sprintf("err: a_img decoding: %s", err))
	}
	b_img, err := png.Decode(bytes.NewReader(b))
	if err != nil {
		return C.CString(fmt.Sprintf("err: b_img decoding: %s", err))
	}

	diff, percent, err := diff.CompareImages(a_img, b_img)
	if err != nil {
		return C.CString(fmt.Sprintf("err: diffing: %s", err))
	}

	f, err := os.Create("diff.png")
	if err != nil {
		return C.CString(fmt.Sprintf("err: error opening result file: %s", err))
	}
	png.Encode(f, diff)
	f.Close()

	return C.CString(fmt.Sprintf("res: %f", percent))
}
