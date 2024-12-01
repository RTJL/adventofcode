```
467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..
```

do inplace update upon reaching a symbol, update the number to a `.`.

prevent double counting.

```
467..114..
...*......
..35..633.
```

read line-by-line, char-by-char.

upon hitting a symbol, find the possible numbers (top-left, top, top-right, left, right, bottom-left, bottom, bottom-right).

if it is a number, then replace the number to a `.`.

will become

1. the top-left number(467), after reading
```
.....114..
...*......
..35..633.
```

2. the bottom number (35), after reading
```
.....114..
...*......
......633.
```
