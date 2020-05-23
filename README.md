# freqs

Linux CPU frequency monitor

## Example

```bash
$ freqs
..........                              | 1974
..........                              | 2024
...........                             | 2063
...........                             | 2124
............                            | 2191
............                            | 2169
...........                             | 2123
..........                              | 1979
............                            | 2197
...........                             | 2062
...........                             | 2115
...........                             | 2138
............                            | 2191
............                            | 2163
..........                              | 2016
............                            | 2157
```

## Installation

```bash
$ go get github.com/quells/freqs
$ go install github.com/quells/freqs
```

## About

Reads from `/proc/cpuinfo`, parses out the frequency of each logical CPU (processors with simultaneous multithreading aka hyperthreading will appear with multiple entries per physical core), and prints out a bar chart in the range 1000 to 5000 MHz. If your processor runs outside that range, the lower and upper bounds will be adjusted accordingly.

