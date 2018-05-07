package main

import (
	"net/http"
	"github.com/lpxxn/protobufferwithgo/gosrc/model"
	"github.com/golang/protobuf/proto"
	"bytes"
	"fmt"
	"io/ioutil"
)

func main() {

	client := http.Client{}
	p1 := model.Person{Name: "li", Id: 123}
	pby, _ := proto.Marshal(&p1)



	r, err := client.Post("http://localhost:8999/", "", bytes.NewBuffer(pby))
	if err != nil {
		fmt.Println(err)
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	addr1 := model.AddressBook{}
	fmt.Println(string(body))
	proto.Unmarshal(body, &addr1)
	fmt.Println("post return addr1 :", addr1)
}
