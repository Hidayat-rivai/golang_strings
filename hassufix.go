package main

import "fmt"
import "strings"

func main(){
	var isThere = strings.HasSuffix("disini adalah rumah","rumah")
	fmt.Println(isThere)
}