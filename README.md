# Medpoint System - Backend

## Introduction

**Medpoint System** adalah sebuah platform yang memungkinkan pengguna untuk melakukan berbagai jenis reservasi medis online, termasuk konsultasi dokter, tes laboratorium, tindakan medis, dan vaksinasi. Sistem ini memudahkan pengguna dengan proses reservasi yang lebih efisien sehingga tidak perlu melakukan pendaftaran ulang secara offline saat menjalani tindakan medis.

### Fitur Utama

- **Reservasi Medis Online:** Pengguna dapat melakukan reservasi untuk konsultasi dokter, tes laboratorium, tindakan medis, dan vaksinasi.
- **Pencarian Dokter dan Fasilitas Kesehatan:** Memudahkan pengguna untuk menemukan dokter dan fasilitas kesehatan terdekat.
- **Penjadwalan Janji Temu:** Pengguna dapat menjadwalkan janji temu dengan dokter atau tindakan medis.

## Deskripsi Sistem (Backend)

**Backend** dari **Medpoint System** dikembangkan menggunakan bahasa pemrograman **Go** dan framework **Raiden**. Backend ini bertanggung jawab untuk menangani logika bisnis, pengelolaan data, dan interaksi dengan database.

## Teknologi yang Digunakan

- **Bahasa Pemrograman:** Go (Golang)
- **Framework:** Raiden
- **Database:** PostgreSql + Supabase

## Dokumentasi Raiden

Untuk informasi lebih lanjut tentang framework Raiden, dapat mengunjungi [Dokumentasi Raiden](https://raiden.sev-2.com/).

## Folder Structure

```javascript
├── configs
│   ├── app.yaml          # App configuration file
│   └── modules/          # Module configuration file
│       ├── module_a.yaml
│       ├── module_b.yaml
│       └── ...
├── cmd
│   └── project-name
│       └── project_name.go    # Main project
│   └── apply/main.go
│   └── import/main.go
├── internal
│   ├── bootstrap
│   │   ├── route.go
│   │   ├── rpc.go
│   │   ├── roles.go
│   │   ├── models.go
│   │   └── storages.go
│   ├── controllers
│   │   └── hello.go    # Example controller
│   ├── roles           # ACL/RLS definition
│   │   ├── members.go
│   │   └── ...
│   ├── models          # All model will auto-sync
│   │   ├── users.go
│   │   └── ...
│   ├── rpc             # RPC
│   │   └── get_user.go
│   ├── topics          # Real-time
│   │   ├── notification.go
│   │   └── inbox.go
│   │
│   └── storages
│       └── avatar.go
├── build
│   └── state      # Bytecode containing raiden state
├── go.sum
└── go.mod
```
