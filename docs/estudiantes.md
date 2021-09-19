# Operaciones con estudiantes

## Create
Para registrar un usuario, no es necesario contar con token,
el path es publico. de forma que cualquiera pueda registrares.
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
:Permisos de administradores son necesarios para esta operacion.
Devuelve todos es estudiantes registrados

http://localhost:90/estudiante

## ByID
:Permisos de administradores son necesarios para esta operacion.
http://localhost:90/estudiante/:estudiante_id

## SessionInfo
Devuelve la informacion del estudiante con session actual en el sistema.
GET: http://localhost:90/estudiante/me
