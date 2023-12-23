from collections import deque

with open('day_23/input.txt') as f:
    lines = [line.strip('\n') for line in f.readlines()]

grid = [[col for col in row] for row in lines]

row_count = len(grid)
col_count = len(grid[0])

def get_directions(row, col):
    return [
        ['^', row - 1, col],
        ['v', row + 1, col],
        ['<', row, col - 1],
        ['>', row, col + 1],
    ]

def run(ignore_slopes = False):
    visited = set()
    for row in range(row_count):
        for col in range(col_count):
            i = 0
            for direction, new_row, new_col in get_directions(row, col):
                if (0 <= new_row < row_count and 0 <= new_col < col_count and grid[new_row][new_col] != '#'):
                    i += 1
            if i > 2 and grid[row][col] != '#':
                visited.add((row, col))

    for col in range(col_count):
        if grid[0][col] == '.':
            visited.add((0, col))
            start = (0, col)
        if grid[row_count-1][col] == '.':
            visited.add((row_count - 1, col))

    hikes = {}
    for (row, col) in visited:
        hikes[(row, col)] = []
        queue = deque([(row, col, 0)])
        seen = set()
        while queue:
            queued_row, queued_col, i = queue.popleft()
            if (queued_row, queued_col) in seen:
                continue
            seen.add((queued_row, queued_col))
            if (queued_row, queued_col) in visited and (queued_row, queued_col) != (row, col):
                hikes[(row, col)].append(((queued_row, queued_col), i))
                continue
            for direction, new_row, new_col in get_directions(queued_row, queued_col):
                if (0 <= new_row < row_count and 0 <= new_col < col_count and grid[new_row][new_col] != '#'):
                    if not ignore_slopes and grid[queued_row][queued_col] in ['>', 'v'] and grid[queued_row][queued_col] != direction:
                        continue
                    queue.append((new_row, new_col, i + 1))

    count = 0
    max_distance = 0
    has_seen = [[False for _ in range(col_count)] for _ in range(row_count)]
    def search(coordinates, i):
        nonlocal count
        nonlocal max_distance
        count += 1
        row, col = coordinates
        if has_seen[row][col]:
            return
        has_seen[row][col] = True
        if row == row_count - 1:
            max_distance = max(max_distance, i)
        for (item, distance) in hikes[coordinates]:
            search(item, distance + i)
        has_seen[row][col] = False

    search(start, 0)
    return max_distance

print('Part 1:', run())
print('Part 2:', run(True))
