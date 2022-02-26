// https://www.w3schools.com/js/js_regexp.asp

console.log('\n---------- String search()');
// Using String search() With a Regular Expression
var str = "Visit W3Schools"
var n = str.search(/w3schools/i)
console.log(n) // 6

// Using String search() With String
var str = "Visit W3Schools!"
var n = str.search("W3Schools")
console.log(n) // 6
console.log(str.search("nono")) // -1

str = "Visit this school!"
console.log(str.search("i")) // 1

console.log('\n---------- String replace()');
// Use String replace() With a Regular Expression
var str = "Visit Microsoft!"
var res = str.replace(/microsoft/i, "W3Schools")
console.log(res)

// Using String replace() With a String
var res = str.replace("Microsoft", "W3Schools")
console.log(res)

console.log('\n---------- String match()');
var re = /(\w+)\s(\w+)/;
var str = 'John Smith';
console.log(str.match(re));
console.log('John'.match(re));

var newstr = str.replace(re, '$2, $1');
console.log(newstr);
if (m = str.match(re), m) console.log(m[1], m)
console.log('\n---------- Array.slice()');
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/slice
var animals = ['ant', 'bison', 'camel', 'duck', 'elephant'];

console.log(animals.slice(2));
// expected output: Array ["camel", "duck", "elephant"]

console.log(animals.slice(2, 4));
// expected output: Array ["camel", "duck"]

console.log(animals.slice(1, 5));
// expected output: Array ["bison", "camel", "duck", "elephant"]


console.log('\n---------- String match() & replace()');
var re = /(\w+)\s(\w+)/;
var str = 'John Smith';
console.log(str.match(re));
console.log('John'.match(re));
var newstr = str.replace(re, '$2, $1');
console.log(newstr);
// "Smith, John"

//let m = str.match(re)
if (m = str.match(re), m) console.log(m[1], m)

var cmd = '/listall LU ';
var reUrls = new RegExp(`listall `);
console.log(reUrls.test(cmd))
console.log(cmd.match(reUrls))
reUrls = new RegExp(`listall (\w+)`);
console.log(reUrls.test(cmd))
console.log(cmd.match(reUrls))

reUrls = new RegExp(`listall (\\w+)`);
console.log(reUrls.test(cmd))
console.log(cmd.match(reUrls))

if (m = cmd.match(reUrls)) console.log(m[1], m)

console.log('\n---------- String emoj replace()');
// https://stackoverflow.com/questions/22935442/javascript-regular-expression-and-replace
var s = "<em>Rs</em> 154,451. Hello world, <em>Rs</em> 15,51.";
console.log(s.replace(/<em>Rs<\/em>([^.]+)/g, '<b>Rs$1</b>'))
console.log(s.replace(/<em>(Rs)<\/em>(.*?)(\.)/g, '<b>$1$2</b>$3'))
// <b>Rs 154,451</b>. Hello world, <b>Rs 15,51</b>.
// <b>Rs 154,451</b>. Hello world, <b>Rs 15,51</b>.


const content = 'ABC <img class="qqemoji qqemoji28" text="[憨笑]_web" src="/zh_CN/htmledition/v2/images/spacer.gif" /> Hello world, <img class="qqemoji qqemoji21" text="[愉快]_web" src="/zh_CN/htmledition/v2/images/spacer.gif" /> <img class="emoji emoji1f633" text="_web" src="/zh_CN/htmledition/v2/images/spacer.gif" /> DDD <img class="qqemoji qqemoji83" text="[抱拳]_web" src="/zh_CN/htmledition/v2/images/spacer.gif" /> <img class="qqemoji qqemoji13" text="[呲牙]_web" src="/zh_CN/htmledition/v2/images/spacer.gif" />'
//console.log(content)

console.log(content.replace(/<img class="[q]*emoji .*?text="(.*?)_web" src=.*?" \/>/g, '$1'))
// ABC [憨笑] Hello world, [愉快]  DDD [抱拳] [呲牙]


console.log('\n---------- RegExp Object');

// Using the RegExp Object
var patt = /e/
console.log(patt.test("The best things in life are free!"))
// Since there is an "e" in the string, the output of the code above will be: true

const config = {
  message: {
    watchRE: /[bf]e/,
    notMatch: /xx/
  }
}
console.log(config.message.watchRE.test("The best things in life are free!"))
console.log(config.message.notMatch.test("The best things in life are free!"))

console.log('\n---------- RegExp exec()');
// Using exec()
// The exec() method is a RegExp expression method.
// It searches a string for a specified pattern, and returns the *first* found text.

