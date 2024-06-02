import pandas as pd
from db import get_db_engine, get_daily_avg_ceyrek, get_daily_avg_yarim
from prophet import Prophet
import matplotlib.pyplot as plt
from config import load_config


def predict_ceyrek():
    config = load_config()
    db_engine = get_db_engine(config['db'])
    avg_ceyrek = get_daily_avg_ceyrek(db_engine)
    if avg_ceyrek is None:
        print("couldn't get average ceyrek value per day from db")
        return

    avg_ceyrek_df = pd.DataFrame(avg_ceyrek, columns=['ds', 'y'])
    avg_ceyrek_df['ds'] = avg_ceyrek_df['ds'].dt.tz_localize(None)

    print(avg_ceyrek_df)

    model = Prophet()
    model.fit(avg_ceyrek_df)
    future = model.make_future_dataframe(periods=30)
    forecast = model.predict(future)

    print(forecast[['ds', 'yhat', 'yhat_lower', 'yhat_upper']].tail(30))

    fig = model.plot(forecast)
    plt.show()

def predict_yarim():
    config = load_config()
    db_engine = get_db_engine(config['db'])
    avg_yarim = get_daily_avg_yarim(db_engine)
    if avg_yarim is None:
        print("couldn't get average ceyrek value per day from db")
        return

    avg_yarim_df = pd.DataFrame(avg_yarim, columns=['ds', 'y'])
    avg_yarim_df['ds'] = avg_yarim_df['ds'].dt.tz_localize(None)

    print(avg_yarim_df)

    model = Prophet()
    model.fit(avg_yarim_df)
    future = model.make_future_dataframe(periods=30)
    forecast = model.predict(future)

    print(forecast[['ds', 'yhat', 'yhat_lower', 'yhat_upper']].tail(30))

    fig = model.plot(forecast)
    plt.show()

def predict_tam():
    config = load_config()
    db_engine = get_db_engine(config['db'])
    avg_yarim = get_daily_avg_yarim(db_engine)
    if avg_yarim is None:
        print("couldn't get average ceyrek value per day from db")
        return

    avg_yarim_df = pd.DataFrame(avg_yarim, columns=['ds', 'y'])
    avg_yarim_df['ds'] = avg_yarim_df['ds'].dt.tz_localize(None)

    print(avg_yarim_df)

    model = Prophet()
    model.fit(avg_yarim_df)
    future = model.make_future_dataframe(periods=30)
    forecast = model.predict(future)

    print(forecast[['ds', 'yhat', 'yhat_lower', 'yhat_upper']].tail(30))

    fig = model.plot(forecast)
    plt.show()