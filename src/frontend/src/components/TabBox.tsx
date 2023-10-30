import * as React from "react";
import Tabs from "@mui/material/Tabs";
import Tab from "@mui/material/Tab";
import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography";
import DataTable from "./DataTable";

interface TabPanelProps {
  children?: React.ReactNode;
  index: number;
  value: number;
  tabName: string;
  hidden?: boolean;
}

interface TabBoxProps {
  children?: React.ReactNode;
  numTabs: number;
  tabNames: string[];
}

function getTabProps(index: number) {
  return {
    id: `simple-tab-${index}`,
    "aria-controls": `simple-tabpanel-${index}`,
  };
}

function getTabPanelProps(index: number) {
  return {
    id: `simple-tabpanel-${index}`,
    "aria-labelledby": `simple-tab-${index}`,
  };
}

function CustomTabPanel(props: TabPanelProps) {
  const { children, value, index, tabName, ...other } = props;
  // use tabName to get the right page to load
  const test_row = DataTable(
    {
      columnNames: [tabName, "col_1", "col_2"],
      rows: [
        [tabName + " val 1", "value 1", "value 2"],
        [tabName + " val 2", "value 3", "value 4"]
      ]
    }
  )
  return (
    <div
      role="tabpanel"
      hidden={value !== index}
      {...getTabPanelProps(index)}
      {...other}
    >
      {value === index && (
        <Box sx={{ p: 3 }}>
          {test_row}
          {/* <Typography>{children}</Typography> */}
        </Box>
      )}
    </div>
  );
}

export default function TabBox(tabBoxProps: TabBoxProps) {
  const { children, numTabs, tabNames, ...other } = tabBoxProps;
  const [value, setValue] = React.useState(0);

  const handleChange = (event: React.SyntheticEvent, newValue: number) => {
    setValue(newValue);
  };

  return (
    <Box sx={{ width: "100%" }}>
      <Box sx={{ borderBottom: 1, borderColor: "divider" }}>
        <Tabs
          value={value}
          onChange={handleChange}
          aria-label="basic tabs example"
        >
          {tabNames.map((tabName, index) => (
            <Tab label={tabName} {...getTabProps(index)} />
          ))}
        </Tabs>
      </Box>
      {tabNames.map((tabName, index) => (
        <CustomTabPanel tabName={tabName} value={value} index={index}>
          {index}
        </CustomTabPanel>
      ))}
    </Box>
  );
}
