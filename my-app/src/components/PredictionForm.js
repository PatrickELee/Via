import { useState } from 'react';
import PredictionResult from './PredictionResult.js';
import "./PredictionForm.css";

function PredictionForm() {
  const [place, setPlace] = useState("");
  const [time, setTime] = useState("");
  const [risk, setRisk] = useState("");

  const handleSubmit = (event) => {
    event.preventDefault();
    let randnum = Math.random() * 4;
    if (randnum < 1)
      setRisk("Very low");
    else if (randnum < 2)
      setRisk("Low");
    else if (randnum < 3)
      setRisk("Medium");
    else
      setRisk("High");
  }

  return (
    <div>
    <div className="prediction-form-container">
      <h2></h2>
      <form onSubmit={handleSubmit} className="prediction-form">
        <div>
        <label className="form-input">
          Location
          <br/>
          <input
            type="text" 
            value={place}
            onChange={(e) => setPlace(e.target.value)}
          />
        </label>
        </div>
        <div>
        <label className="form-input">
          Time
          <br/>
          <input 
            type="text" 
            value={time}
            onChange={(e) => setTime(e.target.value)}
          />
        </label>
        </div>
        <input type="submit" className="submit-button"/>
      </form>
      {/* <p>Place: {place}<br />Time: {time}</p> */}
    </div>
    <PredictionResult risk={risk}/>
    </div>
  )
}

export default PredictionForm;