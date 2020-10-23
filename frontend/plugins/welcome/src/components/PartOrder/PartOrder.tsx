import React, { useState, useEffect } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import { ContentHeader, Content, Header, Page, pageTheme } from '@backstage/core';
import {
  Select,
  MenuItem,
} from '@material-ui/core';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import Button from '@material-ui/core/Button';
import FormControl from '@material-ui/core/FormControl';
import { Alert, AlertTitle } from '@material-ui/lab';
import InputLabel from '@material-ui/core/InputLabel';
import SaveRoundedIcon from '@material-ui/icons/SaveRounded';
import CancelRoundedIcon from '@material-ui/icons/CancelRounded';




import { DefaultApi } from '../../api/apis';
import { EntRepairInvoice } from '../../api/models/EntRepairInvoice'; // import interface Repairinvoice
import { EntEmployee } from '../../api/models/EntEmployee'; // import interface Employee
import { EntPart } from '../../api/models/EntPart'; // import interface Part
import { EntPartorder } from '../../api/models/EntPartorder'; // import interface Partorder

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

interface partorder {

  repairinvoice: number;
  employee: number;
  part: number;
}

export default function Partorder() {
 const classes = useStyles();
 const http = new DefaultApi();
 
 const [partorders, setPartorder] = React.useState<EntPartorder[]>([]);

 const [employees, setEmployees] = React.useState<EntEmployee[]>([]);
 const [repairinvoices, setRepairinvoices] = React.useState<EntRepairInvoice[]>([]);
 const [parts, setParts] = React.useState<EntPart[]>([]);

 const [status, setStatus] = useState(false);
 const [alert, setAlert] = useState(true);

 const [loading, setLoading] = useState(true);

 
 const [employee, setEmployee] = useState(Number);
 const [repairinvoice, setRepairinvoice] = useState(Number);
 const [part, setPart] = useState(Number);
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

  const getParts = async () => {
    const res = await http.listPart({ limit: 10, offset: 0 });
    setLoading(false);
    setParts(res);
	console.log(res);
  };
  getParts();

}, [loading]);



const getPartorder = async () => {
  const res = await http.listPartorder({ limit: 10, offset: 0 });
  setPartorder(res);
};


const EmployeehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
  setEmployee(event.target.value as number);
};

const RepairinvoicehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
  setRepairinvoice(event.target.value as number);
};

const ParthandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
  setPart(event.target.value as number);
};


// create returninvoice
const CreatePartorder = async () => {
  const partorder = {
    employee: employee,
    repairinvoice: repairinvoice,
    part: part,
  };
  console.log(partorder);
  const res: any = await http.createPartorder({ partorder : partorder });
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
        title={`Part Order Record`}
        type="REPAIRING COMPUTER SYSTEMS"> 
      </Header>

      <Content>
        <ContentHeader title="Pard Order"> 
              <Button onClick={() => {CreatePartorder();}} variant="contained"  color="primary" startIcon={<SaveRoundedIcon/>}> Save </Button>
              <Button style={{ marginLeft: 20 }} component={RouterLink} to="/Partordertable" variant="contained" startIcon={<CancelRoundedIcon/>}>  Dismiss </Button>
        </ContentHeader>  
        <div className={classes.root}>
          <form noValidate autoComplete="off">
            <FormControl
              //fullWidth
              //className={classes.margin}
              variant="outlined"
             
            >
               

           
               <div className={classes.paper}><strong>Repair invoice NO.</strong></div>
              <Select className={classes.select}
                //style={{ width: 500 ,marginLeft:7,marginRight:-7,marginTop:10}}
                color="secondary"
                id="repairinvoice"
                value={repairinvoice}
                onChange={RepairinvoicehandleChange}
              >
                <InputLabel className={classes.insideLabel}>NO.(Repairinvoice)</InputLabel>

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
                <InputLabel className={classes.insideLabel}>Name(Employee)</InputLabel>

                {employees.map((item: EntEmployee) => (
                  <MenuItem value={item.id}>{item.employeename}</MenuItem>
                ))}
              </Select>

              <div className={classes.paper}><strong>Part Name</strong></div>
              <Select className={classes.select}
                //style={{ width: 500 ,marginLeft:7,marginRight:-7,marginTop:10}}
                color="secondary"
                id="part"
                value={part}
                onChange={ParthandleChange}
              >
                <InputLabel className={classes.insideLabel}>NO.(Part)</InputLabel>

                {parts.map((item: EntPart) => (
                  <MenuItem value={item.id}>{item.pname}</MenuItem>
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

