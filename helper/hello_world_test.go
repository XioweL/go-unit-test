package helper

import (
	"fmt"
	"testing"
)

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
