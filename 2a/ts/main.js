var fs = require("fs");
var values = fs.readFileSync("input.txt").toString().split("\n");
// position tuple
// forward is positive
// down is positive, up is negative (maybe enum to translate) (cannot have a negative depth, so it must reset there if necessary)
var Position = /** @class */ (function () {
    // Normal signature with defaults
    function Position(x, y) {
        if (x === void 0) { x = 0; }
        if (y === void 0) { y = 0; }
        this.x = x;
        this.y = y;
    }
    Position.prototype.plotCourse = function (direction, magnitude) {
        switch (direction) {
            case "down":
                this.y += magnitude;
                break;
            case "up":
                this.y -= magnitude;
                if (this.y < 0) {
                    this.y = 0;
                }
                break;
            case "forward":
                this.x += magnitude;
                break;
            default:
                console.log("Failed to parse direction");
                break;
        }
    };
    return Position;
}());
var sub = new Position();
for (var i = 0; i < values.length; i++) {
    var tokens = values[i].split(" ");
    if (tokens[0] != "" && tokens[1] != "") {
        console.log(tokens[0] + tokens[1]);
        sub.plotCourse(tokens[0], parseInt(tokens[1]));
    }
}
console.log(sub);
console.log(sub.x * sub.y);
