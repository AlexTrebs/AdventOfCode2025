# RULES

The task is to complete a cephlapods math homework. The issue is, cephlapods have a different way of writing equations:

```
123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +  
```

Above is the example input. Here we see that the numbers are at the top, with the operation at the bottom.

For task one we must calculate the total of all the output from the input, assuming that we are reading the numbers from top to bottom.

So for the example input we have:

- `123 * 45 * 6 = 33210`
- `328 + 64 + 98 = 490`
- `51 * 387 * 215 = 4243455`
- `64 + 23 + 314 = 401`

So the output would be `33210 + 490 + 4243455 + 401 = 4277556`.

Part two now assumes we are reading top right to bottom left, this mean for the example we have:


- The rightmost problem is `4 + 431 + 623 = 1058`
- The second problem from the right is `175 * 581 * 32 = 3253600`
- The third problem from the right is `8 + 248 + 369 = 625`
- Finally, the leftmost problem is `356 * 24 * 1 = 8544`

Which results in: `1058 + 3253600 + 625 + 8544 = 3263827`

# How to run

`go run .`
