// https://es6console.com/jdtnm5eu/

const short = x => x.toISOString()
      .replace(/^.*T/,'')
      .replace(/:[^:]+$/,'')

let t = new Date()
console.log(t)
console.log(short(t))

