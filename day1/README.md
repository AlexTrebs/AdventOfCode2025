# RULES
The task is a to follow some instructions to a dial lock.

The instructions read as:
- `L68`
- `L30`
- `R48`
- `L5`
- `R60`
- `L55`
- `L1`
- `L99`
- `R14`
- `L82`

The exercise starts at 50, and the dial goes from 0 - 99 so `R1` from 99 would be 0.

Following these rotations would cause the dial to move as follows:

- The dial starts by pointing at 50.
- The dial is rotated L68 to point at 82.
- The dial is rotated L30 to point at 52.
- The dial is rotated R48 to point at 0.
- The dial is rotated L5 to point at 95.
- The dial is rotated R60 to point at 55.
- The dial is rotated L55 to point at 0.
- The dial is rotated L1 to point at 99.
- The dial is rotated L99 to point at 0.
- The dial is rotated R14 to point at 14.
- The dial is rotated L82 to point at 32.

For `part one` the instruction is to output the number of times the dial stops and points to 0 in this process.
So for the example above it would be 3.

For `part two` the instruction is to output the number of times the dial passes 0, so if you are at 0 and you turn right 500 times it would cross 0 5 times.
So for the example above it would be 6.

# How to run

`go run main.go`
