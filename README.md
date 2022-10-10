# 压测实例

```shell
goos: darwin
goarch: amd64
pkg: json-benchmark
cpu: Intel(R) Core(TM) i7-1068NG7 CPU @ 2.30GHz
BenchmarkDecodingJsonLarge-8             	    7570	    162292 ns/op	     592 B/op	      16 allocs/op
BenchmarkDecodingJsoniterLarge-8         	   16003	     78893 ns/op	   12483 B/op	    1126 allocs/op
BenchmarkEncodingJsonLarge-8             	 1393060	       875.1 ns/op	     128 B/op	       1 allocs/op
BenchmarkEncodingJsoniterLarge-8         	 2955939	       426.7 ns/op	     136 B/op	       2 allocs/op
BenchmarkDecodeStdStructMedium-8         	   74409	     17190 ns/op	     496 B/op	      11 allocs/op
BenchmarkDecodeJsoniterStructMedium-8    	  244887	      5003 ns/op	     384 B/op	      41 allocs/op
BenchmarkEncodeStdStructMedium-8         	 1659696	       688.6 ns/op	     216 B/op	       2 allocs/op
BenchmarkEncodeJsoniterStructMedium-8    	 1904721	       588.9 ns/op	     216 B/op	       2 allocs/op
BenchmarkDecodingJsoniterStructSmall-8   	 1730190	       628.3 ns/op	     176 B/op	       3 allocs/op
BenchmarkDecodingJsonStructSmall-8       	  526699	      2192 ns/op	     384 B/op	       7 allocs/op
BenchmarkEncodingJsoniterStructSmall-8   	 3632998	       326.8 ns/op	     192 B/op	       2 allocs/op
BenchmarkEncodingJsonStructSmall-8       	 3151880	       382.4 ns/op	     192 B/op	       2 allocs/op
PASS
ok  	json-benchmark	20.600s
```

总结（基础用法情况下）：

1. jsoniter和std json在 encode情况下，性能基本差不多（大字节情况下jsoniter是std json的两倍）
2. 在decode情况下，jsoniter性能基本是std json的3倍左右
