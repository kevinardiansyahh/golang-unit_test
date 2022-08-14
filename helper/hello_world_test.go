package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Kevin")
	}
}

func BenchmarkHelloWorldTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "Kevin",
			request: "Kevin",
		},
		{
			name:    "Ardiansyah",
			request: "Ardiansyah",
		},
		{
			name:    "Hidayat",
			request: "Hidayat",
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

func TestHelloWorldKevin(t *testing.T) {
	result := HelloWorld("Kevin")

	if result != "Hello Kevin" {
		//fail condition
		//panic("Function test HelloWorld gagal") -> bad practice

		//using Fail() -> make fail condition & continuing the process below
		//t.Fail()

		//using Error() -> same as fail, but return log message
		t.Error("Result must be 'Hello Kevin'")
	}

	fmt.Println("TestHelloWorldKevin done")
}

func TestHelloWorldArdi(t *testing.T) {
	result := HelloWorld("Ardi")

	if result != "Hello Ardi" {
		//fail condition
		//panic("Function test HelloWorld gagal") -> bad practice

		//using FailNow() -> make fail condition & terminate process
		//t.FailNow()

		//using Fatal() -> same as failnow, but return log message
		t.Fatal("Result must be 'Hello Ardi'")
	}

	fmt.Println("TestHelloWorldArdi done")
}

func TestMain(m *testing.M) {
	//before
	fmt.Println("BEFORE UNIT TEST")

	m.Run()

	//after
	fmt.Println("AFTER UNIT TEST")
}

func TestSubTest(t *testing.T) {
	t.Run("Kevin", func(t *testing.T) {
		result := HelloWorld("Kevin")
		if result != "Hello Kevin" {
			t.Error("Result must be 'Hello Kevin'")
		}
	})

	t.Run("Ardi", func(t *testing.T) {
		result := HelloWorld("Ardi")
		if result != "Hello Ardi" {
			t.Error("Result must be 'Hello Ardi'")
		}
	})
}

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Kevin",
			request:  "Kevin",
			expected: "Hello Kevin",
		},
		{
			name:     "Ardiansyah",
			request:  "Ardiansyah",
			expected: "Hello Ardiansyah",
		},
		{
			name:     "Hidayat",
			request:  "Hidayat",
			expected: "Hello Hidayat",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}
