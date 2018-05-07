package main

import (
	"github.com/lpxxn/protobufferwithgo/gosrc/model"
	"fmt"
	"github.com/golang/protobuf/proto"
)

func main() {
	person1 := &model.Person{Name: "li", Id: 2}
	ps := make([]*model.Person, 0)
	ps = append(ps, person1)
	ps = append(ps, &model.Person{Name:"peng", Id: 3})

	addr1 := model.AddressBook{People: ps}

	fmt.Println("addr1 : ", addr1)
	by1, err := proto.Marshal(&addr1)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(addr1.String())
	fmt.Println(string(by1))

	addr2 := model.AddressBook{}
	proto.Unmarshal(by1, &addr2)
	fmt.Println(addr2)
}
