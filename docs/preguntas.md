## Lista de preguntas CASM
GET: http://localhost:90/test/:test_id/casm?items=<numero de items>&page=<numero de pagina>
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