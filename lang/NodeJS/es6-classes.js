console.log('\n---------- From constructors to classes');
// 4.13 From constructors to classes
// http://exploringjs.com/es6/ch_core-features.html#sec_from-constr-to-class

// In ES5, you implement constructor functions directly
{
  function Person(name) {
    this.name = name;
  }

  Person.prototype.describe = function () {
    return '(ES5) Person called '+this.name;
  };

  var p1 = new Person('jeff');
  console.log(p1.describe());
}

// In ES6, classes provide slightly more convenient syntax for constructor functions:
{

  class Person {
    constructor(name) {
      this.name = name;
    }
    describe() {
      return '(ES6) Person called '+this.name;
    }
  }

  var p2 = new Person('joe');
  console.log(p2.describe());

}

console.log('\n---------- Object Literal Property Value Shorthand');
// http://jsfiddle.net/GzYJ6/
var a = ['zero', 'one', 'two'];

//console.clear();
(() => console.log('IIFE Arrow Inline'))();

// var document;
// ((win, doc) => {
//   win.console.log('IIFE Arrow Block');
// })(this, document);

a.forEach(n => console.log(n));// Arrow Inline

var fn = () => console.log("fn");
fn();    // Logs 'fn'

var o = {
  _name: 'foo',
  //fn,    // This doesn't work [1]
  square: n => n * n,
  printName: () => { console.log(this._name); },        // this refers to window [2]
  printName2: function () { console.log(this._name); }  // this refers to o
};
console.log(o.square(2)); // Logs '4'
o.printName();            // Logs 'undefined'
o.printName2();           // Logs 'foo'

// [1]: http://ariya.ofilabs.com/2013/02/es6-and-object-literal-property-value-shorthand.html
// [2]: The declaration works but the 'this' syntax is to window and not the object literal


console.log('\n---------- OO');
// https://jsbin.com/duluji/edit?js,console,output
class Entity {
  constructor(name) {
    this.id = Entity.newId();
    //Object.seal(this); // this has the side effect of preventing inheritence 
  }

  persist() {
    console.log('Entity::persist(', this.id, ')');
  }

  // not so sure about this, might be best served by a thing that does this exclusively
  static newId() {
    if (typeof Entity.newId.__value === 'undefined') {
      Entity.newId.__value = 0;
    }

    return ++Entity.newId.__value;
  }
}

class Person extends Entity {
  constructor(name) {
    super();
    this.name = name;
  }
}

class Employee extends Person {
  constructor(name, title) {
    super(name);
    this.title = title;
  }
}

var p1 = new Person('jeff');
var p2 = new Person('joe');
var p3 = new Employee('john', 'ceo');

console.log(p1, p1 instanceof Entity, p1 instanceof Person, p1 instanceof Employee); // Person { id: 1, name: 'jeff' } true true false
console.log(p2); // Person { id: 2, name: 'joe' }
console.log(p3); // Employee { id: 3, name: 'john', title: 'ceo' }


console.log('\n---------- yield & pause');
// https://jsbin.com/wakutos/edit?js,console

var name = "trae";
console.log(`My name is ${name}`);

function *foo() {
  var x = 1;
  console.log("Before Pause " + x)
  x = (yield x);
  console.log("After Pause " + x)
  return x;
}

var it = foo();
console.log(it.next(2));
console.log(it.next(3));
