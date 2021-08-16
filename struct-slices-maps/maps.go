package main

import "fmt"

func main() {
	imap := make(map[string]int)
	imap["key1"] = 10
	imap["key2"] = 11
	fmt.Println(imap)
	anotherMap := map[string]int{
		"key1": 10,
		"key2": 11,
	}
	fmt.Println(anotherMap)
	//will not produce any warning
	delete(anotherMap, "key1")
	delete(anotherMap, "key1")
	delete(anotherMap, "key1")
	fmt.Println("another map", anotherMap)
	//check by comma operator does key exist or not
	_, ok := imap["anyUndefinedKey"]
	if ok {
		fmt.Println("exist")
	} else {
		fmt.Println("does not exist")
	}
}
