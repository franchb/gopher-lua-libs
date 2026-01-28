package crypto

import (
	"testing"

	"github.com/franchb/gopher-lua-libs/filepath"
	"github.com/franchb/gopher-lua-libs/hex"
	"github.com/franchb/gopher-lua-libs/tests"
	"github.com/stretchr/testify/assert"
)

func TestApi(t *testing.T) {
	preload := tests.SeveralPreloadFuncs(
		Preload,
		filepath.Preload,
		hex.Preload,
	)
	assert.NotZero(t, tests.RunLuaTestFile(t, preload, "./test/test_api.lua"))
}
