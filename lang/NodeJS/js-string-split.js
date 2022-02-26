
console.log('\n---------- Ex1');
var str = 'It iS   a 5r&e@@t Day.'
var array = str.split(" ");
console.log(array);
// [ 'It', 'iS', '', '', 'a', '5r&e@@t', 'Day.' ]
//            !!^^^^^^^^!!

var str = ' @botname  /do  with parameters'
var re = /\s+/
var array = str.split(re)
console.log(array)

function getArgs(str) {
  return !/^$|^\/|^@/.test(str)
}

console.log(array.filter(getArgs));

console.log('\n---------- Ex2');
// https://www.dyn-web.com/javascript/strings/split.php

var fruits = 'apple, orange, pear, banana, raspberry, peach';
var ar = fruits.split(', '); // split string on comma space
console.log( ar );
// [ "apple", "orange", "pear", "banana", "raspberry", "peach" ]

// Empty String Separator: If you pass an empty string as a separator, each character in the string will become an element in the returned array:
var str = 'abcdefg';
var ar = str.split(''); // empty string separator
console.log( ar ); // [ "a", "b", "c", "d", "e", "f", "g" ]

// Regular Expression Separator: The separator can be a regular expression:

var str = 'favorite desserts: brownies, banana bread, ice cream, chocolate chip cookies';
// regular expression separator
var re = /:\s|,\s/; // split on colon space or comma space
var ar = str.split(re);
console.log( ar );
// [ "favorite desserts", "brownies", "banana bread", "ice cream", "chocolate chip cookies" ]
