const fs = require("fs");

const values = fs.readFileSync("input.txt").toString().split("\n");

interface power {
  gamma: number;
  epsilon: number;
  result: number;
}

function computerPower(onesCounter: number[], total: number): power {
  let gammaString = "";
  let epsilonString = "";

  for (let i = 0; i < onesCounter.length; i++) {
    if (onesCounter[i] > total / 2) {
      gammaString += "1";
      epsilonString += "0";
    } else {
      gammaString += "0";
      epsilonString += "1";
    }
  }
  let pow: power = {
    gamma: parseInt(gammaString, 2),
    epsilon: parseInt(epsilonString, 2),
    result: parseInt(gammaString, 2) * parseInt(epsilonString, 2),
  };
  return pow;
}

// init onesCounter
let onesCounter = new Array(values[0].length);
for (let i = 0; i < onesCounter.length; i++) {
  onesCounter[i] = 0;
}

for (let i = 0; i < values.length; i++) {
  for (let j = 0; j < values[i].length; j++) {
    onesCounter[j] += parseInt(values[i][j]);
  }
}

let power = computerPower(onesCounter, values.length);
console.log(power);
