"use strict";
const fs = require("fs");
const values = fs.readFileSync("input.txt").toString().split("\n");
let counter = 0;
for (let i = 0; i < values.length; i++) {
    if (parseInt(values[i]) < parseInt(values[i + 1])) {
        counter++;
    }
}
console.log(counter);
