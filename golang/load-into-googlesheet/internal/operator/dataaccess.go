package operator

// Connection is the interface that connects to an data source/sink.
type Connection interface {
	Get(impl interface{})
}

// Iterator is
type Iterator interface {
	// Next stores the next iteration of results in data.
	// data is a reference type that matches the data type of the results.
	// Next returns true if there are more results to read, false otherwise.
	Next(data interface{}) bool
}

// Input is the interface to read data from a data source connection.
type Input interface {
	// Read creates an iterator to iterate over the input data.
	Read(c Connection) (Iterator, error)
}

// Output is the interface to write data to a data sink connection.
type Output interface {
	// Write writes data to a data sink connection.
	// it is an Iterator that provides data that is compatible with the sink.
	Write(c Connection, it Iterator) error
}
