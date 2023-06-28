import logo from './logo.svg';
import './App.css';
import {BrowserRouter as Router, Route, Routes} from 'react-router-dom'
import { Login } from './components/login';
import { Administrador } from './components/administrador';

function App() {
  return(
    <Router>
      <Routes>
        <Route exact path='/' element={<Login/>} />
        <Route exact path='/admin' element={<Administrador/>} />
      </Routes>
    </Router>
  )
}

export default App;
