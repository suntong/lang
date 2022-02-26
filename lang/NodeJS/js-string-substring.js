// https://stackoverflow.com/questions/10338357/string-manipulation-with-javascript-nodejs

// to get the first # characters
console.log("asdf".substring(0,2))
console.log('鸿蚨'.substring(0,1))
console.log('鸿蚨鸿蚨'.substring(0,3))

// to remove the first # characters
console.log("0123456789ABCDEF".slice(6))
console.log('鸿蚨'.slice(1))

// Check if a string contains "world"
// https://www.w3schools.com/jsref/jsref_includes.asp
{
  let text = "Hello world, welcome to the universe.";
  let result = text.includes("world");
  console.log(result)
  console.log(text.substring(6,11)) // NOT: substring(6,5) !
  result = text.includes("world", 6);
  console.log(result)
}

{
  let text = "Hello World, welcome to the universe.";
  let result = text.includes("world", 6);
  console.log(result)
}
