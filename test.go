package main

import "fmt"

type person struct{
    age int
}

func main(){

    a := [5]person{person{16}, person{31}, person{15}, person{18}, person{32}}
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
   // a.sort()
  //  fmt.Println(a)
    m := make(map[string]int)
    m["largest"] = a[0].age;
    m["smallest"] = a[0].age;
for i := 0; i < len(a); i++ {

 if a[i].age>m["largest"] {
        m["largest"]=a[i].age
    } else if a[i].age<m["smallest"] {
         m["smallest"]=a[i].age
    } else {
    }
     if m["largest"]>=2*m["smallest"] {
      return true
    }
}

 return false

}







