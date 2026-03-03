# MiniPOS Payment Gateway API

Bu proje, POS (Nokta Satış) cihazları ve e-ticaret sistemleri için tasarlanmış, Go dili ile geliştirilmiş bir backend (ara katman) ödeme geçidi simülasyonudur. Proje, finansal işlemlerdeki en kritik gereksinim olan **Veri Tutarlılığını (Data Consistency)** sağlamak amacıyla, veritabanı "Transaction" yönetimini merkeze alarak tasarlanmıştır.

## 🚀 Kullanılan Teknolojiler

* **Dil:** Go (Golang)
* **Web Framework:** Gin HTTP Framework
* **Veritabanı:** PostgreSQL (Docker üzerinden)
* **ORM:** GORM (AutoMigrate destekli)
* **Mimari:** Clean Architecture (Handler -> Repository -> Database)

## 💡 Mimari Yaklaşım ve Tasarım Kararları

Projede, kodun sürdürülebilirliğini ve test edilebilirliğini artırmak için **Katmanlı Mimari (Layered Architecture)** kullanılmıştır:
* **Models:** Veritabanı tablolarının Go `struct` karşılıkları.
* **Repository:** Veritabanı ile konuşan tek katman. SQL sorguları ve `Transaction` (Rollback/Commit) mantığı burada izole edilmiştir.
* **Handlers (Controllers):** HTTP isteklerini karşılayan, veriyi doğrulayan ve iş mantığını Repository'e devreden katman.

**Fintech Odaklı Geliştirme:** Ödeme (Payment) ve İade (Refund) işlemlerinde ACID prensiplerine uyulmuştur. Bir işlem sırasında mağaza bakiyesi güncellenirken bir hata oluşursa, sistem otomatik olarak `Rollback` yaparak para kaybını/sızdırmasını engeller.

## ⚙️ Kurulum ve Çalıştırma

Projeyi lokal ortamınızda çalıştırmak için aşağıdaki adımları izleyebilirsiniz.

### 1. Veritabanını Ayağa Kaldırma (Docker)
PostgreSQL veritabanını hızlıca başlatmak için terminalde şu komutu çalıştırın:
`docker run --name minipos-db -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=gizlisifrem -e POSTGRES_DB=minipos -p 5432:5432 -d postgres:latest`

### 2. Projeyi Başlatma
Gerekli Go paketlerini indirdikten sonra sunucuyu başlatın:
`go mod tidy`
`go run main.go`

*Sunucu varsayılan olarak `http://localhost:8080` adresinde çalışacaktır. GORM, tabloları otomatik olarak oluşturacaktır.*

## 📡 API Uç Noktaları (Endpoints)

### Üye İşyeri (Merchant) Yönetimi
| Metot | Endpoint | Açıklama |
| :--- | :--- | :--- |
| `GET` | `/health` | API'nin ayakta olup olmadığını kontrol eder. |
| `POST` | `/merchants` | Yeni bir mağaza ekler. |
| `GET` | `/merchants` | Sistemdeki tüm aktif/pasif mağazaları listeler. |
| `GET` | `/merchants/:id` | ID'si verilen tek bir mağazanın detaylarını ve anlık bakiyesini döner. |

**Örnek Mağaza Ekleme İsteği (POST /merchants):**
```json
{
    "name": "Test Market",
    "api_key": "gizli_anahtar_123",
    "balance": 1500.50
}
