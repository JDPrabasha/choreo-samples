import ballerina/http;

type Greeting record {
    string 'from;
    string to;
    string message;
};

type MssqlConfig record {
    string host;
    int port;
    string username;
};

configurable MssqlConfig mssqlConfig = {host: "localhost", port: 1433, username: "sa"};
configurable string greetingPrefix = "Welcome to";

service / on new http:Listener(8090) {
    resource function get .(string name) returns Greeting {
        Greeting greetingMessage = {"from" : "Choreo", "to" : name, "message" : greetingPrefix + " Choreo! (db: " + mssqlConfig.host + ")"};
        return greetingMessage;
    }
}
