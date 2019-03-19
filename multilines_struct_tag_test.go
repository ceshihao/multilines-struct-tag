package structtag

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStructTagGet(t *testing.T) {
	cases := []struct {
		structTag reflect.StructTag
		tagName   string
		valid     bool
		expected  string
	}{
		// Valid Tags
		{
			`tag1:"value1" 
			tag2:"value2"`,
			"tag2",
			true,
			"value2",
		},
		{
			`tag1:"value1" 
			tag2:"The value of
			tag1 is value2"`,
			"tag2",
			true,
			"The value of tag1 is value2",
		},
		{
			`tag1:"\\"`,
			"tag1",
			true,
			"\\",
		},
		// Invalid Tags
		{
			``,
			"tag1",
			false,
			"",
		},
		{
			` `,
			"tag1",
			false,
			"",
		},
		{
			`tag1:"`,
			"tag1",
			false,
			"",
		},
		{
			`tag1:a`,
			"tag1",
			false,
			"",
		},
	}

	for i, c := range cases {
		res, ok := Lookup(c.structTag, c.tagName)
		assert.Equal(t, c.valid, ok, "%d case fail: structTag is %s and tagName is %s", i, c.structTag, c.tagName)
		assert.Equal(t, c.expected, res, "%d case fail: structTag is %s and tagName is %s", i, c.structTag, c.tagName)
	}
}
