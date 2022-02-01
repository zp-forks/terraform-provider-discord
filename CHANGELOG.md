# Changelog

All notable changes to this project will be documented in this file. See [standard-version](https://github.com/conventional-changelog/standard-version) for commit guidelines.

### [1.0.4](https://github.com/Lucky3028/terraform-provider-discord/compare/v1.0.3...v1.0.4) (2022-02-01)


### Chores

* add validator ([ac33235](https://github.com/Lucky3028/terraform-provider-discord/commit/ac33235df8d625b269839e2e704790459a0f9343))
* remove default values ([b73c973](https://github.com/Lucky3028/terraform-provider-discord/commit/b73c9730fba324259dd9a5d9b73ec0619a654f4f))
* remove validator ([d775e79](https://github.com/Lucky3028/terraform-provider-discord/commit/d775e79e7c81567a62e1ae89d7fd9d9fff5be4c4))

### [1.0.3](https://github.com/Lucky3028/terraform-provider-discord/compare/v1.0.2...v1.0.3) (2022-02-01)


### Bug Fixes

* bitrate must be above or equal to 8000 ([20ea23a](https://github.com/Lucky3028/terraform-provider-discord/commit/20ea23a9ad1ae984debde1d70b89b2d72df1266f))

### [1.0.2](https://github.com/Lucky3028/terraform-provider-discord/compare/v1.0.1...v1.0.2) (2022-02-01)


### Bug Fixes

* d.Get doesn't return pointer ([a04c9f8](https://github.com/Lucky3028/terraform-provider-discord/commit/a04c9f8a6c6ce55a7cd13b0c683b36696a9a5e19))


### Chores

* add converter from int to uint ([9cfa7c0](https://github.com/Lucky3028/terraform-provider-discord/commit/9cfa7c0665e7413ced2e92c24430804508aadc18))

### [1.0.1](https://github.com/Lucky3028/terraform-provider-discord/compare/v1.0.0...v1.0.1) (2022-02-01)


### Bug Fixes

* ch name is not string point ([9d415c9](https://github.com/Lucky3028/terraform-provider-discord/commit/9d415c96d308776d9f7b2bcd54f460022df57250))

## 1.0.0 (2022-01-29)


### ⚠ BREAKING CHANGES

* **deps:** Update all dependencies

### Features

* add discord_news_channel ([9410e36](https://github.com/Lucky3028/terraform-provider-discord/commit/9410e3607b0227355e32b3e86257b11ecd6d2771))
* add manage events permission ([6a33322](https://github.com/Lucky3028/terraform-provider-discord/commit/6a33322a6c02eb94eb59a6a7fe79e12d5570df46))
* add permissions ([ce4b8ea](https://github.com/Lucky3028/terraform-provider-discord/commit/ce4b8eae5e8d77168d217e89394603c096b72326))
* add use application commands permission ([08dcaea](https://github.com/Lucky3028/terraform-provider-discord/commit/08dcaea08e2eea134192ce87e992a54973ea6a87))


### Bug Fixes

* compile error ([1dfe798](https://github.com/Lucky3028/terraform-provider-discord/commit/1dfe79863a2e691332f2f0be2d566c972fbfb5e8))
* repo owner ([0491fae](https://github.com/Lucky3028/terraform-provider-discord/commit/0491faece9dc550394ddd4e8b51dd460d11cb949))
* repo owner ([9ae8690](https://github.com/Lucky3028/terraform-provider-discord/commit/9ae86906912e2a1a8d7ef27cf16a81700c04d8eb))
* repo owner ([dfecabe](https://github.com/Lucky3028/terraform-provider-discord/commit/dfecabeb328c0d3003e364220f0d2e14358e8ed4))
* role position update when old role is unset is not error ([da50123](https://github.com/Lucky3028/terraform-provider-discord/commit/da50123cab19fd3a8e3742d97dbdc67d5d555348))
* server verification level can be set to 4 ([f6138fe](https://github.com/Lucky3028/terraform-provider-discord/commit/f6138fec689790ab5d2f6fc2e4f1c097f76e330a))
* sytem channel idが0の場合をエラーにしない ([bda9111](https://github.com/Lucky3028/terraform-provider-discord/commit/bda91113efc8eac5e6855d8396934f3feb98a70e))
* userlimit is "user_limit" ([2543c22](https://github.com/Lucky3028/terraform-provider-discord/commit/2543c227a0a6d3a41f2256df4977cfb3d7d0ea5d))
* とりあえず動くように ([611707b](https://github.com/Lucky3028/terraform-provider-discord/commit/611707b2387108fc52475bb70fd771a0b63bf575))


### Document Changes

* **readme:** add news_channel ([695e48a](https://github.com/Lucky3028/terraform-provider-discord/commit/695e48a2ad6a50ff7702fcac6ae59a46ee93ea12))
* **roles:** add sync_perms... in arg reference section ([103d9f4](https://github.com/Lucky3028/terraform-provider-discord/commit/103d9f4c4c4233e4d9f523351bc6eb6db73c164c))
* **todo:** add todo ([bdd85b9](https://github.com/Lucky3028/terraform-provider-discord/commit/bdd85b90fd63b808e5f3f12b5eb78d73cb3bc257))
* **todo:** add todo ([f8358c3](https://github.com/Lucky3028/terraform-provider-discord/commit/f8358c35a4b08c1d67a1d671174050a8188a75d2))


### Refactoring

* merge var declaration and if ([82c83f3](https://github.com/Lucky3028/terraform-provider-discord/commit/82c83f33652e37ab5ee6cfcf114f49040daed043))
* remove unnecessary nil check ([89f0a0f](https://github.com/Lucky3028/terraform-provider-discord/commit/89f0a0f61220a21c88975df7b8d4918a19c8445e))
* 変数にまとめられるものをまとめる ([ee2a1b1](https://github.com/Lucky3028/terraform-provider-discord/commit/ee2a1b1da40f113b14e137595ebad4868855b31f))


### Chores

* add .versionrc.json ([abbc5bc](https://github.com/Lucky3028/terraform-provider-discord/commit/abbc5bc305d3cda38ec8aa8eb22a24495b8c035c))
* **ci:** fix typo ([4b1a941](https://github.com/Lucky3028/terraform-provider-discord/commit/4b1a9416cd84d8a69bb74a48b8593dbf19add353))
* **deps:** Update dependencies ([719502f](https://github.com/Lucky3028/terraform-provider-discord/commit/719502f4d0c09b465cdd9873e15fb42b6f143a3b))
* remove makefile ([16c3c2b](https://github.com/Lucky3028/terraform-provider-discord/commit/16c3c2b53a776e61430494bffd6dfb9bcc3e9dd9))
* remove unnecessary option from goreleaser config ([66717c0](https://github.com/Lucky3028/terraform-provider-discord/commit/66717c02a5a29bcf7ad45a252a1cafbada1b9998))
* test output ([56c0c49](https://github.com/Lucky3028/terraform-provider-discord/commit/56c0c4948009c1b67243d299730925017a430373))
