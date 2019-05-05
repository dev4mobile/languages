package method

func Method_Receiver() {

}

// can't use origin type
//func (a int) A()  {
//
//}

type Num int

func (a Num) A() {

}
