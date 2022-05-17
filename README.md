# JDS-Test
Screening test Jabar Digital Service: registration, login, token validation endpoint

## API SPECIFICATION

**ENDPOINT:**
http://localhost:8080/api/v1/

**METHOD:**
|Method         |Fungsi         |
|---------------|---------------|
|users          |Registrasi user|
|sessions       |Login user     |
|users/fetch    |Validasi token |

## 1.1 Method users
| <!-- -->      | <!-- -->                           |
|---------------|------------------------------------|
|Method         |users                               |
|Type           |POST                                |
|URL            |http://localhost:8080/api/v1/users  |
|Fungsi         |Registrasi user                     |

**PARAMETER**
|Input Parameter|Type   |Deskripsi                                    |
|---------------|-------|---------------------------------------------|
|username       |String |Username dari user yang diregistrasikan      |
|role           |String |Role/hak akses dari user yang diregistrasikan|

**RETURN**
|Return         |Type   |Deskripsi                                                          |
|---------------|-------|-------------------------------------------------------------------|
|username       |String |Username dari user yang diregistrasikan                            |
|role           |String |Role/hak akses dari user yang diregistrasikan                      |
|password       |String |Password hasil generate sistem dari user yang diregistrasikan      |

**Contoh Format JSON (Input Parameter)**

```javascript
{
    "username": "user-test",
    "role":"user"
}
```

**Contoh Format JSON (Return)**
**Sukses:**
```javascript
{
    "meta": {
        "message": "Account has been registered",
        "code": 200,
        "status": "success"
    },
    "data": {
        "username": "user-test",
        "role": "user",
        "password": "RJ9U27"
    }
}
```

**Gagal/Error**
```javascript
{
    "meta": {
        "message": "Register account failed",
        "code": 400,
        "status": "error"
    },
    "data": {
        "errors": "Username has been registered!"
    }
}
```

## 1.2 Method sessions
| <!-- -->      | <!-- -->                              |
|---------------|---------------------------------------|
|Method         |users                                  |
|Type           |POST                                   |
|URL            |http://localhost:8080/api/v1/sessions  |
|Fungsi         |Login user                             |

**PARAMETER**
|Input Parameter|Type   |Deskripsi                                        |
|---------------|-------|-------------------------------------------------|
|username       |String |Username dari user yang sudah terdaftar dis sitem|
|password       |String |Password dari user yang sudah terdaftar di sistem|

**RETURN**
|Return         |Type   |Deskripsi                                                          |
|---------------|-------|-------------------------------------------------------------------|
|username       |String |Username dari user yang telah berhasil login                       |
|role           |String |Role/hak akses dari user yang telah berhasil login                 |
|token          |String |Token yang digenerate dari proses login                            |

**Contoh Format JSON (Input Parameter)**

```javascript
{
    "username": "user-test",
    "password": "RJ9U27"
}
```

**Contoh Format JSON (Return)**
**Sukses:**
```javascript
{
    "meta": {
        "message": "Successfuly loggedin",
        "code": 200,
        "status": "success"
    },
    "data": {
        "username": "user-test",
        "role": "user",
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTI3NzEyMDYsImlzX3ZhbGlkIjp0cnVlLCJ1c2VyX2lkIjoxMH0.ZaNuh-d-w6e2bjOSd6OpBAoNKgs7HMXtFNTkBY2_eFM"
    }
}
```

**Gagal/Error**
```javascript
{
    "meta": {
        "message": "Login failed!",
        "code": 400,
        "status": "error"
    },
    "data": {
        "errors": "crypto/bcrypt: hashedPassword is not the hash of the given password"
    }
}
```

## 1.3 Method users/fetch
| <!-- -->      | <!-- -->                                           |
|---------------|----------------------------------------------------|
|Method         |users/fetch                                         |
|Type           |GET                                                 |
|URL            |http://localhost:8080/api/v1/users/fetch            |
|Header         |"Bearer" + token yang didapat dari proses login user|
|Fungsi         |Validasi token                                      |

**RETURN**
|Return         |Type   |Deskripsi                                                          |
|---------------|-------|-------------------------------------------------------------------|
|is_valid       |Bool   |Status token yang divalidasi                                       |
|expired_at     |Date   |Waktu kadaluarsa token                                             |
|username       |String |Username dari user yang login menggunakan token yang divalidasi    |

**Contoh Format JSON (Return)**
**Sukses:**
```javascript
{
    "meta": {
        "message": "Token is validate",
        "code": 200,
        "status": "success"
    },
    "data": {
        "is_valid": true,
        "expired_at": "2022-05-17 14:19:40",
        "username": "user-test"
    }
}
```

**Gagal/Error**
```javascript
{
    "meta": {
        "message": "Unauthorized",
        "code": 401,
        "status": "error"
    },
    "data": null
}
```
