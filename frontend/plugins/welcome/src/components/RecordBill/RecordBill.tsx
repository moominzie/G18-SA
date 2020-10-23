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
import SaveRoundedIcon from '@material-ui/icons/SaveRounded';
import CancelRoundedIcon from '@material-ui/icons/CancelRounded';
import InputAdornment from '@material-ui/core/InputAdornment';

import { DefaultApi } from '../../api/apis';
import { EntBillingstatus } from '../../api/models/EntBillingstatus'; // import interface Billingtatus
import { EntEmployee } from '../../api/models/EntEmployee'; // import interface Employee
import { EntRepairInvoice } from '../../api/models/EntRepairInvoice'; // import interface Repairinvoice
import { EntBill } from '../../api/models/EntBill';

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

interface recordBill {
  price: number;
  time:  number;
  employee: number;
  repairinvoice : number;
  billingstatus: number;
}

export default function RecordBill() {
 const classes = useStyles();
 const http = new DefaultApi();
 
 const [bills, setBill] = React.useState<EntBill[]>([]);


 const [employees, setEmployees] = React.useState<EntEmployee[]>([]);
 const [repairinvoices, setRepairinvoices] = React.useState<EntRepairInvoice[]>([]);
 const [billingstatuss, setBillingstatuss] = React.useState<EntBillingstatus[]>([]);

 const [status, setStatus] = useState(false);
 const [alert, setAlert] = useState(true);

 const [loading, setLoading] = useState(true);

 const [price, setPrice] = useState(Number);
 const [time, setTime] = useState(Number);
 const [employee, setEmployee] = useState(Number);
 const [repairinvoice, setRepairinvoice] = useState(Number);
 const [billingstatus, setBillingstatus] = useState(Number);
 
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

  const getBillingstatuss = async () => {
    const res = await http.listBillingstatus({ limit: 10, offset: 0 });
    setLoading(false);
    setBillingstatuss(res);
    console.log(res);
  };
  getBillingstatuss();

}, [loading]);


const getBill = async () => {
  const res = await http.listBill({ limit: 10, offset: 0 });
  setBill(res);
};




const handlePriceChange = (event: any) => {
  setPrice(event.target.value as number);
};


const handleTimeChange = (event: any) => {
  setTime(event.target.value as number);
};

const EmployeehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
  setEmployee(event.target.value as number);
};
const RepairinvoicehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
  setRepairinvoice(event.target.value as number);
};

const BillingstatushandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
  setBillingstatus(event.target.value as number);
};


// create bill
const CreateBill = async () => {
  const bill = {
    employee: employee,
    price: Number(price),
    time: Number(time),
    billingstatus: billingstatus,
    repairinvoice: repairinvoice,
  };
  console.log(bill);
  const res: any = await http.createBill({ bill: bill });
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
        title={`Payment Billing Record`}
        type="Repairing computer systems"> 
      </Header>

      <Content>
        <ContentHeader title="Payment Billing"> 
              <Button onClick={() => {CreateBill();}} variant="contained"  color="primary" startIcon={<SaveRoundedIcon/>}> Create Bill </Button>
              <Button style={{ marginLeft: 20 }} component={RouterLink} to="/recordbilltable" variant="contained" startIcon={<CancelRoundedIcon/>}>  Dismiss </Button>
        </ContentHeader>  
        <div className={classes.root}>
          <form noValidate autoComplete="off">
            <FormControl
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
              
                {repairinvoices.map((item:EntRepairInvoice) => (
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
              
                {employees.map((item:EntEmployee) => (
                  <MenuItem value={item.id}>{item.employeename}</MenuItem>
                ))}
              </Select>
               <div className={classes.paper}><strong>Price</strong></div>
              <TextField className={classes.textField}
    //          style={{ width: 500 ,marginLeft:7,marginRight:-7}}
              InputProps={{
                startAdornment: (
                  <InputAdornment position="start">
                  </InputAdornment>
                ),
              }}
                id="price"
                label=""
                variant="standard"
                color="secondary"
                type="number"
                size="medium"
                value={price}
                onChange={handlePriceChange}
              />

            <div className={classes.paper}><strong>Time to repair (days)</strong></div>
              <TextField className={classes.textField}
              //style={{ width: 500 ,marginLeft:7,marginRight:-7}}
              InputProps={{
                startAdornment: (
                  <InputAdornment position="start">
                  </InputAdornment>
                ),
              }}
                id="time"
                label=""
                variant="standard"
                color="secondary"
                type="number"
                size="medium"
                value={time}
                onChange={handleTimeChange}
              />


              <div className={classes.paper}><strong>Billing Status</strong></div>
              <Select className={classes.select}
                //style={{ width: 500 ,marginLeft:7,marginRight:-7,marginTop:10}}
                color="secondary"
                id="billingstatus"
                value={billingstatus}
                onChange={BillingstatushandleChange}
              >
                <InputLabel className={classes.insideLabel}>Choose Billing status</InputLabel>

                {billingstatuss.map((item: EntBillingstatus) => (
                  <MenuItem value={item.id}>{item.billingstatusname}</MenuItem>
                ))}
              </Select>

              {status ? ( 
                      <div className={classes.margin} style={{ width: 500 ,marginLeft:30,marginRight:-7,marginTop:16}}>
              {alert ? ( 
                      <Alert severity="success"> <AlertTitle>Success</AlertTitle> Complete data — check it out! </Alert>) 
              : (     <Alert severity="warning"> <AlertTitle>Warining</AlertTitle> Incomplete data — please try again!</Alert>)} </div>
            ) : null}
            
            </FormControl>

          </form>
        </div>
      </Content>
    </Page>
  );
 }

