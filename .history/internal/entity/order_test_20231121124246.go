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
	order := Order{ ID: "123" }
	assert.Error(t, order.Validate(), "price must be greater than zero")
}

func Test_If_It_Gets_An_Error_If_Tax_Is_Blank(t *testing.T) {
	order := Order{ ID: "123," }
	assert.Error(t, order.Validate(), "price must be greater than zero")
}