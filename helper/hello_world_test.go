package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Randi")
	require.Equal(t, "Hello Randi", result, "Result must be 'Hello Randi'")
	fmt.Println("TestHelloWorld with Require Done")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Randi")
	assert.Equal(t, "Hello Randi", result, "Result must be 'Hello Randi'")
	fmt.Println("TestHelloWorld with Assert Done")
}

func TestHelloWorldRandi(t *testing.T) {
	result := HelloWorld("Randi")
	if result != "Hello Randi" {
		// error
		t.Error("Result must be 'Hello Randi'")
	}

	fmt.Println("TestHelloWorldRandiDone")
}

func TestHelloWorldFebriadi(t *testing.T) {
	result := HelloWorld("Febriadi")
	if result != "Hello Febriadi" {
		// error
		t.Fatal("Result must be 'Hello Febriadi'")
	}

	fmt.Println("TestHelloWorldFebriadiDone")
}
