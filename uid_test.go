package uid

import "testing"

func TestNewUID(t *testing.T) {
	uid := NewUID(KSUID)
	if len(uid) == 0 {
		t.Errorf("expected non-empty uid")
	}

	t.Logf("uid: %s", uid)

	uid = NewUID(UUID)
	if len(uid) == 0 {
		t.Errorf("expected non-empty uid")
	}

	t.Logf("uid: %s", uid)

	uid = NewUID(XID)
	if len(uid) == 0 {
		t.Errorf("expected non-empty uid")
	}

	t.Logf("uid: %s", uid)

	uid = NewUID(ULID)
	if len(uid) == 0 {
		t.Errorf("expected non-empty uid")
	}

	t.Logf("uid: %s", uid)
}
