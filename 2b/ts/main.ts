const fs = require("fs");

const values = fs.readFileSync("input.txt").toString().split("\n");

class Position {
  x: number;
  y: number;
  aim: number;

  // Normal signature with defaults
  constructor(x = 0, y = 0, aim = 0) {
    this.x = x;
    this.y = y;
    this.aim = aim;
  }

  plotCourse(direction: string, magnitude: number) {
    switch (direction) {
      case "down":
        this.aim += magnitude;
        break;
      case "up":
        this.aim -= magnitude;
        break;
      case "forward":
        this.x += magnitude;
        this.y += this.aim * magnitude;
        if (this.y < 0) {
          this.y = 0;
        }
        break;
      default:
        console.log("Failed to parse direction");
        break;
    }
  }
}
let sub = new Position();

for (let i = 0; i < values.length; i++) {
  let tokens = values[i].split(" ");
  if (tokens[0] != "" && tokens[1] != "") {
    sub.plotCourse(tokens[0], parseInt(tokens[1]));
  }
}
console.log(sub);
console.log(sub.x * sub.y);
