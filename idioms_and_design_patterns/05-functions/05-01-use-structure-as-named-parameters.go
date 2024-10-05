package main

type MyFuncOpts struct {
	FirstName string
	LastName string
	Age int
   }

func MyFunc(opts MyFuncOpts) error {
	// выполнение каких-либо действий
}

func main() {
	MyFunc(
		MyFuncOpts {
			LastName: "Patel",
			Age: 50,
		})
	MyFunc(
		MyFuncOpts {
		FirstName: "Joe",
		LastName: "Smith",
		})
}