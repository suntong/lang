// ES6 Spread & Destructuring
// https://jsbin.com/tusevob/edit?js,console

console.log('... Spread operator');
console.log([1,2,3]);
console.log(...[1,2,3]);

let first = [1,2,3];
let second = [4,5,6];

console.log('without spread operator : array push');
first.push(second);
console.log(first);

first = [1,2,3];
console.log('with spread operator : array push');
first.push(...second);
console.log(first);

console.log('************************************');
console.log('Spread operator in function parameter');
let addThreeThings = (a,b,c) => console.log(a+b+c);
addThreeThings(...second); // 15


console.log('************************************');
console.log('String expressions');
var salutation = "Hello";
var greeting = `${salutation}, 
You are crazy!!`;
console.log(greeting);


console.log()
console.log('************************************');
console.log('Destructuring Assignment in ES6 : Ex1');
/*
  var obj = {
  color: "blue"
  }
  console.log(obj.color);
*/
var {color} = {
  color: "blue"
}
console.log(color);

console.log('************************************');
console.log('Destructuring Assignment in ES6 : Ex2');
var {color, position} = {
  color: 'blue',
  name: 'John',
  state: 'CA',
  position: 'Forward'
}
console.log(color);
console.log(position);


console.log('************************************');
console.log('Destructuring Assignment in ES6 : Ex3');
var generateObj = function() {
  return {
    color: 'blue',
    name: 'John',
    state: 'CA',
    position: 'Forward'
  }
}
var {name, state} = generateObj();
console.log(name);
console.log(state);
console.log("---")
var {name:fn, state:location} = generateObj();
console.log(fn);
console.log(location);


console.log('************************************');
console.log('Destructuring Assignment in ES6 : Ex4');
var [fst,,,,fth] = ['red', 'yellow', 'green', 'blue', 'orange'];
console.log(fst, fth); // red orange

console.log('***********************************');
console.log('Spread Operator ES6: Ex5');
var webLanguages = ['html', 'css', 'javascript', 'es6', 'angular', 'nodejs'];
var serverLanguages = ['php', 'asp', 'nodejs', 'c#', 'java'];
var fullStackLanguages = [...webLanguages, ...serverLanguages];//ES6
// or
// var fullStackLanguages= webLanguages.push(...serverLanguages);
// or
// var fullStackLanguages = webLanguages.concat(serverLanguages);
console.log(fullStackLanguages);

//Merge two arrays and filter duplicates
webLanguages = ['html', 'css', 'javascript', 'html', 'es6', 'java', 'angular', 'css', 'nodejs'];
serverLanguages = ['php', 'asp', 'nodejs', 'c#', 'java'];
console.log('*******JavaScript ES5 way of merging two arrays and removing duplicates*****');
fullStackLanguages = webLanguages
  .filter(function(item,index) {
    return webLanguages.indexOf(item) === index;
  })
  .concat(serverLanguages
	  .filter(function(item) {
	    return webLanguages.indexOf(item) === -1;
	  })
	 );
console.log(fullStackLanguages);

// ES6 way of merging two array and removing duplicates using Set
console.log('******ES6 way of merging two array and removing duplicates using Set*****');
fullStackLanguages = [...new Set([...webLanguages ,...serverLanguages])];
console.log(fullStackLanguages);

console.log('********Spread arguments of a function*****');
var spreadArgumentsFunc = function(name, id, ...details) {
  console.log(`name: ${name}, id: ${id}, details: ${details}`, details);
}
spreadArgumentsFunc('John', 1234, '123 St, City', 'phone', 'email@domain.com', 'something else');


console.log('*******Default function arguments*********');
var defaultFucArgs = function(arg1, arg2) {
  if(arg2 === undefined) {
    arg2 = 2;
  }
  console.log(arguments);//this will hold only the arugments coming from the function call.
}
defaultFucArgs(1);// Output: [1]

defaultFucArgs = function(arg1, arg2 = 2) {
  console.log(arguments);
}
console.dir(defaultFucArgs);
defaultFucArgs(1);// Output: [1] but arg2 will assigned with default value of 2 from function defination.

console.log('*************Arrow Function***********');

var add = (a,b) => a+b // (a,b) => { return a+b }
console.log('ADD: => ', add(1,2));

var test = (...num) => num;
console.log(test(1,2,3,4));

