package tests

import (
	"testing"
	"tsql/internal"
)

func BenchmarkInitSetting(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := internal.Init("../../.test_folder/")
		if err != nil {
			b.Fatalf("unexpected error: %v", err)
		}
	}

}
