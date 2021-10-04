let address = '0xab1960922acf3dd1ec61413619f6a6c37c1e6435'
let unixTime = 1632972933669
let msg = `${address}:${unixTime}`


let signedTx = '0x71068725fee8010ca8920ed7712d629ba0a6798afed5e9090bc7ffa3fa3a7ac933001e3bf8c3d34842561e6ba18902b287b4e1922ba2da9c393850c16563afb41c'

let utils = require('ethereumjs-utils')


let s1 = utils.keccak256(msg)
console.log(s1.toString())