// /e/.exec("The best things in life are free!")
console.log(/[bfr]e/.exec("The best things in life are free!")); // "be"
console.log(config.message.watchRE.exec("The best things in life are free!")) // "be"

// executes a search for a match in a string. It returns an array of information of the *first* match or null on a mismatch.
const execR = /([bfr]e).*?([bfr]e).+.+([bfr]e)/.exec("The best things in life are free!")
console.log(execR)
// produces: 
// ["best things in life are fre", "be", "fe", "re"]
console.log(execR.slice(1))
// ["be", "fe", "re"]

// https://stackoverflow.com/questions/7735124/
console.log('\n---------- RegExp dynamic value');

{
str1 = "pattern"
const re = new RegExp(str1)

let testStr = "pattern matching ."
let found = re.test(testStr)
console.log(found) // true

var re2 = new RegExp(str1, "g")
console.log("pattern matching .".replace(re2, "regex")) // regex matching .
}

var num = 3;
var patt = '%' + num + ':';
var testVar = '%3:';
var result = patt.match(testVar);
console.log(result) // [ '%3:', index: 0, input: '%3:', groups: undefined ]

console.log('\n---------- JavaScript/Guide/Regular_Expressions');
//
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Guide/Regular_Expressions
//

var re = /(\w+)\s(\w+)/;
var str = 'John Smith';
console.log(str.match(re));
console.log('John'.match(re));
var newstr = str.replace(re, '$2, $1');
console.log(newstr);
// "Smith, John"

var re = /\w+\s/g;
// or var re = new RegExp('\\w+\\s', 'g');
var str = 'fee fi fo fum';
var myArray = str.match(re);
console.log(myArray);
// ["fee ", "fi ", "fo "]

var re = new RegExp('\\w+\\s', 'g')
var xArray; while(xArray = re.exec(str)) console.log(xArray);
// produces: 
// ["fee ", index: 0, input: "fee fi fo fum"]
// ["fi ", index: 4, input: "fee fi fo fum"]
// ["fo ", index: 7, input: "fee fi fo fum"]

// Examples
// The following examples show some uses of regular expressions
// The name string contains multiple spaces and tabs,
// and may have multiple spaces between first and last names.
var names = 'Orange Trump ;Fred Barney; Helen Rigby ; Bill Abel ; Chris Hand ';

var output = ['---------- Original String\n', names + '\n'];

// Prepare two regular expression patterns and array storage.
// Split the string into array elements.

// pattern: possible white space then semicolon then possible white space
var pattern = /\s*;\s*/;

// Break the string into pieces separated by the pattern above and
// store the pieces in an array called nameList
var nameList = names.split(pattern);

// new pattern: one or more characters then spaces then characters.
// Use parentheses to "memorize" portions of the pattern.
// The memorized portions are referred to later.
pattern = /(\w+)\s+(\w+)/;
console.log(str.match(pattern));

// Below is the new array for holding names being processed.
var bySurnameList = [];

// Display the name array and populate the new array
// with comma-separated names, last first.
//
// The replace method removes anything matching the pattern
// and replaces it with the memorized string—the second memorized portion
// followed by a comma, a space and the first memorized portion.
//
// The variables $1 and $2 refer to the portions
// memorized while matching the pattern.

output.push('---------- After Split by Regular Expression');

var i, len;
for (i = 0, len = nameList.length; i < len; i++) {
  output.push(nameList[i]);
  bySurnameList[i] = nameList[i].replace(pattern, '$2, $1');
}

// Display the new array.
output.push('---------- Names Reversed');
for (i = 0, len = bySurnameList.length; i < len; i++) {
  output.push(bySurnameList[i]);
}

// Sort by last name, then display the sorted array.
bySurnameList.sort();
output.push('---------- Sorted');
for (i = 0, len = bySurnameList.length; i < len; i++) {
  output.push(bySurnameList[i]);
}

output.push('---------- End');

console.log(output.join('\n'));

// Escape Newline/Paragraph separators
// https://github.com/expressjs/express/issues/1132

var body = JSON.stringify({
    username: 'thefourtheye',
    password: 'hello\u2028world'
});

let a = `This is a very long 
multi 
line 
string which might be used to 
display some text. 
It has much more than 80 symbols 
so it would take more then one screen 
in your text editor to view it.`

body = a+'hello\n\u2028world'
console.log(body)
console.log(body.replace(/\n/g, '\\n').replace(/\u2028/g, '\\u2028').replace(/\u2029/g, '\\u2029'))

