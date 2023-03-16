# BOOK RENTAL SHOP

### **ความต้องการระบบ**

1.  สามารถจัดการข้อมูลพื้นฐานของระบบได้
2.  สามารถให้ยืม คืนหนังสือได้ (ด้วยระบบบาร์โค๊ด)
3.  สามารถกำหนดอัตราค่าเช่า จำนวนวันให้เช่า และค่าปรับได้
4.  สามารถแสดงรายการข้อมูลหนังสือค้างส่ง หนังสือที่มีการเช่า หรือคืนได้
5.  สามารถลงทะเบียนหนังสือใหม่ได้
6.  สามารถจัดการข้อมูลหนังสือได้
7.  สามารถสรุปรายรับ รายจ่ายประจำวัน เดือน ปีตามลำดับได้
8.  สามารถจัดการบัญชีรายรับรายจ่ายได้
9.  สามารถแสดงรายการหนังสือ ที่มีความนิยมในขณะช่วงเวลานั้น ๆ ได้
10. สามารถทำการจองหนังสือผ่านออนไลน์ได้
11. สามารถออกรายงานได้

- 11.1 รายงานหนังสือค้างส่ง
- 11.2 รายงานหนังสือยอดนิยม
- 11.3 รายงานงานรายจ่าย
- 11.4 รายงานรายรับ

### **Development tool**

- Frontend
  - React
- Backend
  - Go
    - Fiber
- Utilities
  - PostgresSQL
  - Vscode

## ER DIAGRAM

![ER](/ERV1.PNG)

## API ENDPOINT

### Category Service

| Method | URI           | Description     |
| ------ | ------------- | --------------- |
| POST   | /category     | Create category |
| GET    | /category     | Get category    |
| PUT    | /category     | Update category |
| DELETE | /category/:id | Delete category |

### POST /category

Request

```json
{
  "name": "นิยายรัก"
}
```

Response

```json
{
  "message": "Add category successfully."
}
```

### GET /category

Response

```json
[
  {
    "id": "0db-zxs-dx",
    "name": "นิยายรัก"
  },
  {
    "id": "0db-zxs-dxz",
    "name": "นิยายสิบสวน"
  }
]
```

### PUT /category

Request

```json
{
  "id": "0db-zxs-dxz",
  "name": "นิยายรัก"
}
```

Response

```json
{
  "message": "Update category successfully."
}
```

### DELETE /category/:id

Parameter

- id CategoryID

Response

```json
{
  "message": "Delete category successfully."
}
```

---

### User Service

| Method | URI        | Description    |
| ------ | ---------- | -------------- |
| POST   | /users     | Create user    |
| GET    | /users     | Get user       |
| PUT    | /users     | Update user    |
| GET    | /users/:id | Get user by id |

### POST /users

Request

```json
{
  "name": "นิยายรัก"
}
```

Response

```json
{
  "message": "Add category successfully."
}
```

### GET /users

Response

```json
[
  {
    "id": "0db-zxs-dx",
    "name": "นิยายรัก"
  },
  {
    "id": "0db-zxs-dx",
    "name": "นิยายสิบสวน"
  }
]
```
