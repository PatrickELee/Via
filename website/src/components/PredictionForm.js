import { useState } from 'react';
import PredictionResult from './PredictionResult.js';
import "./PredictionForm.css";

function PredictionForm() {
  const [place, setPlace] = useState("");
  const [submittedPlace, setSubmittedPlace] = useState("");
  const [time, setTime] = useState("");
  const [submittedTime, setSubmittedTime] = useState("");
  const [risk, setRisk] = useState("");

  const handleSubmit = async (event) => {
    event.preventDefault();

    const formData = `{"location" : "${place}", "time" : "${time}"}`;

		// formData.append("word", "fsdofjsj");
		const requestOptions = {
			method: 'POST',
			body: formData
		};
		setRisk('Test');
		const res = await fetch('http://localhost:32314/api/dangerprobability', requestOptions)
		const data = await res.json();
    console.log(data);
    setSubmittedTime(time);
    setSubmittedPlace(place);
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