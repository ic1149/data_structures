# Data Structures

My attempt to implement different data structures, e.g. linked lists, in go.

## Linked Lists

- consists of nodes and a head pointer
- each nodes has data and pointer to next node
- null pointer at last item
- double linked list
- circular linked list

### Linked List methods

- output the whole list
- check if an item exists
- delete an item
- append an item

## Queues

- first in first out
- e.g. printer queue
- circular queue
- priority queue: allows jumping

### Queue methods

- enqueue: append item to back of queue
- peek: output first item
- dequeue (output and remove) first item

## Stacks

- last in first out
- e.g. undo redo

### Stack methods

- push: append item on top
- peek: output item on top
- pop (output and remove) item on top

## Graphs

- nodes (verticies)
- edges (arcs)
- directed or undirected
- Dijkstra and A Star
- A Star has heuristics
- Graphs can have loops/cycles but trees can't
- Weights
