with open('day_14/input.txt') as f:
    lines = [line.strip('\n') for line in f.readlines()]


def check_for_move(line, current_index, stored_index):
    if line[current_index] == '.' and stored_index is None:
        stored_index = current_index
    if line[current_index] == '#':
        stored_index = None
    if line[current_index] == 'O' and stored_index is not None:
        line[current_index] = '.'
        line[stored_index] = 'O'
        current_index = stored_index
        stored_index = None

    return current_index, stored_index

def move_rocks(line, direction = 'N'):
    stored_index = None
    if direction in 'NW':
        current_index = 0
        while current_index < len(line):
            current_index, stored_index = check_for_move(line, current_index, stored_index)
            current_index += 1
    if direction in 'SE':
        current_index = len(line) - 1
        while current_index >= 0:
            current_index, stored_index = check_for_move(line, current_index, stored_index)
            current_index -= 1
    return line

def calculate_load(lines):
    load = 0
    for index, line in enumerate(lines):
        load += line.count('O') * (len(lines) - index)
    return load


columns = [list(line) for line in zip(*lines)]

for index, column in enumerate(columns):
    columns[index] = move_rocks(column)
columns = [list(column) for column in zip (*columns)]
print('Part One:', calculate_load(columns))


lookup = {}
iterations = 1000000000
restart = None
load = 0
for i in range(iterations):
    load = 0
    for direction in 'NWSE':
        lines = [list(line) for line in zip(*lines)]
        for index, line in enumerate(lines):
            lines[index] = move_rocks(line, direction)
    load = calculate_load(lines)
    if str(lines) in lookup:
        restart = str(lines)
        break
    lookup[str(lines)] = load

loop_length = len(lookup) - list(lookup.keys()).index(restart)
loop_offset = len(lookup) - loop_length
index = ((iterations - loop_offset) % loop_length) + loop_offset - 1
print('Part Two:', list(lookup.values())[index])

