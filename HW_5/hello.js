Queue = {};
Queue.prototype = {
    add: function (x) {this.data.push(x);},
    remove: function () {return this.data.shift();}
};    
Queue.create = function () {
    var q = Object.create(Queue.prototype);
    q.data = [];
    return q;
};


