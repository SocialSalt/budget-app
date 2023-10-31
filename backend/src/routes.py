import flask
from transactions import get_rows_by_month, get_rows_by_year, get_rows_by_month_and_year, TRANSACTIONS

app = flask.Flask(__name__)


@app.route("/transactions/<int:year>/<int:month>", methods=["GET"])
def query_transactions(year, month):
    if year == 0 and month == 0:
        result_df = TRANSACTIONS
    elif month == 0:
        result_df = get_rows_by_year(TRANSACTIONS, year)
    elif year == 0:
        result_df = get_rows_by_month(TRANSACTIONS, month)
    else:
        result_df = get_rows_by_month_and_year(TRANSACTIONS, month, year)
    columns = result_df.columns.values.tolist()
    index = result_df.index.values.tolist()
    rows = result_df.values.tolist()
    return {"columns": columns, "index": index, "rows": rows}


@app.route("/tables/income", methods=["GET"])
def get_income_table():
    """Get income tables

    Returns:
        list: List of income tables in the form of json. Each table has a header
            row that is the sum of the other rows. Each table has the following
            columns: Excess, Budget, Actual, Expected
            Excess is the extra income from previous months
            Budget is the total expected income this month
            Actual is the total received for that category this month
            Expected is the remaining income expected this month
    """
    pass


@app.route("/tables/expense", methods=["GET"])
def get_expense_table():
    pass


@app.route("/budgets", methods=["GET"])
def query_budgets():
    pass
