// https://es6console.com/jdug5ci1/

// Map object
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Map
// The Map object holds key-value pairs. Any value (both objects and primitive values) may be used as either a key or a value.
// The term "global objects" means standard built-in objects
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects


// == Using the Map object
var myMap = new Map();

var keyString = 'a string',
    keyObj = {},
    keyFunc = function() {};

// setting the values
myMap.set(keyString, "value associated with 'a string'");
myMap.set(keyObj, 'value associated with keyObj');
myMap.set(keyFunc, 'value associated with keyFunc');

myMap.size; // 3

// getting the values
myMap.get(keyString);    // "value associated with 'a string'"
myMap.get(keyObj);       // "value associated with keyObj"
myMap.get(keyFunc);      // "value associated with keyFunc"

myMap.get('a string');   // "value associated with 'a string'"
                         // because keyString === 'a string'
myMap.get({});           // undefined, because keyObj !== {}
myMap.get(function() {}) // undefined, because keyFunc !== function () {}


// == Relation with Array objects
console.log('== Relation with Array objects')

var kvArray = [['key1', 'value1'], ['key2', 'value2']];

// Use the regular Map constructor to transform a 2D key-value Array into a map
var myMap = new Map(kvArray);

myMap.get('key1'); // returns "value1"

// Use the Array.from function to transform a map into a 2D key-value Array
console.log(Array.from(myMap)); // Will show you exactly the same Array as kvArray

// Or use the keys or values iterators and convert them to an array
console.log(Array.from(myMap.keys())); // Will show ["key1", "key2"]

// == Using NaN as Map keys
console.log('== Using NaN as Map keys')

var myMap = new Map();
myMap.set(NaN, 'not a number');

myMap.get(NaN); // "not a number"

var otherNaN = Number('foo');
console.log(myMap.get(otherNaN)); // "not a number"

// == Iterating Maps with for..of
console.log('== Iterating Maps with for..of')

var myMap = new Map();
myMap.set(0, 'zero');
myMap.set(1, 'one');
for (var [key, value] of myMap) {
  console.log(key + ' = ' + value);
}
// 0 = zero
// 1 = one

for (var key of myMap.keys()) {
  console.log(key);
}
// 0
// 1

for (var value of myMap.values()) {
  console.log(value);
}
// zero
// one

for (var [key, value] of myMap.entries()) {
  console.log(key + ' = ' + value);
}
// 0 = zero
// 1 = one

// == Iterating Maps with forEach()
console.log('== Iterating Maps with forEach()')

myMap.forEach(function(value, key) {
  console.log(key + ' = ' + value);
});
// Will show 2 logs; first with "0 = zero" and second with "1 = one"


// == Iterating Maps with Array.from
console.log('== Iterating Maps with Array.from')

Array.from(myMap).map(v => console.log(v))
// Will show you exactly the same Array 
// [0,"zero"] 
// [1,"one"] 

// == 
console.log('== cat & dog')

const defaults = [
  {"cat": 0, "dog": 1},
  {"cat": 1, "dog": 2},
  {"cat": 2, "dog": 3},
]

var keyRe = /regexp/i
var mapa = [
  ['map_name_0', defaults[2]],
  ['map_name_1', {"cat": 100, "dog": 200}],
  ['map_name_2', {"cat": 400, "dog": 600}],
  ['map_name_3', defaults[1]],
  [keyString, defaults[0]] ,
  [keyObj, defaults[1]],
  [keyFunc, defaults[2]],
  [/RegExp/, defaults[2]],
  [keyRe, defaults[0]],
]

var myMap = new Map(mapa);
Array.from(myMap).map(v => console.log(v))

// getting the values
console.log(myMap.get(keyString))
console.log(myMap.get('a string'))
console.log(myMap.get(keyObj))
console.log(myMap.get(keyFunc))

console.log(myMap.get(/RegExp/))
console.log(myMap.get(keyRe))

for (var key of myMap.keys()) {
  console.log(key);
}

//Array.from(myMap).filter(v => /^map_name_/.test(v[0])).map(v => console.log(v))
Array.from(myMap).filter(v => /^map_name_/.test(v[0])).map(v => v[1].dog+=3)
Array.from(myMap).filter(v => /^map_name_/.test(v[0])).map(v => console.log(v))

