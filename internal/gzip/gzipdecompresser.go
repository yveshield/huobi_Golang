package gzip

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
	"sync"
)

var (
	spBuffer sync.Pool
)

func init() {
	// 公共对象池,更极致的优化可以建多个池
	spBuffer = sync.Pool{New: func() interface{} {
		return new(bytes.Buffer)
	}}
}

func GZipDecompress(input []byte) (string, error) {
	buf := spBuffer.Get().(*bytes.Buffer)
	defer func() {
		// 归还buff
		buf.Reset()
		spBuffer.Put(buf)
	}()
	_, berr := buf.Write(input)
	if berr != nil {
		return "", berr
	}
	// buf := bytes.NewBuffer(input)
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
	var buf bytes.Buffer
	gz := gzip.NewWriter(&buf)

	_, err := gz.Write([]byte(input))
	if err != nil {
		return nil, err
	}

	err = gz.Flush()
	if err != nil {
		return nil, err
	}

	err = gz.Close()
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
