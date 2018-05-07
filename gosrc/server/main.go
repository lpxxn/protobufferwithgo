package main

import (
	"github.com/lpxxn/protobufferwithgo/gosrc/model"
	"fmt"
	"github.com/golang/protobuf/proto"
	"net/http"
	"io/ioutil"
)

func main() {
	//person1 := &model.Person{Name: "li", Id: 2}
	//ps := make([]*model.Person, 0)
	//ps = append(ps, person1)
	//ps = append(ps, &model.Person{Name:"peng", Id: 3})
	//
	//addr1 := model.AddressBook{People: ps}
	//
	//fmt.Println("addr1 : ", addr1)
	//by1, err := proto.Marshal(&addr1)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//fmt.Println(addr1.String())
	//fmt.Println(string(by1))
	//
	//addr2 := model.AddressBook{}
	//proto.Unmarshal(by1, &addr2)
	//fmt.Println(addr2)
	http.HandleFunc("/", TestData)

	http.ListenAndServe(":8999", nil)
}
func TestData(w http.ResponseWriter, r *http.Request) {
	fmt.Println("test data api in")
	rby, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	p1 := &model.Person{}
	proto.Unmarshal(rby, p1)
	addr1 := model.AddressBook{People: []*model.Person{p1}}
	fmt.Println(addr1)
	uby, err := proto.Marshal(&addr1)
	if err != nil {
		w.Write([]byte("error"))
	} else {
		w.Write(uby)
	}
}
