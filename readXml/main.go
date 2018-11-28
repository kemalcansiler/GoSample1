package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Users struct {
	XMLName xml.Name `xml:"UserList"`
	Users   []User   `xml:"User"`
}

type User struct {
	XMLName           xml.Name `xml:"User"`
	Identifier        string   `xml:"Identifier"`
	Alias             string   `xml:"Alias"`
	Title             string   `xml:"Title"`
	Type              string   `xml:"Type"`
	FirstCreationTime string   `xml:"FirstCreationTime"`
}

func main() {
	xmlFile, err := os.Open("users.xml")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened users.xml")
	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)

	var users Users
	xml.Unmarshal(byteValue, &users)

	id := ""

	fmt.Scanln(&id)

	for i := 0; i < len(users.Users); i++ {
		if id == users.Users[i].Identifier {
			fmt.Println("User Type: " + users.Users[i].Type)
			fmt.Println("User Identifier: " + users.Users[i].Identifier)
			fmt.Println("User Alias: " + users.Users[i].Alias)
			break
		}
	}
	/*r := gin.Default()

	r.GET("/:id", func(c *gin.Context) {
		identifier := c.Param("identifier")
		user, ok := User[identifier]

		if !ok {
			c.String(http.StatusNotFound, "404 - Page Not Found")
			return
		}

		c.HTML(http.StatusOK, "vacation-overview.html",
			map[string]interface{}{
				"User": user,
			})
	})

	r.Run(":1455")*/
}
