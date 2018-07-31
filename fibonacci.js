// Fibonacci from : https://medium.com/developers-writing/fibonacci-sequence-algorithm-in-javascript-b253dc7e320e

//var bigInt = require("big-integer");

function fibonacciold(num, memo) {
  memo = memo || {};

  if (memo[num]) return memo[num];
  if (num <= 1) return 1;

  return memo[num] = fibonacci(num - 1, memo) + fibonacci(num - 2, memo);
}

function fibonacci(num){
  var a = bigInt.one, b = bigInt.zero, temp;

  while (num >= 0){
    temp = a;
    a = a.plus(b);
    b = temp;
    num--;
  }

  return b.toString();
}
