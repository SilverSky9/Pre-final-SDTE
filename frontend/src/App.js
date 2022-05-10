import './App.css';
import Timer from './page/count'
import GetAPI from './page/getAPI'
import axios from 'axios';

axios.defaults.baseURL = 'http://localhost:3050'
axios.defaults.withCredentials = true


function App() {
  return (
    <div className="App">
      {/* <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          Hello World!!! and Jenkins
        </p>
        <a
          className="App-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn React
        </a>
      </header> */}
      <Timer/>
      <GetAPI/>
    </div>
    
  );
}

export default App;
