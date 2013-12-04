//	http://jsfiddle.net/72SGF/1/ 


var x = 10;

function determine(a, b) {
    alert(x);
}

function half() {
	x = x / 2;
}

function addFour() {
	x = x + 4;
}

determine(addFour(), half());