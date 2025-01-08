namespace go user

struct RegisterRequest {
    1: required string ID
    2: required string Password
    3: required string Name
}

struct RegisterResponse {
    1: required string ID
}



service UserService {
    RegisterResponse Register(1: RegisterRequest req);
}