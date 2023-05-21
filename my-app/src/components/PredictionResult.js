import React from 'react';
import './PredictionResult.css';

let riskLevelMap = {
  "High": "high",
  "Moderate": "moderate",
  "Low": "low",
  "Very low": "very-low"
};

let containerMap = {
  "High": "high-container",
  "Moderate": "moderate-container",
  "Low": "low-container",
  "Very low": "very-low-container"
};

function RiskLevel(props) {

  return (
    <span className={riskLevelMap[props.text]}>{props.text}</span>
  );
}

function PredictionResult(props) {
  var name = props.risk === "" ? "" : " " + containerMap[props.risk];

  return (
    <div className={"prediction-result-container" + name}>
      <h2>Dangerous Crime Prediction</h2>
      <h3><span><RiskLevel text={props.risk}/></span> risk of dangerous crime occurring at 00:00 at 400 Howard Way</h3>
    </div>
  );
}

export default PredictionResult;