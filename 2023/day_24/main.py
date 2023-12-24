from sympy import var, Eq, solve

with open('day_24/input.txt') as f:
    lines = [line.strip('\n').split(' @ ') for line in f.readlines()]

def find_intersection(x1, y1, vx1, vy1, x2, y2, vx2, vy2):
    m1 = vy1 / vx1
    m2 = vy2 / vx2

    if m1 == m2:
        return None

    x = (y2 - y1 + m1 * x1 - m2 * x2) / (m1 - m2)
    y = m1 * (x - x1) + y1

    return (x, y)

def in_past(sc, ic, v):
    return (v < 0 and ic > sc) or (v > 0 and ic < sc)

def in_area(l1, l2, a_min, a_max):
    p1, v1, p2, v2 = l1[0].split(', '), l1[1].split(', '), l2[0].split(', '), l2[1].split(', ')
    x1, y1, vx1, vy1, x2, y2, vx2, vy2 = int(p1[0]), int(p1[1]), int(v1[0]), int(v1[1]), int(p2[0]), int(p2[1]), int(v2[0]), int(v2[1])

    intersection = find_intersection(x1, y1, vx1, vy1, x2, y2, vx2, vy2)
    if not intersection:
        return False
    if in_past(x1, intersection[0], vx1):
        return False
    if in_past(x2, intersection[0], vx2):
        return False
    if in_past(y1, intersection[1], vy1):
        return False
    if in_past(y2, intersection[1], vy2):
        return False

    return a_min <= intersection[0] <= a_max and a_min <= intersection[1] <= a_max

part1 = 0
for i, line in enumerate(lines):
    for j in range(i + 1, len(lines)):
        if in_area(line, lines[j], 200000000000000, 400000000000000):
            part1 += 1
print('Part 1:', part1)

sx = var("sx")
sy = var("sy")
sz = var("sz")
vx = var("vx")
vy = var("vy")
vz = var("vz")
equations = []

for line in lines:
    ps = list(map(int, line[0].split(", ")))
    vs = list(map(int, line[1].split(", ")))

    ts = "t{}".format(len(equations) // 3)
    exec(f'{ts} = var("{ts}")')

    equations.append(Eq(eval(f"{sx} + {vx} * {ts}"), eval(f"{ps[0]} + {vs[0]} * {ts}")))
    equations.append(Eq(eval(f"{sy} + {vy} * {ts}"), eval(f"{ps[1]} + {vs[1]} * {ts}")))
    equations.append(Eq(eval(f"{sz} + {vz} * {ts}"), eval(f"{ps[2]} + {vs[2]} * {ts}")))

    if len(equations) > 9:
        break

ans = solve(equations)[0]
print('Part 2:', ans[sx] + ans[sy] + ans[sz])
