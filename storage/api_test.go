package storage

import (
	"github.com/franchb/gopher-lua-libs/tests"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"

	inspect "github.com/franchb/gopher-lua-libs/inspect"
	time "github.com/franchb/gopher-lua-libs/time"
)

func TestApi(t *testing.T) {
	preload := tests.SeveralPreloadFuncs(
		inspect.Preload,
		time.Preload,
		Preload,
	)
	assert.NoError(t, os.MkdirAll("./test/db/badger/", 0755), "mkdir")
	assert.NotZero(t, tests.RunLuaTestFile(t, preload, "./test/test_api.lua"))
}
