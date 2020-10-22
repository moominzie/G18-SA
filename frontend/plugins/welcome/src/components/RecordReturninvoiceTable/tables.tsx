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
import { EntReturninvoice, EntRepairInvoice, EntUser } from '../../api';
 
const useStyles = makeStyles({
 table: {
   minWidth: 650,
 },
});
 
export default function ComponentsRecordReturninvoiceTable() {
  const classes = useStyles();
  const http = new DefaultApi();
  const [returninvoices, setReturninvoices] = useState<EntReturninvoice[]>([]);
  const [repairinvoice, setRepairinvoices] = useState<EntRepairInvoice[]>([]);
  const [user, setUsers] = useState<EntUser[]>([]);
  const [loading, setLoading] = useState(true);
  
  useEffect(() => {
    const getReturninvoices = async () => {
      const res = await http.listReturninvoice({ limit: 10, offset: 0 });
      setLoading(true);
      setReturninvoices(res);
      console.log(res);
    };
    getReturninvoices();


    const getRepairinvoices = async () => {
      const res = await http.listRepairInvoice({ limit: 10, offset: 0 });
      setLoading(false);
      setRepairinvoices(res);
      console.log(res);
    };
    getRepairinvoices();

    const getUsers = async () => {
      const res = await http.listUser({ limit: 10, offset: 0 });
      setLoading(false);
      setUsers(res);
      console.log(res);
    };
    getUsers();
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
             {repairinvoice.filter((setfilterid:any) => setfilterid.id === item.edges?.repairinvoice?.id).map((item2:any) => (
                  <TableCell align="center">{item2.edges?.user?.personalName}</TableCell>
              ))}
             {repairinvoice.filter((setfilterid:any) => setfilterid.id === item.edges?.repairinvoice?.id).map((item2:any) => (
                  <TableCell align="center">{item2.edges?.device?.dname}</TableCell>
              ))}
            {user.filter((setfilterid:any) => setfilterid.id === item.id).map((item2:EntUser) => (
                  <TableCell align="center">{item2.edges?.room?.rname}</TableCell>
              ))}
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