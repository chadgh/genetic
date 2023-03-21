# genetic

A basic genetic algorithm solution in golang.

There are two reference implementations. The 8-queens problem and a business earnings problem (catdog).

The eightqueens director is not used but was used as a starting point to creating the general solution.


## Usage

genetic provides the following for writing a genetic algorithm solution.

* `github.com/chadgh/genetic/genetic/types`: provides type primitives like Organism, Population, Alphabet for
constructing the genetic algorithm solution.
* `github.com/chadgh/genetic/genetic/strategies`: provides a basic genetic algorithm implementation. This can 
be easily extended to implement more complex or specific variations for a specific problem.
* `github.com/chadgh/genetic/genetic`: provides the main evolve functions for running a genetic algorithm.

