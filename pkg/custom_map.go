package custommap

type CustomMap struct {

}

func New () *CustomMap {
	return &CustomMap{}
}


func (m *CustomMap) Set (key string, value any) error {
	return nil
}

func (m *CustomMap) Get (key string) (any, error) {
	return nil, nil
}