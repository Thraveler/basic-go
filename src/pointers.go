package src

import "fmt"

type pc struct {
	brand string
	ram   int
	disk  int
}

func (myPc pc) printData() {
	fmt.Println("Brand: ", myPc.brand)
	fmt.Println("Memory: ", myPc.ram)
	fmt.Println("Disk: ", myPc.disk)
}

func (myPc *pc) duplicateMemory() {
	myPc.ram = myPc.ram * 2
}

// An example of stringers
func (myPc pc) String() string {
	return fmt.Sprintf("Pc %s, with %d GB of ram and %d GB of sdd", myPc.brand, myPc.ram, myPc.disk)
}

func TestPointer() {
	a := 40
	b := &a

	fmt.Println("Value for a: ", a)
	fmt.Println("Memory direction for b: ", b)
	fmt.Println("Value for b: ", *b)

	myPc := pc{
		brand: "Razer",
		ram:   8,
		disk:  500,
	}

	fmt.Println("MyPc: ", myPc)
	myPc.ram = 16
	fmt.Println("MyPc modified: ", myPc)

	myPc.printData()

	fmt.Println("Afer duplicate")
	myPc.duplicateMemory()
	myPc.printData()

}
