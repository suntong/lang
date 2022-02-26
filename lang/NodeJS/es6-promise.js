
console.log('\n---------- ES6 promise');
// ES6 Promise Example
// https://jsfiddle.net/jpsierens/5g0d0ww7/9/

// server
// -----------------------------------------------------
// Imagine this array lives in your server
const people = [
  {name: 'John Doe', age: '30'},
  {name: 'Jane Doe', age: '24'}
];

// client
// -----------------------------------------------------

// function that simulates an AJAX call to our
// imaginary server, where the people array lives.
// The call may succeed or fail.
const simulateAJAXCall = (index, cb) => {
  let chance = Math.random();
  let isSuccess= (chance <= 0.80) ? true : false;
  
  return setTimeout(() => {
    if (isSuccess) {
      return cb(people[index]);
    }
    else {
      return cb(null);
    }
  }, 1000)
}

// Function that returns a promise to search
// a person according to his index
const getPerson = (index) => {
  
  // The promise resolves if the data comes back
  // successfully and it rejects in case of an error
  return new Promise((resolve, reject) => {
    
    console.log('fetching person...');
    simulateAJAXCall(index, (data) => {
      if (data === null) {
	reject('Something went wrong!');
      }
      resolve(data);
    });
  });
}

// Call getPerson, which will return a promise.
// If the promise resolves, then log the person,
// if there's an error, it will go to the catch
// block and be handled there.
getPerson(1).then(person => {
  console.log(person);
}).catch(error => {
  console.log(error);
});

