with open('day_13/input.txt') as f:
    grids = [grid.splitlines() for grid in f.read().split('\n\n')]

def is_reflection(grid, mirror):
    min_length = min(len(grid), len(mirror))
    return grid[:min_length] == mirror[:min_length]

def check_for_reflections(grid, multiplier, bad = 0):
    for index, line in enumerate(grid[1:], 1):
        if grid[index - 1] == line and is_reflection(grid[index - 1::-1], grid[index:]):
            if index * multiplier != bad:
                return index * multiplier
    return 0

def modify_grid(grid):
    for l_index, line in enumerate(grid):
        for c_index, char in enumerate(line):
            clone = grid[:]
            char = '#' if char == '.' else '.'
            clone[l_index] = line[:c_index] + char + line[c_index + 1:]
            yield clone

sum = 0
for grid in grids:
    reflection = check_for_reflections(grid, 100) or check_for_reflections([line for line in zip(*grid)], 1)
    sum += reflection
print('Part 1:', sum)

sum = 0
for grid in grids:
    for modified_grid in modify_grid(grid):
        reflection = check_for_reflections(grid, 100) or check_for_reflections([line for line in zip(*grid)], 1)
        modified_reflection = check_for_reflections(modified_grid, 100, reflection) or check_for_reflections([line for line in zip(*modified_grid)], 1, reflection)
        if modified_reflection:
            sum += modified_reflection
            break
print('Part 2:', sum)
