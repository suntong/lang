
function update(tgt, src) {
  let target = tgt;
  for (let prop in src) {
    if (src.hasOwnProperty(prop)) {
      target[prop] = src[prop];
    }
  }
  return target;
}


console.log("Test 1")
{
const source1 = {a:1, b:2, c:3}
let target = source1
console.log(source1, target)

const source2 = {b:3, d:5}
target = source2
console.log(source1, source2, target)
// { a: 1, b: 2, c: 3 } { b: 3, d: 5 } { b: 3, d: 5 }
}

console.log("Test 2")
{
const source1 = {a:1, b:2, c:3}
let target = source1
console.log(source1, target)

const source2 = {b:3, d:5}
target = update(target, source2)

console.log(source1, source2, target)
// { a: 1, b: 3, c: 3, d: 5 } { b: 3, d: 5 } { a: 1, b: 3, c: 3, d: 5 }
console.log(source1 == target)
console.log(source1 === target)
// both true
}

// Merging objects (associative arrays)
// https://stackoverflow.com/questions/929776/merging-objects-associative-arrays
// merging two associative arrays

obj1 = {a: 1, b: 2}; obj2 = {a: 4, c: 110};
obj3 = Object.assign({},obj1, obj2); // Object {a: 4, b: 2, c: 110}
console.log(obj3)



