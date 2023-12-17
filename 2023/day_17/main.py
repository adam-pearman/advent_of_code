from heapq import heappop, heappush

with open('day_17/input.txt') as f:
    grid = [list(int(num) for num in line.strip()) for line in f.readlines()]

directions = [(0, 1), (1, 0), (0, -1), (-1, 0)]

def shortest_path(min_travel, max_travel):
    heat_cost, x, y, cantTravel = 0, 0, 0, -1
    queue = [(heat_cost, x, y, cantTravel)]
    visited = set()
    heat_costs = {}
    while len(queue) > 0:
        heat_cost, x, y, cantTravel  = heappop(queue)
        if x == len(grid) - 1 and y == len(grid[0]) - 1:
            return heat_cost
        if (x, y, cantTravel) in visited:
            continue
        visited.add((x, y, cantTravel))
        for direction in range(4):
            heat_cost_increase = 0
            if direction == cantTravel or (direction + 2) % 4 == cantTravel:
                continue
            for distance in range(1, max_travel + 1):
                xx = x + directions[direction][0] * distance
                yy = y + directions[direction][1] * distance
                if xx in range(len(grid)) and yy in range(len(grid[0])):
                    heat_cost_increase += grid[xx][yy]
                    if distance < min_travel:
                        continue
                    new_heat_cost = heat_cost + heat_cost_increase
                    if heat_costs.get((xx, yy, direction), float('inf')) <= new_heat_cost:
                        continue
                    heat_costs[(xx, yy, direction)] = new_heat_cost
                    heappush(queue, (new_heat_cost, xx, yy, direction))

print('Part 1:', shortest_path(1, 3))
print('Part 2:', shortest_path(4, 10))
