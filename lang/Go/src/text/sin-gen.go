////////////////////////////////////////////////////////////////////////////
// Porgram: sin-gen
// Purpose: Social Insurance Number Generator
// Authors: Tong Sun (c) 2014, All rights reserved
////////////////////////////////////////////////////////////////////////////

// Style: gofmt -tabs=false -tabwidth=4 -w

package main

import (
	"fmt"
	"os"
	//"strconv"
	//"strings"
)

var progname string = "sin-gen" // os.Args[0]

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s SIN_Str\n", progname)
	os.Exit(0)
}

// http://play.golang.org/p/j9gwGwr2FU
// Martin Schnabel mb0@mb0.org

func toDigits(s string) []int {
	result := make([]int, 0, len(s))
	// for _, c := range strings.Split(s, "") {
	// 	n, err := strconv.Atoi(c)
	for i, c := range s {
		if c < '0' || c > '9' {
			panic(fmt.Errorf("Character #%d from SIN_Str '%c' is invalid\n",
				i, c))
		}
		result = append(result, int(c-'0'))
	}
	return result
}

/*

Validation Procedure
http://www.ryerson.ca/JavaScript/lectures/forms/textValidation/sinProject.html

Fortunately, the Canadian Government provides social insurance numbers that
can be checked using a fairly straight forward method. Here is an excerpt from
document T4127(E), Payroll Deductions Formulas for Computer Programs (71st
Edition Effective January 1, 2000) published by Revenue Canada that describes
it:



Validation of the Social Insurance Number (SIN)

Human Resources Development Canada uses the SIN and employee information we
provide them to maintain accurate records of employee contributions and
earnings. To minimize the enquiries you receive, we recommend that you include
a SIN validity check as part of your payroll program.

A SIN has nine digits. The first eight digits are the basic number while the
ninth digit is a check digit. You can check whether or not a SIN is valid by
using the following verification method.

Example

The employee provides Social Insurance Number 193-456-787. You can check the
validity of the number by calculating the check digit as follows:

Basic number (first eight digits)	Check digit
193 456 78								7

Make a number from each alternate position to the left
beginning at the second digit

                                                        9  4  6  8

   Add the number to itself

                                                        9  4  6  8

   Sum
                                                        18 8 12 16

   Cross-add the digits in the sum (1 + 8 + 8 + 1 + 2 + 1 + 6) =

                                                                          27

   Add each alternate digit beginning at the first digit (1 + 3 + 5 + 7) =

                                                                          16

   Total

                                                                          43


If the total is a multiple of 10, the check digit should be 0; otherwise, subtract
the total calculated (43) from the next highest number ending in zero (50) 50

The check digit is (50 - 43)

7 = 7

Social Insurance Numbers that do not pass the validation check

If the SIN provided by an individual does not pass the verification check, the
preparer should confirm the SIN with the employer who received the original
number. If you are unable to obtain the correct number for the employee,
please do NOT leave the SIN field on the information slip blank. Instead,
report the SIN that was provided, even if it is not a valid number.
Frequently, even an incorrect number will enable us to find a match so that we
can correct the record and ensure the employee receives proper credit for the
deductions.

Instead of worrying about how the user formats the SIN number we can simply
extract the nine digits they provide and check that it is a proper SIN number
using the method described above.

Validation
http://en.wikipedia.org/wiki/Social_Insurance_Number#Validation

Social Insurance Numbers can be validated through a simple check digit process
called the Luhn Algorithm.

	046 454 286 <--- A fictitious, but valid SIN
	121 212 121 <--- Multiply each top number by the number below it.

In the case of a two-digit number, add the digits together and insert the
result (the digital root). Thus, in the second-to-last column, 8 multiplied by
2 is equal to 16. Add the digits (1 and 6) together (1 + 6 = 7) and insert the
result (7). So the result of the multiplication is:

	  086 858 276

Then, add all of the digits together:

	  0+8+6+8+5+8+2+7+6=50

If the SIN is valid, this number will be evenly divisible by 10.


United States Social Security number
https://sourcegraph.com/github.com/django/django-localflavor-us/symbols/python/django_localflavor_us/forms/USSocialSecurityNumberField

Checks the following rules to determine whether the number is valid:

    * Conforms to the XXX-XX-XXXX format.
    * No group consists entirely of zeroes.
    * The leading group is not "666" (block "666" will never be allocated).
    * The number is not in the promotional block 987-65-4320 through
      987-65-4329, which are permanently invalid.
    * The number is not one known to be invalid due to otherwise widespread
      promotional use or distribution (078-05-1120 or 219-09-9999).

*/

var multiply []int = []int{1, 2, 1, 2, 1, 2, 1, 2, 1}

/*
	validate
	Given the first 8 SIN digits in array,
	return the last SIN digit (9th) that satisfy validation
*/

