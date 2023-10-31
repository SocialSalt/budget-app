import * as React from "react";
import {
  TableContainer,
  Table,
  TableHead,
  TableCell,
  TableBody,
  TableRow,
} from "@mui/material";
import { DataGrid, GridColDef, GridValueGetterParams } from "@mui/x-data-grid";
import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography";
import Paper from "@mui/material/Paper";

interface TableProps {
  children?: React.ReactNode;
  columnNames: string[];
  rows: string[][];
}

function generateKey(name: string, index: number) {
  return {
    key: `${name}-${index}`,
  };
}

export default function DataTable(props: TableProps) {
  const { columnNames, rows } = props;
  return (
    <TableContainer component={Paper}>
      <Table sx={{ minWidth: 650 }} aria-label="simple table">
        <TableHead>
          <TableRow>
            {columnNames.map((columnName, index) => (
              <TableCell {...generateKey(columnName, index)} >{columnName}</TableCell>
            ))}
          </TableRow>
        </TableHead>
        <TableBody>
          {rows.map((row, index) => (
            <TableRow
              {...generateKey(row[0][0], index)}
              sx={{ "&:last-child td, &:last-child th": { border: 0 } }}
            >
              {columnNames.map((_, index) => (
                <TableCell {...generateKey(row[index], index)}>{row[index]}</TableCell>
              ))}
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </TableContainer>
  );
}
