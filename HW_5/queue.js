/*
 * A simple queue data type. Queues are entirely mutable.  The
 * add and remove methods are simple conveniences, but clients
 * are free to mess with the queue in any way they like.
 */

var Queue = {};

Queue.prototype = {
    add: function (x) {this.data.push(x);},
    remove: function () {return this.data.shift();}
};    

module.exports.Queue.create = function () {
    var q = Object.create(Queue.prototype);
    q.data = [];
    return q;
};

test("Queue test", function () {
    var q1 = Queue.create();
    var q2 = Queue.create();
    q1.add(0);
    q1.add(1);
    q1.add(2);
    q2.add(3);
    equal(q1.remove(), 0);
    equal(q2.remove(), 3);
    equal(q1.remove(), 1);
});
