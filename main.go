package main

type MyStruct struct {
	Buf [1 << 16]byte
}

var Obj MyStruct = MyStruct{}

func PBP(myObj *MyStruct) {}
func PBV(myObj MyStruct)  {}

func TestBoth() {
	PBP(&Obj)
	PBV(Obj)
}

func main() {

}
