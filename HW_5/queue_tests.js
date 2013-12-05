var k = require("./queue");
var a = k.create();
a.add("Derp");
console.log(a.remove());

//	Code below will not work, 
//	data can not be accessed 
//	from outer scope.

/* a.data([1,2,3,4,5]);
 * console.log(a.remove());
 *
 * Example error message:
 *	TypeError: Property 'data' of object #<Object> is not a function
 */