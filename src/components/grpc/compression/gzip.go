package compression

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
)

type GzipEncodeDecoder struct{}

func (g *GzipEncodeDecoder) Encode(v []byte) ([]byte, error) {
	var buf bytes.Buffer
	zw := gzip.NewWriter(&buf)
	if _, err := zw.Write(v); err != nil {
		return nil, err
	}
	if err := zw.Close(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (g *GzipEncodeDecoder) Decode(b []byte) ([]byte, error) {
	zr, err := gzip.NewReader(bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}
	buf, err := ioutil.ReadAll(zr)
	if err != nil {
		return nil, err
	}
	if err := zr.Close(); err != nil {
		return nil, err
	}
	return buf, nil
}
