package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIfItGetsAnErrorIfIDIsBlank(t *testing.T) {
	order := Order{}
	// if order.Validate() == nil {
	// 	t.Error("ID is required")
	// }
	assert.Error(t, order.Validate(), "ID is required")
}