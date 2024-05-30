from sqlalchemy import Column, BigInteger, TIMESTAMP
from sqlalchemy.ext.declarative import declarative_base

# Create a base class for the declarative model
Base = declarative_base()

class Prices(Base):
    __tablename__ = 'prices'

    id = Column(BigInteger, primary_key=True, autoincrement=True)
    last_at = Column(TIMESTAMP(timezone=True), nullable=True)
    ayar22_altin = Column(BigInteger, nullable=True)
    ceyrek = Column(BigInteger, nullable=True)
    yarim = Column(BigInteger, nullable=True)
    tam = Column(BigInteger, nullable=True)
    cumhuriyet = Column(BigInteger, nullable=True)
    iab_kapanis = Column(BigInteger, nullable=True)