// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/encodeURI

console.log(decodeURI("%D1%88%D0%B5%D0%BB%D0%BB%D1%8B"));
console.log(decodeURI("%3A"));
console.log(decodeURI("SP%3ASPX"));
// шеллы
// %3A
// SP%3ASPX

console.log(decodeURIComponent("%3A"));
console.log(decodeURIComponent("SP%3ASPX"));
// :
// SP:SPX


