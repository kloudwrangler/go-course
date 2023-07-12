package hello

import (
	"os"
	"testing"
)

func TestGreet(t *testing.T) {
    name := "John"
    expected := loadGoldenFile(t, "testdata/greet.golden")
    actual := Greet(name)
    if actual != expected {
        t.Errorf("Greet(%s) = %s, want %s", name, actual, expected)
    }
}

func loadGoldenFile(t *testing.T, filename string) string {
    golden, err := os.ReadFile(filename)
    if err != nil {
        t.Fatalf("failed to read golden file %s: %s", filename, err)
    }
    return string(golden)
}