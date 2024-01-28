package panic_recover

import (
	"errors"
	"testing"
)

func TestPanicVsExit(t *testing.T) {
	//defer func() {
	//	t.Log("Finally!")
	//}()
	defer func() {
		if err := recover(); err != nil {
			t.Logf("recovered from %v", err)
		}
	}()
	t.Log("Start")
	panic(errors.New("Something wrong!"))
	t.Log("End")
	//os.Exit(-1)
}
