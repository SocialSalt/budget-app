from enum import Enum
from functools import cache

RAW_TRANSACTIONS_CSV = "data/budget-app - Transactions.csv"
# TRANSACTIONS_CSV = "data/transactions.csv"
TRANSACTIONS_CSV = "data/test_data.csv"

RAW_CATEGORIES_CSV = "data/budget-app - Categories.csv"
CATEGORIES_CSV = "data/categories.csv"

BUDGETS_CSV = "data/budgets.csv"


class Accounts(Enum):
    SAPPHIRE = "9635"
    AMAZON = "0635"
    CHECKING = "8237"
    BROKERAGE = "4003"
    SAVINGS = "0639"
    EMILY_401K = "3926"


MONTHS = [
    "January",
    "February",
    "March",
    "April",
    "May",
    "June",
    "July",
    "August",
    "September",
    "October",
    "November",
    "December",
]
CAT_MONTHS = [
    "Jan 2023",
    "Feb 2023",
    "Mar 2023",
    "Apr 2023",
    "May 2023",
    "Jun 2023",
    "Jul 2023",
    "Aug 2023",
    "Sep 2023",
    "Oct 2023",
    "Nov 2023",
    "Dec 2023",
]


@cache
def dollar_to_int(dollar: str):
    return float(dollar.replace(",", "").replace("$", ""))
