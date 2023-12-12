max_cubes = {
    'blue': 14,
    'green': 13,
    'red': 12,
}

def is_possible(cubes):
    return all(amount <= max_cubes[colour] for colour, amount in cubes.items())

def product(cubes):
    result = 1
    for amount in cubes.values():
        result *= amount

    return result

with open('day_02/input.txt') as puzzle_input:
    sum = 0
    power = 0
    for line in puzzle_input:
        game = line.split(': ')
        id = int(game[0].replace('Game ', ''))
        rounds = game[1].replace('\n', '').split('; ')

        most_cubes = {}
        for round in rounds:
            cubes = round.split(', ')
            for cube in cubes:
                cube = cube.split(' ')
                amount = int(cube[0])
                colour = cube[1]

                if colour not in most_cubes or most_cubes[colour] < amount:
                    most_cubes[colour] = amount

        if is_possible(most_cubes):
            sum += id

        power += product(most_cubes)

    print('Sum: ', sum)
    print('Power: ', power)
