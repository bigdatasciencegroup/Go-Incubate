---
title: "Merging two sorted lists in Scala"
tags: 
  - scala 
categories:
  - scala   
last_modified_at: 2018-04-22
---

This post elucidates the power of programming in Scala. We describe how to merge two sorted lists into a third sorted list in Scala. Although there are built-in methods to perform this form of sorting, here, we shall use pattern matching and recursion to solve the problem. 

```scala
def merge(i:List[Int], j:List[Int]):List[Int]={
  (i,j) match {
    case (Nil, Nil) => Nil
    case (x::xs, Nil) => i
    case (Nil, y::ys) => j
    case (x::xs, y::ys) => {
      if (x <= y)
        x :: merge(i.tail,j)
      else
        y :: merge(i,j.tail)
    }
  }
}

val i = List(1,4,7,9)
val j = List(3,4,5,11,12,14)
merge(i,j)
```

Output upon running the above code in Scala:
```scala
i: List[Int] = List(1, 4, 7, 9)
j: List[Int] = List(3, 4, 5, 11, 12, 14)
res0: List[Int] = List(1, 3, 4, 4, 5, 7, 9, 11, 12, 14)
```

Programming in Scala is simple and elegant, isn't it?
