package filter

type Chain struct {
	Name    string
	Filters *[]Filter
}

func NewChain(name string, filters ...Filter) *Chain {
	return &Chain{
		Name:    name,
		Filters: &filters,
	}
}

func (f *Chain) Process(data Request) (Response, error) {
	var ret interface{}
	var err error
	for _, filter := range *f.Filters {
		ret, err = filter.Process(data)
		if err != nil {
			return ret, err
		}
		data = ret
	}
	return ret, err
}
