from collections import defaultdict

with open('day_22/input.txt') as f:
    lines = [line.strip('\n') for line in f.readlines()]

bricks = []
for line in lines:
    start,end = line.split("~")
    start = list(map(int, start.split(",")))
    end = list(map(int, end.split(",")))
    bricks.append((start,end))

bricks.sort(key = lambda x: x[0][2])

top = defaultdict(lambda: (0, -1))
unsafe = set()
graph = [[] for _ in range(len(bricks))]
for i, end in enumerate(bricks):
    offset = -1
    supports = set()
    for j in range(end[0][0], end[1][0] + 1):
        for k in range(end[0][1], end[1][1] + 1):
            if top[j, k][0] + 1 > offset:
                offset = top[j, k][0] + 1
                supports = {top[j, k][1]}
            elif top[j, k][0] + 1 == offset:
                supports.add(top[j,k][1])

    for j in supports:
        if j != -1:
            graph[j].append(i)

    if len(supports) == 1:
        unsafe.add(supports.pop())

    fallen = end[0][2] - offset
    if fallen > 0:
        end[0][2] -= fallen
        end[1][2] -= fallen

    for j in range(end[0][0], end[1][0] + 1):
        for k in range(end[0][1], end[1][1] + 1):
            top[j,k] = (end[1][2], i)

print(len(bricks) - (len(unsafe) - 1))

def count(index, graph):
    lookup = [0 for _ in range(len(bricks))]
    for i in range(len(bricks)):
        for j in graph[i]:
            lookup[j] += 1
    queue = [index]
    count = -1
    while len(queue) > 0:
        count += 1
        x = queue.pop()
        for j in graph[x]:
            lookup[j] -= 1
            if lookup[j] == 0:
                queue.append(j)

    return count

print(sum(count(i, graph) for i in range(len(bricks))))
