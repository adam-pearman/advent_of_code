import networkx as nx

with open('day_25/input.txt') as f:
    lines = [line.strip('\n') for line in f.readlines()]

graph = nx.Graph()
for line in lines:
    node, *edges = line.split()
    graph.add_edges_from((node[:-1], edge) for edge in edges)

graph.remove_edges_from(nx.minimum_edge_cut(graph))
a, b = nx.connected_components(graph)

print('Part 1:', len(a) * len(b))
