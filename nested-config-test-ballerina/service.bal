import ballerina/http;

// Testbed for exercising nested Ballerina configurables: an array of
// primitives, an array of records, a record nested inside a record, and an
// array of records nested inside a record that is itself nested inside
// another record — the shapes `choreo describe component-config` needs to
// enumerate correctly.

type Contact record {|
    string email;
    string phone;
|};

type Address record {|
    string city;
    string[] tags;
    Contact[] contacts;
|};

type Profile record {|
    string name;
    int age;
    string[] roles;
    Address address;
|};

configurable string[] destinations = ["dest-a", "dest-b"];
configurable Profile profile = {
    name: "default",
    age: 0,
    roles: ["viewer"],
    address: {
        city: "colombo",
        tags: ["home"],
        contacts: [{email: "a@example.com", phone: "0000000000"}]
    }
};

service /nested\-config\-test on new http:Listener(9090) {

    resource function get destinations() returns string[] {
        return destinations;
    }

    resource function get profile() returns Profile {
        return profile;
    }
}
