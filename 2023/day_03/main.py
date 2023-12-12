ignored_symbols = ['.', '\n']

def get_lines():
    lines = []
    with open('day_03/input.txt') as input:
        for line in input:
            current_line = []
            number = ''
            min_index = 0
            max_index = 0
            for index, char in enumerate(line):
                if char.isdigit():
                    number += char
                    max_index = index
                else:
                    if number:
                        current_line.append({number: [min_index, max_index]})
                    number = ''
                    min_index = index + 1
                if not char.isdigit() and char not in ignored_symbols:
                    current_line.append({char: index})

            if number:
                current_line.append({number: [min_index, max_index]})
            lines.append(current_line)
    return lines

def check_line(line, symbols, last_added=None):
    sum = 0
    new_added = last_added if last_added else []
    for value in line:
        key = list(value.keys())[0]
        for symbol in symbols:
            symbol_key = list(symbol.keys())[0]
            if key.isdigit(
            ) and value[key][0] - 1 <= symbol[symbol_key] <= value[key][1] + 1:
                if {key: value[key]} in new_added:
                    continue
                new_added.append({key: value[key]})
                sum += int(key)
    return sum, new_added


def check_cog(lines, line_index, cog_index):
    cog_values = []
    if line_index > 0:
        for value in lines[line_index - 1]:
            key = list(value.keys())[0]
            if key.isdigit() and value[key][0] - 1 <= cog_index <= value[key][1] + 1:
                cog_values.append(key)

    for value in lines[line_index]:
        key = list(value.keys())[0]
        if key.isdigit() and value[key][0] - 1 <= cog_index <= value[key][1] + 1:
            cog_values.append(key)

    if line_index < len(lines) - 1:
        for value in lines[line_index + 1]:
            key = list(value.keys())[0]
            if key.isdigit() and value[key][0] - 1 <= cog_index <= value[key][1] + 1:
                cog_values.append(key)
    return int(cog_values[0]) * int(cog_values[1]) if len(cog_values) == 2 else 0


lines = get_lines()
sum = 0
gear_ratio = 0
prev_numbers = []
current_numbers = []
for line_index, line in enumerate(lines):
    symbols = []
    for value in line:
        key = list(value.keys())[0]
        if not key.isdigit():
            symbols.append(value)
        if key == '*':
            gear_ratio += check_cog(lines, line_index, value[key])

    if symbols:
        if line_index > 0:
            prev_sum, _ = check_line(lines[line_index - 1], symbols, prev_numbers)
            sum += prev_sum

        curr_sum, prev_numbers = check_line(line, symbols, current_numbers)
        sum += curr_sum

        if line_index < len(lines) - 1:
            next_sum, current_numbers = check_line(lines[line_index + 1], symbols)
            sum += next_sum

    else:
        prev_numbers = current_numbers

print('Sum:', sum)
print('Gear Ratio:', gear_ratio)
