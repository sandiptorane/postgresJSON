package validate

import (
	"encoding/json"
	"fmt"
	"log"
)

func IsJSON(data string)bool{
	return json.Valid([]byte(data))
}

type User struct {
	User string `json:"user"`
	Pass string `json:"pass"`
}

func Unmarshal(data string){
	user := User{}
	err:= json.Unmarshal([]byte(data),&user)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println(user)
}