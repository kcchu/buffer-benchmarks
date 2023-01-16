package flatbuffers

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"log"
	"testing"
)

func BenchmarkFlatBuffers(b *testing.B) {
	var encoded []byte
	b.Run("Encode", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			encoded = Encode()
		}
	})
	b.Run("Decode", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Decode(encoded)
		}
	})
	fmt.Printf("Wire format size = %d\n", len(encoded))

	var buf bytes.Buffer
	zw := gzip.NewWriter(&buf)
	_, err := zw.Write(encoded)
	if err != nil {
		log.Fatalln("Failed to gzip:", err)
	}
	err = zw.Close()
	if err != nil {
		log.Fatalln("Failed to gzip:", err)
	}
	fmt.Printf("Wire format size (gzip)= %d\n", len(buf.Bytes()))
}
