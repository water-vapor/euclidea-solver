# Eulidea Solver
Euclidea Solver is a brute force search program aimed to search for short solutions of straightedge and compass construction problems in [Euclidea](https://www.euclidea.xyz). 
It performs searches by traveling the whole sensible search space with depth-first search.
### Usage
`euclidea-solver [-p[=parallel level]] <chapter number> <problem number>`

`chapter number` and `problem number` refers to corresponding problems in the website. Sequential search is used by default.
`-p` is equivalent to `-p=1`, this parallel level refers to the level in DFS. The number of go routines spawned is calculated 
automatically and printed once the program starts, this number is not fully predictable but will grow with the parallel level 
specified.

### Adding Problems
All the problems are manually coded into `problem` package, I've only added a few examples (all examples are tested to be solvable), 
but I believe you can figure out how to describe a problem by looking at them. Don't forget to add the function name to `problem.go` 
after adding a custom function. A detailed guide may be added later.

### Performance
The current search algorithm does not include optimization or branch pruning of any kind. If you have any good ideas for 
a certain problem, just include them in the problem statement, and that will greatly reduce the search time.

If you plan to use parallel searching, make sure the total number of go routine spawned is not much larger than your number of processor. 
Otherwise it will be very slow.