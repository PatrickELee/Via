import { useState } from 'react';
import PredictionResult from './PredictionResult.js';
import "./PredictionForm.css";

function PredictionForm() {
  const [place, setPlace] = useState("");
  const [time, setTime] = useState("");
  const [risk, setRisk] = useState("");

  const handleSubmit = async (event) => {
    event.preventDefault();

    const formData = '{"location" : "Test Location", "time" : "1280"}';

		// formData.append("word", "fsdofjsj");
		const requestOptions = {
			method: 'POST',
			body: formData
		};
		setRisk('Test');
		const res = await fetch('http://localhost:32314/api/dangerprobability', requestOptions)
		const data = await res.json();
    console.log(data);
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