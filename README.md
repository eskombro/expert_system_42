# expert_system_42

[Project specifications PDF](./docs/expert-system.en.pdf)

Implementation of a simple backward-chaining inference engine.  

Program accepts one file as parameter, which is the input file.  
This file should contain a list of rules, then a list of initial facts and queries.  
It can contain comments preceded by `#`

## Usage

```
usage: Expert System [-h|--help] [-f|--file "<value>" [-f|--file "<value>" ...]]
                    [-v|--verbose1] [-V|--verbose2]

                    Expert System | 42 Paris

Arguments:

  -h  --help      Print help information
  -f  --file      list of paths to file
  -v  --verbose1  Launch program with verbose level 1
  -V  --verbose2  Launch program with verbose level 2
```
  
## Examples

- Single file  
`go run expert_system -f test/tests/and`  

- Multiple files  
`go run expert_system -f test/tests/and -f test/tests/basic_biconditional`  

- Verbose  
`go run expert_system -f test/tests/parenthesis -V`  

## Input file example  



```
# this is a comment$
# all the required rules and symbols, along with the bonus ones, will be
# shown here. spacing is not important

C          => E      # C implies E
A + B + C  => D      # A and B and C implies D
A | B      => C      # A or B implies C
A + !B     => F      # A and not B implies F
C | !G     => H      # C or not G implies H
V ^ W      => X      # V xor W implies X
A + B      => Y + Z  # A and B implies Y and Z
C | D      => X | V  # C or D implies X or V
E + F      => !V     # E and F implies not V
A + B      <=> C     # A and B if and only if C

=ABG                 # Initial facts : A, B and G are true. All others are false.
                     # If no facts are initially true, then a simple "=" followed
                     # by a newline is used

?GVX                 # Queries : What are G, V and X ?
```



## Rules and features

The rules can use the following logic operations:

- AND conditions [ + ]:  
`A + B => C`  
(if A and B are true, then C is true)

- OR conditions [ | ]:  
`A | B => C`  
(if A or B are true, then C is true)  

- XOR conditions [ ^ ]:  
`A ^ B => C`  
(if either A or B are true, but the other one is false, then C is true)  

- Negation [ ! ]:  
`!A + B => C`  
(if A isn't true, and B is true, then C is true)  

- AND conclusions [ + ]:  
`A => B + C`  
(if A is true, then B and C are true)  


It implements the following features:  

- Parenthesis in expressions:  
`(A + B) | C => D`  
(if A and B are true, or C is true, then D is true)  

 - Multiple rules with same conclusion:  
 `A + B => C`  
 `D | E => C`  
 (if A and B are true, then C is true. Also, if D or E are true, C is true)  

 - Bidirectional rules [ <=> ]:  
`A <=> B`  
(if A is true, then B is true. Also, if B is true, then A is true)  
