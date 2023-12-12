with open('day_11/input.txt') as f:
    lines = [line.strip() for line in f.readlines()]

stars = []
double_rows = []
double_cols = []

for index, line in enumerate(lines):
    has_star = False
    for i, char in enumerate(line):
        if char == '#':
            stars.append((index, i))
            has_star = True
    if not has_star:
        double_rows.append(index)

for i in range(len(lines[0])):
    if i not in [star[1] for star in stars]:
        double_cols.append(i)

# Part 1
# space_size = 1

# Part 2
space_size = 999999

steps = 0
for index, star in enumerate(stars):
    for i in range(index + 1, len(stars)):
        next_star = stars[i]
        row_diff = abs(star[0] - next_star[0])
        col_diff = abs(star[1] - next_star[1])
        extra_rows = sum([space_size for row in double_rows if min(star[0], next_star[0]) < row < max(star[0], next_star[0])])
        extra_cols = sum([space_size for col in double_cols if min(star[1], next_star[1]) < col < max(star[1], next_star[1])])
        steps += row_diff + col_diff + extra_rows + extra_cols

print(steps)
