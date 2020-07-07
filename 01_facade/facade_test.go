package facade

import (
	"fmt"
	"testing"
)

func TestFacade(t *testing.T) {
	api := NewAPI()
	ret := api.DoSomeThing()
	fmt.Println(ret)
}
