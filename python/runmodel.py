import pandas as pd
import joblib
from sklearn.preprocessing import MinMaxScaler

'''
Runs the logistic regression model on the given input values
Returns: Array of 2 values
    Idx 0: Probability of non-dangerous crime
    Idx 1: Probability of dangerous crime
'''
def run_log_model(latitude, longitude, time):
    input = pd.DataFrame([[latitude, longitude, time]], columns=["Latitude", "Longitude", "Time"])

    # Load model and scaler
    loaded_log_model = joblib.load("logreg_model")
    loaded_scaler = joblib.load("scaler")

    # scale the input
    input_scaled = pd.DataFrame(loaded_scaler.transform(input), columns = ['Latitude','Longitude','Time'])

    return loaded_log_model.predict_proba(input_scaled)[0]

'''
Runs the knn model on the given input values
Returns: Array of 2 values
    Idx 0: Probability of non-dangerous crime
    Idx 1: Probability of dangerous crime
'''
def run_log_model(latitude, longitude, time):
    input = pd.DataFrame([[latitude, longitude, time]], columns=["Latitude", "Longitude", "Time"])

    # Load model and scaler
    loaded_knn_model = joblib.load("knn_model")
    loaded_scaler = joblib.load("scaler")

    # scale the sample input
    input_scaled = pd.DataFrame(loaded_scaler.transform(input), columns = ['Latitude','Longitude','Time'])

    return loaded_knn_model.predict_proba(input_scaled)[0]
