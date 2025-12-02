# RULES

The task is to check product ids. The input looks like this:

`11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124`

The task is to check for any invalid inputs in these. An invalid input can have 0 at the start, or be made of some sequence of repeating numbers e.g. 66, 6262, or 123123.

So for the example input, the output would be: 

- `11-22` has two invalid IDs, 11 and 22.
- `95-115` has one invalid ID, 99.
- `998-1012` has one invalid ID, 1010.
- `1188511880-1188511890` has one invalid ID, 1188511885.
- `222220-222224` has one invalid ID, 222222.
- `1698522-1698528` contains no invalid IDs.
- `446443-446449` has one invalid ID, 446446.
- `38593856-38593862` has one invalid ID, 38593859.
- The rest of the ranges contain no invalid IDs.

For part one the instruction is to add up all the invalid id's where there can only be one repetiting numbers, so for the example input it would result in: `1227775554`.

For part two the instruction is to add up all the invalid id's where there can only be more than repetiting numbers, so 121212 would be one, and 13131313, so for the example input it would be:

- `11-22` still has two invalid IDs, 11 and 22.
- `95-115` now has two invalid IDs, 99 and 111.
- `998-1012` now has two invalid IDs, 999 and 1010.
- `1188511880-1188511890` still has one invalid ID, 1188511885.
- `222220-222224` still has one invalid ID, 222222.
- `1698522-1698528` still contains no invalid IDs.
- `446443-446449` still has one invalid ID, 446446.
- `38593856-38593862` still has one invalid ID, 38593859.
- `565653-565659` now has one invalid ID, 565656.
- `824824821-824824827` now has one invalid ID, 824824824.
- `2121212118-2121212124` now has one invalid ID, 2121212121.

This would result in the output being: `4174379265`

# How to run

`go run main.go`
