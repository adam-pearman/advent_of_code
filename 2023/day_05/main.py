with open('day_05/input.txt') as input:
    lines = input.readlines()

def part_one():
    seeds = []
    changed = []

    for index, line in enumerate(lines):
        if index == 0:
            seeds = [int(seed) for seed in line.strip().split(': ')[1].split(' ')]

        if line[0].isdigit():
            line = line.strip().split(' ')
            diff = int(line[2])
            old_min = int(line[1])
            new_min = int(line[0])
            for index, seed in enumerate(seeds):
                if seed in range(old_min, old_min + diff) and index not in changed:
                    changed.append(index)
                    seeds[index] = seed + (new_min - old_min)
        else:
            changed = []

    return min(seeds)

def part_two():
    seed_ranges = []
    changed = []
    for index, line in enumerate(lines):
        if index == 0:
            seeds = [int(seed) for seed in line.strip().split(': ')[1].split(' ')]
            for index, seed in enumerate(seeds):
                if index % 2 == 0:
                    seed_ranges.append([seed])
                else:
                    seed_ranges[index // 2].append(seed_ranges[index // 2][0] + seed - 1)

        if line[0].isdigit():
            line = line.strip().split(' ')
            diff = int(line[2])
            old_min = int(line[1])
            new_min = int(line[0])
            for index, seed_range in enumerate(seed_ranges):
                if index in changed:
                    continue
                min_in_range = seed_range[0] in range(old_min, old_min + diff)
                max_in_range = seed_range[1] in range(old_min, old_min + diff)
                if not min_in_range and not max_in_range:
                    continue
                elif min_in_range and not max_in_range:
                    seed_ranges.append([old_min + diff + 1, seed_range[1]])
                    seed_ranges[index][1] = old_min + diff
                elif not min_in_range and max_in_range:
                    seed_ranges.append([seed_range[0], old_min - 1])
                    seed_ranges[index][0] = old_min
                changed.append(index)
                seed_ranges[index][1] = seed_ranges[index][1] + (new_min - old_min)
                seed_ranges[index][0] = seed_ranges[index][0] + (new_min - old_min)
        else:
            changed = []
    return min(seed_ranges)[0]

print(part_one())
print(part_two())
