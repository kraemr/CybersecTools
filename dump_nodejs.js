// you can use this to write strings to a file
require('fs').writeFileSync("/tmp/pwned","test")
// or this if its an object
require('fs').writeFileSync("/tmp/pwned",require('util').inspect(Object))
