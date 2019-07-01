package react

import "errors"

// ComputeCellImpl2 is an implementation of the ComputeCell interface
type ComputeCellImpl2 struct {
	sourceCell1            Cell
	sourceCell2            Cell
	value                  int
	lastStableValue        int
	computationFunction    func(int, int) int
	callbacks              map[int64]func(int)
	callbackIndex          int64
	subscribers            map[int64]CellObserver
	subscriptionIndex      int64
	childrenNotifiedStable bool // This cell has two parents and will
	// receive one NotifyStable from each of them. No point in calling
	// it more than once per stable steps, this bool lets us keep track
	// of it.
}

// SetValue sets the value of a ComputationCell
func (cell ComputeCellImpl2) Value() int {
	return cell.value
}

// AddCallback appends a callback to the list of computation callback
// functions of the computation cell, and returns a Canceler letting
// the calling program disable said callback.
func (computeCell *ComputeCellImpl2) AddCallback(callbackFunction func(int)) Canceler {
	computeCell.callbacks[computeCell.callbackIndex] = callbackFunction
	var canceler Canceler = &CancelerImpl2{computeCell, computeCell.callbackIndex} // not exactly threadsafe I reckon...
	computeCell.callbackIndex += 1
	return canceler
}

// NotifyChange implements the CellObserver interface, and allows a
// cell to notify this computation cell of a value change.
func (computeCell *ComputeCellImpl2) NotifyChange() {
	computeCell.childrenNotifiedStable = false
	newValue := computeCell.computationFunction(computeCell.sourceCell1.Value(), computeCell.sourceCell2.Value())
	if newValue != computeCell.Value() {
		computeCell.value = newValue
		computeCell.notify()
	}
}

// NotifyStable implements the CellObserver interface and allows a
// parent cell to inform this cell that all its siblings have
// processed the last change
func (computeCell *ComputeCellImpl2) NotifyStable() {
	if computeCell.lastStableValue != computeCell.Value() {
		computeCell.lastStableValue = computeCell.Value()
		for _, callback := range computeCell.callbacks {
			callback(computeCell.Value())
		}
	}

	if !computeCell.childrenNotifiedStable {
		for _, observer := range computeCell.subscribers {
			observer.NotifyStable()
		}
	}
}

// Notify notifies all the computation cells subscribed to this cell.
// That this input cell's value was changed.
func (computeCell *ComputeCellImpl2) notify() {
	for _, observer := range computeCell.subscribers {
		observer.NotifyChange()
	}
}

// Subscribe lets a computation cell subscribe to updates from this cell.
// It returns an identifier that can be used to later unsubscribe.
func (computeCell *ComputeCellImpl2) Subscribe(subscribingCell CellObserver) int64 {
	computeCell.subscriptionIndex += 1
	computeCell.subscribers[computeCell.subscriptionIndex] = subscribingCell
	return computeCell.subscriptionIndex
}

// Unsubscribe lets a computation cell unsubscribe form updates from this cell
// It uses the identifier returned by the Subscribe function.
func (computeCell *ComputeCellImpl2) Unsubscribe(subscriptionId int64) error {
	_, subscribed := computeCell.subscribers[subscriptionId]
	if !subscribed {
		return errors.New("Invalid subscription ID")
	}
	delete(computeCell.subscribers, computeCell.subscriptionIndex)
	return nil
}

// CancelerImpl1 is an implementation of the Canceler interface
type CancelerImpl2 struct {
	cell    *ComputeCellImpl2
	cbIndex int64
}

// Cancel disables a callback.
func (canceler CancelerImpl2) Cancel() {
	delete(canceler.cell.callbacks, canceler.cbIndex)
}
