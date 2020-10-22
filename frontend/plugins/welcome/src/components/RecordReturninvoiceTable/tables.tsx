import React, { useState, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import Button from '@material-ui/core/Button';
import { DefaultApi } from '../../api/apis';
import DeleteIcon from '@material-ui/icons/Delete';
import {
  Content,
  Header,
  Page,
  pageTheme,
} from '@backstage/core';
import PersonAddRoundedIcon from '@material-ui/icons/PersonAddRounded';
import HomeRoundedIcon from '@material-ui/icons/HomeRounded';
import { ControllersReturninvoice, EntReturninvoice } from '../../api';
 
const useStyles = makeStyles({
 table: {
   minWidth: 650,
 },
});
 
export default function ComponentsRecordReturninvoiceTable() {
  const classes = useStyles();
  const http = new DefaultApi();
  const [returninvoices, setReturninvoices] = useState<EntReturninvoice[]>([]);
  const [loading, setLoading] = useState(true);
  
  useEffect(() => {
    const getReturninvoices = async () => {
      const res = await http.listReturninvoice({ limit: 10, offset: 0 });
      setLoading(true);
      setReturninvoices(res);
      console.log(res);
    };
    getReturninvoices();
  }, [loading]);
  
 
  
 return (
  <Page theme={pageTheme.tool}>
    <Header title={`Return invoice`} type="Repairing computer systems" >
    <Button variant="contained" color="default" href="/recordReturninvoice" startIcon={<PersonAddRoundedIcon />}>  New Return invoice</Button>
    <div>&nbsp;&nbsp;&nbsp;</div>
    <Button variant="contained" color="primary" href="/home" startIcon={<HomeRoundedIcon/>}> home</Button>
    </Header>
    
    <Content>
   <TableContainer component={Paper}>
     <Table className={classes.table} aria-label="simple table">
       <TableHead>
         <TableRow>
           <TableCell align="center">No</TableCell>
           <TableCell align="center">Name</TableCell>
           <TableCell align="center">Device</TableCell>
           <TableCell align="center">Address</TableCell>
           <TableCell align="center">Employees</TableCell>
           <TableCell align="center">Status</TableCell>
         </TableRow>
       </TableHead>

       <TableBody>
         {returninvoices.map((item:any) => (
           <TableRow key={item.id}>
             <TableCell align="center">{item.id}</TableCell>
             <TableCell align="center">{item.edges?.repairinvoice?.userid}</TableCell>
             <TableCell align="center">{item.edges?.repairinvoice?.deviceid}</TableCell>
             <TableCell align="center">{item.edges?.repairinvoice?.userid}</TableCell>
             <TableCell align="center">{item.edges?.employee?.employeename}</TableCell>
             <TableCell align="center">{item.edges?.statust?.statustname}</TableCell>
           </TableRow>
         ))}
       </TableBody>
     </Table>
   </TableContainer>
   </Content>
  </Page>
);
}