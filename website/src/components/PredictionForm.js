import { useState } from 'react';
import PredictionResult from './PredictionResult.js';
import "./PredictionForm.css";

function PredictionForm() {
  const [place, setPlace] = useState("");
  const [submittedPlace, setSubmittedPlace] = useState("");
  const [time, setTime] = useState("");
  const [submittedTime, setSubmittedTime] = useState("");
  const [risk, setRisk] = useState("");

  const handleSubmit = (event) => {
    event.preventDefault();
    setSubmittedTime(time);
    setSubmittedPlace(place);
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
    <div>
      <form onSubmit={handleSubmit} className="prediction-form">
        <div>
        <label className="predict-form-label">
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
        <label className="predict-form-label">
          Time (HH:MM)
          <br/>
          <input 
            type="text" 
            value={time}
            onChange={(e) => setTime(e.target.value)}
          />
        </label>
        </div>
        <input type="submit" className="submit-predict-button"/>
      </form>
      {/* <p>Place: {place}<br />Time: {time}</p> */}
    </div>
    {risk != "" && <PredictionResult risk={risk} time={submittedTime} place={submittedPlace}/>}
    </div>
  )
}

export default PredictionForm;