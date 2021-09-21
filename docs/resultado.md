# resultado
El resultado solo estara disponible si todo las test an sido terminados.

GET: http://localhost:90/test/1/resultado

JSON:
```json
    {
        "code": "OK",
        "data": {
            "test_id": 1,
            "resultado_casm": [
                {
                    "per": 80,
                    "res": "LING"
                },
                {
                    "per": 95,
                    "res": "JURI"
                }
            ],
            "resultado_berger": {
                "emotivo": "EMOTIVO",
                "activo": "ACTIVO",
                "orden": "PRIMARIO"
            },
            "resultado_berger_final": "COLERICO",
            "resultado_hea": [
                "USO DEL TIEMPO",
                "AYUDAS PARA EL ESTUDIO"
            ],
            "done": true,
            "resultado": [
                "LINGUISTA",
                "DERECHO PENAL"
            ],
            "created_at": "2021-09-18T17:04:04.691357Z"
        }
    }
```
## Detalle CASM
CCFM: "Ciencias físicas matemáticas",
CCSS: "Ciencias sociales",
CCNA: "Ciencias naturales",
CCCO: "Ciencias de la comunicación",
ARTE: "Artes",
BURO: "Burocracia",
CCLP: "Ciencias económicas políticas",
IIAA: "Fuerzas armadas de Perú",
FINA: "Finanzas",
LING: "Lingüística",
JURI: "Jurisprudencia",