# randmatrix
The task is to fill a matrix (m × n) with random, int, unique numbers from a certain range (for example for the range 1-3 possible numbers are: 1, 2)  

## Examples

```go
func main() {
    m := 5
    n := 5
    min := 0
    max := 100
    res, err := matrix.CreateArray(m, n, min, max)
    if err = nil{
        fmt.Println(err)
    }else{
        fmt.Println(res)
    }
}
```
Notes:     
This solution works only with small dimensions as unique numbers are selected on the fly.
