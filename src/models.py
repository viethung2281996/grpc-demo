
from sqlalchemy import Column, Integer, func
from sqlalchemy.dialects.postgresql import TIMESTAMP

from . import db


class Symbol(db.Model):
    __tablename__ = 'symbol'

    id = Column(Integer, primary_key=True)
    created_at = Column(TIMESTAMP(timezone=True), server_default=func.now())
    updated_at = Column(TIMESTAMP(timezone=True), server_onupdate=func.now())
