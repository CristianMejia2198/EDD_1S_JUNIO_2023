import React, {useState, useEffect} from 'react';
import '../css/administrador.css'
import 'bootstrap/dist/css/bootstrap.min.css'

export const Filtros = () => {
    const [filtro, setFiltro] = useState(0)
    const salir = (e) => {
        e.preventDefault();
        console.log("Listo")
        window.open("/empleado","_self")
    }

    const validar = (data) =>{
        console.log(data)
    }

    const aplicarFiltros = async(e) => {
        e.preventDefault();
        fetch('http://localhost:3001/aplicarfiltro',{
            method: 'POST',
            body: JSON.stringify({
                Tipo: filtro
            }),
            headers:{
                'Content-Type': 'application/json'
            }
        })
        .then(response => response.json())
        .then(data => validar(data))
    }

    const handleChange = (e) => {
        var j = parseInt(e.target.value);
        setFiltro(j)
    }
 

    return(
        <div className="form-signin1">
            <div className="text-center">
                  <form className="card card-body">
                    <h1 className="h3 mb-3 fw-normal">Dashboard Empleado {localStorage.getItem("empleado")}</h1>
                    <br/>
                    <h4 className="h3 mb-3 fw-normal">Elige un Filtro</h4>
                    <br/>
                    <div className="col align-self-center">
                        <select className="form-control" aria-label=".form-select-lg example" onChange={handleChange}>
                            <option value={0}>Elegir....</option>
                            <option value={1}>Negativo</option>
                            <option value={2}>Escala de Grises</option>
                            <option value={3}>Espejo X</option>
                            <option value={4}>Espejo Y</option>
                            <option value={5}>Ambos Espejos</option>
                        </select>
                    </div>
                    <br/>
                    <center><button className="w-50 btn btn-outline-primary" onClick={aplicarFiltros}>Generar Imagen con Filtro</button></center>
                    <br/>
                    <center><button className="w-50 btn btn-outline-success" onClick={salir}>Salir</button></center>
                    <br/>
                    <p className="mt-5 mb-3 text-muted">EDD 201700918</p>
                    <br/>
                  </form>
            </div>
          </div>
    );
}