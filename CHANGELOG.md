# Changelog

## 0.8.0 (2026-03-28)

Full Changelog: [v0.7.0...v0.8.0](https://github.com/AnagramaGames/anagrama-go/compare/v0.7.0...v0.8.0)

### Features

* **internal:** support comma format in multipart form encoding ([68bc51e](https://github.com/AnagramaGames/anagrama-go/commit/68bc51edf1810ad36c0ff713f1f0f835ae92bad9))


### Bug Fixes

* allow canceling a request while it is waiting to retry ([985c69f](https://github.com/AnagramaGames/anagrama-go/commit/985c69f0fd71df033536e9a79d70478a59693240))
* prevent duplicate ? in query params ([f9e26b1](https://github.com/AnagramaGames/anagrama-go/commit/f9e26b13c048b58152231bb8122998c19284f824))


### Chores

* **ci:** skip lint on metadata-only changes ([1753cc6](https://github.com/AnagramaGames/anagrama-go/commit/1753cc67790640c569dce8a0045b755e2b70c705))
* **ci:** skip uploading artifacts on stainless-internal branches ([40a93a6](https://github.com/AnagramaGames/anagrama-go/commit/40a93a6e16ea1213c57c79f123728bec621ba737))
* **ci:** support opting out of skipping builds on metadata-only commits ([759deb7](https://github.com/AnagramaGames/anagrama-go/commit/759deb72308fbc7c238df4bc26f9cf131f2659a6))
* **client:** fix multipart serialisation of Default() fields ([821d4fa](https://github.com/AnagramaGames/anagrama-go/commit/821d4fae4db695036bfd4e7d6e5d05ffdb0dd2f7))
* **internal:** codegen related update ([3f160ba](https://github.com/AnagramaGames/anagrama-go/commit/3f160ba93e43d9013ff1f1eceaa6fe0ef959cfc1))
* **internal:** codegen related update ([4467db4](https://github.com/AnagramaGames/anagrama-go/commit/4467db4c403e02aaa22f35b75057d39d03174a49))
* **internal:** minor cleanup ([3d34b69](https://github.com/AnagramaGames/anagrama-go/commit/3d34b6920c7d2505fc8178090f10a6833556217e))
* **internal:** move custom custom `json` tags to `api` ([ec5a4e9](https://github.com/AnagramaGames/anagrama-go/commit/ec5a4e93f25beaa11cb6c65b165b6a176ffac907))
* **internal:** remove mock server code ([5d44348](https://github.com/AnagramaGames/anagrama-go/commit/5d4434819d9665ffbe860bfa7f63bd96d05dd2cc))
* **internal:** support default value struct tag ([637fbdc](https://github.com/AnagramaGames/anagrama-go/commit/637fbdc83fc565d0a6c5514fa10c2c68aaa9741d))
* **internal:** tweak CI branches ([7acb27c](https://github.com/AnagramaGames/anagrama-go/commit/7acb27c6a2da3f26cccf946c7dee13f259d1ac1e))
* **internal:** update gitignore ([ebfdbbf](https://github.com/AnagramaGames/anagrama-go/commit/ebfdbbf56550b0dcbbb4a1e626a443aa5b64ee60))
* **internal:** use explicit returns ([7a8928e](https://github.com/AnagramaGames/anagrama-go/commit/7a8928ea25b54f4a9db1684929c2f7f4f7791892))
* **internal:** use explicit returns in more places ([08eabc9](https://github.com/AnagramaGames/anagrama-go/commit/08eabc9dfed597bf104082d818191e8a6200b430))
* remove unnecessary error check for url parsing ([7c55f9e](https://github.com/AnagramaGames/anagrama-go/commit/7c55f9e7e369e70d2478fdc8d00a18eb07f1b67b))
* update docs for api:"required" ([8e455cd](https://github.com/AnagramaGames/anagrama-go/commit/8e455cd90ede72d06bb41d73e2d2d316f36002a3))
* update mock server docs ([ae1d3f2](https://github.com/AnagramaGames/anagrama-go/commit/ae1d3f2d00ef1ea75553158c56dc1cd87be55db2))

## 0.7.0 (2026-02-16)

Full Changelog: [v0.6.0...v0.7.0](https://github.com/AnagramaGames/anagrama-go/compare/v0.6.0...v0.7.0)

### Features

* **api:** api update ([4891a5b](https://github.com/AnagramaGames/anagrama-go/commit/4891a5bf87c3f08834bd1dcd65a1d7a77607e8f3))

## 0.6.0 (2026-02-16)

Full Changelog: [v0.5.0...v0.6.0](https://github.com/AnagramaGames/anagrama-go/compare/v0.5.0...v0.6.0)

### Features

* **api:** api update ([5be807c](https://github.com/AnagramaGames/anagrama-go/commit/5be807c7d719fa2e6c5b3d5a55ec0f1c1fa69117))
* **api:** api update ([660e64e](https://github.com/AnagramaGames/anagrama-go/commit/660e64ec48aafdeeaaf6a88fefbaaa2175dcb6a5))

## 0.5.0 (2026-02-16)

Full Changelog: [v0.4.1...v0.5.0](https://github.com/AnagramaGames/anagrama-go/compare/v0.4.1...v0.5.0)

### Features

* **api:** api update ([ddc03e3](https://github.com/AnagramaGames/anagrama-go/commit/ddc03e3be582ecd31d43e3ae6f1c7f428822642d))

## 0.4.1 (2026-02-11)

Full Changelog: [v0.4.0...v0.4.1](https://github.com/AnagramaGames/anagrama-go/compare/v0.4.0...v0.4.1)

### Bug Fixes

* **encoder:** correctly serialize NullStruct ([db9d5ea](https://github.com/AnagramaGames/anagrama-go/commit/db9d5eaccec16bd859ceb0ff2dc615c916709626))

## 0.4.0 (2026-02-09)

Full Changelog: [v0.3.0...v0.4.0](https://github.com/AnagramaGames/anagrama-go/compare/v0.3.0...v0.4.0)

### Features

* **api:** api update ([3bc539b](https://github.com/AnagramaGames/anagrama-go/commit/3bc539bc47666f6a7fdc03fb3030b9aebc4190ef))

## 0.3.0 (2026-02-09)

Full Changelog: [v0.2.0...v0.3.0](https://github.com/AnagramaGames/anagrama-go/compare/v0.2.0...v0.3.0)

### Features

* **api:** api update ([b25b628](https://github.com/AnagramaGames/anagrama-go/commit/b25b628b8b5abebab63379230bab1edc2e4b7926))
* **api:** api update ([096eb7b](https://github.com/AnagramaGames/anagrama-go/commit/096eb7b3af372e291e26fcd7d1ee2098106427d3))
* **api:** api update ([c9e6302](https://github.com/AnagramaGames/anagrama-go/commit/c9e6302d689aafafd846f351aee036370ca12652))

## 0.2.0 (2026-02-09)

Full Changelog: [v0.1.0...v0.2.0](https://github.com/AnagramaGames/anagrama-go/compare/v0.1.0...v0.2.0)

### Features

* **api:** api update ([ab6987b](https://github.com/AnagramaGames/anagrama-go/commit/ab6987b5bc234ab9397074db9fc0e57094bfd331))


### Chores

* update SDK settings ([071552a](https://github.com/AnagramaGames/anagrama-go/commit/071552a6deda2627dfd93aad79979dfc6452d749))

## 0.1.0 (2026-02-09)

Full Changelog: [v0.0.1...v0.1.0](https://github.com/AnagramaGames/anagrama-go/compare/v0.0.1...v0.1.0)

### Features

* **api:** api update ([4daa746](https://github.com/AnagramaGames/anagrama-go/commit/4daa746c14c3837f2434d73ee25c5af0b816d74e))


### Chores

* update SDK settings ([65e57b8](https://github.com/AnagramaGames/anagrama-go/commit/65e57b8df9752a4a91618bf51ccea1490e61a4ae))
