package entity

import (
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/assert"
)

func TestIfItGetsAnErrorIfIDIsBlank(t *testing.T) {
	order := Order{}
	// if order.Validate() == nil {
	// 	t.Error("ID is required")
	// }
	assert.Error
}