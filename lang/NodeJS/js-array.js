console.log('\n---------- Array append');
// Appending to array
// http://stackoverflow.com/questions/351409/appending-to-array
// Use the `push()` function to append to an array:


// initialize array

var arr = [
  "Hi",
  "Hello",
  "Bonjour"
];


// append new value to the array
arr.push("Hola");

// display all values
for (var i = 0; i < arr.length; i++) {
  console.log(arr[i]);
}
// Hi
// Hello
// Bonjour
// Hola

console.log('\n---------- Array concat');
// JavaScript Array concat() Method
// http://www.w3schools.com/jsref/jsref_concat_array.asp

var hege = ["Cecilie", "Lone"];
var stale = ["Emil", "Tobias", "Linus"];
var children = hege.concat(stale);
console.log(children);
// [ 'Cecilie', 'Lone', 'Emil', 'Tobias', 'Linus' ]

console.log('\n---------- Array filter');
// Javascript Array filter() Method
// http://www.tutorialspoint.com/javascript/array_filter.htm
// https://www.w3schools.com/jsref/jsref_filter.asp

function checkAdult(age) {
  return age >= 18;
}

var ages = [32, 33, 16, 40];
console.log(ages.filter(checkAdult));
// [ 32, 33, 40 ]


console.log('\n---------- Array filter with variable parameters');
// Filter a javascript array with variable parameters
// http://stackoverflow.com/questions/17099029/how-to-filter-a-javascript-object-array-with-variable-parameters

arr = [
  { name: "joe",  age21: 1 },
  { name: "nick", age21: 0 },
  { name: "blast", age21: 1 }
];

// specific
arr.filter(function(item) {
  return (item.name === "nick" && item.age21 === 1);
});

// generalized function that work for any numbers of properties given the object
function filter(arr, criteria) {
  return arr.filter(function(obj) {
    return Object.keys(criteria).every(function(c) {
      return obj[c] == criteria[c];
    });
  });
}

var arr = [
  { name: 'Steve', age: 18, color: 'red' },
  { name: 'Louis', age: 21, color: 'blue' }, //*
  { name: 'Mike', age: 20, color: 'green' },
  { name: 'Greg', age: 21, color: 'blue' }, //*
  { name: 'Josh', age: 28, color: 'red' }
];
console.log(filter(arr, { age: 21, color: 'blue' }));
// [
//   { name: 'Louis', age: 21, color: 'blue' },
//   { name: 'Greg', age: 21, color: 'blue' }
// ]


// generalized regular expression matching
function filter(arr, criteria) {
  return arr.filter(function(obj) {
    return Object.keys(criteria).every(function(c) {
      return new RegExp(criteria[c]).test(obj[c]);
    });
  });
}
console.log(filter(arr, { age: /^2\d$/, color: /^(red|blue)$/ }));
//^ Louis, Greg               --^-- twenty-something
//                                               ----^---- red OR blue

