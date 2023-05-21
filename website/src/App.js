import './App.css';
import { useState } from 'react';
import PredictionForm from './components/PredictionForm.js';
import ReportForm from './components/ReportForm';

function App() {
  const [form, setFormChoice] = useState("predict");

  function handleChange(event) {
    setFormChoice(event.target.value);
  }

  return (
    <div>
      <h1 className="title">Via</h1>
      <div className="animated-line"></div>
      <div className="form-container">
      <select className="select-form" onChange={handleChange}>
        <option value="predict">Predict</option>
        <option value="report">Report</option>
      </select>
        {form == "predict" ? <PredictionForm /> : <ReportForm />}
      </div>
    </div>
  );
}

export default App;
