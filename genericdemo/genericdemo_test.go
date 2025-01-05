package genericdemo

import (
	"testing"
)

func Test_Add(t *testing.T) {
	Add(1, 2)
	Add(1.0, 2)
}

func TestReflectDemo(t *testing.T) {
	ReflectDemo()
}

func TestChanDemo(t *testing.T) {
	ChanDemo()
}
func TestChanDemo2(t *testing.T) {
	ChanDemo2()
}
