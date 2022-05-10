import './App.css';
import axios from 'axios';
import Select from './page/selection'

axios.defaults.baseURL = 'http://localhost:3050'
axios.defaults.withCredentials = true


function App() {
  return (
    <div className="App">
      <Select/>
    </div>
    
  );
}

export default App;
