import * as React from "react";
import Box from "@mui/material/Box";
import InputLabel from "@mui/material/InputLabel";
import MenuItem from "@mui/material/MenuItem";
import FormControl from "@mui/material/FormControl";
import Select, { SelectChangeEvent } from "@mui/material/Select";

type HandleCallbackType = (a: string) => void;
interface MenuProps {
  children?: React.ReactNode;
  menuLabel: string;
  menuItems: string[];
  menuValues: any[];
  handleChangeCallback: HandleCallbackType;
}

function getMenuItemProps(index: number) {
  return {
    key: `menu_item_${index}`
  };
}

export default function DropDownMenu(props: MenuProps) {
  const [selection, setSelection] = React.useState("");
  const { children, menuLabel, menuItems, menuValues, handleChangeCallback } = props;
  const exampleHandleChange = (event: SelectChangeEvent) => {
    setSelection(event.target.value as string);
    handleChangeCallback(event.target.value as string);
  };

  return (
    <Box sx={{ minWidth: 120 }}>
      <FormControl fullWidth>
        <InputLabel id="demo-simple-select-label">{menuLabel}</InputLabel>
        <Select
          labelId="demo-simple-select-label"
          id="demo-simple-select"
          value={selection}
          label={menuLabel}
          onChange={exampleHandleChange}
        >
          {menuItems.map((menuItem, index) => (
            <MenuItem {...getMenuItemProps(index)} value={menuValues[index]}>{menuItem}</MenuItem>
          ))}
        </Select>
      </FormControl>
    </Box>
  );
}
