//Exercise 2.3: Rewrite PopCount to use a loop instead of a single expression.
//Compare the performance of the two versions. (Section 11.4 shows how to compare
//the performance of different implementations systematically.)
// byte is an alias for uint8

package main
import "fmt"
// pc[i] is the population count of i.
var pc [256]byte

func init() {
    for i := range pc {
        pc[i] = pc[i/2] + byte(i&1)
    }
}

func PopCount(x uint64) int {
    return int(
        pc[byte(x>>(0*8))] +
                pc[byte(x>>(1*8))] +
                pc[byte(x>>(2*8))] +
                pc[byte(x>>(3*8))] +
                pc[byte(x>>(4*8))] +
                pc[byte(x>>(5*8))] +
                pc[byte(x>>(6*8))] +
                pc[byte(x>>(7*8))])
}

func PopCount2(x uint64) int {
    var result uint64
    for i:=0; i < 8; i++{
        result += uint64(pc[byte(x>>(i*8))])
    }
    return int(result)
}
func main() {
    fmt.Println(PopCount(255))
    fmt.Println(PopCount(255))
}
