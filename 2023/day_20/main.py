from collections import deque
from math import lcm

with open('day_20/input.txt') as f:
    config = [line.strip('\n') for line in f.readlines()]

flip_flops = {}
conjunctions = {}
vd_inputs = {}

def set_modules():
    broadcaster = []
    for line in config:
        line = line.split(' -> ')
        module = line[0]
        if module == 'broadcaster':
            broadcaster = line[1].split(', ')
        elif module[0] == '%':
            flip_flops[module[1:]] = {
                'state': 'off',
                'destinations': line[1].split(', ')
            }
        elif module[0] == '&':
            conjunctions[module[1:]] = {
                'input_frequencies': {},
                'destinations': line[1].split(', ')
            }

    for module in flip_flops:
        for destination in flip_flops[module]['destinations']:
            if destination in conjunctions:
                conjunctions[destination]['input_frequencies'][module] = 'low'
            if destination == 'vd':
                vd_inputs[module] = 0

    for module in conjunctions:
        for destination in conjunctions[module]['destinations']:
            if destination in conjunctions:
                conjunctions[destination]['input_frequencies'][module] = 'low'
            if destination == 'vd':
                vd_inputs[module] = 0

    return broadcaster

q = deque()

def update_flip_flop(module, frequency, _):
    if frequency == 'high':
        return
    if flip_flops[module]['state'] == 'off':
        flip_flops[module]['state'] = 'on'
        frequency = 'high'
    else:
        flip_flops[module]['state'] = 'off'

    for destination in flip_flops[module]['destinations']:
        q.append([destination, frequency, module])

def update_conjunction(module, frequency, input):
    conjunctions[module]['input_frequencies'][input] = frequency
    output = 'high' if 'low' in conjunctions[module]['input_frequencies'].values() else 'low'

    for destination in conjunctions[module]['destinations']:
        q.append([destination, output, module])

def run(iterations):
    button_presses = 0
    low_pulses = 0
    high_pulses = 0
    for _ in range(iterations):
        button_presses += 1
        for module in broadcaster:
            q.append([module, 'low', 'broadcaster'])

        while len(q) > 0:
            module, frequency, input = q.popleft()
            if frequency == 'low':
                low_pulses += 1
            else:
                high_pulses += 1
                if module == 'vd':
                    vd_inputs[input] = button_presses
                    if 0 not in vd_inputs.values():
                        return lcm(*vd_inputs.values())
            if module in flip_flops:
                update_flip_flop(module, frequency, input)
            elif module in conjunctions:
                update_conjunction(module, frequency, input)
    return (button_presses + low_pulses) * high_pulses

broadcaster = set_modules()
print('Part 1: ', run(1000))

broadcaster = set_modules()
print('Part 2: ', run(5000))
