package src

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	fmt.Println("INITIAL - Everything before doing unit testing")
	m.Run()
	fmt.Println("END - Unit Test")
}

func TestTableHelloWorld(t *testing.T) {
	testScenario := []struct {
		name    string
		request string
		expect  string
	}{
		{
			name:    "Should Return Hi Andre",
			request: "Andre",
			expect:  "Hi Andre",
		},
		{
			name:    "Should be Return Not Equal",
			request: "Budi",
			expect:  "Hi Andre",
		},
	}

	for _, test := range testScenario {
		t.Run(test.name, func(t *testing.T) {
			result, _ := HelloWorld(test.request)
			assert.Equal(t, test.expect, result)
		})
	}

}

func TestHelloWorld(t *testing.T) {
	result, _ := HelloWorld("Andre")
	assert.Equal(t, "Hi Andre", result, "Should Be Return Hi Andre")
}
