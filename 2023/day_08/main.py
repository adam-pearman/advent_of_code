import math

with open('day_08/input.txt') as f:
    lines = f.readlines()

directions = {
    "L": 0,
    "R": 1,
}

map = {}
for line in lines[2:]:
    line = line.strip().split(' = ')
    key = line[0]
    map[key] = line[1].replace('(', '').replace(')', '').replace(' ', '').split(',')

instructions = lines[0].strip().replace(' ', '')

# Part 1
# count = 0
# lost = True
# current_key = 'AAA'
# while lost:
#     for instruction in instructions:
#         current_key = map[current_key][directions[instruction]]
#         count += 1
#         if current_key == 'ZZZ':
#             lost = False
#             break
# print(count)

# Part 2
current_keys = [[key, 0] for key in map if key[2] == 'A']

for index, current_key in enumerate(current_keys):
    lost = True
    while lost:
        for instruction in instructions:
            current_keys[index][0] = map[current_key[0]][directions[instruction]]
            current_keys[index][1] += 1
            if current_keys[index][0][2] == 'Z':
                lost = False
                break

print(math.lcm(*[current_key[1] for current_key in current_keys]))
