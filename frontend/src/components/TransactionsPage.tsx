import * as React from "react";
import { Box } from "@mui/material";
import DropDownMenu from "./DropDownMenu";
import DataTable from "./DataTable";

const base_url = process.env.REACT_APP_HOST_URI

async function getTransactions(year: number, month: number) {
  // const response = await fetch(base_url + `/transactions/${year}/${month}`);
  // const data = await response.json();
  // console.log(data);
  // console.log(base_url + `/transactions/${year}/${month}`)
  // console.log(process.env.NODE_ENV)
  return;
}


export default function TransactionsPage() {
  const [selection, setSelection] = React.useState({year: 0, month: 0});

  const handleYearChange = (year: string) => {
    setSelection({...selection, year: parseInt(year)});
  };

  const handleMonthChange = (month: string) => {
    setSelection({...selection, month: parseInt(month)});
  };

  let transactions = getTransactions(selection.year, selection.month)

  return (
    <div>
    <Box>
      <DropDownMenu menuLabel={"Year"} menuItems={["0", "1", "2"]} menuValues={[0, 1, 2]} handleChangeCallback={handleYearChange}/>
      <DropDownMenu menuLabel={"Month"} menuItems={["0", "1", "2"]} menuValues={[0, 1, 2]} handleChangeCallback={handleMonthChange}/>
    </Box>
    <Box>Current year: {selection.year}</Box>
    <Box>Current month: {selection.month}</Box>
    {/* <DataTable /> */}
    </div>
  );
}
