package mytool

func init(){
	println("mytool init")
}

func Add(xPtr *int) {
	*xPtr = *xPtr + 5
}
