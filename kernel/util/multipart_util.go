package util

import (
	"bytes"
	"path"
)

const (
	boundaryPre       = "\r\n--"
	boundaryPreEnd    = "\r\n"
	boundaryPreMinLen = len(boundaryPre) + len(boundaryPreEnd)
	boundaryEnd       = "--\r\n"
	boundaryEndMinLen = len(boundaryPre) + len(boundaryEnd)
)

func GetEntryBoundary(boundary string) []byte {
	bb := &bytes.Buffer{}
	bb.Grow(boundaryPreMinLen + len(boundary))
	bb.WriteString(boundaryPre)
	bb.WriteString(boundary)
	bb.WriteString(boundaryPreEnd)
	return bb.Bytes()
}

func GetEndBoundary(boundary string) []byte {
	return []byte("\r\n--" + boundary + "--\r\n")
}

const (
	cd               = "Content-Disposition:form-data;name=\""
	ct               = "\"\r\nContent-Type:"
	tp               = "text/plain\r\n\r\n"
	minLenTextEntry  = len(cd) + len(ct) + len(tp)
	aos              = "application/octet-stream\r\n\r\n"
	filenameAttr     = "\";filename=\""
	minLentFileEntry = len(cd) + len(filenameAttr) + len(ct) + len(aos)
)

func GetTextEntry(fieldName, fieldValue string) []byte {
	bb := &bytes.Buffer{}
	bb.Grow(minLenTextEntry + len(fieldName) + len(fieldValue))
	bb.WriteString(cd)
	bb.WriteString(fieldName)
	bb.WriteString(ct)
	bb.WriteString(tp)
	bb.WriteString(fieldValue)
	return bb.Bytes()
}

func GetFileEntry(fieldName, filePath string) []byte {
	fileName, _ := path.Split(filePath)
	bb := &bytes.Buffer{}
	bb.Grow(minLentFileEntry + len(fieldName) + len(fileName))
	bb.WriteString(cd)
	bb.WriteString(fieldName)
	bb.WriteString(filenameAttr)
	bb.WriteString(fileName)
	bb.WriteString(ct)
	bb.WriteString(aos)
	return bb.Bytes()
}
