import { Operation, type Config } from "../config";
import Projector from "../projector";


function getData() {
    return {
        projector: {
            "/": {
                "foo": "bar",
                "fem": "boy"
            },
            "/foo": {
                "foo": "bar2",
            },
            "/foo/bar": {
                "foo": "bar3",
            },

        }
    }
}

function getProjector(pwd: string, data = getData()): Projector {
    return new Projector({
        args: [],
        operation: Operation.Print,
        pwd: pwd,
        config: "Hello,something"
    }, data);
}

test("getValueAll", function() {
    const projector = getProjector("/foo/bar");

    expect(projector.getValueAll()).toEqual({
        "fem": "boy",
        "foo": "bar3"
    })
})

test("getValueKey", function() {
    const projector = getProjector("/foo/bar");
    expect(projector.getValue("foo")).toEqual("bar3")

    const projector1 = getProjector("/foo");
    expect(projector1.getValue("foo")).toEqual("bar2")
    const projector2 = getProjector("/");
    expect(projector2.getValue("foo")).toEqual("bar")
})
test("setValue", function() {
    let data = getData();
    let projector = getProjector("/foo/bar", data);
    projector.setValue("foo", "barzzz")
    expect(projector.getValue("foo")).toEqual("barzzz");
    projector = getProjector("/foo", data);
    projector.setValue("foo", "barzzz2");
    expect(projector.getValue("foo")).toEqual("barzzz2");
    projector = getProjector("/", data);
    projector.setValue("foo", "barzzz3");
    expect(projector.getValue("foo")).toEqual("barzzz3");

})

test("removeValue", function() {
    const projector = getProjector("/foo/bar");
    projector.removeValue("foo");
    expect(projector.getValue("foo")).toEqual("bar2");
})



