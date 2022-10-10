package go_benchmark

import (
	"encoding/json"
	"github.com/json-iterator/go"
	"testing"
)

//
//func BenchmarkJsonParserSmall(b *testing.B) {
//	b.ReportAllocs()
//	paths := [][]string{
//		[]string{"uuid"},
//		[]string{"tz"},
//		[]string{"ua"},
//		[]string{"st"},
//	}
//
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		var data SmallPayload
//
//		jsonparser.EachKey(smallFixture, func(idx int, value []byte, vt jsonparser.ValueType, err error) {
//			switch idx {
//			case 0:
//				data.Uuid, _ = jsonparser.ParseString(value)
//			case 1:
//				v, _ := jsonparser.ParseInt(value)
//				data.Tz = int(v)
//			case 2:
//				data.Ua, _ = jsonparser.ParseString(value)
//			case 3:
//				v, _ := jsonparser.ParseInt(value)
//				data.St = int(v)
//			}
//		}, paths...)
//
//	}
//}
//
//func BenchmarkJsnoiterPullSmall(b *testing.B) {
//	b.ReportAllocs()
//	iter := jsoniter.ParseBytes(jsoniter.ConfigDefault, smallFixture)
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		var data SmallPayload
//		iter.ResetBytes(smallFixture)
//		for field := iter.ReadObject(); field != ""; field = iter.ReadObject() {
//			switch field {
//			case "uuid":
//				data.Uuid = iter.ReadString()
//			case "tz":
//				data.Tz = iter.ReadInt()
//			case "ua":
//				data.Ua = iter.ReadString()
//			case "st":
//				data.St = iter.ReadInt()
//			default:
//				iter.Skip()
//			}
//		}
//	}
//}
func BenchmarkDecodingJsoniterStructSmall(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var data SmallPayload
		jsoniter.Unmarshal(smallFixture, &data)
	}
}

func BenchmarkDecodingJsonStructSmall(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var data SmallPayload
		json.Unmarshal(smallFixture, &data)
	}
}

func BenchmarkEncodingJsoniterStructSmall(b *testing.B) {
	var data SmallPayload
	jsoniter.Unmarshal(smallFixture, &data)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		jsoniter.Marshal(data)
	}
}

/*
   encoding/json
*/

func BenchmarkEncodingJsonStructSmall(b *testing.B) {
	var data SmallPayload
	json.Unmarshal(smallFixture, &data)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		json.Marshal(data)
	}
}

//
//func BenchmarkEasyJsonSmall(b *testing.B) {
//	b.ReportAllocs()
//	for i := 0; i < b.N; i++ {
//		lexer := &jlexer.Lexer{Data: smallFixture}
//		data := new(SmallPayload)
//		data.UnmarshalEasyJSON(lexer)
//	}
//}
