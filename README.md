# gudoku
Sudoku builder/solver

# How to use :

```

    rand := rand.New(rand.NewSource(time.Now().UnixMicro()))
	puzzle, solution := sdk.Build(rand)

	fmt.Println("Puzzle (replace all zeroes !) :")
	puzzle.Dump()
	fmt.Println("Solution :")
	fmt.Println(solution)
    
```
