package main

import (
   "fmt"
   "os"
   "strconv"
   "time"
   "math"
   "flag"
)

func main() {
   debug := flag.Bool("debug", false, "If you want to see what's happening under the hood...")
   printList := flag.Bool("print", false, "Shows the prime numbers' list")

   primes := make([]uint64, 1)
   primes[0] = 2
   
   var limit, candidate uint64
   var err error

   if len(os.Args) > 1 {
      limit, err = strconv.ParseUint(os.Args[1], 10, 64)
      if err != nil {
         fmt.Println("Error in conversion, using default value (100)")
         fmt.Println(err)
         limit = 100
      }
   } else {
      fmt.Println("No limit specified, using default value (100)")
      limit = 100
   }
   
   flag.Parse()
   fmt.Println("Debug", *debug)
   fmt.Println("Print list", *printList)

   
   
   t0 := time.Now()
   
   for candidate = 3; candidate < limit; candidate++ {
      if candidate % 10000 == 0 {
         fmt.Printf("Evaluating %v - found so far %v - elapsed %v\n", candidate, len(primes), time.Now().Sub(t0))
      }
      
      // prime factors are bound to sqrt of candidate
      candidateSqrt := uint64(math.Floor(math.Sqrt(float64(candidate))))
      isPrime := true
      for j := 0; j < len(primes) && isPrime && primes[j] <= candidateSqrt; j++ {
         //fmt.Println(candidate, j, primes[j], candidate % primes[j], isPrime)
         isPrime = (candidate % primes[j]) != 0
      }
      if isPrime {
         primes = append (primes, candidate)
      }
   }

   t1 := time.Now()
   fmt.Println(primes)

   fmt.Printf("There are %v prime numbers below %v\n", len(primes), limit)
   fmt.Printf("It took %v to find them all.\n", t1.Sub(t0))

}
