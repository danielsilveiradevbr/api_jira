package helper

import (
	"encoding/hex"
	"fmt"
	"math/rand"
	"strings"
)

func Cripto(action, src, key string) string {
	if src == "" {
		return ""
	}

	var (
		keyLen = len(key) - 1
		keyPos = -1
		offset int
		dest   string
		srcPos int
		srcAsc int
		tmpAsc int
		rng    = 255
	)
	if strings.ToUpper(action) == "E" { // Encrypt
		offset = rand.Intn(rng)
		println(offset)
		dest = fmt.Sprintf("%02X", offset)
		println(dest)
		for srcPos = 0; srcPos < len(src); srcPos++ {
			srcAsc = (int(src[srcPos]) + offset) % 255
			println(srcAsc)
			if keyPos < keyLen-1 {
				keyPos++
			} else {
				keyPos = 0
			}

			srcAsc ^= int(key[keyPos])
			println(srcAsc)
			dest += fmt.Sprintf("%02X", srcAsc)
			println(dest)
			offset = srcAsc
		}
	} else if strings.ToUpper(action) == "D" { // Decrypt
		offset = parseHex(src[:2])
		srcPos = 2

		for srcPos < len(src) {
			srcAsc = parseHex(src[srcPos : srcPos+2])
			if keyPos < keyLen-1 {
				keyPos++
			} else {
				keyPos = 0
			}

			tmpAsc = srcAsc ^ int(key[keyPos])
			if tmpAsc <= offset {
				tmpAsc = 255 + tmpAsc - offset
			} else {
				tmpAsc -= offset
			}

			dest += string(tmpAsc)
			offset = srcAsc
			srcPos += 2
		}
	}
	return dest
}

func parseHex(hexStr string) int {
	val, err := hex.DecodeString(hexStr)
	if err != nil || len(val) == 0 {
		return 0
	}
	return int(val[0])
}
