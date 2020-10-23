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
import { EntRepairInvoice } from '../../api/models/EntRepairInvoice';
 
const useStyles = makeStyles({
 table: {
   minWidth: 650,
 },
});
 
export default function ComponentsRecordRepairInvoiceTable() {
  const classes = useStyles();
  const http = new DefaultApi();
  const [repairinvoices, setRepairInvoices] = useState<EntRepairInvoice[]>([]);
  const [loading, setLoading] = useState(true);
  
  useEffect(() => {
    const getRepairInvoices = async () => {
      const res = await http.listRepairInvoice({ limit: 5, offset: 0 });
      setLoading(true);
      setRepairInvoices(res);
      console.log(res);
    };
    getRepairInvoices();
  }, [loading]);
  
  const deleteRepairInvoices = async (id: number) => {
    const res = await http.deleteRepairInvoice({ id: id });
    setLoading(true);
  };
 
  
 return (
  <Page theme={pageTheme.tool}>
    <Header title={`Repair Invoice`} type="Repairing computer systems" >
    <Button variant="contained" color="default" href="/recordrepairinvoice" startIcon={<PersonAddRoundedIcon />}> New Repaired Invoice</Button>
    <div>&nbsp;&nbsp;&nbsp;</div>
    <Button variant="contained" color="primary" href="/home" startIcon={<HomeRoundedIcon/>}> home</Button>
    </Header>
    
    <Content>
   <TableContainer component={Paper}>
     <Table className={classes.table} aria-label="simple table">
       <TableHead>
         <TableRow>
           <TableCell align="center">Queue</TableCell>
           <TableCell align="center">Repaired Invoice ID</TableCell>
           <TableCell align="center">Owner</TableCell>
           <TableCell align="center">Device</TableCell>
           <TableCell align="center">Symptom</TableCell>
           <TableCell align="center">StatusR</TableCell>
         </TableRow>
       </TableHead>

       <TableBody>
         {repairinvoices.map((item:any) => (
           <TableRow key={item.id}>
             <TableCell align="center">{item.id}</TableCell>
             <TableCell align="center">{item.rename}</TableCell>
             <TableCell align="center">{item.edges?.user?.personalName}</TableCell>
             <TableCell align="center">{item.edges?.device?.dname}</TableCell>
             <TableCell align="center">{item.edges?.symptom?.syname}</TableCell>
             <TableCell align="center">{item.edges?.status?.sname}</TableCell>

             <TableCell align="center">
             <Button
                 onClick={() => {
                   deleteRepairInvoices(item.id);
                 }}
                 style={{ marginLeft: 10 }}
                 variant="contained"
                 color="secondary"
                 startIcon={<DeleteIcon/>}
                 href="/recordrepairinvoicetable"
               >
                 Delete
               </Button>
 
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