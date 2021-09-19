## Crear Test
Para crear un test, solo se necesitar realizar una peticion `GET`:
Esta peticion tiene que ser realizada por el rol estudiante.

GET: http://localhost:90/prueba/create
Response:
```json
    {
        "code": "OK",
        "data": {
            "test_id": 7,
            "resultado_casm": "",
            "resultado_berger": "",
            "resultado_lea": "",
            "done": false,
            "created_at": "2021-09-19T15:15:13-05:00"
        }
    }
```
Los campos faltantes seran actualizados despues de terminar el test.

## Cargar Todo los test
En esta caso, si la peticion es hecha por un admin, se devuelve todo, 
de lo contrario, solo lo correspondiente al estudiante.

GET: http://localhost:90/prueba


## Get By ID
GET: http://localhost:90/prueba/:test_id

## Eliminar
Solo se puede eliminar, los test que aun no an sido iniciados.
DELETE: http://localhost:90/prueba/:test_id