with open('day_16/input.txt') as f:
    lines = [line.strip('\n') for line in f.readlines()]

def move(position, direction):
    if direction == 'U':
        return [position[0] - 1, position[1]]
    if direction == 'D':
        return [position[0] + 1, position[1]]
    if direction == 'L':
        return [position[0], position[1] - 1]
    if direction == 'R':
        return [position[0], position[1] + 1]

def hit_mirror(mirror, direction):
    mirrors = {
        '/': {
            'U': 'R',
            'D': 'L',
            'L': 'D',
            'R': 'U',
        },
        '\\': {
            'U': 'L',
            'D': 'R',
            'L': 'U',
            'R': 'D',
        }
    }
    return mirrors[mirror][direction]

def run(starting_position, starting_direction):
    mirrors_hit = { 'U': [], 'D': [], 'L': [], 'R': []}
    splitters_hit = { 'U': [], 'D': [], 'L': [], 'R': []}
    energised_tiles = []

    def light_beam(position, direction):
        while 0 <= position[0] < len(lines) and 0 <= position[1] < len(lines[0]):
            tile = lines[position[0]][position[1]]
            if position not in energised_tiles:
                energised_tiles.append(position)
            if tile in '/\\':
                if position in mirrors_hit[direction]:
                    break
                mirrors_hit[direction].append(position)
                direction = hit_mirror(tile, direction)
            elif tile == '-' and direction in 'UD':
                if position in splitters_hit[direction]:
                    break
                splitters_hit[direction].append(position)
                direction = 'L'
                light_beam(position, 'R')
            elif tile == '|' and direction in 'LR':
                if position in splitters_hit[direction]:
                    break
                splitters_hit[direction].append(position)
                direction = 'U'
                light_beam(position, 'D')
            position = move(position, direction)

    light_beam(starting_position, starting_direction)

    return len(energised_tiles)

print('Part 1:', run([0, 0], 'R'))

results = {}

for index in range(len(lines)):
    line_length = len(lines[0])
    results[f'{index}0R'] = run([index, 0], 'R')
    results[f'{len(lines) - index}{line_length}L'] = run([len(lines) - index, line_length], 'L')

for index in range(len(lines[0])):
    lines_length = len(lines)
    results[f'0{index}D'] = run([0, index], 'D')
    results[f'{lines_length}{len(lines[0]) - index}L'] = run([lines_length, len(lines[0]) - index], 'U')

print('Part 2:', max(results.values()))
