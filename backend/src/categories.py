import pandas as pd
from backend.src.utils import *


def load_raw_categories(
    filename=RAW_CATEGORIES_CSV,
) -> tuple[pd.DataFrame, pd.DataFrame]:
    raw_categories = pd.read_csv(filename)
    raw_categories.loc[:, CAT_MONTHS] = raw_categories.loc[:, CAT_MONTHS].applymap(
        dollar_to_int
    )

    categories = raw_categories.loc[:, ["Category", "Group", "Type"]]
    rows = []
    for m_y in CAT_MONTHS:
        for category in categories.Category:
            month = filter(lambda m: m_y.split(" ")[0] in m, MONTHS)
            year = int(m_y.split(" ")[1])
            value = raw_categories.loc[category].values[0]
            rows.append([category, month, year, f"{month} {year}", value])

    budget = pd.DataFrame(
        rows, columns=["Category", "Month", "Year", "MonthYear", "Value"]
    )

    return categories, budget


def load_categories(filename=CATEGORIES_CSV) -> pd.DataFrame:
    return pd.read_csv(filename)


def write_to_csv(df: pd.DataFrame, filename=CATEGORIES_CSV):
    df.to_csv(filename, index=False)


def get_categories_by_type(categories_df: pd.DataFrame, Type: str):
    return categories_df[categories_df.Type == Type]


@cache
def income_categories() -> pd.Series:
    categories = load_categories()
    return categories[categories.Type == "Income"].loc[:, "Category"]


@cache
def expense_categories() -> pd.Series:
    categories = load_categories()
    return categories[categories.Type == "Expense"].loc[:, "Category"]


@cache
def income_groups() -> pd.Series:
    categories = load_categories()
    categories = categories[categories.Type == "Income"]
    return categories.loc[:, "Group"].unique()


@cache
def expense_groups() -> pd.Series:
    categories = load_categories()
    categories = categories[categories.Type == "Expense"]
    return expense_categories().loc[:, "Group"].unique()
