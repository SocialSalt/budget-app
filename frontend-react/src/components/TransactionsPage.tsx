import * as React from "react";
import { Box, Typography } from "@mui/material";
import DropDownMenu from "./DropDownMenu";
import DataTable from "./DataTable";
import Grid from '@mui/material/Grid'; // Grid version 1

const base_url = process.env.REACT_APP_HOST_URI


async function getTransactions(year: number, month: number, setTransactionsCallback: (col: string[], row: string[][]) => void) {
  const response = await fetch(base_url + `/transactions/${year}/${month}`);
  const data = await response.json();
  setTransactionsCallback(data.columns, data.rows)
}


export default function TransactionsPage() {
  const [selection, setSelection] = React.useState({year: 0, month: 0});
  const [transactions, setTransactions] = React.useState({columns: [""], rows: [[""]]});
  
  const handleTransactionsChange = (columns: string[], rows: string[][]) => {
    setTransactions({...transactions, columns: columns, rows: rows});
  }

  const handleYearChange = (year: string) => {
    setSelection({...selection, year: parseInt(year)});
    getTransactions(selection.year, selection.month, handleTransactionsChange);
  };

  const handleMonthChange = (month: string) => {
    setSelection({...selection, month: parseInt(month)});
    getTransactions(selection.year, selection.month, handleTransactionsChange);
  };


  React.useEffect(() => {
    getTransactions(selection.year, selection.month, handleTransactionsChange);
  }, []);

  return (
    <div>
    <Grid container spacing={1}>
      <Grid item xs={8} />
      <Grid item xs={2}>
        <DropDownMenu menuLabel={"Year"} menuItems={["0", "1", "2"]} menuValues={[0, 1, 2]} handleChangeCallback={handleYearChange}/>
      </Grid>
      <Grid item xs={2}>
        <DropDownMenu menuLabel={"Month"} menuItems={["0", "1", "2"]} menuValues={[0, 1, 2]} handleChangeCallback={handleMonthChange}/>
      </Grid>
    </Grid>
    <DataTable columnNames={transactions.columns} rows={transactions.rows} />
    </div>
  );
}
