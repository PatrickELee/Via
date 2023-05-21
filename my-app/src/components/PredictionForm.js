import { useState } from 'react';
import PredictionResult from './PredictionResult.js';
import "./PredictionForm.css";

function PredictionForm() {
  const [place, setPlace] = useState("");
  const [formPlace, setFormPlaace] = useState("");
  const [time, setTime] = useState("");
  const [formTime, setFormTime] = useState("");
  const [risk, setRisk] = useState("");

  const handleSubmit = (event) => {
    event.preventDefault();
    setFormTime(time);
    setFormPlaace(place);
    let randnum = Math.random() * 4;
    if (randnum < 1)
      setRisk("Very low");
    else if (randnum < 2)
      setRisk("Low");
    else if (randnum < 3)
      setRisk("Moderate");
    else
      setRisk("High");
  }

  return (
    <div>
    <div className="prediction-form-container">
      <h2></h2>
      <form onSubmit={handleSubmit} className="prediction-form">
        <div>
        <label className="form-label">
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
        <label className="form-label">
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
    {risk != "" && <PredictionResult risk={risk} time={formTime} place={formPlace}/>}
    </div>
  )
}

export default PredictionForm;