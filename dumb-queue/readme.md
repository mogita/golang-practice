Just a dumb and simple queue implementation in go.

Inspired by an interview question: implement a FIFO queue without using an array or list, with these methods:

- enqueue: add a new item to the end of the queue
- dequeue: remove the first item of the queue
- size: return the length of the queue

I added an extra show() method to reveal the data of the queue. It's equivalent to accessing `q.data` but by exposing a method rather than accessing the field directly makes the API more stable in spite of any breaking changes to the implementation of the Queue struct. Yep, future proof.