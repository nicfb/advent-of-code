horizontal = 0
depth = 0
aim = 0

with open("day2.txt") as file:
    for line in file:
        cmd = line.rstrip("\n").split(" ")
        direction = cmd[0]
        amount = int(cmd[1])
        
        if direction == "forward":
            horizontal += amount
            depth += aim * amount
        elif direction == "up":
            aim -= amount
        elif direction == "down":
            aim += amount
    
    print(f"horizontal: {horizontal}, depth: {depth}")
