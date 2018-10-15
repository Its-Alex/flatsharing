package helper_test

import (
	"testing"

	"github.com/Its-Alex/flatsharing/src/core/helper"

	"github.com/stretchr/testify/require"
)

func TestGenUlid(t *testing.T) {
	require.NotEqual(t, helper.GenUlid(), helper.GenUlid())
}
