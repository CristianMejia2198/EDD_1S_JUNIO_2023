import logo from './logo.svg';
import './App.css';
import {BrowserRouter as Router, Route, Routes} from 'react-router-dom'
import { Login } from './components/login';
import { Administrador } from './components/administrador';
import { Reportes } from './components/reporte';
import { Empleado } from './components/empleado';
import { Filtros } from './components/filtros';

function App() {
  return(
    <Router>
      <Routes>
        <Route exact path='/' element={<Login/>} />
        <Route exact path='/admin' element={<Administrador/>} />
        <Route exact path='/reportes' element={<Reportes/>} />
        <Route exact path='/empleado' element={<Empleado/>} />
        <Route exact path='/filtros' element={<Filtros/>} />
      </Routes>
    </Router>
  )
}

export default App;
