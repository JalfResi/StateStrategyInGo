package main

//go:generate dot -Tpng -oDoorStateDiagram.png door.dot

import "errors"

// Door is our model. It has state.
type Door struct {
	state DoorState
}

// DoorState is the state and possible transitions
type DoorState interface {
	Can(DoorState) bool
	Open() (DoorState, error)
	Close() (DoorState, error)
	Lock() (DoorState, error)
	Unlock() (DoorState, error)
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
	s, err := d.state.Open()
	if err != nil {
		return err
	}
	d.state = s

	return nil
}

// Close is a helper function to call the Close transition
// on the current state
func (d *Door) Close() error {
	s, err := d.state.Close()
	if err != nil {
		return err
	}
	d.state = s

	return nil
}

// Lock is a helper function to call the Lock transition
// on the current state
func (d *Door) Lock() error {
	s, err := d.state.Lock()
	if err != nil {
		return err
	}
	d.state = s

	return nil
}

// Unlock is a helper function to call the Unlock transition
// on the current state
func (d *Door) Unlock() error {
	s, err := d.state.Unlock()
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

func (d abstractDoorState) Open() (DoorState, error) {
	return &abstractDoorState{}, ErrIllegalStateTransition
}

func (d abstractDoorState) Close() (DoorState, error) {
	return &abstractDoorState{}, ErrIllegalStateTransition
}

func (d abstractDoorState) Lock() (DoorState, error) {
	return &abstractDoorState{}, ErrIllegalStateTransition
}

func (d abstractDoorState) Unlock() (DoorState, error) {
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
func (s OpenDoorState) Close() (DoorState, error) {
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
func (s ClosedDoorState) Open() (DoorState, error) {
	return OpenDoorState{}, nil
}

// Lock transitions from ClosedDoorState to LockedDoorState
func (s ClosedDoorState) Lock() (DoorState, error) {
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
func (s LockedDoorState) Unlock() (DoorState, error) {
	return ClosedDoorState{}, nil
}
