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

	defer func() {
		if err := recover(); err != nil { //注意必须要判断
			fmt.Println(err)
		}
	}() //用来调用此匿名函数
	sourceFile, _ := ioutil.ReadFile("./demo.png")

	imageOut := unsafe.Pointer(&sourceFile[0])

	width := C.uint(0)
	height := C.uint(0)

	datalen := C.ulong(len(sourceFile))

	var imageIn *C.uchar = nil

	loadStatus := C.lodepng_decode32(&imageIn, &width, &height, (*C.uchar)(imageOut), datalen)

	if loadStatus > 0 {
		fmt.Println("Error:", C.GoString(C.lodepng_error_text(loadStatus)))
		os.Exit(99)
	}

	attrHandle := C.liq_attr_create()
	image := C.liq_image_create_rgba(attrHandle, unsafe.Pointer(imageIn), C.int(width), C.int(height), 0)

	var imageResult *C.liq_result
	if C.liq_image_quantize(image, attrHandle, &imageResult) != C.LIQ_OK {
		fmt.Println("Error:", "liq_image_quantize Failed")
		os.Exit(99)
	}

	C.liq_set_dithering_level(imageResult, 1.0)

	pixelsSize := width * height
	raw8bitPixels := make([]byte, pixelsSize)
	rawPoint := unsafe.Pointer(&raw8bitPixels[0])
	C.liq_write_remapped_image(imageResult, image, rawPoint, (C.ulong)(pixelsSize))
	palette := C.liq_get_palette(imageResult)

	var state C.LodePNGState
	C.lodepng_state_init(&state)
	state.info_raw.colortype = C.LCT_PALETTE
	state.info_raw.bitdepth = 8
	state.info_png.color.colortype = C.LCT_PALETTE
	state.info_png.color.bitdepth = 8

	for i := 0; i < int(palette.count); i++ {
		C.lodepng_palette_add(&state.info_png.color, palette.entries[i].r, palette.entries[i].g, palette.entries[i].b, palette.entries[i].a)
		C.lodepng_palette_add(&state.info_raw, palette.entries[i].r, palette.entries[i].g, palette.entries[i].b, palette.entries[i].a)
	}

	var imageOutData *C.uchar
	var size uint64
	outStatus := C.lodepng_encode(&imageOutData, (*C.ulong)(&size), (*C.uchar)(rawPoint), width, height, &state)
	if outStatus > 0 {
		fmt.Println("Can't encode image: %s\n", C.lodepng_error_text(outStatus))
		os.Exit(99)
	}

	err := ioutil.WriteFile("./demo1.png", ([]byte)(string(C.GoBytes(unsafe.Pointer(imageOutData), C.int(size)))), 0666)
	if err != nil {
		fmt.Println(err)
	}

}
