package main

import "fmt"

func main() {
	fmt.Println("Section 2 : Data types and control structure")

	//int
	var m1 int32 = 2
	var m2 int64 = 3

	fmt.Println(m1+int32(m2))

	//string
	s1 := "string value"
	fmt.Println(s1)


	//arrays
	array()

	//struct
	structs()

	//Pointers 
	
	x:= 1
	y:=2
	swap(&x,&y)

	// Control flows
	// if else, for, switch case, continue, break
	// There is no while loop, but for loop in Golang is very powerful


	//if
	flag := true
	if  flag {
		fmt.Println(" flag is true")
		flag = false
	}else if !flag {
		fmt.Println(" flag is false")
	}else {
		fmt.Println(" flag is nil")
	}

	//for

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	arr := [] string {"Waheb","Firas","Maria"}
	for i,names := range(arr) {
		fmt.Println(i)
		fmt.Println(names)
	}

	//continue
	fmt.Print("continue")
	for i := 0; i < 10; i++ {
		if i%2==0 {
			continue // in this case, if the condition is true, the loop will start from the begining
		}
		fmt.Println(i)
	}

	//break
	fmt.Println("break")
	for i := 1; i < 10; i++ {
		if i%2==0 {
			break // in this case, if the condition is true, the loop will stop ever
		}
		fmt.Println(i)
	}


	//Switch case
	fmt.Println("Switch case")

	day := "Fri"
	switch day {
	case "Fri":
		fmt.Println("This is Friday")
	case "Mon","Thu","Wed":
		fmt.Println("prin anything else")

	default:
		fmt.Println(day)	
	}	
}

func array()  {
	arr := [] int {1,2,3}
	fmt.Println(arr)
}

func structs ()  {
	c := Car{ "Audi", 1 ,"12"}
	fmt.Println(c)
}

type Car struct{
	Name 	string
	Age 	int
	MdelNo 	string 
}

func swap(m1, m2 *int)  {
	//before sawp
	fmt.Println(*m1, *m2)
	//after swap
	temp:= m1
	m1=m2
	m2=temp

	fmt.Println(*m1, *m2)
} 
