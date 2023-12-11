package common

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ID entity ID
type ID = uuid.UUID

// NewID create a new entity ID
func NewID() ID {
	return ID(uuid.New())
}

// StringToID convert a string to an entity ID
func StringToID(s string) (ID, error) {
	id, err := uuid.Parse(s)
	return ID(id), err
}

func StringToIDHex(s string) (IDHex, error) {
	id, err := primitive.ObjectIDFromHex(s)
	return IDHex(id), err
}

// ID Hex entity ID
type IDHex = primitive.ObjectID

// NewIDHex create a new entity ID
func NewIDHex() IDHex {
	return IDHex(primitive.NewObjectID())
}

// UUID or HexID string checker
func IsUUIDOrHexID(v string) bool {
	_, errUUID := StringToID(v)
	if errUUID != nil {
		_, errHex := StringToIDHex(v)
		if errHex != nil {
			return false
		}
	}
	return true
}
