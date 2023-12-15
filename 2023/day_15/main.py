with open('day_15/input.txt') as f:
    steps = [step for step in f.read().strip('\n').split(',')]

def convert_step(step):
    total = 0
    for char in step:
        total += ord(char)
        total *= 17
        total %= 256
    return total

sum = 0
for step in steps:
    sum += convert_step(step)
print('Part 1:', sum)

boxes = {}
def add_to_box(step):
    label, focal_length = step.split('=')
    box = str(convert_step(label))
    if box not in boxes:
        boxes[box] = {label: int(focal_length)}
        return
    boxes[box][label] = int(focal_length)

def remove_from_box(step):
    label, _ = step.split('-')
    box = str(convert_step(label))
    if box not in boxes or label not in boxes[box]:
        return
    del boxes[box][label]

for step in steps:
    if step.find('=') != -1:
        add_to_box(step)
        continue
    remove_from_box(step)

sum = 0
for key, lenses in boxes.items():
    for index, focal_length in enumerate(lenses.values()):
        box = int(key) + 1
        slot = index + 1
        focusing_power = box * slot * focal_length
        sum += focusing_power
print('Part 2:', sum)
