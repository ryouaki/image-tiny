package main

/*
#cgo CFLAGS: -I.
#cgo LDFLAGS: -L./ -limagequant
#cgo LDFLAGS: -L./ -llodepng

#include "libimagequant.h"
#include "./lodepng.h"
*/
import "C"
import (
	"fmt"
	"io/ioutil"
	"os"
	"unsafe"
)

func main() {
	sourceFile, _ := ioutil.ReadFile("./demo.png")

	imageOut := unsafe.Pointer(&sourceFile[0])

	width := C.uint(0)
	height := C.uint(0)

	datalen := C.ulong(len(sourceFile))

	var imageDate []byte
	imageIn := unsafe.Pointer(&imageDate)

	loadStatus := C.lodepng_decode32((**C.uchar)(imageIn), &width, &height, (*C.uchar)(imageOut), datalen)

	if loadStatus > 0 {
		fmt.Println("Error:", C.GoString(C.lodepng_error_text(loadStatus)))
		os.Exit(99)
	}

	attrHandle := C.liq_attr_create()
	image := C.liq_image_create_rgba(attrHandle, imageIn, C.int(width), C.int(height), 0)

	var imageResult *C.liq_result

	if C.liq_image_quantize(image, attrHandle, &imageResult) != C.LIQ_OK {
		fmt.Println("Error:", "liq_image_quantize Failed")
		os.Exit(99)
	}

	C.lig_set_dithering_level(imageResult, 1.0)

	fmt.Println(loadStatus, width, height)
}
