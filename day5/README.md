# RULES

The task is to find the fresh food.

The example input is:

```
3-5
10-14
16-20
12-18

1
5
8
11
17
32
```

The ranges at the top are ranges of fresh food. Then below we have the ids of the food.

For task 1 we are told to find how much of the food is fresh, and there for lies within the ranges. This follows for the example:


- Ingredient ID 1 is spoiled because it does not fall into any range.
- Ingredient ID 5 is fresh because it falls into range 3-5.
- Ingredient ID 8 is spoiled.
- Ingredient ID 11 is fresh because it falls into range 10-14.
- Ingredient ID 17 is fresh because it falls into range 16-20 as well as range 12-18.
- Ingredient ID 32 is spoiled.

So for part one the output would be 3.

For part two we only want to find out how many foodIds are fresh, so we are only looking at the ranges to find how many ids this spans.

This means for the example we have `3, 4, 5, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, and 20.`, so there are 14 fresh food ids.

# How to run

`go run .`
