import React, { useState, useEffect } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import { ContentHeader, Content, Header, Page, pageTheme } from '@backstage/core';
import {
  Select,
  MenuItem,
} from '@material-ui/core';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import FormControl from '@material-ui/core/FormControl';
import { Alert, AlertTitle } from '@material-ui/lab';
import InputLabel from '@material-ui/core/InputLabel';
import AddCircleOutlinedIcon from '@material-ui/icons/AddCircleOutlined';
import SaveRoundedIcon from '@material-ui/icons/SaveRounded';
import CancelRoundedIcon from '@material-ui/icons/CancelRounded';
import InputAdornment from '@material-ui/core/InputAdornment';
import PersonIcon from '@material-ui/icons/Person';

import { DefaultApi } from '../../api/apis';
import { EntEmployee } from '../../api/models/EntEmployee'; // import interface Employee
import { EntRepairInvoice } from '../../api/models/EntRepairInvoice'; // import interface Repairinvoice
import { EntStatust } from '../../api/models/EntStatust'; // import interface Statust
import { EntReturninvoice } from '../../api/models/EntReturninvoice'; // import interface Returninvoice

// css style 
const useStyles = makeStyles((theme: Theme) =>
 createStyles({
   root: {
     display: 'flex',
     flexWrap: 'wrap',
     justifyContent: 'center',
   },
   margin: {
      margin: theme.spacing(2),
   },
   insideLabel: {
    margin: theme.spacing(1),
  },
   button: {
    marginLeft: '40px',
  },
   textField: {
    width: 500 ,
    marginLeft:7,
    marginRight:-7,
   },
    select: {
      width: 500 ,
      marginLeft:7,
      marginRight:-7,
      //marginTop:10,
    },
    paper: {
      marginTop: theme.spacing(1),
      marginBottom: theme.spacing(1),
      marginLeft: theme.spacing(1),
    },
  }),
);

interface recordReturninvoice {
  addedtime: Date;
  repairinvoice: number;
  employee: number;
  statust: number;
}

export default function RecordReturninvoice() {
 const classes = useStyles();
 const http = new DefaultApi();
 
 const [returninvoices, setReturninvoice] = React.useState<EntReturninvoice[]>([]);

 const [employees, setEmployees] = React.useState<EntEmployee[]>([]);
 const [repairinvoices, setRepairinvoices] = React.useState<EntRepairInvoice[]>([]);
 const [statusts, setStatusts] = React.useState<EntStatust[]>([]);

 const [status, setStatus] = useState(false);
 const [alert, setAlert] = useState(true);

 const [loading, setLoading] = useState(true);

 const [addedtime, setAddedtime] = useState(Date);
 const [employee, setEmployee] = useState(Number);
 const [repairinvoice, setRepairinvoice] = useState(Number);
 const [statust, setStatust] = useState(Number);

 useEffect(() => {
  const getEmployees = async () => {
    const res = await http.listEmployee({ limit: 10, offset: 0 });
    setLoading(false);
    setEmployees(res);
	console.log(res);
  };
  getEmployees();

  const getRepairinvoices = async () => {
    const res = await http.listRepairInvoice({ limit: 10, offset: 0 });
    setLoading(false);
    setRepairinvoices(res);
    console.log(res);
  };
  getRepairinvoices();

  const getStatusts = async () => {
    const res = await http.listStatust({ limit: 10, offset: 0 });
    setLoading(false);
    setStatusts(res);
	console.log(res);
  };
  getStatusts();

}, [loading]);



const getReturninvoice = async () => {
  const res = await http.listReturninvoice({ limit: 10, offset: 0 });
  setReturninvoice(res);
};

const EmployeehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
  setEmployee(event.target.value as number);
};

const RepairinvoicehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
  setRepairinvoice(event.target.value as number);
};

const StatusthandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
  setStatust(event.target.value as number);
};


// create returninvoice
const CreateReturninvoice = async () => {
  const returninvoice = {
    addedtime: addedtime,
    employee: employee,
    repairinvoice: repairinvoice,
    statust: statust,
  };
  console.log(returninvoice);
  const res: any = await http.createReturninvoice({ returninvoice: returninvoice });
  setStatus(true);
  
  if (res.id != '') {
    setAlert(true);
  } else {
    setAlert(false);
  }
  const timer = setTimeout(() => {
    setStatus(false);
  }, 3000);
};

  return (
    <Page theme={pageTheme.tool}>

      <Header
        title={`Create return invoice`}
        type="Repairing computer systems"> 
      </Header>

      <Content>
        <ContentHeader title="Return invoice"> 
              <Button onClick={() => {CreateReturninvoice();}} variant="contained"  color="primary" startIcon={<SaveRoundedIcon/>}> Create new return invoice </Button>
              <Button style={{ marginLeft: 20 }} component={RouterLink} to="/RecordReturninvoiceTable" variant="contained" startIcon={<CancelRoundedIcon/>}>  Dismiss </Button>
        </ContentHeader>  
        <div className={classes.root}>
          <form noValidate autoComplete="off">
            <FormControl
              //fullWidth
              //className={classes.margin}
              variant="outlined"
            >

              <div className={classes.paper}><strong>Repair Invoice NO.</strong></div>
              <Select className={classes.select}
                //style={{ width: 500 ,marginLeft:7,marginRight:-7,marginTop:10}}
                color="secondary"
                id="repairinvoice"
                value={repairinvoice}
                onChange={RepairinvoicehandleChange}
              >
                <InputLabel className={classes.insideLabel}>Choose Repair invoice</InputLabel>

                {repairinvoices.map((item: EntRepairInvoice) => (
                  <MenuItem value={item.id}>{item.rename}</MenuItem>
                ))}
              </Select>

              <div className={classes.paper}><strong>Employee</strong></div>
              <Select className={classes.select}
                //style={{ width: 500 ,marginLeft:7,marginRight:-7,marginTop:10}}
                color="secondary"
                id="employee"
                value={employee}
                onChange={EmployeehandleChange}
              >
                <InputLabel className={classes.insideLabel}>Choose Employee</InputLabel>

                {employees.map((item: EntEmployee) => (
                  <MenuItem value={item.id}>{item.employeename}</MenuItem>
                ))}
              </Select>

              <div className={classes.paper}><strong>Transport status</strong></div>
              <Select className={classes.select}
                //style={{ width: 500 ,marginLeft:7,marginRight:-7,marginTop:10}}
                color="secondary"
                id="statust"
                value={statust}
                onChange={StatusthandleChange}
              >
                <InputLabel className={classes.insideLabel}>Choose Transport status</InputLabel>

                {statusts.map((item: EntStatust) => (
                  <MenuItem value={item.id}>{item.statustname}</MenuItem>
                ))}
              </Select>
            
              {status ? ( 
                      <div className={classes.margin} style={{ width: 500 ,marginLeft:30,marginRight:-7,marginTop:16}}>
              {alert ? ( 
                      <Alert severity="success"> <AlertTitle>Success</AlertTitle> Complete data _ check it out! </Alert>) 
              : (     <Alert severity="warning"> <AlertTitle>Warining</AlertTitle> Incomplete data _ please try again!</Alert>)} </div>
            ) : null}
            
            </FormControl>

          </form>
        </div>
      </Content>
    </Page>
  );
 }