import { solve, getFile } from "./2";

test("d1 p1", () => {
    expect(solve(">")).toBe(2);
    expect(solve("^>v<")).toBe(4);
    expect(solve("^v^v^v^v^v")).toBe(2);
    expect(solve(getFile("input.txt"))).toBe(2081);
});
