import re

with open('day_10/input.txt') as f:
    lines = f.readlines()

lines = [line.strip() for line in lines]

def direction_from(oldPosition, newPosition):
    if oldPosition[0] == newPosition[0]:
        return 'L' if oldPosition[1] < newPosition[1] else 'R'
    return 'U' if oldPosition[0] < newPosition[0] else 'D'

def move(position, direction):
    if direction == 'U':
        return [position[0] - 1, position[1]]
    if direction == 'D':
        return [position[0] + 1, position[1]]
    if direction == 'L':
        return [position[0], position[1] - 1]
    if direction == 'R':
        return [position[0], position[1] + 1]

def find_next_move(pipe, direction_from):
    if pipe == '|':
        return 'U' if direction_from == 'D' else 'D'
    if pipe == '-':
        return 'R' if direction_from == 'L' else 'L'
    if pipe == 'L':
        return 'R' if direction_from == 'U' else 'U'
    if pipe == 'J':
        return 'U' if direction_from == 'L' else 'L'
    if pipe == 'F':
        return 'R' if direction_from == 'D' else 'D'
    if pipe == '7':
        return 'D' if direction_from == 'L' else 'L'

valid_starting_moves = {
    '|': ['U', 'D'],
    '-': ['L', 'R'],
    'L': ['D', 'L'],
    'J': ['D', 'R'],
    'F': ['U', 'L'],
    '7': ['U', 'R'],
    '.': []
}

def get_starting_moves(position):
    l = lines[position[0]][position[1] - 1] if position[1] >= 0 else None
    u = lines[position[0] - 1][position[1]] if position[0] >= 0 else None
    r = lines[position[0]][position[1] + 1] if position[1] < len(lines[position[0]]) else None
    d = lines[position[0] + 1][position[1]] if position[0] < len(lines) else None

    moves = []
    if l and 'L' in valid_starting_moves[l]:
        moves.append('L')
    if u and 'U' in valid_starting_moves[u]:
        moves.append('U')
    if r and 'R' in valid_starting_moves[r]:
        moves.append('R')
    if d and 'D' in valid_starting_moves[d]:
        moves.append('D')

    return moves

old_positions = []
current_positions = []
steps = 0

pipe_positions = []

for index, line in enumerate(lines):
    s_index = line.find('S')
    if s_index != -1:
        current_positions = [[index, s_index], [index, s_index]]
        old_positions = [[index, s_index], [index, s_index]]
        pipe_positions = [[index, s_index]]

starting_moves = get_starting_moves(current_positions[0])

s_pipe = ''
if starting_moves == ['L', 'U']: s_pipe = 'J'
elif starting_moves == ['L', 'R']: s_pipe = '-'
elif starting_moves == ['L', 'D']: s_pipe = '7'
elif starting_moves == ['U', 'R']: s_pipe = 'L'
elif starting_moves == ['U', 'D']: s_pipe = '|'
elif starting_moves == ['R', 'D']: s_pipe = 'F'

lines[current_positions[0][0]] = lines[current_positions[0][0]].replace('S', s_pipe)

for index, starting_move in enumerate(starting_moves):
    current_positions[index] = move(current_positions[index], starting_move)
    pipe_positions.append(current_positions[index])
steps += 1

while current_positions[0] != current_positions[1]:
    for index, position in enumerate(current_positions):
        pipe = lines[position[0]][position[1]]
        direction = direction_from(old_positions[index], position)
        next_move = find_next_move(pipe, direction)
        newPosition = move(position, next_move)
        old_positions[index] = position
        current_positions[index] = newPosition
        pipe_positions.append(current_positions[index])
    steps += 1

print('Steps to farthest point:', steps)

row_counts = []
for row_index, line in enumerate(lines):
    line = [char if [row_index, col_index] in pipe_positions else '.' for col_index, char in enumerate(line)]
    line = ''.join(line)

    line = re.sub(r'L-*7', '|', line)
    line = re.sub(r'L-*J', '||', line)
    line = re.sub(r'F-*7', '||', line)
    line = re.sub(r'F-*J', '|', line)

    cross = 0
    inside = 0

    for char in line:
        if char == '.' and cross % 2:
            inside += 1
        if char in 'F7LJ|':
            cross += 1
    row_counts.append(inside)

print('Enclosed spaces:', sum(row_counts))
