from sqlalchemy import create_engine, func
from sqlalchemy.exc import SQLAlchemyError
from prices import Prices
from sqlalchemy.orm import sessionmaker


def get_db_engine(db_config):
    if len(db_config) < 5:
        print('some of the values of the db config is missing.')
        return

    db_host = db_config['host']
    db_port = db_config['port']
    db_name = db_config['dbName']
    db_user = db_config['user']
    db_password = db_config['password']

    try:
        engine = create_engine("postgresql+psycopg2://{user}:{pw}@{host}:{port}/{db}"
                           .format(user=db_user, pw=db_password, host=db_host, port=db_port, db=db_name))
    except SQLAlchemyError as e:
        print(f"error occurred while creating engine: {e}")
        return

    return engine

# returns average ceyrek price(try) per day
def get_daily_avg_ceyrek(db_engine):
    try:
        Session = sessionmaker(bind=db_engine)
        session = Session()
        daily_avg_ceyrek = session.query(
            func.date_trunc('day', Prices.last_at).label('day'),
            func.avg(Prices.ceyrek).label('average_ceyrek')).group_by(
            func.date_trunc('day', Prices.last_at)
        ).all()
    except SQLAlchemyError as e:
        print(f"error while getting average ceyrek price: {e}")
    finally:
        session.close()
        return daily_avg_ceyrek

def get_daily_avg_yarim(db_engine):
    try:
        Session = sessionmaker(bind=db_engine)
        session = Session()
        daily_avg_yarim = session.query(
            func.date_trunc('day', Prices.last_at).label('day'),
            func.avg(Prices.yarim).label('average_yarim')).group_by(
            func.date_trunc('day', Prices.last_at)
        ).all()
    except SQLAlchemyError as e:
        print(f"error while getting average yarim prices: {e}")
    finally:
        session.close()
        return daily_avg_yarim

def get_daily_avg_tam(db_engine):
    try:
        Session = sessionmaker(bind=db_engine)
        session = Session()
        daily_avg_tam = session.query(
            func.date_trunc('day', Prices.last_at).label('day'),
            func.avg(Prices.tam).label('average_tam')).group_by(
            func.date_trunc('day', Prices.last_at)
        ).all()
    except SQLAlchemyError as e:
        print(f"error while getting all the prices: {e}")
    finally:
        session.close()
        return daily_avg_tam

def get_daily_avg_cumhuriyet(db_engine):
    try:
        Session = sessionmaker(bind=db_engine)
        session = Session()
        daily_avg_tam = session.query(
            func.date_trunc('day', Prices.last_at).label('day'),
            func.avg(Prices.cumhuriyet).label('average_cumhuriyet')).group_by(
            func.date_trunc('day', Prices.last_at)
        ).all()
    except SQLAlchemyError as e:
        print(f"error while getting all cumhuriyet prices: {e}")
    finally:
        session.close()
        return daily_avg_tam
