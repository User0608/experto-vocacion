## Lista de preguntas CASM
GET: http://localhost:90/test/:test_id/casm?items=<num_items>&page=<num_pagina>
Recupera la lista de preguntas segun el numero de items y pagina
Ejemplo de un elemento del array de preguntas entregados como respuesta:

```json
        {
            "id": 1,
            "question_a": "",
            "question_b": "",
            "answer_a": false,
            "answer_b": false,
            "done": true
        }
```
La pregunta CASM devuelve el estado done el cual indica si, esta desarrollada o no.

## Lista de preguntas CASM
GET: http://localhost:90/test/:test_id/berger?items=<num_items>&page=<num_pagina>

```json
        {
            "id": 1,
            "question_a": "...",
            "question_b": "...",
            "answer": 9,
            "done": true
        },
```

## Lista de preguntas HEA
GET: http://localhost:90/test/:test_id/berger?items=<num_items>&page=<num_pagina>

```json
        {
            "id": 1,
            "question": "...",
            "answer": "S",
            "done": true
        }
```

 
## Registro de respuestas CASM
Inviar un arreglo no nullo, siguiendo la estructura del ejemplo.
Los datos o preguntas que ya esten respondidas seran omitidas!
POST: http://localhost:90/test/casm
```json
    [
        {
            "test_id":1,
            "casm_id":1,
            "answer_a":false,
            "answer_b":true
        },
        {
            "test_id":1,
            "casm_id":2,
            "answer_a":false,
            "answer_b":true
        },
    ]
```
Response: `num_created`, sera la candid de registro que fueron aceptados y creado!
por el contrario `num_omitted` son los registros que no se tomaron en cuenta, porque ya estan registrados, finalmente, `created` son los registros aceptados y registrados.

```json
    {
        "code": "OK",
        "data": {
            "num_created": 1,
            "num_omitted": 3,
            "created": [
                {
                    "test_id": 1,
                    "casm_id": 4,
                    "answer_a": true,
                    "answer_b": false
                }
            ]
        }
    }
```