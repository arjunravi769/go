package main

import "fmt"

type person struct{
    age int
}

func main(){

    a := [5]person{person{21}, person{42}, person{43}, person{24}, person{25}}

    fmt.Println(function1(a))
    fmt.Println(function2(a))


}


func function1(a [5] person)bool {
for i := 0; i < len(a); i++ {
for j := 0; i < len(a); i++ {
        if a[i].age==2*a[j].age {
            return true
        }
    }
}
return false
}

func function2(a [5] person)bool {
for i := 0; i < len(a); i++ {
for j := 0; i < len(a); i++ {
        if a[i].age>=2*a[j].age {
            return true
        }
    }
}
return false
}
