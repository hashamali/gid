package gid

import (
	"errors"
	"fmt"
	"math/big"
	"strings"

	uuid "github.com/google/uuid"
)

// GetShortID will convert a UUID to a base62 encoded short ID.
func GetShortID(id uuid.UUID) string {
	var i big.Int
	i.SetString(strings.Replace(id.String(), "-", "", 4), 16)
	shortID := i.Text(62)
	for len(shortID) < 22 {
		shortID = "0" + shortID
	}

	return shortID
}

// GetUUIDFromShortID will convert a base62 encoded short ID to a UUID.
func GetUUIDFromShortID(shortID string) (uuid.UUID, error) {
	var i big.Int
	i.SetString(shortID, 62)
	e := i.Text(16)

	for len(e) < 32 {
		e = "0" + e
	}

	return GetUUIDFromCollapsedID(e)
}

// GetCollapsedID converts a UUID object to a string without '-'.
func GetCollapsedID(id uuid.UUID) string {
	return strings.Replace(id.String(), "-", "", -1)
}

// GetUUIDFromCollapsedID will convert a 32 character alphanumeric string into a UUID.
func GetUUIDFromCollapsedID(collapsedID string) (uuid.UUID, error) {
	if len(collapsedID) != 32 {
		return uuid.Nil, errors.New("invalid length")
	}

	uuid1 := collapsedID[0:8]
	uuid2 := collapsedID[8:12]
	uuid3 := collapsedID[12:16]
	uuid4 := collapsedID[16:20]
	uuid5 := collapsedID[20:32]

	uuidString := fmt.Sprintf("%s-%s-%s-%s-%s", uuid1, uuid2, uuid3, uuid4, uuid5)
	return uuid.Parse(uuidString)
}
