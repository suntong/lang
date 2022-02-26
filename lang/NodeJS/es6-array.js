console.log('\n---------- Array copy');
// Copying array by value in JavaScript
// https://stackoverflow.com/questions/7486085/copying-array-by-value-in-javascript

{
var arr1 = ['a','b','c']; var arr2 = arr1; arr2.push('d'); //Now, arr1 = ['a','b','c','d']
console.log(arr1, arr2)
}
// How to copy the array to get two independent arrays?
// use newArray = oldArray.slice();
{
  var arr1 = ['a','b','c']
  var arr3 = arr1.slice(); arr3.push('d')
  console.log(arr1, arr3)
}

// https://es6console.com/jdv1tpko/

// https://code.tutsplus.com/tutorials/how-to-use-map-filter-reduce-in-javascript--cms-26209
var monday = [
  {
    'name'     : 'Write a tutorial',
    'duration' : 180
  },
  {
    'name'     : 'Some web development',
    'duration' : 120
  }
];

var tuesday = [
  {
    'name'     : 'Keep writing that tutorial',
    'duration' : 240
  },
  {
    'name'     : 'Some more web development',
    'duration' : 180
  },
  {
    'name'     : 'A whole lot of nothing',
    'duration'  : 240
  }
];
var tasks = [monday, tuesday];
console.log(tasks);
// [[{"name":"Write a tutorial"...,"duration":120}],[{"name":"Keep writing...:240}]]

{
const transformers = [
  {
    name: 'Optimus Prime',
    form: 'Freightliner Truck',
    team: 'Autobot'
  },
  {
    name: 'Megatron',
    form: 'Gun',
    team: 'Decepticon'
  },
]

const transformers2 = [
 {
    name: 'Bumblebee',
    form: 'VW Beetle',
    team: 'Autobot'
  },
  {
    name: 'Soundwave',
    form: 'Walkman',
    team: 'Decepticon'
  }
];


const second = transformers2

// https://stackoverflow.com/questions/7486085/copying-array-by-value-in-javascript

console.log('without spread operator : array push');
var first = transformers.slice()
console.log(first);
first.push(second);
console.log(first);


console.log('with spread operator : array push');
var first = transformers.slice()
first.push(...second);
console.log(first);
}


console.log('\n---------- Array append/merge');
// Merge Arrays in one with ES6 Array spread
// https://gist.github.com/yesvods/51af798dd1e7058625f4
// This utilizes the ES6 rest parameters syntax

{
const arr1 = [1,2,3]
const arr2 = [3,4,5,6]
const arr3 = [...arr1, ...arr2]
console.log(arr3) //arr3 ==> [1,2,3,3,4,5,6]
}

console.log('\n---------- Array append to each array element');
// Merge objects in array
// Merge objects to each array element
// https://es6console.com/jdv0uf0p/

{
  var obj1 = {a: 1, b: 2}; 
  var obj2 = {a: 4, c: 110}; 
  var obj3 = Object.assign({},obj1, obj2);
  console.log(obj3);
}

let transformers = [
  {
    name: 'Optimus Prime',
    form: 'Freightliner Truck',
    team: 'Autobot'
  },
  {
    name: 'Megatron',
    form: 'Gun',
    team: 'Decepticon'
  },
 // {
 //    name: 'Bumblebee',
 //    form: 'VW Beetle',
 //    team: 'Autobot'
 //  },
 //  {
 //    name: 'Soundwave',
 //    form: 'Walkman',
 //    team: 'Decepticon'
 //  }
];
const extras = {
  team: 'Decepticon',
  collected: true,
  favorite: true
}

{
  transformers = transformers.map(t => Object.assign({}, t, extras))
  console.log(transformers)
}

// transformers.map(t => t = Object.assign({}, t, extras))
transformers.map(t => {r = Object.assign({}, t, extras); console.log(r); t = r; })

console.log('\n---------- Array iterations');
// https://www.digitalocean.com/community/tutorials/how-to-use-array-methods-in-javascript-iteration-methods

let fish = [ "piranha", "barracuda", "cod", "eel" ];
// Loop through the length of the array
for (let i = 0; i < fish.length; i++) {
  console.log(fish[i]);
}
console.log('-------- forEach');
// Print out each item in the array
fish.forEach(individualFish => {
  console.log(individualFish);
})

console.log('-------- forEach with filter');
// https://stackoverflow.com/questions/31399411/
// to skip the current iteration, simply return
fish.forEach(individualFish => {
  if (individualFish > "d") { return }
  console.log(individualFish);
})



console.log('-------- map');
// Print out each item in the array
let printFish = fish.map(individualFish => {
  console.log(individualFish);
});
printFish;
// Pluralize all items in the fish array
let pluralFish = fish.map(individualFish => {
  return `${individualFish}s`;
});

console.log('-------- filter');
let seaCreatures = [ "shark", "whale", "squid", "starfish", "narwhal" ];
// Filter all creatures that start with "s" into a new list
let filteredList = seaCreatures.filter(creature => {
  return creature[0] === "s";
});
console.log(filteredList);

