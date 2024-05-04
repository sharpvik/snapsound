package main

import (
	"encoding/binary"
	"errors"
	"image"
	"image/png"
	"io"
	"log"
	"math"
	"os"
)

func encodeBytesAsImage(bytes []byte) *image.RGBA {
	log.Println("ORIG LEN:", len(bytes))
	binLength := make([]byte, 8)
	binary.BigEndian.PutUint64(binLength, uint64(len(bytes)))
	bytes = append(binLength, bytes...)

	sideLength := int(math.Ceil(math.Sqrt(float64(len(bytes)) / 4)))

	img := image.NewRGBA(image.Rect(0, 0, sideLength, sideLength))
	log.Println("LEN+MP3 LEN:", len(bytes))
	copy(img.Pix, bytes)

	return img
}

func originalBytes(r io.Reader) (bs []byte, err error) {
	img, err := png.Decode(r)
	if err != nil {
		return nil, err
	}

	nrgba, ok := img.(*image.NRGBA)
	if !ok {
		return nil, errors.New("failed to revert to original: expected NRGBA")
	}

	originalLength := binary.BigEndian.Uint64(nrgba.Pix[:8])
	log.Println("ORIG LEN:", originalLength)
	return nrgba.Pix[8 : originalLength+8], nil
}

func saveImage(name string, img *image.RGBA) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()
	e := png.Encoder{
		CompressionLevel: png.NoCompression,
	}
	return e.Encode(f, img)
}
