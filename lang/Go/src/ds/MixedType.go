/*

Q: Mixed type array in Go

I need a data struct / solution that I can store mixed data-types into an array. How to architect that?

Details -- Consider the checkout point at the cashier, the checkout process receives the following string of instructions

- purchase
- coupon
- purchase
- purchase

I want to store all the above request data into an array, so I don't lost the sequence how they arrive, which in turn requests the array element to be either purchase or coupon.

How to make it happen?


A: From Josh Humphries @bluegosling.com
https://groups.google.com/d/msg/golang-nuts/VOGJ8su7KjU/dNqm0ZO9AAAJ

You would typically define the array type as an array of some interface that is implemented by both purchase and coupon. You can then use a type-switch or type-assertion to determine/assert the actual type at runtime, on a per element basis.

If the two types do not share some interface (e.g. common methods), you could define the array as []interface{}. But this also allows code to accidentally add any kind of value to the array (not just a purchase or a coupon). So it may be better to create a marker interface -- an unexported interface with a no-op unexported marker method -- and make purchase and coupon both implement that interface.

For example, the following code:

*/

package main

import "fmt"

type checkoutObject interface {
	checkoutObject()
}

type Purchase struct {
	// yadda yadda
	item string
}

func (p *Purchase) checkoutObject() {
	// no-op marker method
}

type Coupon struct {
	// yadda yadda
	discount int // the discount percent
}

func (c *Coupon) checkoutObject() {
	// no-op marker method
}

// assert that Purchase and Coupon implement checkoutObject
var _ checkoutObject = (*Purchase)(nil)
var _ checkoutObject = (*Coupon)(nil)

// Now you can define your array like so:
var checkoutItems []checkoutObject

// The compiler will only let you add *Purchase and *Coupon
// values to the array.

func main() {

	// init
	checkoutItems = append(checkoutItems, &Purchase{"Orange"})
	checkoutItems = append(checkoutItems, &Coupon{10})
	checkoutItems = append(checkoutItems, &Purchase{"Apple"})
	checkoutItems = append(checkoutItems, &Purchase{"Banana"})

	// fmt.Printf("%#v\n", checkoutItems)
	// for _, e := range checkoutItems {
	// 	fmt.Printf("%#v\n", e)
	// }

	// You can use the values at runtime like so:
	for _, e := range checkoutItems {
		switch e := e.(type) {
		case *Purchase:
			// handle purchase
			fmt.Printf("%#v\n", e)
		case *Coupon:
			// handle coupon
			fmt.Printf("%#v\n", e)
		default:
			panic(fmt.Sprintf("unsupported type: %T", e))
		}
	}

}

/*

$ go run MixedType.go
&main.Purchase{item:"Orange"}
&main.Coupon{discount:10}
&main.Purchase{item:"Apple"}
&main.Purchase{item:"Banana"}

*/
