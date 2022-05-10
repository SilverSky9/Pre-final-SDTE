import './App.css';
import axios from 'axios';
import Select from './page/selection'
import Card from './page/card'

axios.defaults.baseURL = 'http://localhost:3050'
axios.defaults.withCredentials = true


function App() {
  return (
    <div className="App">
      <Select/>
      <Card/>
    </div>
    
  );
}

export default App;
