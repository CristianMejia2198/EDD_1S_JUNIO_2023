import React, {useState, useEffect} from 'react';
import '../css/maqueta.css'
import 'bootstrap/dist/css/bootstrap.min.css'
import swal from 'sweetalert'
import axios from 'axios';

export const ArbolAVL = () =>{
    const [valorNuevo, setValorNuevo] = useState(0)
    const [imagen, setImagen] = useState('https://yakurefu.com/wp-content/uploads/2020/02/Chi_by_wallabby.jpg')
    const handleSubmit = async(e) => {
        e.preventDefault();
        await fetch('http://localhost:3001/agregar-arbol',{
            method: 'POST',
            mode: 'no-cors',
            body: JSON.stringify({
                Valor: parseInt(valorNuevo)
            }),
            headers:{
                'Access-Control-Allow-Origin': '*',
                'Content-Type': 'application/json'
            }
        })
    }

    const validar = (data) =>{
        console.log(data)
        setImagen(data.Imagenbase64)
        /*if (data.mensaje === "OK"){
            swal({
                title: "Session Iniciada",
                text: "Datos Correctos",
                icon: "success",
                button: "aceptar"
            }).then(respuesta => {
                if(respuesta){
                    console.log(respuesta);
                    localStorage.setItem('current',respuesta.data);
                    window.open("/inicio","_self");
                }
            })
            
        }else{
            swal({
                title: "Error en Credenciales",
                text: "Su usuario o contraseña son incorrectos",
                icon: "error",
                timer: "4000",
                buttons: false
            })
        }*/
    }

    const pedirReporte = (e) => {
        e.preventDefault();
        fetch('http://localhost:3001/reporte-arbol',{
        })
        .then(response => response.json())
        .then(data => validar(data));
    }

    return(
        <div className="form-signin">
            <div className="text-center">
                <form onSubmit={handleSubmit} className="card card-body">
                    <h1 className="h3 mb-3 fw-normal">Arbol AVL</h1>
                    <label htmlFor="inputEmail" className="visually-hidden">Valor</label>
                    <input type="text" id="valorN" className="form-control" placeholder="100" required 
                    onChange={e => setValorNuevo(e.target.value)} 
                    value={valorNuevo} 
                    autoFocus/>
                    <br/>
                    <button className="w-100 btn btn-lg btn-primary" type="submit">Enviar Valor</button>
                    <br/>
                    <button className="w-100 btn btn-lg btn-primary" onClick={pedirReporte}>Ver Reporte</button>
                    <br/>
                    <img src={imagen} width="250" height="250" alt='some value' />
                </form>
            </div>
          </div>
    );
}