// https://jsbin.com/cubarot/edit?js,console,output

function getCoordinates(element) {
  let x, y, z;
  // X: { x, y, z } = { 1, 2, 3 }
  x = 1
  y = 2
  z = 3
  // do stuff to get coordinates
  return [x, y, z];
}

{
  const [x, y, z] = getCoordinates();
  console.log(x, y, z) // 1 2 3
}

{
  const [x1, y2, z3] = getCoordinates();
  console.log(x1, y2, z3) // 1 2 3
  // .: OK to rename
}

function getXYZ() {
  let x, y, z;
  // X: { x, y, z } = { 1, 2, 3 }
  [ x, y, z ] = [ 1, 2, 3 ]
  // do stuff to get coordinates
  return {x, y, z};
  // same as 
  //return {x:x, y:y, z:z};
}

{
  const {x, y, z} = getXYZ();
  console.log(x, y, z) // 1 2 3
}

{
  const {x, y1, z} = getXYZ();
  console.log(x, y1, z) // 1 undefined 3
  // .: NOK to rename
}

{
  const {x, y:y1, z} = getXYZ();
  console.log(x, y1, z) // 1 2 3
  // proper way to rename
}

