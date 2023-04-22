package tests

import (
	"testing"
	"tsql/internal"
)

func BenchmarkCheckTypeMatch(b *testing.B) {
	set, err := internal.Init("../../.test_folder/")
	if err != nil {
		b.Fatalf("unexpected error: %v", err)
	}

	f := internal.Field{
		Name: "age",
		Type: "int",
	}
	value := 42

	for i := 0; i < b.N; i++ {
		_, err := set.CheckTypeMatch(f, value)
		if err != nil {
			b.Fatalf("unexpected error: %v", err)
		}
	}
}
