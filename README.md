# TO-DO App (Backend – Go + Fiber)

Bu proje, backend staj başvurusu kapsamında geliştirilmiş bir RESTful TO-DO uygulamasıdır. Uygulama Go dilinde [Fiber](https://gofiber.io) framework'ü kullanılarak yazılmıştır. Kullanıcılar yapılacaklar listeleri oluşturabilir, bu listelere adımlar ekleyebilir ve tamamlayarak ilerleyebilirler.

---

## Teknolojiler

- Go (Golang)
- Fiber v2
- JWT (kimlik doğrulama)
- UUID
- Mock JSON veri yapısı (veritabanı yerine)

---

## Kullanıcı Rolleri

| Rol    | Açıklama                              |
|--------|----------------------------------------|
| `user` | Sadece kendi TO-DO verilerini görür    |
| `admin`| Tüm kullanıcıların TO-DO verilerini görür |

### Hazır Kullanıcılar

```json
{
  "username": "normaluser",
  "password": "123456",
  "role": "user"
},
{
  "username": "adminuser",
  "password": "admin123",
  "role": "admin"
}
```

---

## Kimlik Doğrulama

Tüm korumalı route’lara erişim için `Authorization: Bearer <JWT Token>` başlığı kullanılmalıdır.

### Giriş için:
**POST** `/api/auth/login`

#### Örnek Body:

```json
{
  "username": "normaluser",
  "password": "123456"
}
```

---

## API Endpoint'leri

### TO-DO Listeleri

| Metot | URL                | Açıklama                                         |
|-------|--------------------|--------------------------------------------------|
| GET   | `/api/todos/mine`  | Kullanıcının kendi TO-DO listelerini getirir    |
| GET   | `/api/todos/all`   | Admin kullanıcı için tüm TO-DO listeleri getirir|
| POST  | `/api/todos`       | Yeni TO-DO listesi oluşturur                    |
| DELETE| `/api/todos/:id`   | TO-DO listesini siler (soft delete)             |

---

### TO-DO Adımları

| Metot | URL                                  | Açıklama                                        |
|-------|---------------------------------------|-------------------------------------------------|
| GET   | `/api/steps/:todoId`                 | Belirli bir TO-DO listesine ait adımları getirir|
| POST  | `/api/steps/:todoId`                 | Yeni bir adım oluşturur                         |
| PUT   | `/api/steps/complete/:stepId`        | Adımı tamamlandı olarak işaretler               |
| DELETE| `/api/steps/:stepId`                 | Adımı siler (soft delete)                       |

---

## Projeyi Çalıştırma

```bash
git clone https://github.com/kullanici-adin/todo-app-backend.git
cd todo-app-backend
go mod tidy
go run main.go
```

---

## Notlar

- Veritabanı yerine `mockdata/` klasöründeki `users.json` ve `todos.json` dosyaları kullanılır.
- Silme işlemleri kalıcı değildir. `DeletedAt` alanı atanarak veri listelerde görünmez hale getirilir.
- Güncellemelerde `UpdatedAt` alanı otomatik olarak güncellenir.
- Projede Clean Architecture ve MVC prensiplerine uygun modüler yapı kullanılmıştır.

---