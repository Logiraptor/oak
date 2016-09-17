package main

import (
	"jobtracker/app/tests"

	"github.com/stretchr/testify/assert"

	"testing"
)

func TestStructToFields(t *testing.T) {
	tests.Describe(t, "StructToFields", func(c *tests.Context) {
		c.It("Works for primitive types", func() {
			type Example struct {
				A int
				B string
				C bool
			}
			f := structToFields(Example{})
			assert.EqualValues(t, []Field{
				{Name: "A", Type: Int},
				{Name: "B", Type: String},
				{Name: "C", Type: Bool},
			}, f)
		})
	})
}
