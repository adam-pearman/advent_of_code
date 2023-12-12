def convert_numbers(line):
    numbers = {
        'one': '1',
        'two': '2',
        'three': '3',
        'four': '4',
        'five': '5',
        'six': '6',
        'seven': '7',
        'eight': '8',
        'nine': '9',
    }

    first = {}
    last = {}

    for number in numbers:
        if number in line:
            first_index = line.find(number)
            last_index = line.rfind(number)

            if not first or first_index < list(first.values())[0]:
                first = {number: first_index}
            if not last or last_index > list(last.values())[0]:
                last = {number: last_index}

    if first:
        number = list(first.keys())[0]
        index = list(first.values())[0]
        line_list = list(line)
        line_list.insert(index, numbers[number])
        line = ''.join(line_list)

    if last:
        number = list(last.keys())[0]
        index = list(last.values())[0]
        line_list = list(line)
        line_list.insert(index + len(number) + 1, numbers[number])
        line = ''.join(line_list)

    return line

with open('day_01/input.txt') as puzzle_input:
    sum = 0
    for line in puzzle_input:
        line = convert_numbers(line)

        first = None
        last = None

        for char in line:
            if char.isdigit() and first is None:
                first = char
                last = char
            elif char.isdigit():
                last = char

        if first and last:
            sum += int(str(first) + str(last))

    print(sum)
