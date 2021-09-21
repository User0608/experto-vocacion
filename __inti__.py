from flask import Flask, request
from flask_restful import Resource, Api
from core import PrologMT

app = Flask(__name__)
api = Api(app)

prolog = PrologMT()
prolog.consult(filename="zexperto/reglas.pl")


def consultar(area, habito, caracter):
    try:
        prolog.retractall("area(_)")
        prolog.retractall("habito(_)")
        prolog.retractall("caracter(_)")
    except:
        print("paso error")
    prolog.assertz(f"area('{area}')")
    prolog.assertz(f"habito('{habito}')")
    prolog.assertz(f"caracter('{caracter}')")
    result = list(prolog.query("vocacion(Y)"))
    if len(result) > 0:
        return result[0]["Y"]
    else:
        return "RESPUESTA NO ENCONTRADA"


def consultarBerger(emotivo, activo, orden):
    try:
        prolog.retractall("emotivo(_)")
        prolog.retractall("activo(_)")
        prolog.retractall("orden(_)")
    except:
        print("paso error")
    prolog.assertz(f"emotivo('{emotivo}')")
    prolog.assertz(f"activo('{activo}')")
    prolog.assertz(f"orden('{orden}')")
    result = list(prolog.query("personalidad(Y)"))
    if len(result) > 0:
        return result[0]["Y"]
    else:
        return "RESPUESTA NO ENCONTRADA"


class Experto(Resource):
    def post(self):
        json_data = request.get_json(force=True)
        try:
            area = json_data['area']
            habito = json_data['habito']
            caracter = json_data['caracter']
        except:
            return {"code": "Error", "message": "Datos inconsistentes"}
        try:
            res = consultar(area, habito, caracter)
        except:
            return {"code": "IN_ERR", "message": "Algo paso"}
        return {"code": "OK", "respuesta": f"{res}"}


class ExpertoBerger(Resource):
    def post(self):
        json_data = request.get_json(force=True)
        try:
            emotivo = json_data['emotivo']
            activo = json_data['activo']
            orden = json_data['orden']
        except:
            return {"code": "Error", "message": "Datos inconsistentes"}
        try:
            res = consultarBerger(emotivo, activo, orden)
        except:
            return {"code": "IN_ERR", "message": "Algo paso"}
        return {"code": "OK", "respuesta": f"{res}"}


api.add_resource(Experto, '/experto')
api.add_resource(ExpertoBerger, '/berger')

if __name__ == '__main__':
    app.run(port=81, debug=False)
