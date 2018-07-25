package react

import "errors"

// InputCellImpl is an implementation of the InputCell Interface.
type InputCellImpl struct {
	value             int
	subscribers       map[int64]CellObserver
	subscriptionIndex int64
}

// SetValue sets the value of an InputCell.
func (cell *InputCellImpl) Value() int {
	return cell.value
}

// SetValue sets the value of an InputCell.
func (inputCell *InputCellImpl) SetValue(value int) {
	if value != inputCell.value {
		defer inputCell.notifyChange()
	}
	inputCell.value = value
}

// Notify notifies all the computation cells subscribed to this cell.
// That this input cell's value was changed. Once the NotifyChanges
// were called recursively on the whole graph, trigger the NotifyStable
// recursive calls on the graph/
func (inputCell *InputCellImpl) notifyChange() {
	for _, observer := range inputCell.subscribers {
		observer.NotifyChange()
	}
	inputCell.notifyStable()
}

// notifystable calls the NotifyStable observer function on all the
// cell's subscribers
func (inputCell *InputCellImpl) notifyStable() {
	for _, observer := range inputCell.subscribers {
		observer.NotifyStable()
	}
}

// Subscribe lets a computation cell subscribe to updates from this cell.
// It returns an identifier that can be used to later unsubscribe.
func (inputCell *InputCellImpl) Subscribe(subscribingCell CellObserver) int64 {
	inputCell.subscriptionIndex += 1
	inputCell.subscribers[inputCell.subscriptionIndex] = subscribingCell
	return inputCell.subscriptionIndex
}

// Unsubscribe lets a computation cell unsubscribe form updates from this cell
// It uses the identifier returned by the Subscribe function.
func (inputCell *InputCellImpl) Unsubscribe(subscriptionId int64) error {
	_, subscribed := inputCell.subscribers[subscriptionId]
	if !subscribed {
		return errors.New("Invalid subscription ID")
	}
	delete(inputCell.subscribers, inputCell.subscriptionIndex)
	return nil
}
