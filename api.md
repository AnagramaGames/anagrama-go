# Words

Response Types:

- <a href="https://pkg.go.dev/github.com/AnagramaGames/anagrama-go">anagramasdk</a>.<a href="https://pkg.go.dev/github.com/AnagramaGames/anagrama-go#WordGetDailyResponse">WordGetDailyResponse</a>
- <a href="https://pkg.go.dev/github.com/AnagramaGames/anagrama-go">anagramasdk</a>.<a href="https://pkg.go.dev/github.com/AnagramaGames/anagrama-go#WordGetRandomResponse">WordGetRandomResponse</a>

Methods:

- <code title="get /api/v1/words/daily">client.Words.<a href="https://pkg.go.dev/github.com/AnagramaGames/anagrama-go#WordService.GetDaily">GetDaily</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/AnagramaGames/anagrama-go">anagramasdk</a>.<a href="https://pkg.go.dev/github.com/AnagramaGames/anagrama-go#WordGetDailyResponse">WordGetDailyResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/v1/words/random">client.Words.<a href="https://pkg.go.dev/github.com/AnagramaGames/anagrama-go#WordService.GetRandom">GetRandom</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/AnagramaGames/anagrama-go">anagramasdk</a>.<a href="https://pkg.go.dev/github.com/AnagramaGames/anagrama-go#WordGetRandomParams">WordGetRandomParams</a>) (\*<a href="https://pkg.go.dev/github.com/AnagramaGames/anagrama-go">anagramasdk</a>.<a href="https://pkg.go.dev/github.com/AnagramaGames/anagrama-go#WordGetRandomResponse">WordGetRandomResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Wordsets

Response Types:

- <a href="https://pkg.go.dev/github.com/AnagramaGames/anagrama-go">anagramasdk</a>.<a href="https://pkg.go.dev/github.com/AnagramaGames/anagrama-go#WordsetResponse">WordsetResponse</a>
- <a href="https://pkg.go.dev/github.com/AnagramaGames/anagrama-go">anagramasdk</a>.<a href="https://pkg.go.dev/github.com/AnagramaGames/anagrama-go#WordsetWord">WordsetWord</a>

Methods:

- <code title="get /api/v1/wordsets/{wordsetId}">client.Wordsets.<a href="https://pkg.go.dev/github.com/AnagramaGames/anagrama-go#WordsetService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, wordsetID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/AnagramaGames/anagrama-go">anagramasdk</a>.<a href="https://pkg.go.dev/github.com/AnagramaGames/anagrama-go#WordsetGetParams">WordsetGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/AnagramaGames/anagrama-go">anagramasdk</a>.<a href="https://pkg.go.dev/github.com/AnagramaGames/anagrama-go#WordsetResponse">WordsetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Cli

## Auth

Response Types:

- <a href="https://pkg.go.dev/github.com/AnagramaGames/anagrama-go">anagramasdk</a>.<a href="https://pkg.go.dev/github.com/AnagramaGames/anagrama-go#CliAuthCompleteResponse">CliAuthCompleteResponse</a>
- <a href="https://pkg.go.dev/github.com/AnagramaGames/anagrama-go">anagramasdk</a>.<a href="https://pkg.go.dev/github.com/AnagramaGames/anagrama-go#CliAuthPollResponse">CliAuthPollResponse</a>
- <a href="https://pkg.go.dev/github.com/AnagramaGames/anagrama-go">anagramasdk</a>.<a href="https://pkg.go.dev/github.com/AnagramaGames/anagrama-go#CliAuthStartResponse">CliAuthStartResponse</a>

Methods:

- <code title="post /cli/auth/complete">client.Cli.Auth.<a href="https://pkg.go.dev/github.com/AnagramaGames/anagrama-go#CliAuthService.Complete">Complete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/AnagramaGames/anagrama-go">anagramasdk</a>.<a href="https://pkg.go.dev/github.com/AnagramaGames/anagrama-go#CliAuthCompleteParams">CliAuthCompleteParams</a>) (\*<a href="https://pkg.go.dev/github.com/AnagramaGames/anagrama-go">anagramasdk</a>.<a href="https://pkg.go.dev/github.com/AnagramaGames/anagrama-go#CliAuthCompleteResponse">CliAuthCompleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /cli/auth/poll">client.Cli.Auth.<a href="https://pkg.go.dev/github.com/AnagramaGames/anagrama-go#CliAuthService.Poll">Poll</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/AnagramaGames/anagrama-go">anagramasdk</a>.<a href="https://pkg.go.dev/github.com/AnagramaGames/anagrama-go#CliAuthPollParams">CliAuthPollParams</a>) (\*<a href="https://pkg.go.dev/github.com/AnagramaGames/anagrama-go">anagramasdk</a>.<a href="https://pkg.go.dev/github.com/AnagramaGames/anagrama-go#CliAuthPollResponse">CliAuthPollResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /cli/auth/start">client.Cli.Auth.<a href="https://pkg.go.dev/github.com/AnagramaGames/anagrama-go#CliAuthService.Start">Start</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/AnagramaGames/anagrama-go">anagramasdk</a>.<a href="https://pkg.go.dev/github.com/AnagramaGames/anagrama-go#CliAuthStartParams">CliAuthStartParams</a>) (\*<a href="https://pkg.go.dev/github.com/AnagramaGames/anagrama-go">anagramasdk</a>.<a href="https://pkg.go.dev/github.com/AnagramaGames/anagrama-go#CliAuthStartResponse">CliAuthStartResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
