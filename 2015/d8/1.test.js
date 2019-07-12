import { getChInCode, getChInMemory, matchsticks } from "./1";

import { readFileSync } from "fs";
export const getFile = path => readFileSync(`${__dirname}/${path}`, "utf8");

const [case1, case2, case3, case4] = getFile("fixture.txt").split("\n");

test("d8 p1 get number of characters in code", () => {
    expect(getChInCode(case1)).toBe(2);
    expect(getChInCode(case2)).toBe(5);
    expect(getChInCode(case3)).toBe(10);
    expect(getChInCode(case4)).toBe(6);
});

test("d8 p1 get number of characters in memory", () => {
    expect(getChInMemory(case1)).toBe(0);
    expect(getChInMemory(case2)).toBe(3);
    expect(getChInMemory(case3)).toBe(7);
    expect(getChInMemory(case4)).toBe(1);
});

test("d8 p1 get diff (ch in code - ch in mem)", () => {
    expect(matchsticks(getFile("fixture.txt"))).toBe(12);
    expect(matchsticks(getFile("input.txt"))).toBe(1350);
});
