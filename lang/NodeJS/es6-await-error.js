// Error handling with Async/Await in JS
// https://itnext.io/error-handling-with-async-await-in-js-26c3f20bc06a

console.log('\n---------- simple try...catch');
function thisThrows() {
  throw new Error("Thrown from thisThrows()");
}

try {
  thisThrows();
} catch (e) {
  console.error(e);
} finally {
  console.log('We do cleanup here');
}
// Output:
// Error: Thrown from thisThrows()
//   ...stacktrace
// We do cleanup here


console.log('\n---------- throw from try');
// https://stackoverflow.com/questions/38050857/what-happen-to-return-statement-in-catch-block
function example0() {
  try {
    throw new Error()
    return 1;
  }
  catch(e) {
    return 2;
  }
  console.log(5)
  return 4;
}

console.log(example0()); // 2! not 4

function example() {
  try {
    throw new Error()
    return 1;
  }
  catch(e) {
    return 2;
  }
  finally { 
    return 3;
  } 
  console.log(5)
  return 4;
}

console.log(example());
// only outputs 3! 5 is never printed since after finally block value is returned


console.log('\n---------- rejects with a promise');
// Now we modify thisThrows() so it actually rejects with a promise instead of a regular error
async function asyncThrows() {
  throw new Error("Thrown from thisThrows()");
}

try {
  asyncThrows();
} catch (e) {
  console.log('NEVER reaches here!!!!!');
  console.error(e);
} finally {
  console.log('We do cleanup here');
}
// output:
// We do cleanup here
// UnhandledPromiseRejectionWarning: Error: Thrown from asyncThrows()

// Now we have the classic problem, thisThrows returns a rejecting promise, so the regular try...catch is not able to catch the error.

console.log('\n---------- *await promise rejection*');
async function arun() {
  let varInTry;
  try {
    // let // ReferenceError: varInTry is not defined
    varInTry = "Val from within try";
    await thisThrows();
    return 1;
  } catch (e) {
    console.error(e);
    return 2;
  } finally {
    console.log(varInTry);
    console.log('We do cleanup here');
    return 3;
  }
}

console.log(arun());
// Output:
// Error: Thrown from thisThrows()
//   ...stacktrace
// We do cleanup here

console.log('\n---------- with a .catch() call');
asyncThrows() // 89:1
    .catch(console.error)
    .then(() => console.log('We do cleanup after a .catch() call at 89:1\n'));
// Output:
// Error: Thrown from thisThrows()
//   ...stacktrace
// We do cleanup after a .catch() call

// Both solutions work fine, but the async/await one is easier to reason about
console.log("No output here but at program ends");

// console.log('\n---------- await with a .catch() call');
// SyntaxError: await is only valid in async function
// const r = await asyncThrows() // 81:1
//     .catch(console.error)
//     .then(() => console.log('We do cleanup after a .catch() call at 81:1\n'));
// console.log("Return from asyncThrows()", r);

console.log('\n---------- !returning from an async function!');
async function myFunctionThatCatches0() {
  try {
    return thisThrows();
  } catch (e) {
    console.log('We got a UnhandledPromiseRejection from an async function');
    console.error(e);
  } finally {
    console.log('We do cleanup from async function0');
  }
  return "Nothing found 0";
}

async function run0() {
  const myValue = await myFunctionThatCatches0();
  console.log(`From run0: ${myValue}`);
}

run0();
// we got a UnhandledPromiseRejection !


console.log('\n---------- *fixed UnhandledPromiseRejection*');
async function myFunctionThatCatches() {
  try {
    return await thisThrows();
  } catch (e) {
    console.error(e);
  } finally {
    console.log('We do cleanup from an async function');
  }
  return "Nothing found";
}

async function run() {
  const myValue = await myFunctionThatCatches();
  console.log(`From run: ${myValue}`);
}

run();

console.log('\n---------- !wraps in a new error!');
function myFunctionThatCatches1() {
  try {
    return thisThrows();
  } catch (e) {
    throw new TypeError(e.message);
  } finally {
    console.log('We do cleanup from myFunctionThatCatches1');
  }
}

async function run1() {
  try {
    await myFunctionThatCatches1();
  } catch (e) {
    console.error(e);
  }
}

run1();
// Outputs:
// We do cleanup here
// TypeError: Error: Thrown from thisThrows()
//    at myFunctionThatCatches (/repo/error_stacktrace_1.js:9:15) <-- Error points to our try catch block
//    at run (/repo/error_stacktrace_1.js:17:15)
//    at Object.<anonymous> (/repo/error_stacktrace_1.js:23:1)

console.log('\n---------- *re-throw original error*');
function myFunctionThatCatches2() {
  try {
    return thisThrows();
  } catch (e) {
    // Maybe do something else here first.
    throw e;
  } finally {
    console.log('We do cleanup from myFunctionThatCatches2');
  }
}

async function run2() {
  try {
    await myFunctionThatCatches2();
  } catch (e) {
    console.error(e);
  }
}

run2();
// Outputs:
// We do cleanup here
// Error: Thrown from thisThrows()
//     at thisThrows (/repo/error_stacktrace_2.js:2:11) <-- Notice we now point to the origin of the actual error
//     at myFunctionThatCatches (/repo/error_stacktrace_2.js:7:16)
//     at run (/repo/error_stacktrace_2.js:18:15)
//     at Object.<anonymous> (/repo/error_stacktrace_2.js:24:1)


console.log('\n---------- End');

/*

Summary

- We can use try...catch for synchronous code.
- We can use try...catch (in combination with async functions) and the .catch() approaches to handle errors for asynchronous code.
- When returning a promise within a try block, make sure to await it if you want the try...catch block to catch the error.
- Be aware when wrapping errors and rethrowing, that you lose the stack trace with the origin of the error.

*/
