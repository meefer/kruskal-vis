## kruskal-vis
![ci](https://github.com/meefer/kruskal-vis/workflows/ci/badge.svg)

kruskal-vis is a command line tool to generate a gif demo for Kruskal's algorithm based on Euclidean distance.

Demo can be found [here](https://meefer.github.io/kruskal-vis) ([Wasm](https://en.wikipedia.org/wiki/WebAssembly) runtime).

To install kruskal-vis locally run
```bash
go get -u github.com/meefer/kruskal-vis
```

`kruskal-vis -h`
```
kruskal-vis generates a random graph with a fixed number of nodes
and runs it through Kruskal's minimum-spanning-tree algorithm
producing three files in a working directory:
  - graph.png -- original graph
  - kruskal_anim.gif -- visualization of the Kruskal's algorithm
  - kruskal.png -- minimum spanning tree of the original graph
Usage:
  -N int
        number of graph nodes (default 10)
  -t    if set, a textual graph presentation will be written to the standard output
```
