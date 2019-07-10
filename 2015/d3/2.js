// https://adventofcode.com/2015/day/3

import { readFileSync } from "fs";

export const getFile = path => readFileSync(`${__dirname}/${path}`, "utf8");
// ---
const getLastPoint = arr => arr[arr.length - 1];
const getX = point => parseInt(point.split(":")[0]);
const getY = point => parseInt(point.split(":")[1]);
const makePoint = (x, y) => `${x}:${y}`;

export const solve = commands => {
    const r = (acc, el) => {
        const lastPosition = getLastPoint(acc);
        const x = getX(lastPosition);
        const y = getY(lastPosition);

        switch (el) {
            case ">":
                return [...acc, makePoint(x + 1, y)];
            case "v":
                return [...acc, makePoint(x, y - 1)];
            case "<":
                return [...acc, makePoint(x - 1, y)];
            case "^":
                return [...acc, makePoint(x, y + 1)];

            default:
                return acc;
        }
    };

    return new Set(commands.split("").reduce(r, ["0:0"])).size;
};

export default solve;
