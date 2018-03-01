package helper_test

import (
	"flatsharing/core/helper"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenUlid(t *testing.T) {
	require.NotEqual(t, helper.GenUlid(), helper.GenUlid())
}
