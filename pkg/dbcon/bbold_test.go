package dbcon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBboldConnect(t *testing.T) {
	db, err := BboltConnect("../../pota/pota.db")
	defer db.Close()
	assert.NotEmpty(t, db)
	assert.NoError(t, err)
}
