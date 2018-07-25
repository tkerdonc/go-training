package react

import "errors"

// ComputeCellImpl1 is an implementation of the ComputeCell interface
type ComputeCellImpl1 struct {
	sourceCell      Cell // The paren cell
	value           int  // The cell's value
	lastStableValue int  // The cell's value in the last stable state
	// used to determine wether calling callbacks
	// is required
	computationFunction func(int) int
	callbacks           map[int64]func(int)
	// Im a storing callbacks in this map. I Could
	// have used a slice, and just appended them.
	// However, when canceling a callback (and using
	// its index to do so, It would have required a
	// dummy replacement value (func(int){} that
	// would still have been called. Would it have
	// been better than this map with silly ints in
	// range as key ?
	callbackIndex int64                  // Key to use for the the next callback
	subscribers   map[int64]CellObserver // Children cells
	// (subscribed  to this one)
	subscriptionIndex int64 // Key to use for the next subscriber
}

// SetValue sets the value of a ComputationCell
func (computeCell *ComputeCellImpl1) Value() int {
	return computeCell.value
}

// AddCallback appends a callback to the list of computation callback
// functions of the computation cell, and returns a Canceler letting
// the calling program disable said callback.
func (computeCell *ComputeCellImpl1) AddCallback(callbackFunction func(int)) Canceler {
	computeCell.callbacks[computeCell.callbackIndex] = callbackFunction
	var canceler Canceler = &CancelerImpl1{computeCell, computeCell.callbackIndex} // not exactly threadsafe I reckon...
	computeCell.callbackIndex += 1
	return canceler
}

// NotifyChange implements the CellObserver interfaces, and allows an input
// cell to notify this computation cell of a value change.
func (computeCell *ComputeCellImpl1) NotifyChange() {
	newValue := computeCell.computationFunction(computeCell.sourceCell.Value())
	if newValue != computeCell.Value() {
		computeCell.value = newValue
		computeCell.notify()
	}
}

// NotifyStable implements the CellObserver interface and allows a
// parent cell to inform this cell that all its siblings have
// processed the last change
func (computeCell *ComputeCellImpl1) NotifyStable() {
	if computeCell.lastStableValue != computeCell.Value() {
		computeCell.lastStableValue = computeCell.Value()
		for _, callback := range computeCell.callbacks {
			callback(computeCell.Value())
		}
	}

	for _, observer := range computeCell.subscribers {
		observer.NotifyStable()
	}
}

// Notify notifies all the computation cells subscribed to this cell.
// That this input cell's value was changed.
func (computeCell *ComputeCellImpl1) notify() {
	for _, observer := range computeCell.subscribers {
		observer.NotifyChange()
	}
}

// Subscribe lets a computation cell subscribe to updates from this cell.
// It returns an identifier that can be used to later unsubscribe.
func (computeCell *ComputeCellImpl1) Subscribe(subscribingCell CellObserver) int64 {
	computeCell.subscriptionIndex += 1
	computeCell.subscribers[computeCell.subscriptionIndex] = subscribingCell
	return computeCell.subscriptionIndex
}

// Unsubscribe lets a computation cell unsubscribe form updates from this cell
// It uses the identifier returned by the Subscribe function.
func (computeCell *ComputeCellImpl1) Unsubscribe(subscriptionId int64) error {
	_, subscribed := computeCell.subscribers[subscriptionId]
	if !subscribed {
		return errors.New("Invalid subscription ID")
	}
	delete(computeCell.subscribers, computeCell.subscriptionIndex)
	return nil
}

// CancelerImpl1 is an implementation of the Canceler interface
type CancelerImpl1 struct {
	cell    *ComputeCellImpl1
	cbIndex int64
}

// Cancel disables a callback.
func (canceler CancelerImpl1) Cancel() {
	delete(canceler.cell.callbacks, canceler.cbIndex)
}
