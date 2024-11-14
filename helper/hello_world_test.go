package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// var result string

func TestMain(m *testing.M) {

	fmt.Println("BEFORE RUNNING")
	// result = HelloWorld("Faiz")

	m.Run()

	// result = ""
	fmt.Println("AFTER RUNNING")

}

func TestTableHelloWorld(t *testing.T) {

	data := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Faiz",
			request:  "Faiz",
			expected: "Hello Faiz",
		},
		{
			name:     "Lagi",
			request:  "Lagi",
			expected: "Hello Lagi",
		},
	}

	for _, datas := range data {

		t.Run(datas.name, func(t *testing.T) {
			result := HelloWorld(datas.request)
			require.Equal(t, datas.expected, result)
		})
	}

}
func TestSubTest(t *testing.T) {
	t.Run("Faiz", func(t *testing.T) {
		result := HelloWorld("Faiz")
		require.Equal(t, "Hello Faiz", result, "Harusnya Hello Faiz")
	})
	t.Run("Fajar", func(t *testing.T) {
		result := HelloWorld("Faiz")
		require.Equal(t, "Hello Faiz", result, "Harusnya Hello Faiz")
	})

}

func TestSkip(t *testing.T) {

	if runtime.GOOS == "darwin" {
		t.Skip("gabisa dijalankan")
	}

	fmt.Println("skip done")
}
func TestHelloWorldRequire(t *testing.T) {

	result := HelloWorld("Faiz")
	require.Equal(t, "Hello Faiz", result, "Harusnya Hello Faiz")

	fmt.Println("kode dibawah")

}

func TestHelloWorldAssert(t *testing.T) {

	result := HelloWorld("Faiz")
	assert.Equal(t, "Hello Faiz", result, "Harusnya Hello Faiz")

	fmt.Println("kode dibawah")

}

func TestHelloWorld(t *testing.T) {

	result := HelloWorld("Faiz")

	if result != "Hello Faiz" {
		assert.Equal(t, "Hello Faiz", result, "Harusnya Hello Faiz")
	}

	fmt.Println("kode dibawah")

}

func TestHelloWorldLagi(t *testing.T) {

	result := HelloWorld("lagi")

	if result != "Hello lagi" {
		t.Fatal("Harusnya Hello lagi")
	}

	fmt.Println("hello world lagi done")

}

func BenchmarkTable(b *testing.B) {
	data := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:    "Faiz",
			request: "Faiz",
		},
		{
			name:    "Lagi",
			request: "Lagi",
		},
	}

	for _, datas := range data {

		b.Run(datas.name, func(b *testing.B) {
			HelloWorld(datas.request)
		})
	}
}

func BenchmarkTestingSub(b *testing.B) {
	b.Run("Faiz", func(b *testing.B) {
		HelloWorld("Faiz")
	})
	b.Run("Fajar", func(b *testing.B) {
		HelloWorld("Fajar")
	})
}

func BenchmarkTesting(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Faiz")
	}
}

func BenchmarkTestingLagi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Faiz")
	}
}
