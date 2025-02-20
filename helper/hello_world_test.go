package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		excepted string
	}{
		{
			name:     "HelloWorld(Randi)",
			request:  "Randi",
			excepted: "Hello Randi",
		},
		{
			name:     "HelloWorld(Febriadi)",
			request:  "Febriadi",
			excepted: "Hello Febriadi",
		},
		{
			name:     "HelloWorld(Fauzan)",
			request:  "Fauzan",
			excepted: "Hello Fauzan",
		},
		{
			name:     "HelloWorld(Fatwa)",
			request:  "Fatwa",
			excepted: "Hello Fatwa",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			assert.Equal(t, test.excepted, result)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Randi", func(t *testing.T) {
		result := HelloWorld("Randi")
		require.Equal(t, "Hello Randi", result, "Result must be 'Hello Randi'")
	})
	t.Run("Febriadi", func(t *testing.T) {
		result := HelloWorld("Febriadi")
		require.Equal(t, "Hello Febriadi", result, "Result must be 'Hello Febriadi'")
	})
}

func TestMain(m *testing.M) {
	// before
	fmt.Println("BEFORE UNIT TEST")

	m.Run()

	// after
	fmt.Println("AFTER UNIT TEST")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Can not run on Windows")
	}
	result := HelloWorld("Randi")
	require.Equal(t, "Hello Randi", result, "Result must be 'Hello Randi'")
}

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
