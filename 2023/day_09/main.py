with open('day_09/input.txt') as f:
    lines = f.readlines()

def calc_differences(numbers):
    differences = []
    for index, number in enumerate(numbers):
        if index == len(numbers) - 1:
            break
        differences.append(number - numbers[index + 1])
    if len(set(differences)) == 1:
        return numbers[0] + differences[0]
    return numbers[0] + calc_differences(differences)

end_sum = 0
beginning_sum = 0
for line in lines:
    # Part 1
    numbers = [int(x) for x in line.strip().split()][::-1]
    end_sum += calc_differences(numbers)
    # Part 2
    numbers = [int(x) for x in line.strip().split()]
    beginning_sum += calc_differences(numbers)

print('Sum of Last Digits:', end_sum)
print('Sum of First Digits:', beginning_sum)
