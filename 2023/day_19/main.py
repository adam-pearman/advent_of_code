from copy import deepcopy

with open('day_19/input.txt') as f:
    instructions, parts = [block.splitlines() for block in f.read().split('\n\n')]

instructions = {instruction.split('{')[0]: instruction.split('{')[1][:-1].split(',') for instruction in instructions}
parts = [part.strip('{}').split(',') for part in parts]

new_parts = []
for i, part in enumerate(parts):
    new_parts.append({})
    for attribute in part:
        new_parts[i][attribute[:1]] = int(attribute[2:])
parts = new_parts

def follow_instruction_for_part(part, instruction):
    for step in instruction:
        if ':' not in step:
            if step == 'A':
                return sum(list(part.values()))
            if step == 'R':
                return 0
            return follow_instruction_for_part(part, instructions[step])
        value, destination = step[2:].split(':')
        if (step[1] == '>' and part[step[0]] > int(value)) or (step[1] == '<' and part[step[0]] < int(value)):
            if destination == 'A':
                return sum(list(part.values()))
            if destination == 'R':
                return 0
            return follow_instruction_for_part(part, instructions[destination])

answer = 0
for part in parts:
    answer += follow_instruction_for_part(part, instructions['in'])
print('Part 1:', answer)

def follow_instruction_for_possibilities(possibilities, instruction, ranges):
    output = 0
    for step in instruction:
        if ':' not in step:
            if step == 'A':
                output += possibilities
                break
            if step == 'R':
                break
            output += follow_instruction_for_possibilities(possibilities, instructions[step], ranges)
            break
        value, destination = step[2:].split(':')
        split_ranges = deepcopy(ranges)
        if step[1] == '<':
            if int(value) <= ranges[step[0]][1]:
                value = int(value) - 1
                percentage = (value - ranges[step[0]][0]) / (ranges[step[0]][1] - ranges[step[0]][0])
                split_ranges[step[0]][1] = value
                ranges[step[0]][0] = value
                split = possibilities * percentage
                possibilities -= split
            else:
                split = possibilities
        elif step[1] == '>':
            if int(value) > ranges[step[0]][0]:
                value = int(value)
                percentage = 1 - ((value - ranges[step[0]][0]) / (ranges[step[0]][1] - ranges[step[0]][0]))
                split_ranges[step[0]][0] = value
                ranges[step[0]][1] = value
                split = possibilities * percentage
                possibilities -= split
            else:
                split = possibilities
        if destination == 'A':
            output += split
        elif destination == 'R':
            continue
        else:
            output += follow_instruction_for_possibilities(split, instructions[destination], split_ranges)
    return output

ranges = {'x': [0, 4000], 'm': [0, 4000], 'a': [0, 4000], 's': [0, 4000]}
print('Part2:', int(follow_instruction_for_possibilities(4000**4, instructions['in'], ranges)))
