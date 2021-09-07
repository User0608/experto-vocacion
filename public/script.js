const form = document.getElementById("mi_form")
const responseBox = document.getElementById("response_box")

const fetchData = async (url, data) => {
    const res = await fetch(url, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)
    })
    const jData = await res.json()
    return jData
}
const processData = async (data) => {
    const res = await fetchData(`${location.origin}/experto`, data)
    if (res.code === "OK") {
        responseBox?.classList.remove('d-none')
        responseBox.innerHTML = `
        <h2>Respuesta:</h2>
            <p>${res.respuesta}</p>
        `
    }
}
form?.addEventListener('submit', (e) => {
    e.preventDefault()
    let area = form.area.value
    let habito = form.habito.value
    let caracter = form.caracter.value
    if (area === "" || habito === "" || caracter === "") {
        alert("No se a ingresado datos.")
        return
    }
    processData({ area, habito, caracter })
})

const inputs = document.getElementsByTagName("select")
for (let i = 0; i < inputs.length; i++) {
    inputs[i].addEventListener('change', () => {
        responseBox?.classList.add('d-none')
    })
}