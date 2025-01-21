namespace go user

struct RegisterRequest {
    1: required string Name
    2: required string Password

}

struct RegisterResponse {
    1: required string ID
    2: required string Name
}

struct LoginRequest {
    1: required string ID
    2: required string Password
}

struct LoginResponse {
    1: required string ID
    2: required string Name
}

service UserService {
    RegisterResponse Register(1: RegisterRequest req);
    LoginResponse Login(1: LoginRequest req);
}