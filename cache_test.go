package golru_test

import (
	"testing"

	"github.com/jaysonesmith/golru"
	"github.com/stretchr/testify/assert"
)

func TestAddToCache(t *testing.T) {
	testCases := []struct {
		name, key, value string
		expected         interface{}
	}{
		{name: "Empty", key: "", value: "", expected: nil},
		{name: "A", key: "foo", value: "bar", expected: ""},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cache := golru.New()

			err := cache.Add(tc.key, tc.value)
			actual, _ := cache.Store.Peek(tc.key)

			assert.Nil(t, err)
			assert.Equal(t, tc.value, actual)
		})
	}
}
