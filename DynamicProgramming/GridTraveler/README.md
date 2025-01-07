# The Grid Traveler

Considering a grid of size (n, m).

How many ways are there to reach the bottom right corner, moving only right and down ?

With dynamic programming, this problem can be broken down into 2 base cases.

If n = 1 and m = 1, we don't need to move, we can then return 1
If either n or m = 0, the grid is invalid, and thus, we return 0

With caching, the time complexity goes from O(2^(n+m)) to O(n*m)
The space complexity is, in both case O(n+m)

We can even optimize further by recognizing that f(n,m) = f(m,n)
