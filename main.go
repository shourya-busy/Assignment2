package main

import(
	"fmt"
	"reflect"
	"encoding/json"
)


//Prints JSON recursively
func printJson(data interface{}) {

	//reflect.Value is extracted from the data
	//Indirect() extracts the value if data is a pointer
	dataReflected := reflect.Indirect(reflect.ValueOf(data))
	

	//Cases to check the kind of reflected data
	switch dataReflected.Kind() {

	//Iterate over all the keys and handle the printing of values through recursive calls
	case reflect.Map :
		fmt.Println("\n========//=========")
		for _,key := range dataReflected.MapKeys(){
			val := dataReflected.MapIndex(key).Interface()
            fmt.Printf("%v (key) contaning %T: \n",key,val)
			printJson(val)
		}
		fmt.Println("=========//========\n")
	
    //Iterate over the len of slice and pack each element into interface{} to be printed in subsequent calls  
	case reflect.Slice:
		for i := 0; i < dataReflected.Len(); i++{
			fmt.Printf("Item %v : ",i+1)
			printJson(dataReflected.Index(i).Interface())
		}
		fmt.Println("")
    
	//Iterate over all the fields of struct and let recursive calls handle the values
	case reflect.Struct:
		fmt.Println("\n========//=========")
        for i := 0; i < dataReflected.NumField(); i++ {
			val := dataReflected.Field(i).Interface()
			fmt.Printf("%v : \n",dataReflected.Type().Field(i).Name)
            printJson(val)
        }
		fmt.Println("\n========//=========")
	
	//If the data does not belong to any special data structure simply print it 
	default:
		fmt.Printf("------ `%v` is of Type : %v and Kind :  %v\n\n",dataReflected,dataReflected.Type(),dataReflected.Kind())
	}


}

func main() {

	//Given input string
	input := `{
		"name" : "Tolexo Online Pvt. Ltd",
		"age_in_years" : 8.5,
		"origin" : "Noida",
		"head_office" : "Noida, Uttar Pradesh",
		"address" : [
		{
		"street" : "91 Springboard",
		"landmark" : "Axis Bank",
		"city" : "Noida",
		"pincode" : 201301,
		"state" : "Uttar Pradesh"
		},
		{
		"street" : "91 Springboard",
		"landmark" : "Axis Bank",
		"city" : "Noida",
		"pincode" : 201301,
		"state" : "Uttar Pradesh"
		}
		],
		"sponsers" : {
		"name" : "One"
		},
		"revenue" : "19.8 million$",
		"no_of_employee" : 630,
		"str_text" : ["one","two"],
		"int_text" : [1,3,4]
	}`

	//Interface type variable to store unmarshalled data
    var data interface{}

    //Unmarshalling the input string to data
	err := json.Unmarshal([]byte(input), &data)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
	}

	//Printing the data and it's types
	printJson(data)	
}