func validate(da []int) int {
	if len(da) != 8 {
		panic(fmt.Errorf("Internal error: func validate need 8 SIN digits in array as input\n"))
	}
	sum := 0
	for i, d := range da {
		sum += (d * multiply[i]) % 9
	}
	return (10 - sum%10) % 10
}

func pow10(e int) int {
	if e == 0 {
		return 1
	}

	ret := 10
	for ii := 1; ii < e; ii++ {
		ret *= 10
	}
	return ret
}

func main() {
	// There will be only one command line argument
	if len(os.Args) != 2 {
		usage()
	}

	// the first command line argument is SIN# prefix
	sinStr := os.Args[1]

	padlen := 8 - len(sinStr)
	for ii := 0; ii < pow10(padlen); ii++ {
		// Pad leading zero with the length from a varible
		fmtstr := fmt.Sprintf("%%s%%0%dd", padlen)
		// in case the SIN# is 8 digits already, use those 8 digits
		fullstr := fmt.Sprintf(fmtstr, sinStr, ii)[:8]
		digits := toDigits(fullstr)

		d9 := validate(digits)
		fmt.Printf("%s%d\n", fullstr, d9)
	}

}

/*
Ref:

http://golang.org/pkg/os/

Canadian Social Insurance Number (SIN) Validation
http://www.runnersweb.com/running/sin_check.html

function clear(str) {
        var esum = 0;
        var enumbers = "";
        var checknum = 0;
        var ch_sum = "";
        var checkdigit = 0;
        var sin = "";
        var lastdigit = 0;
        }

function isNum(text) {

	  if(text == "") {
	        alert("You left the SIN field blank.");
               	return false;
	        }
	    	inStr = text;
                sin = text;
                inLen = inStr.length;

		if (inLen > 11 || inLen < 11) {
        		alert("SIN must be 11 characters long");
			return false;
			}

        	 for (var i = 0; i < text.length; i++) {
	 		var ch = text.substring(i, i + 1)

			if ((ch < "0" || "9" < ch) && (ch != "-"))  {
	               		alert("You must enter a 9 digits and two dashes.\nFormat 999-999-999.")
				return false;
		              	}
                        if ((i == 3 || i == 7) && (ch != "-")) {
                                alert("Invalid character in position 4 or 8;\nMust be a dash!");
                                return false;
                                }
			}
                        lastdigit = text.substring(10, 10 + 1);
                        // add numbers in odd positions; IE 1, 3, 6, 8
			var odd = ((text.substring(0,0 + 1)) * (1.0)  + (text.substring(2,2 + 1)) * (1.0)
			+(text.substring(5, 5+1)) * (1.0) + (text.substring(8,8 + 1)) * (1.0));

                        // form texting of numbers in even positions IE 2, 4, 6, 8
                        var enumbers =  (text.substring(1,1 + 1)) + (text.substring(4,4 + 1))+
                        (text.substring(6,6 + 1)) + (text.substring(9,9 + 1));

                        // add together numbers in new text string
                        // take numbers in even positions; IE 2, 4, 6, 8
                        // and double them to form a new text string
                        // EG if numbers are 2,5,1,9 new text string is 410218
                        for (var i = 0; i < enumbers.length; i++) {
                                var ch = (enumbers.substring(i, i + 1) * 2);
                                ch_sum = ch_sum + ch;
                                }

                        for (var i = 0; i < ch_sum.length; i++) {
                                var ch = (ch_sum.substring(i, i + 1));
                                esum = ((esum * 1.0) + (ch * 1.0));
                                }


			checknum = (odd + esum);

                        // subtextact checknum from next highest multiple of 10
                        // to give check digit which is last digit in valid SIN
			if (checknum <= 10) {
        			(checdigit = (10 - checknum));
				}
			if (checknum > 10 && checknum <= 20) {
				(checkdigit = (20 - checknum));
				}
                        if (checknum > 20 && checknum <= 30) {
				(checkdigit = (30 - checknum));
				}
                        if (checknum > 30 && checknum <= 40) {
				(checkdigit = (40 - checknum));
				}
                        if (checknum > 40 && checknum <= 50) {
				(checkdigit = (50 - checknum));
				}
                        if (checknum > 50 && checknum <= 60) {
				(checkdigit = (60 - checknum));
				}

                        if (checkdigit != lastdigit) {
                                alert(sin + "  is an invalid SIN; \nCheck digit incorrect!\nShould be: " + checkdigit);
                                history.go(0);
                                return false;
                                }
	               	return true;
	}

    function validate(textfield) {
                var esum = 0;
                var enumbers = "";
                var checknum = 0;
                var ch_sum = "";
                var checkdigit = 0;
                var sin = "";
                var lastdigit = 0;
                if (isNum(textfield.value))
    			alert(textfield.value + ' is a valid SIN');
                history.go(0);
	}

*/
