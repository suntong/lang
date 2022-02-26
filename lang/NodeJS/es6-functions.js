
console.log('\n---------- ES6 Arrow Function');
// ES6 Arrow Function
// https://jsbin.com/weyacex/edit?js,console

// Normal JavaScript function
var createGreeting = function(message, name) {
  return message + name;
}

// Also same as below
/*
  var arrowGreeting = (message, name) => {
  return message + name;
  }
*/

var arrowGreeting = (message, name) => message + name;


console.log(createGreeting('Hi, ', 'John'));
console.log(arrowGreeting('Hi, ', 'John'));

//returns square of x. ex: square(2); console: 4
var square = x => x * x;
console.log(square(3)); // 9


//this context.
var deliveryBoy = {
  name: 'John',
  handleMessage: function(message, handler) {
    handler(message);
  },
  receive: function() {
    /*
      var that = this;
      this.handleMessage("Hello, ", function(message) {
      that.name;//get the proper name
      console.log(message + that.name);
      });*/
    this.handleMessage("Hello, ", message => console.log(message + this.name)); //Here this refers to deliveryBoy context. Lexical scope.
  }
}

deliveryBoy.receive();

// https://stackoverflow.com/questions/28029523/
var test = {
  firstname: 'David',
  fn: function() {
    return ['one', 'two', 'tree'].map(() => this.firstname)
  }
}
console.log(test.fn())
// ["David", "David", "David"]


console.log('\n---------- ES6 higher order function');
// ES6 higher order function example
// https://es6console.com/jdtnm5eu/
// https://news.ycombinator.com/item?id=12096522
// https://jsfiddle.net/jpsierens/obgh1uc7/2/

const add = (x, y) => x+y;
const multiply = (x, y) => x*y;
const withLogging = (wrapped) => (x, y) => console.log(wrapped(x, y));

const addWithLogging = withLogging(add);
const multWithLogging = withLogging(multiply);

addWithLogging(2, 3); //5
multWithLogging(2, 3); //6
