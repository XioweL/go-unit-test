package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSubTest(t *testing.T) {
	t.Run("Ferdi", func(t *testing.T) {
		result := HelloWorld("Ferdi")
		require.Equal(t, "Hello Ferdi", result)
	})
	t.Run("Alvanza", func(t *testing.T) {
		result := HelloWorld("Alvanza")
		require.Equal(t, "Hello Alvanza", result)
	})
}
func TestMain(m *testing.M) {
	// before
	fmt.Println("Sebelum Unit Test")
	m.Run()
	fmt.Println("After Unit Test")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Unit test tidak bisa jalan di Windows")
	}
	result := HelloWorld("Ferdi")
	require.Equal(t, "Hello Ferdi", result)
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Ferdi")
	require.Equal(t, "Hello Ferdi", result, "Result must be 'Hello Ferdi'")
	fmt.Println("TestHelloWorld with Require Done")
}
func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Ferdi")
	assert.Equal(t, "Hello Ferdi", result, "Result must be 'Hello Ferdi'")
	fmt.Println("TestHelloWorld with Assert Done")
}
func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Ferdi")
	if result != "Hello Ferdi" {
		t.Error("Result is not Hello Ferdi")
	}
	fmt.Println("Dieksekusi")
}
func TestHelloWorldFatal(t *testing.T) {
	result := HelloWorld("Ferdi")
	if result != "Hello Ferdi" {
		t.Fatal("Result must be Hello Ferdi")
	}
	fmt.Println("Tidak Dieksekusi")
}
