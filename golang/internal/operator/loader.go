package operator

type Loader interface {
	Load(record []string) error
	Flush()
}
