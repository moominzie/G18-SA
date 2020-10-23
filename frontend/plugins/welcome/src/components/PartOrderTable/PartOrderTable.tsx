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
import { EntRepairInvoice, EntPartorder} from '../../api';
 
const useStyles = makeStyles({
 table: {
   minWidth: 650,
 },
});
 
export default function ComponentsPartOrderTable() {
  const classes = useStyles();
  const http = new DefaultApi();
  const [partorders, setPartorders] = useState<EntPartorder[]>([]);
  const [repairinvoice, setRepairinvoices] = useState<EntRepairInvoice[]>([]);
  const [loading, setLoading] = useState(true);
  
  useEffect(() => {
    const getUsers = async () => {
      const res = await http.listPartorder({ limit: 10, offset: 0 });
      setLoading(true);
      setPartorders(res);
      console.log(res);
    };
    getUsers();

    const getRepairinvoices = async () => {
      const res = await http.listRepairInvoice({ limit: 10, offset: 0 });
      setLoading(true);
      setRepairinvoices(res);
      console.log(res);
    };
    getRepairinvoices();
  }, [loading]);
  
  const deletePartorders = async (id: number) => {
    const res = await http.deletePartorder({ id: id });
    setLoading(true);
  };
 
  
 return (
  <Page theme={pageTheme.tool}>
    <Header title={`Part Order`} type="Repairing computer systems" >
    <Button variant="contained" color="default" href="/PartOrder" startIcon={<PersonAddRoundedIcon />}></Button>
    <div>&nbsp;&nbsp;&nbsp;</div>
    <Button variant="contained" color="primary" href="/home" startIcon={<HomeRoundedIcon/>}> home</Button>
    </Header>
    
    <Content>
   <TableContainer component={Paper}>
     <Table className={classes.table} aria-label="simple table">
       <TableHead>
         <TableRow>
           <TableCell align="center">No</TableCell>
           <TableCell align="center">Device Name</TableCell>
           <TableCell align="center">Employee Name</TableCell>
           <TableCell align="center">Part Name</TableCell>
         </TableRow>
       </TableHead>

       <TableBody>
         {partorders.map((item:any) => (
           <TableRow key={item.id}>
             <TableCell align="center">{item.id}</TableCell>
             {repairinvoice.filter((setfilterid:any) => setfilterid.id === item.edges?.repairinvoice?.id).map((item2:any) => (
                  <TableCell align="center">{item2.edges?.device?.dname}</TableCell>
              ))}
             <TableCell align="center">{item.edges?.employee?.employeename}</TableCell>
             <TableCell align="center">{item.edges?.part?.pname}</TableCell>

             <TableCell align="center">
             <Button
                 onClick={() => {
                   deletePartorders(item.id);
                 }}
                 style={{ marginLeft: 10 }}
                 variant="contained"
                 color="secondary"
                 startIcon={<DeleteIcon/>}
                 href="/recordusertable"
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