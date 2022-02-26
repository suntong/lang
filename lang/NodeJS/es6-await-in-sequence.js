
console.log('\n---------- await in sequence');
// https://jsbin.com/somusav/1/edit?js,console
// https://stackoverflow.com/questions/37576685/using-async-await-with-a-foreach-loop/

// https://zellwk.com/blog/async-await-in-loops/
const sleep = ms => {
  return new Promise(resolve => setTimeout(resolve, ms))
}

const fruitBasket = {
  apple: 27,
  grape: 0,
  pear: 14
}

const getNumFruit = fruit => {
  return sleep(10).then(v => fruitBasket[fruit])
}

getNumFruit('apple')
  .then(num => console.log(num))
// 27


console.log("= = = = = = = = = = = = ")
console.log("a", "b")

const fruitsToGet = ['apple', 'grape', 'pear']

const showNumFruit = fruit => {
  const w = Math.random() * 1000
  return sleep(w).then(v => console.log(`${fruit} ${fruitBasket[fruit]} ${w}`))
}

console.log("= = = = = = = = = = = = ")
console.log("reduce")
// https://stackoverflow.com/a/24985483/2125837
function showNum() {
  const parts = fruitsToGet

  parts.reduce(function(promise, part) {
    return promise.then(function() {
      console.log(part)
      showNumFruit(part)
    });
  }, Promise.resolve());
  return sleep(1)
}
showNum().catch((error) => { console.log(error) })


const retNumFruit = fruit => {
  const w = Math.random() * 1000
  return sleep(w).then(v => `<= ${fruit} ${fruitBasket[fruit]} ${w}`)
}

async function showSeq() {
  const parts = fruitsToGet
  for (const part of parts) {
    const r = await retNumFruit(part)
    console.log(part, r)
  }
}

sleep(6000)
console.log("= = = = = = = = = = = = ")
console.log("for loop")
showSeq()

/*

---------- await in sequence
= = = = = = = = = = = = 
a b
= = = = = = = = = = = = 
reduce
= = = = = = = = = = = = 
for loop
apple
grape
pear
27
grape 0 352.6851929560979
apple 27 438.00829178416524
pear 14 777.3110776126011
apple <= apple 27 994.801058969091
grape <= grape 0 503.13152208625047
pear <= pear 14 861.8505720860405

*/
