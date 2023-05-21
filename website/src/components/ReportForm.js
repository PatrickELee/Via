import { useState } from 'react';
import "./ReportForm.css";

function ReportForm() {
  const [place, setPlace] = useState("");
  const [submittedPlace, setSubmittedPlace] = useState("");
  const [time, setTime] = useState("");
  const [submittedTime, setSubmittedTime] = useState("");

  const handleSubmit = (event) => {
    event.preventDefault();
    setSubmittedTime(time);
    setSubmittedPlace(place);
  }

  return (
    <div >
      <form onSubmit={handleSubmit} className="report-form">
        <div>
        <label className="report-form-label">
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
        <label className="report-form-label">
          Time (HH:MM)
          <br/>
          <input 
            type="text" 
            value={time}
            onChange={(e) => setTime(e.target.value)}
          />
        </label>
        </div>
        <input type="submit" className="submit-report-button"/>
      </form>
    </div>
  )
}

export default ReportForm;