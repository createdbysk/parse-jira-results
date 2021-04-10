package operator

type MockConnection struct{}

func (c *MockConnection) Execute(q Query, data interface{}) {}

type MockQuery struct{}

func (q *MockQuery) Get(result interface{}) error {
	return nil
}
