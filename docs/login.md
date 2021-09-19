## Usuario admin
POST: http://localhost:90/login/admin
```json
    {
    "username":"<username>",
    "password":"<password>"
}
```
Response:
```json
    {
        "code": "OK",
        "token": "<token>"
    }
```
## Login Estudiante
POST: http://localhost:90/login
```json
{
    "username":"<dni>",
    "password":"<password>"
}
```
Response:
```json
{
    "code": "OK",
    "usuario": {
        "estudiante_id": 1,
        "nombre": "<name>",
        "apellido_paterno": "<apellido_paterno>",
        "apellido_materno": "<apellido_materno>",           
        "fecha_nacimiento": "1998-06-08T00:00:00Z",
        "dni": "<dni>"    
    },
    "token": "<token>"
}
```