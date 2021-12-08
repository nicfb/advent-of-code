horizontal = 0
depth = 0

with open("day2.txt") as file:
    for line in file:
        cmd = line.rstrip("\n").split(" ")
        direction = cmd[0]
        amount = cmd[1]
        
        if direction == "forward":
            horizontal += int(amount)
        elif direction == "up":
            depth -= int(amount)
        elif direction == "down":
            depth += int(amount)
    
    print(f"horizontal: {horizontal}, depth: {depth}")
