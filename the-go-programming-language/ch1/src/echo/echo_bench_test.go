package echo_test

import (
	"echo"
	"fmt"
	"testing"
)

var fixture = func() []string {
	var result []string
	for i := 0; i < 1000; i++ {
		result = append(result, "asdfasdfasdf")
	}
	return result
}()

func BenchmarkEcho1(b *testing.B) {
	fmt.Println()
	for i := 0; i < b.N; i++ {
		echo.Echo1(fixture)
	}
}

func BenchmarkEcho2(b *testing.B) {
	fmt.Println()
	for i := 0; i < b.N; i++ {
		echo.Echo2(fixture)
	}
}
