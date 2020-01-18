package gid

import (
	"testing"

	uuid "github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestShortID(t *testing.T) {
	// Standard UUIDs.
	for i := 0; i < 1000; i++ {
		id := uuid.New()
		shortID := GetShortID(id)
		newID, err := GetUUIDFromShortID(shortID)
		assert.NoError(t, err)
		assert.Equal(t, id.String(), newID.String())
	}

	// Nil UUID.
	shortNilID := GetShortID(uuid.Nil)
	nilID, err := GetUUIDFromShortID(shortNilID)
	assert.NoError(t, err)
	assert.Equal(t, uuid.Nil.String(), nilID.String())

	// Broken short ID.
	_, err = GetUUIDFromShortID("abcdefghijklmnopqrstuvwxyz1234567890")
	assert.Error(t, err)
}

func TestCollapsedID(t *testing.T) {
	// Standard UUIDs.
	for i := 0; i < 1000; i++ {
		id := uuid.New()
		collapsedID := GetCollapsedID(id)
		newID, err := GetUUIDFromCollapsedID(collapsedID)
		assert.NoError(t, err)
		assert.Equal(t, id.String(), newID.String())
	}

	// Nil UUID.
	collapsedNilID := GetCollapsedID(uuid.Nil)
	nilID, err := GetUUIDFromCollapsedID(collapsedNilID)
	assert.NoError(t, err)
	assert.Equal(t, uuid.Nil.String(), nilID.String())

	// Broken collapsed ID.
	_, err = GetUUIDFromCollapsedID("abcdefghijklmnopqrstuvwxyz1234567890")
	assert.Error(t, err)
}
