package main

//go:generate dot -Tpng -oDoorStateDiagram.png door.dot
//go:generate dot -Tpng -oDoorObjectHeirarchy.png object_heirarchy.dot

import "errors"

// Door is our model. It has state.
type Door struct {
	state DoorState
	VisitorCount int
	LockCount int
}

// DoorState is the state and possible transitions
type DoorState interface {
	Can(DoorState) bool
	Open(*Door) (DoorState, error)
	Close(*Door) (DoorState, error)
	Lock(*Door) (DoorState, error)
	Unlock(*Door) (DoorState, error)
}

// New creates a new instance of Door initialised with a state
func New(s DoorState) *Door {
	return &Door{state: s}
}

// Can is a helper function which returns true if we can transition
// from the current state to the desired state
func (d *Door) Can(n DoorState) bool {
	return d.state.Can(n)
}

// Open is a helper function to call the Open transition
// on the current state
func (d *Door) Open() error {
	s, err := d.state.Open(d)
	if err != nil {
		return err
	}
	d.state = s

	return nil
}

// Close is a helper function to call the Close transition
// on the current state
func (d *Door) Close() error {
	s, err := d.state.Close(d)
	if err != nil {
		return err
	}
	d.state = s

	return nil
}

// Lock is a helper function to call the Lock transition
// on the current state
func (d *Door) Lock() error {
	s, err := d.state.Lock(d)
	if err != nil {
		return err
	}
	d.state = s

	return nil
}

// Unlock is a helper function to call the Unlock transition
// on the current state
func (d *Door) Unlock() error {
	s, err := d.state.Unlock(d)
	if err != nil {
		return err
	}
	d.state = s

	return nil
}

// IsOpen returns true if we are in the OpenDoorState
func (d *Door) IsOpen() bool {
	_, ok := d.state.(OpenDoorState)
	return ok
}

// IsClosed returns true if we are in the ClosedDoorState
func (d *Door) IsClosed() bool {
	_, ok := d.state.(ClosedDoorState)
	return ok
}

// IsLocked returns true if we are in the LockedDoorState
func (d *Door) IsLocked() bool {
	_, ok := d.state.(LockedDoorState)
	return ok
}

// ErrIllegalStateTransition is raised if we cannot transiton
var ErrIllegalStateTransition = errors.New("Illegal state transition")

// abstractDoorState is an embedable type that provides default
// implementations for transitions
type abstractDoorState struct{}

func (d abstractDoorState) Can(DoorState) bool {
	return false
}

func (d abstractDoorState) Open(ctx *Door) (DoorState, error) {
	return &abstractDoorState{}, ErrIllegalStateTransition
}

func (d abstractDoorState) Close(ctx *Door) (DoorState, error) {
	return &abstractDoorState{}, ErrIllegalStateTransition
}

func (d abstractDoorState) Lock(ctx *Door) (DoorState, error) {
	return &abstractDoorState{}, ErrIllegalStateTransition
}

func (d abstractDoorState) Unlock(ctx *Door) (DoorState, error) {
	return &abstractDoorState{}, ErrIllegalStateTransition
}

// OpenDoorState is the open door state
type OpenDoorState struct {
	abstractDoorState
}

// Can returns true if we can transition to the desired
// state from the current OpenDoorState
func (s OpenDoorState) Can(n DoorState) bool {
	switch n.(type) {
	case ClosedDoorState:
		return true
	default:
		return false
	}
}

// Close transitions from OpenDoorState to ClosedDoorState
func (s OpenDoorState) Close(door *Door) (DoorState, error) {
	return ClosedDoorState{}, nil
}

// ClosedDoorState is the closed door state
type ClosedDoorState struct {
	abstractDoorState
}

// Can returns true if we can transition to the desired
// state from the current ClosedDoorState
func (s ClosedDoorState) Can(n DoorState) bool {
	switch n.(type) {
	case OpenDoorState:
		return true
	case LockedDoorState:
		return true
	default:
		return false
	}
}

// Open transitions from ClosedDoorState to OpenDoorState
func (s ClosedDoorState) Open(door *Door) (DoorState, error) {
	door.VisitorCount++
	return OpenDoorState{}, nil
}

// Lock transitions from ClosedDoorState to LockedDoorState
func (s ClosedDoorState) Lock(door *Door) (DoorState, error) {
	door.LockCount++
	return LockedDoorState{}, nil
}

// LockedDoorState is the locked door state
type LockedDoorState struct {
	abstractDoorState
}

// Can returns true if we can transition to the desired
// state from the current LockedDoorState
func (s LockedDoorState) Can(n DoorState) bool {
	switch n.(type) {
	case ClosedDoorState:
		return true
	default:
		return false
	}
}

// Unlock transitions from LockedDoorState to ClosedDoorState
func (s LockedDoorState) Unlock(door *Door) (DoorState, error) {
	return ClosedDoorState{}, nil
}
