# Technical Test - Application Engineer [Muhammad Yusuf]

Proyek ini adalah implementasi dari soal Technical Test untuk posisi Application Engineer di Pt Praweda

```javascript
Kandidat = {
  Nama: "Muhammad Yusuf",
  Email: "yusufsmd58@gmail.com",
};
```

## 🚀 Instalasi dan Setup

1. Clone repository ini

```
git clone https://github.com/M16Yusuf/test-praweda-app-eng.git
```

2. buat/tambahkan env di backend `backend/.env`

```
DBUSER=<your_user>
DBPASS=<your_pass>
DBNAME=<your_db>
DBHOST=<your_host>
DBPORT=<your_port>


JWT_SECRET=<your_jwt>
```

3. pindah ke backend, jalankan go mod tidy, dan mgirate database

```sh
cd backend
go mod tidy
```

4. Migrate database

```sh
migrate -database YOUR_DATABASE_URL -path ./db/migrations up

# jika menggunakan makefile
make migrate-createUp
```

5. Jalankan program backend

```sh
go run ./cmd/main.go
```

6. Jalankan program frontend

```sh
cd .. # make sure berada di root repository
cd frontend # masuk ke dalam frontend
npm install # install dependencies
npm run dev # jalankan frontend
```

## 📋 Implementasi Tugas

### 1️⃣ . Membuat API ecommerce menghitung diskon

| Method | Endpoint                        | Body (Input)                                   | Response (Output Data)                                                                                                                                                                         |
| ------ | ------------------------------- | ---------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| POST   | 127.0.0.1:8080/ecommerce/diskon | {"harga_awal": 5000000, "pake_voucher": true}" | {"is_success": true, "code": 200, "message": "Anda menggunakan voucher..", "data": {"harga_awal": 5000000, "pake_voucher": true, "harga_setelah_diskon": 2500000, "points_didapatkan": 50000}} |

### 2️⃣ . Membuat API Oauth JWT dan Cookies

| Method | Path                       | body                    | Deskripsi                                                                        | Authentication |
| ------ | -------------------------- | ----------------------- | -------------------------------------------------------------------------------- | -------------- |
| POST   | 127.0.0.1:8080/auth/login  | { "username": "yusuf" } | Membuat sesi baru dan mengautentikasi pengguna (membuat cookies dengan isi jwt). | None           |
| GET    | 127.0.0.1:8080/auth/me     | { "username": "yusuf" } | Memverifikasi sesi mencocokan username dengan key cookies.                       | Session Cookie |
| DELETE | 127.0.0.1:8080/auth/logout | { "username": "yusuf" } | Mengakhiri sesi pengguna yang aktif (menghapus key cookies).                     | Session Cookie |

### 3️⃣ . Membuat query join 2 table berdasarkan soal

![image_no3](/assets/soal_no3.png)

jawaban:

```sql
SELECT c.User_id, c.Id AS Company_id, u.Nama, u.Email, u.Telp, c.Company_code, c.Company_name
FROM 'User' u
RIGHT JOIN 'Company' c ON c.User_id = u.Id;

```

### 4️⃣ . Membuat fetch internal di backend, manipulasi responsenya, dan jadikan hasilnya menjadi API sendiri yang dinamis

API yang dimanupalsi :

```
https://randomuser.me/api?results=10&page=1
```

Hasil dari pengerjaan manipulasi API:

```
http://127.0.0.1:8080/manipulasi?page=1&results=10
```

**Request**
Method: GET

URL: http://127.0.0.1:8080/manipulas

Query Parameters:

| Parameter | Required | Type    | Default | Description                                         |
| --------- | -------- | ------- | ------- | --------------------------------------------------- |
| page      | Optional | Integer | 1       | Nomor halaman yang ingin diambil.                   |
| results   | Optional | Integer | 10      | Jumlah item maksimum yang dikembalikan per halaman. |

Lalu implementasikan API diatas dengan frontend yang diminta menggunakan react, typescript, & ant Design

Hasil yang didapatkan :

![image_demo](/assets/demo-frontend.png)

### 5️⃣ . Manipulasi Array

Manipulasi array, menjadi dinamis ketika array `warna` bertambah datanya

```typescript
const warna: string[] = ["merah", "kuning", "hijau", "pink", "ungu"];
const pakaian: string[] = ["baju", "celana", "topi", "jaket", "sepatu"];
const statusDiskon: string[] = ["Diskon", "Sale", "Diskon", "Sale", "Sale"];

const manipulasiArray: string[] = [];

warna.push("maroon");

for (let i = 0; i < warna.length; i++) {
  const pakaianItem = pakaian[i] || pakaian[pakaian.length - i];
  const diskonItem = statusDiskon[i] || statusDiskon[statusDiskon.length - i];

  manipulasiArray.push(`${warna[i]} ${pakaianItem} ${diskonItem}`);
}

console.log(manipulasiArray);

export default manipulasiArray;
```
