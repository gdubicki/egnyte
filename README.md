# Egnyte Go SDK

| Author | Email |
| --- | --- |
| Dmytro Soloviov | dmytro.soloviov@gmail.com |

Unofficial [Egnyte](https://www.egnyte.com) Go SDK.

## Usage

[Support page](https://developers.egnyte.com/docs)

```go
token := egnyte.GetAccessToken("domain", "clientId", "clientSecret", "username", "password")
egnyte := egnyte.Connect(token, "domain")
fmt.Println(egnyte.User.Current(egnyte))
```

where `domain` is your Egnyte domain. For example, for `https://example.egnyte.com`, acceptable values are:

- `https://example.egnyte.com`
- `example.egnyte.com`
- `example`

Currently the only suuported auth method is [internal app](https://developers.egnyte.com/docs/read/Public_API_Authentication#Internal-Applications). Hovewer, if you already have Egnyte Public API access token, you can omit using `GetAccessToken()` function.

### User Management API

[Support page](https://developers.egnyte.com/docs/read/User_Management_API_Documentation)

```go
// Get currently authenticated user
user := egnyte.User.Me(egnyte)
// Get user details by ID
user := egnyte.User.Details(egnyte, 42)
```