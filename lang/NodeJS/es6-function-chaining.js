console.log('\n---------- Function composition / pipe');
// https://es6console.com/jjnj87up/

// Function Composition in JavaScript with Pipe
// https://vanslaars.io/post/create-pipe-function/

const inc = (num) => num+1
const dbl = (num) => num*2
const sqr = (num) => num*num

// Pipe
const _pipe = (f, g) => (...args) => g(f(...args))
const pipe = (...fns) => fns.reduce(_pipe)

const incDblSqr = pipe(inc, dbl, sqr)
const result = incDblSqr(2)

console.log(result); // 36


console.log('\n---------- Function chainable / cascading');
// Chainable (Cascading) Methods
// https://es6console.com/jjnisk0b/

// http://javascriptissexy.com/beautiful-javascript-easily-create-chainable-cascading-methods-for-expressiveness/

// The data store:
const usersData = [
  {firstName:"tommy", lastName:"MalCom", email:"test@test.com", id:102},
  {firstName:"peTer", lastName:"brecHt", email:"test2@test2.com", id:103},
  {firstName:"RoHan", lastName:"sahu", email:"test3@test3.com", id:104}
];

function titleCaseName(str)
{
  return str.replace(/\w\S*/g, function(txt){return txt.charAt(0).toUpperCase() + txt.substr(1).toLowerCase();});
}

// Our object with the chainable methods
const userController = {

  currentUser:"",
  
  // https://medium.com/javascript-scene/javascript-factory-functions-with-es6-4d224591a8b1
  findUser (userEmail) {
    var arrayLength = usersData.length, i;
    for (i = arrayLength - 1; i >= 0; i--) {
      if (usersData[i].email === userEmail) {
	this.currentUser = usersData[i];
	break;
      }
    }
    return this;
  },

  formatName () {
    if (this.currentUser) {
      this.currentUser.fullName = titleCaseName (this.currentUser.firstName) + " " + titleCaseName (this.currentUser.lastName);
    }
    return this;

  },

  createLayout () {
    if (this.currentUser) {
      this.currentUser.viewData = "<h2>Member: " + this.currentUser.fullName + "</h2>"  + "<p>ID: " + this.currentUser.id + "</p>" + "<p>Email: " + this.currentUser.email + "</p>";
    }
    return this;
  },

  displayUser () {
    if (!this.currentUser) return;

    //$(".members-wrapper").append(this.currentUser.viewData);
    console.log(this.currentUser.viewData);
  }
};

// Now, use the chaninable methods with expressiveness:
userController.findUser("test2@test2.com").formatName().createLayout().displayUser();
// <h2>Member: Peter Brecht</h2><p>ID: 103</p><p>Email: test2@test2.com</p>
