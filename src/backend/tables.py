import categories
import pandas as pd
import transactions
from utils import *


def group_tables(income_df, expense_df) -> dict:
    return_val = {}
    dfs = tuple(income_df, expense_df)
    categories_df = categories.load_categories()
    for i, Type in enumerate(["Income", "Expense"]):
        header = dfs[i].sum(axis=0)
        header = [
            Type,
            header.loc["Excess"],
            header.loc["Budget"],
            header.loc["Actual"],
            header.loc["Remaining"],
        ]
        return_val[f"{Type}Header"] = header
        table_list = []
        groups = categories_df[categories_df.Type == Type].Group.unique()
        for group in groups:
            cats = categories_df[categories_df.Group == group].Category
            header = dfs[i].loc[cats].sum()
            header = [
                group,
                header.loc["Excess"],
                header.loc["Budget"],
                header.loc["Actual"],
                header.loc["Remaining"],
            ]
            table = [header]
            for cat in cats:
                table.append(
                    [
                        cat,
                        dfs[i].loc[cat, "Excess"],
                        dfs[i].loc[cat, "Budget"],
                        dfs[i].loc[cat, "Actual"],
                        dfs[i].loc[cat, "Expected"],
                    ]
                )
            table_list.append(table)

        return_val[f"{Type}Groups"] = table_list

    return return_val


def generate_tables(transaction_df, categories_df, budgets_df, year, up_to_month):
    transaction_df = transactions.get_rows_by_year(transaction_df, year)
    transaction_df = transaction_df.pivot_table(
        values="Amount", index="Category", columns="Month", aggfunc=sum
    ).fillna(0)
    for cat in categories_df.index:
        if cat not in transaction_df.index:
            transaction_df.loc[cat] = 0

    budgets_df = budgets_df[budgets_df.Year == year].pivot_table(
        values="Value", index="Category", columns="Month"
    )

    expense_df = transaction_df.loc[categories_df.Type == "Expense"].fillna(0)
    income_df = transaction_df.loc[categories_df.Type == "Income"].fillna(0)
    income_budget = budgets_df.loc[categories_df.Type == "Income"]
    expense_budget = budgets_df.loc[categories_df.Type == "Expense"]

    previous_expense = expense_df.loc[:, MONTHS[: up_to_month - 1]].sum(axis=1).abs()
    previous_expense_budget = expense_budget.loc[:, MONTHS[: up_to_month - 1]].sum(
        axis=1
    )
    previous_income = income_df.loc[:, MONTHS[: up_to_month - 1]].sum(axis=1).abs()
    previous_income_budget = income_budget.loc[:, MONTHS[: up_to_month - 1]].sum(axis=1)

    current_expense = expense_df.loc[:, MONTHS[up_to_month - 1]].abs()
    current_expense_budget = expense_budget.loc[:, MONTHS[up_to_month - 1]]
    current_income = income_df.loc[:, MONTHS[up_to_month - 1]].abs()
    current_income_budget = income_budget.loc[:, MONTHS[up_to_month - 1]]

    income_excess = previous_income - previous_income_budget
    income_remaining = current_income_budget - current_income

    expense_left_over = previous_expense_budget - previous_expense
    expense_remaining = expense_left_over + current_expense_budget - current_expense

    income_df = pd.DataFrame(
        {
            "Excess": income_excess,
            "Budget": current_income_budget,
            "Actual": current_income,
            "Remaining": income_remaining,
        }
    )

    expense_df = pd.DataFrame(
        {
            "Excess": expense_left_over,
            "Budget": current_expense_budget,
            "Actual": current_expense,
            "Remaining": expense_remaining,
        }
    )
    return income_df, expense_df
