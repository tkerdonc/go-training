// This package implements an interface over the io.Reader and io.Writer
// interface, in order to keep track of the operations ran over this
// interface. More precisely, the amount of bytes written and read from
// them, as well as the cound of read/write operations

package paasio

import "io"
import "sync"

// ReadCounterImpl in an implementation of the ReadCounter interface,
// storing an io.Reader instance, two counter, and a mutex
type ReadCounterImpl struct {
	reader      io.Reader
	opCounter   int
	byteCounter int64
	mutex       *sync.Mutex
}

// Read implements the Read function of the Reader interface. It calls
// the actual implementation from the io.Reader instance and increments
// counter under the protection of the mutex.
func (readCounter *ReadCounterImpl) Read(bytes []byte) (int, error) {
	readCounter.mutex.Lock()
	defer readCounter.mutex.Unlock()
	var numBytes, rError = readCounter.reader.Read(bytes)
	readCounter.opCounter += 1
	readCounter.byteCounter += int64(numBytes)
	return numBytes, rError
}

// ReadCount returns the values of the instance's counters
func (readCounter *ReadCounterImpl) ReadCount() (n int64, nops int) {
	return readCounter.byteCounter, readCounter.opCounter
}

// NewReadCounter creates a new instance of ReadCounterImpl
func NewReadCounter(r io.Reader) ReadCounter {
	var mutex = &sync.Mutex{}
	var counter ReadCounter = &ReadCounterImpl{r, 0, 0, mutex}
	return counter
}

// WriteCounterImpl in an implementation of the WriteCounter interface,
// storing an io.Reader instance, two counter, and a mutex
type WriteCounterImpl struct {
	writer      io.Writer
	opCounter   int
	byteCounter int64
	mutex       *sync.Mutex
}

// Write implements the Write function of the Writer interface. It calls
// the actual implementation from the io.Writer instance and increments
// counter under the protection of the mutex.
func (writeCounter *WriteCounterImpl) Write(bytes []byte) (int, error) {
	writeCounter.mutex.Lock()
	defer writeCounter.mutex.Unlock()
	numBytes, wError := writeCounter.writer.Write(bytes)
	writeCounter.opCounter += 1
	writeCounter.byteCounter += int64(numBytes)
	return numBytes, wError
}

// WriteCount returns the values of the instance's counters
func (writeCounter *WriteCounterImpl) WriteCount() (n int64, nops int) {
	return writeCounter.byteCounter, writeCounter.opCounter
}

// NewWriteCounter creates a new instance of WriteCounterImpl
func NewWriteCounter(w io.Writer) WriteCounter {
	var mutex = &sync.Mutex{}
	var counter WriteCounter = &WriteCounterImpl{w, 0, 0, mutex}
	return counter
}

// ReadWriteCounterImpl in an implementation of the ReadWriteCounter
// interface, storing an instance of io.Reader and io/Writer, four
// counters, and a mutex
type ReadWriteCounterImpl struct {
	reader       io.Reader
	writer       io.Writer
	rOpCounter   int
	wOpCounter   int
	rByteCounter int64
	wByteCounter int64
	mutex        *sync.Mutex
}

// Read implements the Read function of the Reader interface. It calls
// the actual implementation from the io.Reader instance and increments
// counter under the protection of the mutex.
func (readWriteCounter *ReadWriteCounterImpl) Read(bytes []byte) (int, error) {
	readWriteCounter.mutex.Lock()
	defer readWriteCounter.mutex.Unlock()
	numBytes, rError := readWriteCounter.reader.Read(bytes)
	readWriteCounter.rOpCounter += 1
	readWriteCounter.rByteCounter += int64(numBytes)
	return numBytes, rError
}

// ReadCount returns the values of the instance's counters
func (readWriteCounter *ReadWriteCounterImpl) ReadCount() (n int64, nops int) {
	return readWriteCounter.rByteCounter, readWriteCounter.rOpCounter
}

// Write implements the Write function of the Writeer interface. It calls
// the actual implementation from the io.Reader instance and increments
// counter under the protection of the mutex.
func (readWriteCounter *ReadWriteCounterImpl) Write(bytes []byte) (int, error) {
	readWriteCounter.mutex.Lock()
	defer readWriteCounter.mutex.Unlock()
	numBytes, wError := readWriteCounter.writer.Write(bytes)
	readWriteCounter.wOpCounter += 1
	readWriteCounter.wByteCounter += int64(numBytes)
	return numBytes, wError
}

// WriteCount returns the values of the instance's counters
func (readWriteCounter *ReadWriteCounterImpl) WriteCount() (n int64, nops int) {
	return readWriteCounter.wByteCounter, readWriteCounter.wOpCounter
}

// NewReadWriteCounter creates a new instance of ReadWriteCounterImpl
func NewReadWriteCounter(rw interface{}) ReadWriteCounter {
	if pReadWriter, isReadWriter := rw.(readWriter); isReadWriter {
		var mutex = &sync.Mutex{}
		var counter ReadWriteCounter = &ReadWriteCounterImpl{pReadWriter.Reader, pReadWriter.Writer, 0, 0, 0, 0, mutex}
		return counter
	} else if pNopReadWriter, isNopReadWriter := rw.(nopReadWriter); isNopReadWriter {
		var mutex = &sync.Mutex{}
		var counter ReadWriteCounter = &ReadWriteCounterImpl{pNopReadWriter.nopReader, pNopReadWriter.nopWriter, 0, 0, 0, 0, mutex}
		return counter
	}
	// I would have rather returned an error, but is it even possible
	// without changing the function's signature ?
	panic("Inappropriate readwriter type")
}
