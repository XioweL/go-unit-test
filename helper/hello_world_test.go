package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
