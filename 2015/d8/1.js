// https://adventofcode.com/2015/day/8

const getAscii = hex => String.fromCharCode(parseInt(hex, 16));

export const getChInCode = line => line.length;

const replaceEscapes = line => {
    const cleanLine = [];
    line = line.slice(1, -1);
    for (let i = 0; i < line.length; i++) {
        if (line[i] === "\\") {
            switch (line[i + 1]) {
                case '"':
                    cleanLine.push('"');
                    i = i + 1;
                    continue;

                case "\\":
                    cleanLine.push("\\");
                    i = i + 1;
                    continue;

                case "x":
                    cleanLine.push(getAscii(line.slice(i, i + 2)));
                    i = i + 3;
                    continue;

                default:
                    break;
            }
        } else {
            cleanLine.push(line[i]);
        }
    }
    return cleanLine.join("");
};

export const getChInMemory = line => replaceEscapes(line).length;

export const matchsticks = lines => {
    lines = lines.split("\n");
    return (
        lines.reduce((acc, el) => acc + getChInCode(el), 0) -
        lines.reduce((acc, el) => acc + getChInMemory(el), 0)
    );
};

export default matchsticks;
