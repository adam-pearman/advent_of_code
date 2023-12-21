with open('day_21/input.txt') as f:
    lines = [line.strip('\n') for line in f.readlines()]

grid = {}
start = ()

for y, row in  enumerate(lines):
    for x, col in enumerate(row):
        if col == "S":
            start = x + 1j * y
        grid[x + 1j * y] = col

max_x, max_y = int(max(point.real for point in grid)), int(max(point.imag for point in grid))

def infinite_grid(grid, point, max_x, max_y):
    x = point.real % (max_x + 1)
    y = point.imag % (max_y + 1)
    return grid[x + y * 1j]

def get_history(grid, start, max_x, max_y, steps):
    odds, evens, queue = set(), {start}, {start}
    closest_points = lambda point: [point - 1j, point - 1, point + 1, point + 1j]
    odd_history, even_history = [0], [1]
    for i in range(1, steps + 1):
        new_points = set()
        for point in queue:
            for closest in [closest for closest in closest_points(point) if closest not in evens and closest not in odds]:
                if infinite_grid(grid, closest, max_x, max_y) in ".S":
                    new_points.add(closest)
        if i % 2:
            odds |= new_points
        else:
            evens |= new_points
        odd_history.append(len(odds))
        even_history.append(len(evens))
        queue = new_points
    return odd_history, even_history

def part_1(grid, start, max_x, max_y):
    _, even_history = get_history(grid, start, max_x, max_y, 64)
    return even_history[-1]

def part_2(grid, start, max_x, max_y):
    history, _ = get_history(grid, start, max_x, max_y, 3 * 262 + 65)
    steps = 101150
    a = history[2 * 262 + 65]
    b = history[2 * 262 + 65] - history[262 + 65]
    c = history[3 * 262 + 65] - 2 * history[2 * 262 + 65] + history[262 + 65]
    return a + b * (steps - 2) + c * ((steps-2) * (steps - 1) // 2)

print('Part 1:', part_1(grid, start, max_x, max_y))
print('Part 2:', part_2(grid, start, max_x, max_y))
