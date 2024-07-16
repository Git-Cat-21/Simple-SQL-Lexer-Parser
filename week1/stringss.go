/*package main
import (
	"fmt"

	s "strings" //alias name
)

func main() {
	p := fmt.Println
	p(s.HasPrefix("abc", "a"))
}

//they are functions from the package npt methos therefore we need to pass the string itself as the 1st arg
*/
//to know the type %T

//text templates in Go are powerful mechanism for generating textual output based on template format
//allow us to define a template string that includes placeholders for dynamic values and execute that template with specific data to produce the final output
//they have a static text along with dynamic values enclosed within {{.Placeholder}}
//to produce an output from a template we need to execute it with specific data and that is done by parsing the template string

/*package main

import (
	"os"
	"text/template"
)

func main() {
	const tpl = "Hello, {{.Name}}! You are {{.Age}} years old\n"
	type data struct {
		Name string
		Age  int
	}
	ok := data{"Khushi P", 19}
	t, err := template.New("Hello").Parse(tpl)
	if err != nil {
		panic(err)
	}
	err = t.Execute(os.Stdout, ok)
	if err != nil {
		panic(err)
	}
}*/

// go data structures into JSON format encode and decode to go data structures
/*package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name`
	Age  int    `json:age`
}

func main() {
	person := Person{"Khushi ", 19}
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error encoding JSON", err)
		return
	}
	fmt.Println("Encoded JSON: ", string(jsonData))
	jsonStr := `{"name":"Bob","age":25}`
	var decodedPerson Person
	err = json.Unmarshal([]byte(jsonStr), &decodedPerson)
	if err != nil {
		fmt.Println("Error decoding JSON", err)
		return
	}
	fmt.Println("Decoded Person: ", decodedPerson)

}*/

/*package main

import (
	"encoding/xml"
	"fmt"
)

type Person struct {
	Name string `xml:"name"`
	Age  int    `xml:"age"`
}

func main() {
	person := Person{"khushi ", 19}
	xmlData, err := xml.Marshal(person)
	if err != nil {
		fmt.Println("Error encoding XML", err)
		return
	}
	fmt.Println("Encoded XML: ", string(xmlData))
	xmlStr := `<Person><name>Bob</name><age>25</age></Person>`
	var decodedPerson Person
	err = xml.Unmarshal([]byte(xmlStr), &decodedPerson)
	if err != nil {
		fmt.Println("error decoding XML", err)
		return
	}
	fmt.Println("deocded person", decodedPerson)

}*/


