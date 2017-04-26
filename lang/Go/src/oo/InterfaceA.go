////////////////////////////////////////////////////////////////////////////
// Porgram: Interface.go
// Purpose: Go interface usage examples
// Authors: Tong Sun (c) 2013, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"time"
)

func main() {
	mainJM()
	mainWL()
	mainYH()
	mainSB()
	mainCC()
	mainJY()
}

////////////////////////////////////////////////////////////////////////////
// assign interface into struct with type assertion
// Jan Mercl, https://play.golang.org/p/adNNjpau-O

type t struct {
	f int
}

func f(v interface{}) {
	fmt.Println(v.(t))
}

func mainJM() {
	v := t{42}
	f(v)
	fmt.Println()
	// {42}
}

////////////////////////////////////////////////////////////////////////////
// Usage of interface in Go
// Worlock, http://stackoverflow.com/a/14509524/2125837

type Info interface {
	Noofchar() int
	Increment()
}

type Testinfo struct {
	noofchar int
}

func (x *Testinfo) Noofchar() int {
	return x.noofchar
}
func (x *Testinfo) Increment() {
	x.noofchar++
}

func mainWL() {
	// t := &Testinfo{noofchar:1} // or
	// t := Testinfo{noofchar:1} // or
	var t Info = &Testinfo{noofchar: 1}
	// var t Info = Testinfo{noofchar:1}
	// cannot use Testinfo literal (type Testinfo) as type Info in assignment:
	//	Testinfo does not implement Info (Increment method has pointer receiver)

	fmt.Println("No of char ", t.Noofchar())
	t.Increment()
	fmt.Println("No of char ", t.Noofchar())
	fmt.Println()
	// No of char  1
	// No of char  2
}

////////////////////////////////////////////////////////////////////////////
// Example 1
// By: qyuhen
// From: github.com/qyuhen

type User struct {
	Name string
}

type Student struct {
	Name string "T_NAME"
	Age  int    "T_AGE"
}

func mainYH() {
	i := 100
	d := map[*int]struct{ x, y float64 }{&i: {1.0, 2.0}}
	fmt.Println(d, d[&i], d[&i].y)

	d2 := map[string]interface{}{"a": 1, "b": User{"user1"}}
	fmt.Printf("b: %#v\n", d2["b"])
	fmt.Println(d2, d2["b"].(User).Name)

	u2 := User{"user2"}
	up := &u2
	d3 := map[string]interface{}{"a": 2, "b": up}
	fmt.Printf("b: %#v\n", d3["b"])
	fmt.Println(d3["b"].(*User).Name)

	fmt.Println("\nReflect:")
	s := Student{"Tom", 23}
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fmt.Printf("%s (%s = %v)\n", f.Tag, f.Name, v.Field(i).Interface())
	}
	fmt.Println()

	/*
		map[0xc82000a5d8:{1 2}] {1 2} 2
		b: main.User{Name:"user1"}
		map[a:1 b:{user1}] user1
		b: &main.User{Name:"user2"}
		user2

		Reflect:
		T_NAME (Name = Tom)
		T_AGE (Age = 23)

	*/
}

////////////////////////////////////////////////////////////////////////////
// Example 2
// By:   Steven Blenkinsop
// From: http://play.golang.org/p/vs7Smk_cbz

// This is an interface. We can assign values of any type
// that has a method called "String" that returns a `string`:
type X interface {
	String() string
}

// This is a slice of `X`. We can put a bunch of `X`s in here and
// get the out again based on an index:
type Xs []X

// Let's define some types that we can assign to `X`
type T1 string
type T2 int
type T3 struct{}

// Okay, so we need to define a `String() string` method for each of them.

// We need to convert `T1` to `string` in order to return is from the method:
func (t T1) String() string { return string(t) }

// strconv.Itoa returns the base ten string representation of an integer:
func (t T2) String() string { return strconv.Itoa(int(t)) }

// We're just going to make up a string representation for this one:
func (t T3) String() string { return "I am a T3" }

// This ^ is a receiver.  It's like a normal parameter, except it turns the
// function into a method. If I write `t.String()`, `t` will get passed as
// the receiver of the method.

func mainSB() {
	// Let's create some values:
	var t1 T1 = "I am a T1"
	var t2 T2 = 2
	var t3 T3 = T3{}

	// I could rewrite that as:
	//     t1 := T1("I am a T1")
	//     t2 := T2(2)
	//     t3 := T3{}
	// Or, for `t3`:
	//     var t3 T3
	// since its value is automatically `T3{}`

	// We can call the method on them and print the result:
	fmt.Println(t1.String())
	fmt.Println(t2.String())
	fmt.Println(t3.String())

	// Now we can try assigning them to an `X`:
	var x X
	x = t1
	fmt.Println(x.String())
	x = t2
	fmt.Println(x.String())
	x = t3
	fmt.Println(x.String())

	// Now let's use that `Xs` slice type:
	xs := Xs{t1, t2, t3}

	// And we can loop over it and print it:
	for i := range xs {
		fmt.Printf("xs: %#v\n", xs[i])
		fmt.Println(xs[i].String())
	}
	fmt.Println()

	/*
	  I am a T1
	  2
	  I am a T3
	  I am a T1
	  2
	  I am a T3
	  (main.T1) I am a T1
	  I am a T1
	  (main.T2) 2
	  2
	  (main.T3) I am a T3
	  I am a T3
	*/
}

