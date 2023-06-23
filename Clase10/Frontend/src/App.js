import logo from './logo.svg';
import './App.css';
import {BrowserRouter as Router, Route, Routes} from 'react-router-dom';
import { ArbolAVL } from './components/arbolAVL';
import { Matriz } from './components/matriz';

function App() {
  /*return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          Edit <code>src/App.js</code> and save to reload.
        </p>
        <a
          className="App-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn React
        </a>
      </header>
    </div>
  );*/
  return (
    <Router>
        <Routes>
          <Route exact path="/" element={<ArbolAVL/>}/>
          <Route exact path="/matriz" element={<Matriz/>}/>
        </Routes>
    </Router>
  );
}

export default App;
