module.exports = function random(start, end, fixed = 0){
  let differ = end - start
  console.log(differ)
  let random = Math.random()
  return (start + differ * random).toFixed(fixed)
}