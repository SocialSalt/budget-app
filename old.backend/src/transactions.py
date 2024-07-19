from datetime import datetime

import pandas as pd
from utils import *


def load_raw_transactions(filename=RAW_TRANSACTIONS_CSV) -> pd.DataFrame:
    df = pd.read_csv(filename).drop("Unnamed: 0", axis=1)
    df.Date = pd.to_datetime(df.Date, format="%m/%d/%Y")
    df["Account #"] = df["Account #"].str.strip("x")
    df.Amount = df.Amount.map(dollar_to_int)
    df.MonthYear = df.Date.map(
        lambda x: MONTHS[getattr(x, "month") - 1] + f" {getattr(x, 'year')}"
    )
    return df

# RAW_TRANSACTIONS = load_raw_transactions()
# def reload_transactions() -> None:
#     global RAW_TRANSACTIONS
#     RAW_TRANSACTIONS = load_raw_transactions()


def load_transactions(filename=TRANSACTIONS_CSV) -> pd.DataFrame:
    return pd.read_csv(filename)


TRANSACTIONS = load_transactions()
def reload_transactions() -> None:
    global TRANSACTIONS
    TRANSACTIONS = load_transactions()


def write_to_csv(df: pd.DataFrame, filename: str = TRANSACTIONS_CSV):
    df.to_csv(filename, index=False)


def get_rows_by_category(transaction_df: pd.DataFrame, category: str):
    return transaction_df.loc[transaction_df.Category == category]


def get_rows_by_month(transaction_df: pd.DataFrame, month: int = datetime.now().month):
    month_mask = transaction_df.Date.map(lambda x: getattr(x, "month")) == month
    return transaction_df.loc[month_mask]


def get_rows_by_year(transaction_df: pd.DataFrame, year: int = datetime.now().year):
    year_mask = transaction_df.Date.map(lambda x: getattr(x, "year")) == year
    return transaction_df.loc[year_mask]


def get_rows_by_month_and_year(transaction_df: pd.DataFrame, month: int, year: int):
    month_mask = transaction_df.Date.map(lambda x: getattr(x, "month")) == month
    year_mask = transaction_df.Date.map(lambda x: getattr(x, "year")) == year
    return transaction_df.loc[month_mask & year_mask]


def get_rows_by_account(transaction_df: pd.DataFrame, acct_number: Accounts):
    mask = transaction_df["Account #"].map(lambda x: x[-4:]) == acct_number
    return transaction_df.loc[mask]

