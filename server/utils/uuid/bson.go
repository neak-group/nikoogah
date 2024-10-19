package uuid

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/x/bsonx/bsoncore"
)

var (
	_ bson.ValueMarshaler   = (*UUID)(nil)
	_ bson.ValueUnmarshaler = (*UUID)(nil)
)

func (u UUID) MarshalBSONValue() (bsontype.Type, []byte, error) {
	return bson.TypeBinary, bsoncore.AppendBinary(nil, bson.TypeBinaryUUID, []byte(u.UUID[:])), nil
}

func (u *UUID) UnmarshalBSONValue(typ bsontype.Type, value []byte) error {
	if typ != bson.TypeBinary {
		return fmt.Errorf("cannot unmarshal %v into a Binary", typ)
	}
	subtype, bin, rem, ok := bsoncore.ReadBinary(value)
	if subtype != bson.TypeBinaryUUID {
		return fmt.Errorf("cannot unmarshal binary subtype %v into a UUID", subtype)
	}
	if len(rem) > 0 {
		return fmt.Errorf("value has extra data: %v", rem)
	}
	if !ok {
		return errors.New("value does not have enough bytes")
	}
	*u = UUID{UUID: uuid.UUID(bin)}
	return nil
}
