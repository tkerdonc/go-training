// This package impleements somthing. Surely.
package react

// CellObserver is an interface implemented by the computation cell,
// allowing them to receive notification from an input cell they
// subscribed to.
// The Notifystable lets an observabel signify its observer that is has
// finished sending the change notifications to all its observers.
type CellObserver interface {
	NotifyChange()
	NotifyStable()
}

// ObservableCell is an interface that allows computation cells
// implementing the CellObserver interface to subscribe to another cell
type ObservableCell interface {
	Subscribe(CellObserver) int64
	Unsubscribe(int64) error
}

// New creates a new reactor instance
func New() Reactor {
	var react Reactor = ReactorImpl{}
	return react
}

// ReactorImpl is the implementation struct of the Reactor Interface
type ReactorImpl struct{}

// CreateInput creates a new input cell, and sets its initial value to
// the one passed as a parameter.
func (reactor ReactorImpl) CreateInput(i int) InputCell {
	var inputCell InputCell = &InputCellImpl{i, make(map[int64]CellObserver), 0}
	return inputCell
}

// CreateCompute1 creates a new computation cell, connects it to a cell
// and sets its computation function
func (reactor ReactorImpl) CreateCompute1(cell Cell,
	computationFunction func(v int) int) ComputeCell {

	var computeCellImpl *ComputeCellImpl1 = &ComputeCellImpl1{
		cell,
		computationFunction(cell.Value()),
		computationFunction(cell.Value()),
		computationFunction,
		make(map[int64]func(int), 0),
		0,
		make(map[int64]CellObserver),
		0,
	}

	var computeCell ComputeCell = computeCellImpl
	var cellObserver CellObserver = computeCellImpl
	observableCell, isObservableCell := cell.(ObservableCell)
	if !isObservableCell {
		return nil
	}
	observableCell.Subscribe(cellObserver)

	return computeCell
}

// CreateCompute2 creates a new computation cell, connects it to two
// cells and sets its computation function
func (reactor ReactorImpl) CreateCompute2(
	cell1 Cell,
	cell2 Cell,
	computationFunction func(v1, v2 int) int) ComputeCell {

	var computeCellImpl *ComputeCellImpl2 = &ComputeCellImpl2{
		cell1,
		cell2,
		computationFunction(
			cell1.Value(),
			cell2.Value(),
		),
		computationFunction(
			cell1.Value(),
			cell2.Value(),
		),
		computationFunction,
		make(map[int64]func(int), 0),
		0,
		make(map[int64]CellObserver),
		0,
		true,
	}

	var computeCell ComputeCell = computeCellImpl
	var cellObserver CellObserver = computeCellImpl
	observableCell1, isObservableCell1 := cell1.(ObservableCell)
	observableCell2, isObservableCell2 := cell2.(ObservableCell)
	if !isObservableCell1 || !isObservableCell2 {
		return nil
	}
	observableCell1.Subscribe(cellObserver)
	observableCell2.Subscribe(cellObserver)

	return computeCell
}
