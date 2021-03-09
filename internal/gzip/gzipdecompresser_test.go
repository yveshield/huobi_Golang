package gzip

import "testing"

func Test_Decompress_Success(t *testing.T) {
	buf, _ := GZipCompress("huobi")

	result, _ := GZipDecompress(buf)

	expected := "huobi"
	if result != expected {
		t.Errorf("expected: %s, actual: %s", expected, result)
	}
}

func Benchmark_Decompress_Success(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buf, _ := GZipCompress("huobi")

		_, _ = GZipDecompress(buf)
	}
}
