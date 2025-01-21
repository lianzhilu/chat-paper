struct ResponseMetadata {
    1: required string RequestID
    2: required string Action
    3: optional Error Error (go.tag = "json:\"Error,omitempty\"")
}

struct Error {
    1: required string Code
    2: required string Message
}