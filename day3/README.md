# RULES

The instructions are to create the largest amount of joltage out of rows of batteries. 

Joltage is calculated by taking the largest numbers in a row of batteries, then adding the sum together.

Our example input is:

```
987654321111111
811111111111119
234234234234278
818181911112111
```

Part one you take the largest 2 numbers and sum the total for each row.

So for out example input you would do:

- In `987654321111111`, you can make the largest joltage possible, 98, by turning on the first two batteries.
- In `811111111111119`, you can make the largest joltage possible by turning on the batteries labeled 8 and 9, producing 89 jolts.
- In `234234234234278`, you can make 78 by turning on the last two batteries (marked 7 and 8).
- In `818181911112111`, the largest joltage you can produce is 92.

Which would result in `98 + 89 + 78 + 92 = 357`.

For part two you take the largest 12 numbers and sum the total for each row.

So for out example input you would do:

- In 987654321111111, the largest joltage can be found by turning on everything except some 1s at the end to produce 987654321111.
- In the digit sequence 811111111111119, the largest joltage can be found by turning on everything except some 1s, producing 811111111119.
- In 234234234234278, the largest joltage can be found by turning on everything except a 2 battery, a 3 battery, and another 2 battery near the start to produce 434234234278.
- In 818181911112111, the joltage 888911112111 is produced by turning on everything except some 1s near the front.

Which would result in: `987654321111 + 811111111119 + 434234234278 + 888911112111 = 3121910778619`.

# How to run

`go run main.go`
