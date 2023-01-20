package helper

import (
	"github.com/gabriel-vasile/mimetype"
)

func GetFileExt(byt []byte) string {
	return mimetype.Detect(byt).Extension()
}

func CheckImageExt(jsnbyt []byte) bool {
	if GetFileExt(jsnbyt) == ".png" || GetFileExt(jsnbyt) == ".gif" || GetFileExt(jsnbyt) == ".jpg" || GetFileExt(jsnbyt) == ".svg" || GetFileExt(jsnbyt) == ".tiff" || GetFileExt(jsnbyt) == ".bmp" {
		return true
	}
	return false
}

func CheckDocumentExt(jsnbyt []byte) bool {
	if GetFileExt(jsnbyt) == ".doc" || GetFileExt(jsnbyt) == ".docx" || GetFileExt(jsnbyt) == ".odt" || GetFileExt(jsnbyt) == ".rtf" || GetFileExt(jsnbyt) == ".tiff" || GetFileExt(jsnbyt) == ".bmp" {
		return true
	}
	return false
}