console.log('-------- reduce');
let numbers = [ 42, 23, 16, 15, 4, 8 ];
// Get the sum of all numerical values
let sum = numbers.reduce((a, b) => {
	return a + b;
});
console.log(sum); // 108

console.log('-------- find');
// Check if a given value is a cephalopod
seaCreatures = [ "whale", "octopus", "shark", "cuttlefish", "flounder" ];
const isCephalopod = cephalopod => {
	return [ "cuttlefish", "octopus" ].includes(cephalopod);
}
console.log(seaCreatures.find(isCephalopod)); // octopus
console.log(seaCreatures.findIndex(isCephalopod)); // 1
// Since octopus was the first entry in the array to satisfy the test in the isCephalopod() function, it is the first value to be returned.


console.log('-------- findIndex');
const isThereAnEel = eel => {
    return [ "eel" ].includes(eel);
}
console.log(seaCreatures.findIndex(isThereAnEel)); // test is not satisfied, -1


console.log('\n---------- Array find');
// https://jsbin.com/luxecur/edit?js,console,output
//
// JavaScript Array find
// https://www.w3schools.com/jsref/jsref_find.asp

// Get the value of the *first* element in the array that has a value of 18 or more:

var ages = [3, 10, 18, 38, 20];
function checkAdult(age) {
    return age > 12;
}
console.log(ages.find(checkAdult)) // 18

// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/find
var array1 = [5, 12, 8, 130, 44];
var found = array1.find(function(element) {
  return element > 10;
});
console.log(found);
// expected output: 12

var inventory = [
    {name: 'apples', quantity: 2},
    {name: 'bananas', quantity: 0},
    {name: 'cherries', quantity: 5}
];
function isCherries(fruit) { 
    return fruit.name === 'cherries';
}
console.log(inventory.find(isCherries)); 
// { name: 'cherries', quantity: 5 }

// https://www.linkedin.com/pulse/javascript-find-object-array-based-objects-property-rafael
var objArray = [
    { id: 0, name: 'Object 0', otherProp: '321' },
    { id: 1, name: 'O1', otherProp: '648' },
    { id: 2, name: 'Another Object', otherProp: '850' },
    { id: 3, name: 'Almost There', otherProp: '046' },
    { id: 4, name: 'Last Obj', otherProp: '984' }
];

var obj = objArray.find(function (obj) { return obj.id === 3; });
console.log(obj) // id: 3 ...
var obj = objArray.find( obj => obj.id === 2 );
console.log(obj) // id: 2 ...

// http://2ality.com/2014/05/es6-array-methods.html
console.log([6, -5, 8].find(x => x < 0)) // -5
console.log([6, -5, 8].findIndex(x => x < 0)) // 1

console.log('\n---------- Array find arg');
console.log("https://github.com/coderbunker/candobot/blob/master/support.js")

/*

const actions = [
{
  action: (_data, _message) => 'DEFAULT ACTION',
  regexp: /^$/i,
  reply: (message, _output) => `Yes ${message.userName}? Ask me for help if you need me.`,
},
{
  action: (_data, _message) => help(),
  regexp: /help/i,
  reply: (message, output) => `\n${output}`,
},
]

 // * Callback for Array.find.
 // * @param action action object
 // * @this object with content and result
 // * @returns true if match regexp, false otherwise
function searchRegexp(action) {
  if(action && action.regexp) {
    const exec = action.regexp.exec(this.content)
    if(exec) {
      this.result = exec.slice(1)
      return true
    }
  }
  return false
}
const findParams = {
    content: message.content,
    result: null
  }
const action = actions.find(searchRegexp, findParams)

*/


console.log('\n---------- Array map / filter / reduce');

{
const numbers = [0, 1, 2, 3, 4, 5, 6];
const doubledNumbers = numbers.map(n => n * 2); // [0, 2, 4, 6, 8, 10, 12]
const evenNumbers = numbers.filter(n => n % 2 === 0); // [0, 2, 4, 6]
const sum = numbers.reduce((prev, next) => prev + next, 0); // 21

console.log(doubledNumbers); // [0, 2, 4, 6, 8, 10, 12]
console.log(evenNumbers); // [0, 2, 4, 6]
console.log(sum) //21

console.log(numbers.filter(n => n % 2 === 1).map(n => n * 2)) // [2,6,10] 

// Compute total grade sum for students with grades 10 or above
const students = [
  { name: "Nick", grade: 10 },
  { name: "John", grade: 15 },
  { name: "Julia", grade: 19 },
  { name: "Nathalie", grade: 9 },
];

const aboveTenSum = students
  .map(student => student.grade) // we map the students array to an array of their grades
  .filter(grade => grade >= 10) // we filter the grades array to keep those 10 or above
  .reduce((prev, next) => prev + next, 0); // we sum all the grades 10 or above one by one

console.log(aboveTenSum) // 44 -- 10 (Nick) + 15 (John) + 19 (Julia), Nathalie below 10 is ignored
}


