package gzip

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
	"sync"
)

var (
	spBuffer sync.Pool
	spWriter sync.Pool
)

func init() {
	// 公共对象池,更极致的优化可以建多个池
	spBuffer = sync.Pool{New: func() interface{} {
		return new(bytes.Buffer)
	}}
	spWriter = sync.Pool{New: func() interface{} {
		buf := new(bytes.Buffer)
		return gzip.NewWriter(buf)
	}}
}

func GZipDecompress(input []byte) (string, error) {
	buf := spBuffer.Get().(*bytes.Buffer)
	defer func() {
		// 归还buff
		buf.Reset()
		spBuffer.Put(buf)
	}()
	_, err := buf.Write(input)
	if err != nil {
		return "", err
	}
	reader, gzipErr := gzip.NewReader(buf)
	if gzipErr != nil {
		return "", gzipErr
	}
	defer reader.Close()

	result, readErr := ioutil.ReadAll(reader)
	if readErr != nil {
		return "", readErr
	}

	return string(result), nil
}

func GZipCompress(input string) ([]byte, error) {
	buf := spBuffer.Get().(*bytes.Buffer)
	gw := spWriter.Get().(*gzip.Writer)
	gw.Reset(buf)
	defer func() {
		// 归还buff
		buf.Reset()
		spBuffer.Put(buf)
		// 归还Writer
		spWriter.Put(gw)
	}()
	_, err := gw.Write([]byte(input))
	if err != nil {
		return nil, err
	}

	err = gw.Flush()
	if err != nil {
		return nil, err
	}

	err = gw.Close()
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
