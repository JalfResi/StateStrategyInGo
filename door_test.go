package main

import (
	"testing"
)

func TestCan(t *testing.T) {

	var testTable = []struct {
		initial  DoorState
		attempt  DoorState
		expected bool
	}{
		{OpenDoorState{}, OpenDoorState{}, false},
		{OpenDoorState{}, ClosedDoorState{}, true},
		{OpenDoorState{}, LockedDoorState{}, false},

		{ClosedDoorState{}, OpenDoorState{}, true},
		{ClosedDoorState{}, ClosedDoorState{}, false},
		{ClosedDoorState{}, LockedDoorState{}, true},

		{LockedDoorState{}, OpenDoorState{}, false},
		{LockedDoorState{}, ClosedDoorState{}, true},
		{LockedDoorState{}, LockedDoorState{}, false},
	}

	for _, tt := range testTable {
		d := New(tt.initial)

		if d.Can(tt.attempt) != tt.expected {
			t.Errorf("Guard failure. Expected: %t got: %t", tt.expected, d.Can(tt.attempt))
		}
	}
}

func TestIs(t *testing.T) {

	var testTable = []struct {
		initial  DoorState
		fnIs     func(d *Door) bool
		expected bool
	}{
		{OpenDoorState{}, func(d *Door) bool { return d.IsOpen() }, true},
		{OpenDoorState{}, func(d *Door) bool { return d.IsClosed() }, false},
		{OpenDoorState{}, func(d *Door) bool { return d.IsLocked() }, false},

		{ClosedDoorState{}, func(d *Door) bool { return d.IsOpen() }, false},
		{ClosedDoorState{}, func(d *Door) bool { return d.IsClosed() }, true},
		{ClosedDoorState{}, func(d *Door) bool { return d.IsLocked() }, false},

		{LockedDoorState{}, func(d *Door) bool { return d.IsOpen() }, false},
		{LockedDoorState{}, func(d *Door) bool { return d.IsClosed() }, false},
		{LockedDoorState{}, func(d *Door) bool { return d.IsLocked() }, true},
	}

	for _, tt := range testTable {
		d := New(tt.initial)

		if tt.fnIs(d) != tt.expected {
			t.Errorf("Guard failure. Expected: %t got: %t", tt.expected, tt.fnIs(d))
		}
	}
}

func TestTransitions(t *testing.T) {

	var testTable = []struct {
		initial      DoorState
		fnTransition func(d *Door) error
		fnExpected   error
	}{
		{OpenDoorState{}, func(d *Door) error { return d.Open() }, ErrIllegalStateTransition},
		{OpenDoorState{}, func(d *Door) error { return d.Close() }, nil},
		{OpenDoorState{}, func(d *Door) error { return d.Lock() }, ErrIllegalStateTransition},
		{OpenDoorState{}, func(d *Door) error { return d.Unlock() }, ErrIllegalStateTransition},

		{ClosedDoorState{}, func(d *Door) error { return d.Open() }, nil},
		{ClosedDoorState{}, func(d *Door) error { return d.Close() }, ErrIllegalStateTransition},
		{ClosedDoorState{}, func(d *Door) error { return d.Lock() }, nil},
		{ClosedDoorState{}, func(d *Door) error { return d.Unlock() }, ErrIllegalStateTransition},

		{LockedDoorState{}, func(d *Door) error { return d.Open() }, ErrIllegalStateTransition},
		{LockedDoorState{}, func(d *Door) error { return d.Close() }, ErrIllegalStateTransition},
		{LockedDoorState{}, func(d *Door) error { return d.Unlock() }, nil},
		{LockedDoorState{}, func(d *Door) error { return d.Lock() }, ErrIllegalStateTransition},
	}

	for _, tt := range testTable {
		d := New(tt.initial)

		if err := tt.fnTransition(d); err != tt.fnExpected {
			t.Errorf("Transition failure. Expected: %v got: %v", tt.fnExpected, err.Error())
		}
	}
}
