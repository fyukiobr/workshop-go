package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {
	user := User{Name: "Fernando", Age: 31}
	assert.Equal(t, "Name: Fernando", user.String(), "Deveria ser igual")
}
