import React, {useState, useEffect} from 'react';
import '../css/maqueta.css'
import 'bootstrap/dist/css/bootstrap.min.css'
import swal from 'sweetalert'
import axios from 'axios';

export const Matriz = () =>{
    const [empleado, setEmpleado] = useState('Empleado ' + window.localStorage.getItem("Empleado"))
    const [cliente, setCliente] = useState('')
    const [imagenCliente, setImagenCliente] = useState('')
    const [filtros, setFiltros] = useState('')
    const [posFiltro, setPosFiltro] = useState(0)
    const [listadoFiltros, setListadoFiltros] = useState(["Negativo", "Escala de grises", "Espejo en X", "Espejo en Y", "Ambos Espejos"])

    const [imagen, setImagen] = useState('https://yakurefu.com/wp-content/uploads/2020/02/Chi_by_wallabby.jpg')
    const handleSubmit = (e) => {
        e.preventDefault();
        fetch('http://localhost:3001/agregar-matriz',{
            method: 'POST',
            body: JSON.stringify({
                Padre: empleado,
                Cliente: cliente,
                Imagen: imagenCliente,
                Filtros: filtros
            }),
            headers:{
                'Access-Control-Allow-Origin': '*',
                'Content-Type': 'application/json'
            }
        })
        .then(response => response.json())
        .then(data => validar(data));
    }

    const validar = (data) =>{
        console.log(data)
        //setImagen(data.Imagenbase64)
    }

    const pedirReporte = (e) => {
        e.preventDefault();
        fetch('http://localhost:3001/',{
        })
        .then(response => response.json())
        .then(data => validar(data));
    }

    const handleChange = (e) => {
        var j = e.target.value
        setPosFiltro(j)
    }

    const concatenarFiltros = (e) => {
        e.preventDefault();
        var aux = filtros + listadoFiltros[posFiltro] + ", "
        setFiltros(aux)
    }

    return(
        <div className="form-signin">
            <div className="text-center">
                <form className="card card-body">
                    <h1 className="h3 mb-3 fw-normal">Grafo</h1>
                    <label htmlFor="inputEmail" className="visually-hidden">Empleado</label>
                    <input type="text" id="valorN" className="form-control" placeholder="100" required 
                    onChange={e => setEmpleado(e.target.value)} 
                    value={empleado} 
                    autoFocus/>
                    <br/>
                    <label htmlFor="inputEmail" className="visually-hidden">Cliente</label>
                    <input type="text" id="valorN" className="form-control" placeholder="1853" required 
                    onChange={e => setCliente(e.target.value)} 
                    value={cliente} 
                    autoFocus/>
                    <br/>
                    <label htmlFor="inputEmail" className="visually-hidden">Imagen</label>
                    <input type="text" id="valorN" className="form-control" placeholder="mario" required 
                    onChange={e => setImagenCliente(e.target.value)} 
                    value={imagenCliente} 
                    autoFocus/>
                    <br/>
                    <label htmlFor="inputEmail" className="visually-hidden">filtros</label>
                    <div className="col align-self-center">
                        <select className="form-control" aria-label=".form-select-lg example" onChange={handleChange}>
                            {
                                listadoFiltros.map((item, i) =>(
                                    <option value={i} key={"filtro"+i}>{item}</option>
                                ))
                            }
                        </select>
                    </div>
                    <br/>
                    <button className="w-100 btn btn-lg btn-primary" onClick={concatenarFiltros}>Aplicar Filtro</button>
                    <br/>
                    <button className="w-100 btn btn-lg btn-primary" onClick={pedirReporte}>Ver Reporte</button>
                    <br/>
                    <button className="w-100 btn btn-lg btn-primary" onClick={handleSubmit}>Enviar Valor</button>
                    <br/>
                    <img src={imagen} width="250" height="250" alt='some value' />
                </form>
            </div>
          </div>
    );
}