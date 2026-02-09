# Words

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/anagrama-go">anagramasdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/anagrama-go#WordGetDailyResponse">WordGetDailyResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/anagrama-go">anagramasdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/anagrama-go#WordGetRandomResponse">WordGetRandomResponse</a>

Methods:

- <code title="get /api/v1/words/daily">client.Words.<a href="https://pkg.go.dev/github.com/stainless-sdks/anagrama-go#WordService.GetDaily">GetDaily</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/anagrama-go">anagramasdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/anagrama-go#WordGetDailyResponse">WordGetDailyResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/v1/words/random">client.Words.<a href="https://pkg.go.dev/github.com/stainless-sdks/anagrama-go#WordService.GetRandom">GetRandom</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/anagrama-go">anagramasdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/anagrama-go#WordGetRandomParams">WordGetRandomParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/anagrama-go">anagramasdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/anagrama-go#WordGetRandomResponse">WordGetRandomResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Wordsets

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/anagrama-go">anagramasdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/anagrama-go#WordsetResponse">WordsetResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/anagrama-go">anagramasdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/anagrama-go#WordsetWord">WordsetWord</a>

Methods:

- <code title="get /api/v1/wordsets/{wordsetId}">client.Wordsets.<a href="https://pkg.go.dev/github.com/stainless-sdks/anagrama-go#WordsetService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, wordsetID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/anagrama-go">anagramasdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/anagrama-go#WordsetGetParams">WordsetGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/anagrama-go">anagramasdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/anagrama-go#WordsetResponse">WordsetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Cli

## Auth

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/anagrama-go">anagramasdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/anagrama-go#CliAuthCompleteResponse">CliAuthCompleteResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/anagrama-go">anagramasdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/anagrama-go#CliAuthPollResponse">CliAuthPollResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/anagrama-go">anagramasdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/anagrama-go#CliAuthStartResponse">CliAuthStartResponse</a>

Methods:

- <code title="post /cli/auth/complete">client.Cli.Auth.<a href="https://pkg.go.dev/github.com/stainless-sdks/anagrama-go#CliAuthService.Complete">Complete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/anagrama-go">anagramasdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/anagrama-go#CliAuthCompleteParams">CliAuthCompleteParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/anagrama-go">anagramasdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/anagrama-go#CliAuthCompleteResponse">CliAuthCompleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /cli/auth/poll">client.Cli.Auth.<a href="https://pkg.go.dev/github.com/stainless-sdks/anagrama-go#CliAuthService.Poll">Poll</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/anagrama-go">anagramasdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/anagrama-go#CliAuthPollParams">CliAuthPollParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/anagrama-go">anagramasdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/anagrama-go#CliAuthPollResponse">CliAuthPollResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /cli/auth/start">client.Cli.Auth.<a href="https://pkg.go.dev/github.com/stainless-sdks/anagrama-go#CliAuthService.Start">Start</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/anagrama-go">anagramasdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/anagrama-go#CliAuthStartParams">CliAuthStartParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/anagrama-go">anagramasdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/anagrama-go#CliAuthStartResponse">CliAuthStartResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
