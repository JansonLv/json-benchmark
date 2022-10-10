package go_benchmark

import (
	"encoding/json"
	jsoniter "github.com/json-iterator/go"
	"testing"
)

//func BenchmarkJsonParserLarge(b *testing.B) {
//	b.ReportAllocs()
//	for i := 0; i < b.N; i++ {
//		count := 0
//		jsonparser.ArrayEach(largeFixture, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
//			count++
//		}, "topics", "topics")
//	}
//}

func BenchmarkDecodingJsonLarge(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		payload := &LargePayload{}
		json.Unmarshal(largeFixture, payload)
	}
}

func BenchmarkDecodingJsoniterLarge(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		payload := &LargePayload{}
		jsoniter.Unmarshal(largeFixture, payload)
	}
}

func BenchmarkEncodingJsonLarge(b *testing.B) {
	payload := &LargePayload{}
	json.Unmarshal(largeFixture, payload)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		json.Marshal(payload)
	}
}

func BenchmarkEncodingJsoniterLarge(b *testing.B) {
	payload := &LargePayload{}
	jsoniter.Unmarshal(largeFixture, payload)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		jsoniter.Marshal(payload)
	}
}
