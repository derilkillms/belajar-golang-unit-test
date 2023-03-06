package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Deril")
	}
}

func BenchmarkHelloWorldMDeril(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("MDeril")
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("Deril", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Deril")
		}
	})
	b.Run("MDeril", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("MDeril")
		}
	})
}

func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "Deril",
			request: "Deril",
		},
		{
			name:    "Pranata",
			request: "Pranata",
		},
		{
			name:    "MuhammadDerilAkbarPranata",
			request: "MuhammadDerilAkbarPranata",
		},
	}
	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}

func TestMain(m *testing.M) {
	//before
	fmt.Println("Before unit test")
	m.Run()
	//after
	fmt.Println("After Unit Test")
}

func TestSubTest(t *testing.T) {
	t.Run("Deril", func(t *testing.T) {
		result := HelloWorld("Deril")
		require.Equal(t, "Hello Deril", result, "Result must be Hello Deril")

	})
	t.Run("MDeril", func(t *testing.T) {
		result := HelloWorld("MDeril")
		require.Equal(t, "Hello MDeril", result, "Result must be Hello MDeril")

	})
}

func TestHelloWorldTable(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Deril",
			request:  "Deril",
			expected: "Hello Deril",
		},
		{
			name:     "MDeril",
			request:  "MDeril",
			expected: "Hello MDeril",
		},
		{
			name:     "Joko",
			request:  "Joko",
			expected: "Hello Joko",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("can not run on Mac Os")
	}
	result := HelloWorld("Deril")
	require.Equal(t, "Hello Deril", result, "Result must be Hello Deril")

}

// assert manggil Fail()
func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Deril")
	assert.Equal(t, "Hello Deril", result, "Result must be Hello Deril")
	fmt.Println("test hello world with assert")
}

// require manggin FailNow()
func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Deril")
	require.Equal(t, "Hello Deril", result, "Result must be Hello Deril")
	fmt.Println("test hello world with assert")
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Deril")
	if result != "Hello Deril" {
		//error
		//panic
		// t.Fail()
		t.Error("result must be hello deril")
	}

	fmt.Println("Test deril done")
}

func TestHelloWorldDeril(t *testing.T) {
	result := HelloWorld("MDeril")
	if result != "Hello MDeril" {
		// error
		//panic
		t.Fatal("result must be hello mderil")
		// t.FailNow()
	}
	fmt.Println("test mderil done")
}
