# Operaciones con estudiantes

## Create
POST: http://localhost:90/estudiante 
```json
    {
        "nombre": "<name>",
        "apellido_paterno": "<string>",
        "apellido_materno": "<string>",
        "fecha_nacimiento": "1996-04-18T00:00:00Z",
        "dni": "<len = 8>",
        "password":"<min len 6>"
    }
```
Respuesta:
```json
    {
        "code": "OK",
        "data": {
            "estudiante_id": 2,
            "nombre": "<string>",
            "apellido_paterno": "<string>",
            "apellido_materno": "<string>",
            "fecha_nacimiento": "1996-04-18T00:00:00Z",
            "dni": "string"
        }
    }
```

## Update
PUT: http://localhost:90/estudiante 
```json
    {
        "estudiante_id":1,
        "nombre": "<name>",
        "apellido_paterno": "<string>",
        "apellido_materno": "<string>",
        "fecha_nacimiento": "1996-04-18T00:00:00Z",
        "dni": "<len = 8>",
        "password":"<min len 6>"
    }
```
Respuesta:
```json
    {
        "code": "OK",
        "data": {
            "estudiante_id": 2,
            "nombre": "<string>",
            "apellido_paterno": "<string>",
            "apellido_materno": "<string>",
            "fecha_nacimiento": "1996-04-18T00:00:00Z",
            "dni": "string"
        }
    }
```
## Listar

## ByID

## SessionInfo
