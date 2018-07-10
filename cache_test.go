package golru_test

import (
	"testing"

	"github.com/jaysonesmith/golru"
	"github.com/stretchr/testify/assert"
)

func TestAddItems(t *testing.T) {
	testCases := []struct {
		name        string
		items       golru.Collection
		key         string
		expected    interface{}
		expectedLen int
	}{
		{
			name:        "Empty",
			items:       golru.Collection{},
			key:         "",
			expected:    nil,
			expectedLen: 0,
		},
		{
			name: "Single item",
			items: golru.Collection{
				Items: []golru.Item{
					golru.Item{
						ID:    "foo",
						Type:  "foo_type",
						Dur:   "30",
						Group: "foo_group",
						URL:   "https://foo.com/",
					},
				},
			},
			key: "foo",
			expected: golru.Item{
				ID:    "foo",
				Type:  "foo_type",
				Dur:   "30",
				Group: "foo_group",
				URL:   "https://foo.com/",
			},
			expectedLen: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cache := golru.New(100)

			err := cache.AddItems(tc.items)
			actual, _ := cache.Store.Peek(tc.key)

			assert.Nil(t, err)
			assert.Equal(t, tc.expectedLen, cache.Store.Len())
			assert.Equal(t, tc.expected, actual)
		})
	}
}
