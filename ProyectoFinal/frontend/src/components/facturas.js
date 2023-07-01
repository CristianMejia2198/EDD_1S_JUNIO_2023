import React, {useState, useEffect} from 'react';
import '../css/administrador.css'
import 'bootstrap/dist/css/bootstrap.min.css'

export const Factura = () => {
    const idEmpleado = localStorage.getItem("empleado")
    const [facturas, setFacturas] = useState([])
    const salir = (e) => {
        e.preventDefault();
        console.log("Listo")
        window.open("/empleado","_self")
    }

    useEffect(() => {
        peticion()
    },[])

    const peticion = () => {
        fetch('http://localhost:3001/facturaempleado',{
        })
        .then(response => response.json())
        .then(data => validar(data))
    }

    const validar = (data) =>{
        console.log(data.factura)
        setFacturas(data.factura) 
    }

    return(
        <div className="form-signin2">
            <div className="text-center">
                  <form className="card card-body">
                    <h1 className="h3 mb-3 fw-normal">Facturas Generadas <br/> Empleado {localStorage.getItem("empleado")}</h1>
                    <br/>
                    <table className="table table-dark table-striped">
                        <thead>
                            <tr>
                                <th scope="col">#</th>
                                <th scope="col">ID Cliente</th>
                                <th scope="col">ID Factura</th>
                            </tr>
                        </thead>
                        <tbody>
                            {
                                facturas.map((element, j) => {
                                    if (element.Id_Cliente != '') {
                                        return <>
                                        <tr key={"fact"+j}>
                                            <th scope="row">{j+1}</th>
                                            <td>{element.Id_Cliente}</td>
                                            <td>{element.Id_Factura}</td>
                                        </tr>
                                    </>
                                    }
                                })
                            }
                        </tbody>
                    </table>
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