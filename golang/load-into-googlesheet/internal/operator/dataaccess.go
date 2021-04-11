package operator

type Connection interface {
	Get(impl interface{})
}
