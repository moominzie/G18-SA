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

import {
  Content,
  Header,
  Page,
  pageTheme,
} from '@backstage/core';
import PersonAddRoundedIcon from '@material-ui/icons/PersonAddRounded';
import HomeRoundedIcon from '@material-ui/icons/HomeRounded';
import { ControllersBill, EntBill } from '../../api';
 
const useStyles = makeStyles({
 table: {
   minWidth: 650,
 },
});
 
export default function ComponentsRecordBillTable() {
  const classes = useStyles();
  const http = new DefaultApi();
  const [bills, setBills] = useState<EntBill[]>([]);
  const [loading, setLoading] = useState(true);
  
  useEffect(() => {
    const getBills = async () => {
      const res = await http.listBill({ limit: 10, offset: 0 });
      setLoading(true);
      setBills(res);
      console.log(res);
    };
    getBills();
  }, [loading]);
  
 return (
  <Page theme={pageTheme.tool}>
    <Header title={`Bill record`} type="Repairing computer systems" >
    <Button variant="contained" color="default" href="/recordbill" startIcon={<PersonAddRoundedIcon />}> New Bill</Button>
    <div>&nbsp;&nbsp;&nbsp;</div>
    <Button variant="contained" color="primary" href="/home" startIcon={<HomeRoundedIcon/>}> home</Button>
    </Header>
    
    <Content>
   <TableContainer component={Paper}>
     <Table className={classes.table} aria-label="simple table">
       <TableHead>
         <TableRow>
           <TableCell align="center">No</TableCell>
           <TableCell align="center">Repair invoice NO.</TableCell>
           <TableCell align="center">Price</TableCell>
           <TableCell align="center">Time</TableCell>
           <TableCell align="center">Empolyee</TableCell>
           <TableCell align="center">BillingStatus</TableCell>
         </TableRow>
       </TableHead>

       <TableBody>
         {bills.map((item:any) => (
           <TableRow key={item.id}>
             <TableCell align="center">{item.id}</TableCell>
             <TableCell align="center">{item.edges?.repairinvoice?.rename}</TableCell>
             <TableCell align="center">{item.price}</TableCell>
             <TableCell align="center">{item.time}</TableCell>
             <TableCell align="center">{item.edges?.employee?.employeename}</TableCell>
             <TableCell align="center">{item.edges?.billingstatus?.billingstatusname}</TableCell>
             <TableCell align="center">
 
             </TableCell>

           </TableRow>
         ))}
       </TableBody>
     </Table>
   </TableContainer>
   </Content>
  </Page>
);
}