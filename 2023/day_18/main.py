with open('day_18/input.txt') as f:
    moves = [line.strip().split(' ') for line in f.readlines()]

def get_area(vertices):
    n = len(vertices)
    sum_1 = 0
    sum_2 = 0
    edge_points = 0
    for i in range(n - 1):
        sum_1 += vertices[i][0] * vertices[i + 1][1]
        sum_2 += vertices[i][1] * vertices[i + 1][0]
        edge_points += abs(vertices[i][0] - vertices[i + 1][0]) + abs(vertices[i][1] - vertices[i + 1][1])
    sum_1 += vertices[n - 1][0] * vertices[0][1]
    sum_2 += vertices[0][0] * vertices[n - 1][1]
    area = abs(sum_1 - sum_2) / 2
    inside_points = (((2 * area) - edge_points + 2) / 2)
    return int(inside_points + edge_points)


vertices = [(0, 0)]
for i, move in enumerate(moves):
    last_vertex = vertices[i]
    if move[0] == 'U':
        vertices.append((last_vertex[0], last_vertex[1] + int(move[1])))
    elif move[0] == 'D':
        vertices.append((last_vertex[0], last_vertex[1] - int(move[1])))
    elif move[0] == 'L':
        vertices.append((last_vertex[0] - int(move[1]), last_vertex[1]))
    elif move[0] == 'R':
        vertices.append((last_vertex[0] + int(move[1]), last_vertex[1]))
print('Part 1:', get_area(vertices))

directions = ('R', 'D', 'L', 'U')
vertices = [(0, 0)]
for i, move in enumerate(moves):
    direction = directions[int(move[2][-2])]
    distance = int(move[2][1:-2].replace('#', ''), 16)
    last_vertex = vertices[i]
    if direction == 'U':
        vertices.append((last_vertex[0], last_vertex[1] + distance))
    elif direction == 'D':
        vertices.append((last_vertex[0], last_vertex[1] - distance))
    elif direction == 'L':
        vertices.append((last_vertex[0] - distance, last_vertex[1]))
    elif direction == 'R':
        vertices.append((last_vertex[0] + distance, last_vertex[1]))
print('Part 2:', get_area(vertices))
