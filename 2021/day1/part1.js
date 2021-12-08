const fs = require("fs");

const input = fs.readFileSync("day1.txt", "utf-8").split("\n").map(Number);

let prev, totalIncreased = 0;
input.map((i) => {
    curr = parseInt(i);
    if (curr > prev) totalIncreased++;
    prev = curr;
});

console.log(`increased ${totalIncreased} times`);