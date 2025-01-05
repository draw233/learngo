package variabledemo

import "testing"

func Test_var1(t *testing.T) {
	VarType()
}

func Test_Condition(t *testing.T) {
	Condition()
}

func Test_Switch(t *testing.T) {
	TypeSwitchDemo()
}

func Test_Goto(t *testing.T) {
	GotoDemo()
}

func Test_Loop(t *testing.T) {
	LoopDemo()
}

func Test_String(t *testing.T) {
	StringDemo()
	ArrDemo()
	SliceDemo()
}

func Test_Map(t *testing.T) {
	MapDemo()
}
