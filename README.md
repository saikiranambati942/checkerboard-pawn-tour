# Checkerboard-problem

### Description:
This program that finds at least one path for a pawn to visit all squares on an international checkerboard using the following pawn movement rules:

1) Jumps 3 squares horizontally and vertically.
2) Jumps 2 squares diagonally.
3) Jumps to squares already visited or outside of the checkerboard are NOT allowed

### Solutions:
For the above mentioned problem I have implemented two solution using two different algorithms.
### Solution1: Warnsdorff's Algorithm
#### Warnsdorff’s Rule:
We can start from any initial position of the pawn on the board.
We always move to an adjacent, unvisited square with minimal degree (minimum number of unvisited adjacent).

````
For example:
1. Set P to be a random initial position on the board
2. Mark the board at P with the move number “1”
3. Do the following for each move number from 2 to the number of squares on the board.
4. Let S be the set of positions accessible from P.
5. Set P to be the position in S with minimum accessibility
6. Mark the board at P with the current move number.
7. Return the marked board — each square will be marked with the move number on which it is visited.
````


### Solution2: BackTracing Algorithm
Backtracking works in an incremental way to attack problems. Typically, we start from an empty solution vector and 
one by one add items (Meaning of item varies from problem to problem. In context of Knight’s tour problem, an item 
is a Knight’s move). When we add an item, we check if adding the current item violates the problem constraint, if 
it does then we remove the item and try other alternatives. If none of the alternatives work out then we go to 
previous stage and remove the item added in the previous stage. If we reach the initial stage back then we say that 
no solution exists. If adding an item doesn’t violate constraints then we recursively add items one by one. If the 
solution vector becomes complete then we print the solution.



To run the program use the below commands:

````
go run Solution1/WarnsdorffsAlgorithmicSolution.go
````

Sample result is: 
````
16      38      53      63      37      50      62      43      
27      3       56      26      2       57      25      1       
52      64      17      51      54      42      36      49      
15      39      28      58      45      33      61      44      
18      4       55      41      11      48      24      8       
29      59      20      34      60      21      35      32      
14      40      10      13      46      9       12      47      
19      5       30      22      6       31      23      7  
````

```
go run Solution2/BackTracingAlgorithmicSolution.go
```

Sample result is: 
````
0       33      59      62      51      41      63      50      
57      45      37      56      44      38      55      43      
35      61      48      34      60      49      52      40      
1       32      58      6       54      42      5       28      
22      46      36      21      47      39      20      25      
14      7       10      13      18      27      53      17      
2       31      23      3       30      24      4       29      
9       12      15      8       11      16      19      26    
````