////////////////////////////////////////////////////////////////////////////
// Carlos Castillo,
// https://groups.google.com/forum/#!topic/golang-nuts/Wavir7gbTvk

/*

if maximizing speed or minimizing string creation is important than you can
use type switches / type assertions and / or reflect. Here I've created a
toString method which, if possible, uses a method / feature of the type to
avoid the extra work and the extra string conversion that fmt.SPrint*
functions have to do. http://play.golang.org/p/fHs2ABo1c0

*/

func toString(x interface{}) string {
	switch y := x.(type) {

	// Handle dates with special logic
	// This needs to come above the fmt.Stringer
	// test since time.Time's have a .String()
	// method
	case time.Time:
		return y.Format("A Monday")

	// Handle type string
	case string:
		return y

	// Handle type with .String() method
	case fmt.Stringer:
		return y.String()

	// Handle type with .Error() method
	case error:
		return y.Error()

	}

	// Handle named string type
	if v := reflect.ValueOf(x); v.Kind() == reflect.String {
		return v.String()
	}

	// Fallback to fmt package for anything else like numeric types
	return fmt.Sprint(x)
}

type astring string

type istringer int

func (is istringer) String() string {
	return strconv.FormatInt(int64(is)*-42, 10)
}

func mainCC() {
	as := astring("foo")
	err := errors.New("Nothing to see here!")
	is := istringer(32)
	flux := time.Date(1965, 11, 5, 0, 0, 0, 0, time.Local) // Day time travel was invented

	fmt.Println(toString(as))    // named string type (reflect code)
	fmt.Println(toString("bar")) // string type
	fmt.Println(toString(err))   // error interface
	fmt.Println(toString(is))    // Stringer interface
	fmt.Println(toString(5))     // number (uses Sprint fallback)
	fmt.Println(toString(flux))  // Custom code in toString to print dates.
	fmt.Println()
	// foo
	// bar
	// Nothing to see here!
	// -1344
	// 5
	// A Friday
}

////////////////////////////////////////////////////////////////////////////
// Example 3
// By:   Jianfeng Ye
// From: https://groups.google.com/forum/?fromgroups=#!topic/golang-china/B-Ky3U0BllY

type DateAccessObject struct {
	Parent interface{}
}

func (self *DateAccessObject) Connect() {
	parent := reflect.ValueOf(self.Parent)
	method := parent.MethodByName("Connect")
	if method.IsValid() {
		method.Call(nil)
	} else {
		fmt.Println("dao connect")
	}
}

func (self *DateAccessObject) Select() {
	parent := reflect.ValueOf(self.Parent)
	method := parent.MethodByName("Select")
	if method.IsValid() {
		method.Call(nil)
	} else {
		fmt.Println("dao select")
	}
}

func (self *DateAccessObject) Disconnect() {
	parent := reflect.ValueOf(self.Parent)
	method := parent.MethodByName("Disconnect")
	if method.IsValid() {
		method.Call(nil)
	} else {
		fmt.Println("dao disconnect")
	}
}

func (self *DateAccessObject) Run() {
	self.Connect()
	self.Select()
	self.Disconnect()
}

type MysqlDao struct {
	dao DateAccessObject
}

func NewMysqlDao() *MysqlDao {
	mysqlDao := new(MysqlDao)
	mysqlDao.dao.Parent = mysqlDao
	return mysqlDao
}

func (self *MysqlDao) Connect() {
	fmt.Println("mysql dao connect")
}

func (self *MysqlDao) Select() {
	fmt.Println("mysql dao select")
}

func (self *MysqlDao) Disconnect() {
	fmt.Println("mysql dao disconnect")
}

type RedisDao struct {
	dao DateAccessObject
}

func NewRedisDao() *RedisDao {
	redisDao := new(RedisDao)
	redisDao.dao.Parent = redisDao
	return redisDao
}

func (self *RedisDao) Connect() {
	fmt.Println("redis dao connect")
}

func (self *RedisDao) Select() {
	fmt.Println("redis dao select")
}

func (self *RedisDao) Disconnect() {
	fmt.Println("redis dao disconnect")
}

type OtherDao struct {
	dao DateAccessObject
}

func NewOtherDao() *OtherDao {
	otherDao := new(OtherDao)
	otherDao.dao.Parent = otherDao
	return otherDao
}

func mainJY() {
	mysqlDao := NewMysqlDao()
	mysqlDao.dao.Run()

	redisDao := NewRedisDao()
	redisDao.dao.Run()

	otherDao := NewOtherDao()
	otherDao.dao.Run()

	/*
	  mysql dao connect
	  mysql dao select
	  mysql dao disconnect
	  redis dao connect
	  redis dao select
	  redis dao disconnect
	  dao connect
	  dao select
	  dao disconnect
	*/
}
