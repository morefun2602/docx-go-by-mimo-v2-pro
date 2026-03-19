package image

import (
	"encoding/binary"
	"fmt"
	"io"
)

func DecodeJPEGSize(r io.Reader) (width, height int, err error) {
	var buf [2]byte
	for {
		if _, err := io.ReadFull(r, buf[:]); err != nil {
			return 0, 0, err
		}
		if buf[0] != 0xFF {
			continue
		}
		marker := buf[1]
		if marker == 0xC0 || marker == 0xC2 {
			break
		}
		if marker >= 0xC0 && marker <= 0xCF && marker != 0xC8 {
			continue
		}
		if marker == 0xD8 || marker == 0xD9 {
			continue
		}
		var lenBuf [2]byte
		if _, err := io.ReadFull(r, lenBuf[:]); err != nil {
			return 0, 0, err
		}
		length := binary.BigEndian.Uint16(lenBuf[:])
		if length < 2 {
			return 0, 0, fmt.Errorf("invalid JPEG segment length")
		}
		remaining := int(length) - 2
		if remaining > 0 {
			if _, err := io.CopyN(io.Discard, r, int64(remaining)); err != nil {
				return 0, 0, err
			}
		}
	}

	var lenBuf [2]byte
	if _, err := io.ReadFull(r, lenBuf[:]); err != nil {
		return 0, 0, err
	}
	length := binary.BigEndian.Uint16(lenBuf[:])
	if length < 2 {
		return 0, 0, fmt.Errorf("invalid JPEG segment length")
	}

	remaining := int(length) - 2
	var skipBuf [1]byte
	for i := 0; i < remaining; i++ {
		if _, err := io.ReadFull(r, skipBuf[:]); err != nil {
			return 0, 0, err
		}
	}

	var sizeBuf [5]byte
	if _, err := io.ReadFull(r, sizeBuf[:]); err != nil {
		return 0, 0, err
	}

	height = int(binary.BigEndian.Uint16(sizeBuf[1:3]))
	width = int(binary.BigEndian.Uint16(sizeBuf[3:5]))

	return width, height, nil
}

func DecodePNGSize(r io.Reader) (width, height int, err error) {
	var header [8]byte
	if _, err := io.ReadFull(r, header[:]); err != nil {
		return 0, 0, err
	}

	var ihdrBuf [25]byte
	if _, err := io.ReadFull(r, ihdrBuf[:]); err != nil {
		return 0, 0, err
	}

	width = int(binary.BigEndian.Uint32(ihdrBuf[4:8]))
	height = int(binary.BigEndian.Uint32(ihdrBuf[8:12]))

	return width, height, nil
}

func DecodeGIFSize(r io.Reader) (width, height int, err error) {
	var header [10]byte
	if _, err := io.ReadFull(r, header[:]); err != nil {
		return 0, 0, err
	}

	width = int(binary.LittleEndian.Uint16(header[6:8]))
	height = int(binary.LittleEndian.Uint16(header[8:10]))

	return width, height, nil
}

func DecodeBMPSize(r io.Reader) (width, height int, err error) {
	var header [26]byte
	if _, err := io.ReadFull(r, header[:]); err != nil {
		return 0, 0, err
	}

	width = int(int32(binary.LittleEndian.Uint32(header[18:22])))
	height = int(int32(binary.LittleEndian.Uint32(header[22:26])))
	if height < 0 {
		height = -height
	}

	return width, height, nil
}

func DecodeTIFFSize(r io.Reader) (width, height int, err error) {
	var header [8]byte
	if _, err := io.ReadFull(r, header[:]); err != nil {
		return 0, 0, err
	}

	isLittleEndian := header[0] == 'I'
	var byteOrder binary.ByteOrder
	if isLittleEndian {
		byteOrder = binary.LittleEndian
	} else {
		byteOrder = binary.BigEndian
	}

	offset := byteOrder.Uint32(header[4:8])

	if _, err := io.CopyN(io.Discard, r, int64(offset)-8); err != nil {
		return 0, 0, err
	}

	var numEntriesBuf [2]byte
	if _, err := io.ReadFull(r, numEntriesBuf[:]); err != nil {
		return 0, 0, err
	}
	numEntries := byteOrder.Uint16(numEntriesBuf[:])

	var tagBuf [12]byte
	for i := uint16(0); i < numEntries; i++ {
		if _, err := io.ReadFull(r, tagBuf[:]); err != nil {
			return 0, 0, err
		}
		tagID := byteOrder.Uint16(tagBuf[0:2])
		value := byteOrder.Uint32(tagBuf[8:12])

		if tagID == 256 {
			width = int(value)
		} else if tagID == 257 {
			height = int(value)
		}
	}

	return width, height, nil
}
