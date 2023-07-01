import logo from './logo.svg';
import './App.css';
import {BrowserRouter as Router, Route, Routes} from 'react-router-dom'
import { Login } from './components/login';
import { Administrador } from './components/administrador';
import { Reportes } from './components/reporte';
import { Empleado } from './components/empleado';
import { Filtros } from './components/filtros';
import { GenerarFactura } from './components/generarfactura';
import { Factura } from './components/facturas';

function App() {
  return(
    <Router>
      <Routes>
        <Route exact path='/' element={<Login/>} />
        <Route exact path='/admin' element={<Administrador/>} />
        <Route exact path='/reportes' element={<Reportes/>} />
        <Route exact path='/empleado' element={<Empleado/>} />
        <Route exact path='/filtros' element={<Filtros/>} />
        <Route exact path='/factura' element={<GenerarFactura/>} />
        <Route exact path='/verfactura' element={<Factura/>} />
      </Routes>
    </Router>
  )
}

export default App;
