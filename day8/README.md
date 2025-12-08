# RULES

We are given a list of junctions that require to be connected. 

An example input is: 

```
162,817,812
57,618,57
906,360,560
592,479,940
352,342,300
466,668,158
542,29,236
431,825,988
739,650,466
52,470,668
216,146,977
819,987,18
117,168,530
805,96,715
346,949,466
970,615,88
941,993,340
862,61,35
984,92,344
425,690,689
```

Each line is a coordinate, with the x, y, and z coordinates.

For part one we need to connect the 1000 closest connectors, and create circuits with them. Once we have completed that we return the multiplication between the 3 largest circuits (number of junctions joined).

So for the example input you have: 

`162,817,812` and `425,690,689` for the closest two, then if we continue this 10 times we have a 5 junction, 4 junction, and two 2 junctions circuits.

This then returns `5*4*2=40`.

For part two we are looking for the last two junctions that require to be connected for a full connections with all junctions. 

Once we have found this we multiply the two x coordinates 
For the example it would be `216,146,977` and `117,168,530`, which would result in `25272`.

# How to run

`go run .`
