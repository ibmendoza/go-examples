package main

import (
	"fmt"
	"github.com/antonholmquist/jason"
	"os"
)

func processSlc(slc []*jason.Object) {
	lenSlcKeys := len(slc)
	var err error
	var title, url string
	for i := 0; i < lenSlcKeys; i++ {

		//print object
		obj, err1 := slc[i].GetObject()
		fmt.Println("OBJECT")
		if err1 == nil {
			fmt.Println(obj)

			fmt.Println("TITLE")
			title, err = obj.GetString("title")

			if err == nil {
				fmt.Println(title)
			}

			fmt.Println("URL")
			url, err = obj.GetString("url")

			if err == nil {
				fmt.Println(url)
			}
		}
	}
}

func main() {
	exampleJSON :=
		`
{
	"a1" : {
		"name" : "HEADING 1",
		"b2" : {
			"name" : "Heading 1.1",
			"entries" : [{
					"title" : "Title1",
					"url" : "/url1"
				}, {
					"title" : "Title2",
					"url" : "/url2"
				}
			]
		}
	},

	"a2" : {
		"name" : "HEADING 2",
		"entries" : [{
				"title" : "Title1",
				"url" : "/url1"
			}, {
				"title" : "Title2",
				"url" : "/url2"
			}
		],
		"b2" : {
			"name" : "Heading 2.1",
			"entries" : [{
					"title" : "Title1",
					"url" : "/url1"
				}, {
					"title" : "Title2",
					"url" : "/url2"
				}
			]
		}
	}
} 
`

	objA, _ := jason.NewObjectFromBytes([]byte(exampleJSON))

	slcKeys := objA.GetKeys()
	lenSlcKeys := len(slcKeys)

	if lenSlcKeys == 0 {
		fmt.Println("No object key")
		os.Exit(1)
	}

	if lenSlcKeys == 1 {
		fmt.Println(slcKeys)
	}

	if lenSlcKeys > 1 {
		fmt.Println("MORE THAN 1", slcKeys)
	}

	slc2 := make([]*jason.Object, 0)

	for i := 0; i < lenSlcKeys; i++ {
		obj, err1 := objA.GetObject(slcKeys[i])
		if err1 == nil {
			slc2 = append(slc2, obj)
		}

		//get heading name
		str, err2 := obj.GetString("name")
		if err2 == nil {
			fmt.Println("NAME")
			fmt.Println(str)
		}

		//get entries array
		objArr, err3 := obj.GetObjectArray("entries")
		if err3 == nil {
			fmt.Println("ENTRIES")
			fmt.Println(objArr)

			processSlc(objArr)
		}
	}
}
