package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_If_It_Gets_An_Error_If_ID_Is_Blank(t *testing.T) {
	order := Order{}
	// if order.Validate() == nil {
	// 	t.Error("ID is required")
	// }
	assert.Error(t, order.Validate(), "ID is required")
}

func Test_If_It_Gets_An_Error_If_Price_Is_Blank(t *testing.T) {
	order := Order{}
	assert.Error(t, order.Validate(), "")
}