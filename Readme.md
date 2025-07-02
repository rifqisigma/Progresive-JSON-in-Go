# Progressive JSON Streaming in Go

## Apa itu Progressive JSON?

**Progressive JSON** adalah teknik mengirim data JSON secara **bertahap (streaming)** daripada langsung semua sekaligus. Ini sangat berguna saat:
- Dataset besar
- Client ingin mulai konsumsi data secepat mungkin
- Respons time harus cepat dan real-time

Contoh umum:
- API real-time feed
- Streaming log data
- Partial response untuk pengalaman pengguna lebih responsif

---

## Implementasi di Go (Tanpa Library Tambahan)

Go punya support **native streaming** lewat kombinasi:
- `json.Encoder`
- `http.ResponseWriter`
- `Flush()` dari `http.Flusher`


## Cara Build & Jalankan

1. jalankan endpoint di cmd

     ```bash
     curl -N http://localhost:8080/stream
