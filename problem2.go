package main

import "fmt"

// The values should be
// x, y, n >= 0
func solution(x, y, n uint32) (s uint32) {
    var i uint32;
    // Iterate from 2 to n-1
    for i = 2; i < n; i++ {
        // Sum if i is divisibly by x
        if (i % x) == 0 {
            s += i;
        } else if (i % y) == 0 { // Sum if i is divisibly by y
            s += i;
        }
    }
    // Return the sum
    return;
}

func reverse(n uint32) (r uint32) {
    // Iteratively create the
    // reversed form
    for ; n != 0; {
        // Get the last digit
        ld := n % 10;
        // Push the last digit
        r = ld + (r * 10);
        // Chop off the last digit
        n /= 10;
    }
    // Return the reversed form
    return;
}

func main() {
    var x, y, n uint32;
    // Test code
    fmt.Print("Enter x: ");
    fmt.Scanf("%d", &x);
    fmt.Print("Enter y: ");
    fmt.Scanf("%d", &y);
    fmt.Print("Enter n: ");
    fmt.Scanf("%d", &n);
    r := solution(3, 5, 10);
    fmt.Println("The sum of multiples =", r);
    fmt.Println("Reversed form =", reverse(r));
}