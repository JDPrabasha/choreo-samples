import ballerina/http;
import ballerina/lang.runtime;
import ballerina/log;
import ballerina/time;

service /timeout\-test on new http:Listener(9090) {

    // GET /timeout-test/sleep?seconds=N sleeps for N seconds (default 90) before
    // responding, so a caller can observe whether the gateway/enforcer cuts the
    // request off before the response is actually sent.
    resource function get sleep(int seconds = 90) returns string {
        time:Utc startTime = time:utcNow();
        log:printInfo(string `Sleep request received, sleeping for ${seconds}s`);
        runtime:sleep(<decimal>seconds);
        decimal elapsed = time:utcDiffSeconds(time:utcNow(), startTime);
        string respondedAt = time:utcToString(time:utcNow());
        return string `Slept for ${elapsed}s (requested ${seconds}s), responded at ${respondedAt}`;
    }
}
