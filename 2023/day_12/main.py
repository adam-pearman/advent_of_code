with open('day_12/input.txt') as f:
    lines = [line.strip() for line in f.readlines()]

visited = {}

def check_springs(springs, sizes):
    variations = 0

    key = springs + ' ' + ','.join(sizes)

    if key in visited:
        return visited[key]

    if len(springs) == 0:
        return 1 if len(sizes) == 0 else 0

    if springs[0] == '.':
        springs = springs.replace('.', '', 1)
        variations += check_springs(springs, sizes)
    elif springs[0] == '?':
        springs = springs.replace('?', '.', 1)
        variations += check_springs(springs, sizes)
        springs = springs.replace('.', '#', 1)
        variations += check_springs(springs, sizes)
    elif springs[0] == '#' and len(sizes) > 0:
        size = int(sizes[0])
        if len(springs) >= size and springs[:size].count('.') == 0:
            springs = springs[size:]
            if len(springs) > 0 and springs[0] != '#':
                springs = springs.replace(springs[0], '.', 1)
            elif len(springs) > 0:
                variations += 0
                return variations
            sizes = sizes[:]
            sizes.pop(0)
            key = springs + ' ' + ','.join(sizes)
            variations += check_springs(springs, sizes)
            visited[key] = variations

    return variations

sum = 0
for line in lines:
    springs = line.split(' ')[0]
    sizes = line.split(' ')[1].split(',')

    sum += check_springs((springs + '?') * 4 + springs, sizes * 5)
print(sum)
