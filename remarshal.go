package main

import (
	"fmt"
	"github.com/ibmendoza/remarshal"
)

func main() {
	str :=
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

	c, _ := remarshal.Convert([]byte(str), "JSON", "TOML")

	fmt.Println(c)

	c, _ = remarshal.Convert([]byte(str), "JSON", "YAML")

	fmt.Println(c)
}
