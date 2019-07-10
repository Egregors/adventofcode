import {
    getHousesNumberForSanta,
    getHousesNumberForSantaAndRobot,
    getFile
} from "./2";

test("d3 p1", () => {
    expect(getHousesNumberForSanta(">")).toBe(2);
    expect(getHousesNumberForSanta("^>v<")).toBe(4);
    expect(getHousesNumberForSanta("^v^v^v^v^v")).toBe(2);
    expect(getHousesNumberForSanta(getFile("input.txt"))).toBe(2081);
});

test("d3 p2", () => {
    expect(getHousesNumberForSantaAndRobot("^v")).toBe(3);
    expect(getHousesNumberForSantaAndRobot("^>v<")).toBe(3);
    expect(getHousesNumberForSantaAndRobot("^v^v^v^v^v")).toBe(11);
    expect(getHousesNumberForSantaAndRobot(getFile("input.txt"))).toBe(2341);
});
