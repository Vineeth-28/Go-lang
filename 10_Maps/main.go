package main

import (
	"fmt"
	"maps"
)

//maps -> hashtable
func main(){
   m:= make(map[string]string)
   //setting an element
   m["name"]="golang"
   m["area"] ="backend"
   
   //get an element 
   fmt.Println(m["name"] , m["area"])

   //getting not present element  , if the key doesnot exits it returns zero value 
   
   fmt.Println(m["phone"])
   number := make(map[string]int )
   number["age"]=30
   number["phone"]=211;
   fmt.Println(number["age"], number["phone"])

   //len of map
   fmt.Println(len(m) ,len(number))

   //delete from map 
   delete(m,"area");
   fmt.Println(m)
   //clear map
   clear(number);

   //another way 
    dict := map[string]int{"price": 40, "phone": 3}
    fmt.Println(dict) 

	//check element present in map ok is idiom in go
	ok:=dict["price"]
	fmt.Println(ok);

	//check map is equal 

	dict2 := map[string]int{"price": 40, "phone": 3}
    fmt.Println(dict2)
	fmt.Println(maps.Equal(dict,dict2))

}