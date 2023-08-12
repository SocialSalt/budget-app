import pandas as pd
from utils import *


def load_budgets(filename=BUDGETS_CSV):
    return pd.read_csv(filename)


def write_to_csv(df: pd.DataFrame, filename=BUDGETS_CSV):
    df.to_csv(filename, index=False)


def get_budgets_table(
    budget_df: pd.DataFrame, years: list[int] = None, months: list[str] = None
):
    if years:
        budget_df = budget_df[budget_df.Year in years]
    if months:
        budget_df = budget_df[budget_df.Month in months]

    return budget_df.pivot_table(values="Value", index="Category", columns="MonthYear")
