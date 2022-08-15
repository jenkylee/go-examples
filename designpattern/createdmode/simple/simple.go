package simplefactory

import "fmt"

type API interface {

	Say(name string) string
}

// NewAPI return Api instance by type
func NewAPI(t int) API {
	if t == 2 {

		return &helloAPI{}
	} 

	return &hiAPI{}
}

// hiApi is one of API implement
type hiAPI struct{}

// Say hi to name
func (*hiAPI) Say(name string) string {

	return fmt.Sprintf("Hi, %s", name)
}

type helloAPI struct{}

func (*helloAPI) Say(name string) string {

	return fmt.Sprintf("Hello, %s", name)
}