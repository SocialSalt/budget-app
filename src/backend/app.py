import flask

app = flask.Flask(__name__)


@app.route("/hello")
def hello():
    return "hello world\n"


@app.route("/transactions", methods=["GET"])
def query_transactions():
    pass


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


if __name__ == "__main__":
    app.run(debug=True, port=8000)
