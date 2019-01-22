# wercker status

[![wercker status](https://app.wercker.com/status/52664147f077abcb9657ac9831ec0211/s/master "wercker status")](https://app.wercker.com/project/byKey/52664147f077abcb9657ac9831ec0211)



## Cacoo go client


## Install

you can use `go get`
```bash
go get github.com/mehmetalisavas/mongo-cache
```

or you can clone into specific path with `git clone`
```bash
git clone git@github.com:mehmetalisavas/cacoo.git
```

## Description about package
This package implements the Cacoo api, It's written in golang.
No recommended version is specified for this package. But also, still better to use latest version of go if posssible.
If there is a misunderstanding points you can check the test cases.
If you find bug or missing points, feel free to open a pull request. If you won't be able to do that, then please feel free to open an issue in repo.

## Usage

```go
// you can create client with token or without token(you won't be able to use many methods that requires api key)
client := NewClient("your token")

// you can fetch your account information with this
account, _, err := client.Account.MyAccountInformation(context.Background())

// If you want to use account name, you can get it like this after using above method
myName := account.Name



Granted Client options are listed below:
- func OptionHttpClient(*http.Client) Option
- func OptionUserAgent(agent string) Option // use custom user agent
- func OptionBaseURL(url string) Option // sets the custom url for client

```
