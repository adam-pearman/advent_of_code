def calculate_score(numbers):
    score = 0
    for _ in numbers:
        if not score:
            score = 1
        else:
            score *= 2

    return score

def check_for_winner(lines, index):
    sum = 0
    count = 0
    line = lines[index]

    numbers = line.split(': ')[1]
    seperated_numbers = numbers.split(' | ')
    winning_numbers = set(seperated_numbers[0].split())
    player_numbers = set(seperated_numbers[1].split())
    player_winning_numbers = winning_numbers & player_numbers

    sum += calculate_score(player_winning_numbers)

    for i, _ in enumerate(player_winning_numbers):
        new_index = index + i + 1
        if new_index < len(lines):
            count += 1
            _, cards = check_for_winner(lines, new_index)
            count += cards

    return sum, count

with open('day_04/input.txt') as input:
    sum = 0
    count = 0
    lines = []
    for line in input:
        lines.append(line.replace('\n', ''))
    for index, _ in enumerate(lines):
        count += 1
        score, cards = check_for_winner(lines, index)
        sum += score
        count += cards
    print('Sum:', sum)
    print('Count:', count)
