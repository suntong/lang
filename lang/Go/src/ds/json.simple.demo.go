package main

import (
	"fmt"

	"github.com/bitly/go-simplejson" // for json get
)

var jsonStr = `
       {
           "person": [{
              "name": "piao",
              "age": 30,
              "email": "piaoyunsoft@163.com",
              "phoneNum": [
                  "13974999999",
                  "13984999999"
              ]
           }, {
              "name": "aaaaa",
              "age": 20,
              "email": "aaaaaa@163.com",
              "phoneNum": [
                  "13974998888",
                  "13984998888"
              ]
           }, {
              "name": "bbbbbb",
              "age": 10,
              "email": "bbbbbb@163.com",
              "phoneNum": [
                  "13974997777",
                  "13984997777"
              ]
           }]
       }-->
            <div bla bla bla
       `

func main() {
	js, err := simplejson.NewJson([]byte(jsonStr))

	if err != nil {
		panic(err.Error())
	}

	personArr, err := js.Get("person").Array()
	fmt.Printf("personArr(%d): %v\n", len(personArr), personArr)

	// 遍历
	for i, v := range personArr {
		fmt.Println(i, v)
		person := js.Get("person").GetIndex(i)
		name := person.Get("name").MustString()
		age := person.Get("age").MustInt()
		email := person.Get("email").MustString()

		fmt.Printf("name=%s, age=%d, email=%s\n", name, age, email)

		// 读取手机号
		phoneNumArr, _ := person.Get("phoneNum").Array()
		for ii, vv := range phoneNumArr {
			fmt.Println(ii, vv)

		}
	}

}
