# Introduction

## Priority

Dalam membuat program, prioritaskan hal-hal berikut.

1. *Correctness*, program harus berjalan sesuai dengan tujuan & fungsi
   bisnisnya. Percuma menggunakan teknologi terbaru, library populer tapi
   program tidak berjalan sesuai ekspektasi fungsi & requirement. Meskipun
   demikian, bugs logic & error pada program pasti akan kita jumpai.
2. *Readability*, program mudah dibaca oleh anggota tim lainnya. Oleh
   karena itu, perlu standar yg sama dalam hal penamaan route path, variable,
   file, struktur & organisasi foldernya. Termasuk logging, karena terkait
   readability dalam proses debugging.
4. *Performace*, jika program sudah berjalan dengan benar & source code-nya
   rapih, baru kita pikirkan performance optimization. Dengan menggunakan
   GoLang, 90% kita sudah terbantu dalam hal performa & utilisasi resource
   server.

## Teknologi

Berikut adalah teknologi yang akan kita gunakan dalam pengembangan backend
services Pemuda Peduli.

* HTTP JSON API dari mobile app menggunakan `fasthttp`
  project bisa dilihat di folder [gcp-pubsub-example](../examples/gcp-pubsub/main.go)
* Cache menggunakan redis
* RDBMS menggunakan postgres

# GoLang Service Design

## Data Flow

1. Request user masuk ke REST endpoint
2. Validasi input user di level handler/domain. Validasi tidak meliputi
   pengecekan ke external system (db, redis, http). Murni hanya pengecekan
   berdasarkan payload request-nya.
3. Build request untuk service object (bila perlu), lalu call fungsi service
   tsb.
4. Berdasarkan response dari service object, handler/domain menentukan:
    * Jika sukses, build & return ok response / payload ke user.
    * Jika gagal, return error-nya langsung.

## Service method naming guide

Gunakan hanya [Create, Get, Put, Delete, Update, Find] prefix.
e.g. :

* `Create` untuk membuat resource baru (baik menulis ke storage / tidak)
    * contoh: `CreateUser(ctx context.Context, ID string) (User, error)`
* `Get` untuk mendapatkan _*single*_ resource yang sudah ada sebelumnya di
  datasource
    * contoh: `GetUser(ctx context.Context, ID string) (User, error)`
* `Update` untuk mengupdate *single* resource yang sudah ada sebelumnya secara
  keseluruhan (replace existing resource)
    * contoh: `UpdateUser(ctx context.Context, u *User)eerror`
* `Put` untuk mengupdate *single* resource secara parsial.
* `Find` untuk mendapatkan *many* resources yang sudah ada sebelumnya di
  datasrouce
    * `FindUser(ctx context.Context, options FilterOptions) ([]User, error)`
* `Delete` untuk menghapus *single* resource di datasource
    * `DeleteUser(ctx context.Context, ID string) error`

2. Gunakan `context.Context` sebagai parameter pertama untuk fungsi yang berinteraksi dengan sistem external.
    * External call ke redis ? gunakan `context.Context`
    * query ke postgres / mongodb ? gunakan `context.Context`
    * External call ke redis ? gunakan `context.Context`

# Go Package Layout

Secara garis besar terbagi menjadi 2 package folder utama
* `cmd` berisi main program
* `src` app business logic

## Service Package Design

```console
user_service_repo
    cmd/ <- boleh import package
        main.go <- bisa run rest di port yg sama, run /internal/{grpc,http}
    src/ <-- boleh import /pkg, tidak boleh diimport secara go-module dari repo luar
        user/ <- boleh import repository, api_client, tidak boleh import /internal/http & internal/grpc
            application/
                user_app <- untuk route path default method (dalam case ini untuk module user)
                user_payload <- untuk handle payload (request) dan mapping response payload untuk user
            domain/ <- boleh import repository, api_client, tidak boleh import /internal/http & internal/grpc
                entity/ <- entity struct model untuk handle struct response repository
                interfaces/ <- interfaces untuk repository
                create_user <- handle/domain untuk service user (create dalam case)
            infrastructure/ <- boleh import repository, api_client, tidak boleh import /internal/http & internal/grpc
                repository/ <- repository database connection 
        common/ <- boleh import repository, api_client, tidak boleh import /internal/http & internal/grpc
            constants/ <- menyimpan constant variable yang digunakan secara global
            handler/ <- untuk handler response payload (success / gagal)
            infrastructure/ <- connection db / web (client)
            interfaces/ <- interfaces application / service route
            middleware/ <- middleware handle
            utility/ <- utility constants global function
```

# Context Parameter

## Syntax



## Kapan menggunakan context parameter

Selalu gunakan `context.Context` untuk fungsi yang berinteraksi dengan system
external.

* query ke postgres / mongodb ? gunakan `context.Context`
* call API dari 3rd party ? gunakan `context.Context`
* dan lain2

> Context. Untuk. Semua. External. Call.

## Hati-hati menggunakan hardcoded timeout context

Perhatikan contoh berikut :

```go
func UpdateSomething(ctx context.Context, param string) error {
  // add span to ctx...
  span, ctx := apm.StartSpan(ctx, "myspan", spanType)
  defer span.End()

  rCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
  defer cancel()

  // continue function...
}
```

`context.WithTimeout` akan membuat copy dari context yg sudah ada, dengan waktu timeout tertentu.
Line diatas berarti untuk 1 blok fungsi tsb, timeout-nya 10 detik.

Hal ini menjadi masalah jika timeout-nya di create di tiap2 fungsi. Misal:
* `controller.UpdatePhone -> service.UpdatePhone -> mongo.UpdatePhone`
* Jika tiap2 fungsi tersebut timeout 10s, artinya total request memungkinkan untuk berjalan selama 30s.
* Padahal timeout dari user / android maksimal 10s
* Ada kemungkinan di hp user timeout, tetapi di server program-nya masih di execute / berlangsung.

Idealnya timeout di set dari sisi echo / grpc server middleware, jadi cukup 1x. Selanjutnya fungsi2 yg membutuhkan context tinggal ambil dari parameter fungsi sebelumnya.

Jika sudah ada parameter `context.Context` dari fungsi, tidak perlu membuat context baru dengan timeout yang berbeda.

# Deploy

PM2 dapat mengelola beberapa aplikasi kita dan memantaunya dengan baik dengan perintah - perintah yang mempermudah hidup kita. Tidak hanya itu, kita pun dapat menggunakan REST API yang disediakan PM2 untuk mengintegrasikan monitoring ke aplikasi web kita yang lain atau aplikasi mobile yang kita kembangkan untuk mengawasi aplikasi Node.js kita. Beberapa fitur utama yang disediakan PM2 antara lain:

* Konfigurasi tingkah laku (behavior)
* Kompatibel dengan PaaS
* Watch dan reload
* Manajemen log
* Monitoring
* Module system
* Max memory reload
* Cluster mode
* Hot reload
* Integrasi dengan Keymetrics
* Development dan Deployment workflow

deploy using PM2
```
$ npm install express
$ npm install pm2
```

build application to pm2
```
go build -o application cmd/main.go 
```

run application to pm2
```
pm2 start application
```

## Referensi

* https://golang.org/pkg/context/#WithTimeout
* [golang context guide]( https://golangbyexample.com/using-context-in-golang-complete-guide/ )
* [Go Module](https://blog.golang.org/using-go-modules)
* [Go Protobuff](https://developers.google.com/protocol-buffers/docs/gotutorial)

* [CRUD-y By Design](https://github.com/adamwathan/laracon2017)
