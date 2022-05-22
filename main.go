package main

import(
	"fmt" 
)

func main() {
	var arr [3]int
	arr[0] =1
	arr[1] =2
	arr[2] =3
	fmt.Println(arr)
	slice := arr[:]
	arr[0]=42
	slice[1] = 23
	fmt.Println(slice , arr)
	//you will see that the array and the slice are the same because the slice is just pointing to the same array 
	dslice := []int{}
	dslice = append(dslice,1,2,3,4,5,6,7,8,9,10,11,12,13)
	fmt.Println(dslice)
	m :=map[string]int{"foo":42,"doo":23}
	m["poo"]=55
	delete(m, "foo")
	fmt.Println(m)
	type User struct {
		ID int 
		FirstName string
		LastName string
		IsMarried bool
	}
	var user User
	fmt.Println(user)
	user.IsMarried=true
	user.LastName="salah"
	user.FirstName="Mahmoud"
	user.ID=1
	fmt.Println(user)
	user2 := User{ID:1,FirstName:"Ahmed",LastName:"Salah",IsMarried:false}
	fmt.Println(user2)

}
