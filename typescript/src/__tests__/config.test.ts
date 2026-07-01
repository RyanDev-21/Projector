import getConfig, { Operation } from "../config";

test("simple test print", function() {
    const config = getConfig({});
    expect(config.operation).toEqual(Operation.Print);
    expect(config.args).toEqual([]);
});
test("simple test print with argument 1", function() {
    const config = getConfig({ args: ["foo"] });
    expect(config.operation).toEqual(Operation.Print);
    expect(config.args).toEqual(["foo"]);
});

test("simple test add", function() {
    const config = getConfig({
        args: ["add", "foo", "bar"]
    });
    expect(config.operation).toEqual(Operation.Add);
    expect(config.args).toEqual(["foo", "bar"]);
});

test("simple test remove", function() {
    const config = getConfig({
        args: ["rm", "foo"]
    });
    expect(config.operation).toEqual(Operation.Remove);
    expect(config.args).toEqual(["foo"]);
});
