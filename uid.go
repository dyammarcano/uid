package uid

import (
	"github.com/google/uuid"
	"github.com/oklog/ulid/v2"
	"github.com/rs/xid"
	"github.com/segmentio/ksuid"
)

type mode int

const (
	// KSUID ModeKSUID is the default mode for generating unique identifiers.
	KSUID mode = iota

	// UUID ModeUUID generates a UUID.
	UUID

	// XID ModeXID generates a XID.
	XID

	// ULID ModeULID generates a ULID.
	ULID
)

type ActionID []byte

func NewUID(m mode) ActionID {
	switch m {
	case KSUID:
		return ksuid.New().Bytes()
	case UUID:
		return ActionID(uuid.NewString())
	case XID:
		return ActionID(xid.New().String())
	case ULID:
		return ActionID(ulid.MustNew(ulid.Now(), nil).String())
	default:
		return ksuid.New().Bytes()
	}
}

func (a ActionID) String() string {
	return ksuid.KSUID(a).String()
}

func (a ActionID) Bytes() []byte {
	return a[:]
}

func Get(id ActionID) {

}
