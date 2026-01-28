package json

import (
	"github.com/franchb/gopher-lua-libs/strings"
	"github.com/franchb/gopher-lua-libs/tests"
	"github.com/stretchr/testify/assert"
	"testing"

	inspect "github.com/franchb/gopher-lua-libs/inspect"
)

func TestApi(t *testing.T) {
	preload := tests.SeveralPreloadFuncs(
		Preload,
		inspect.Preload,
		strings.Preload,
	)
	assert.NotZero(t, tests.RunLuaTestFile(t, preload, "./test/test_api.lua"))
}
