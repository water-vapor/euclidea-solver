[![Go Report Card](https://goreportcard.com/badge/github.com/water-vapor/euclidea-solver)](https://goreportcard.com/report/github.com/water-vapor/euclidea-solver)
[![Build Status](https://travis-ci.org/water-vapor/euclidea-solver.svg?branch=master)](https://travis-ci.org/water-vapor/euclidea-solver)
[![GoDoc](https://godoc.org/github.com/water-vapor/euclidea-solver?status.svg)](https://godoc.org/github.com/water-vapor/euclidea-solver)
# Eulidea Solver
Euclidea Solver is a brute force search program aimed to search for short solutions of straightedge and compass construction problems in [Euclidea](https://www.euclidea.xyz). 
It performs searches by traveling the whole sensible search space with depth-first search.
### Usage
`euclidea-solver [-p] [-l level] [-t thread] [-v version] chapter_number problem_number`

`chapter number` and `problem number` refers to corresponding problems in the website. `version` refers to `E` or `L` goal of a problem, the default is `E`. Sequential search is used by default.
Flag `-p` turns on parallel searching. Flag `-l` refers to the level in DFS, on which the go routines will be spawned in parallel mode. The number of go routines spawned will grow with level but 
not predictable. Flag `-t` limits the number of go routines running simultaneously, preventing the large overhead if too many go routines are spawned when level is deep. 

### Adding Problems
All the problems should be manually coded into `problem` package, I've only added a few examples (all examples are tested to be solvable), 
but I believe you can figure out how to describe a problem by looking at them. Don't forget to add the function name to `problem.go` 
after adding a custom function. A detailed guide may be added later.

### Performance
The current search algorithm does not include optimization or branch pruning of any kind. If you have any good ideas for 
a certain problem, just include them in the problem statement, and that will greatly reduce the search time.

### Example
Here is an example to solve problem 15.10: Angle of 3 Degree.
```$bash
$./euclidea-solver -p -l=4 15 10
13 go routines deployed.
13 go routines deployed.
3 go routines deployed.
22 go routines deployed.
8 go routines deployed.
Solution found!
Number of boards searched: 112902489
Took 2m27.098416518s
```
A folder `15.10 Angle of 3 Degree_E_1561528019` is created, with the following images inside:
<img src="images/15.10&#32;Angle&#32;of&#32;3&#32;Degree_E_1561528019/Step_1.png" alt="step_1" width="500"/>
<img src="images/15.10&#32;Angle&#32;of&#32;3&#32;Degree_E_1561528019/Step_2.png" alt="step_2" width="500"/>
<img src="images/15.10&#32;Angle&#32;of&#32;3&#32;Degree_E_1561528019/Step_3.png" alt="step_3" width="500"/>
<img src="images/15.10&#32;Angle&#32;of&#32;3&#32;Degree_E_1561528019/Step_4.png" alt="step_4" width="500"/>
<img src="images/15.10&#32;Angle&#32;of&#32;3&#32;Degree_E_1561528019/Step_5.png" alt="step_5" width="500"/>
<img src="images/15.10&#32;Angle&#32;of&#32;3&#32;Degree_E_1561528019/Step_6.png" alt="step_6" width="500"/>
<img src="images/15.10&#32;Angle&#32;of&#32;3&#32;Degree_E_1561528019/Step_7.png" alt="step_7" width="500"/>
