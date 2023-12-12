with open('day_06/input.txt') as f:
    lines = f.readlines()

data = []

def part_one():
    times = lines[0].split(':')[1].strip().split(' ')
    for time in times:
        if time:
            data.append({'time': int(time)})
    records = lines[1].split(':')[1].strip().split(' ')
    count = 0
    for record in records:
        if record:
            data[count]['record'] = int(record)
            count += 1
def part_two():
    time = lines[0].split(':')[1].strip().replace(' ', '')
    data.append({'time': int(time)})
    record = lines[1].split(':')[1].strip().replace(' ', '')
    data[0]['record'] = int(record)

# part_one()
part_two()

product = 1
for datum in data:
    count = 0
    is_even = datum['time'] % 2 == 0
    time = datum['time'] // 2
    while True:
        personal_record = time * (datum['time'] - time)
        if personal_record > datum['record']:
            count += 1
            time -= 1
        else:
            break
    count *= 2
    if is_even:
        count -= 1
    product *= count

print(product)
