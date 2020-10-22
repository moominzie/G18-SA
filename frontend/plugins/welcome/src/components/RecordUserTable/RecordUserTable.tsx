import React, { useState, useEffect } from 'react';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
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
import { EntUser } from '../../api';


const useStyles = makeStyles((theme: Theme) =>
 createStyles({
  table: {
    minWidth: 650,
  },
  buttonRight: {
    marginLeft: theme.spacing(150),
    marginBottom: theme.spacing(2),
  },
  }),
);
 
export default function RecordUserTable() {
  const classes = useStyles();
  const http = new DefaultApi();
  const [users, setUsers] = useState<EntUser[]>([]);
  const [loading, setLoading] = useState(true);
  
  // get users
  useEffect(() => {
    const getUsers = async () => {
      const res = await http.listUser({ limit: 10, offset: 0 });
      setLoading(true);
      setUsers(res);
      console.log(res);
    };
    getUsers();
  }, [loading]);
  
  // delete button
  const deleteUsers = async (id: number) => {
    const res = await http.deleteUser({ id: id });
    setLoading(true);
  };

    // clear input form
    function clear() {
      setUsers([]);
    }
  
  // ui 
 return (
  <Page theme={pageTheme.tool}>
    <Header title={`User information record`} type="Repairing computer systems" >
    <div>&nbsp;&nbsp;&nbsp;</div>
    <Button variant="contained" color="default" href="/recorduser" startIcon={<PersonAddRoundedIcon />}> New user</Button>
    <div>&nbsp;&nbsp;&nbsp;</div>
    <Button variant="contained" color="primary" href="/home" startIcon={<HomeRoundedIcon/>}> home</Button>
    </Header>
    
    <Content>
   <TableContainer component={Paper}>
     <Table className={classes.table} aria-label="simple table">
       <TableHead>
         <TableRow>
           <TableCell align="center">No</TableCell>
           <TableCell align="center">Personal ID</TableCell>
           <TableCell align="center">Name</TableCell>
           <TableCell align="center">Faculty</TableCell>
           <TableCell align="center">Branch</TableCell>
           <TableCell align="center">Building</TableCell>
           <TableCell align="center">Room</TableCell>
         </TableRow>
       </TableHead>

       <TableBody>
         {users.map((item:any) => (
           <TableRow key={item.id}>
             <TableCell align="center">{item.id}</TableCell>
             <TableCell align="center">{item.personalID}</TableCell>
             <TableCell align="center">{item.personalName}</TableCell>
             <TableCell align="center">{item.edges?.faculty?.fname}</TableCell>
             <TableCell align="center">{item.edges?.branch?.brname}</TableCell>
             <TableCell align="center">{item.edges?.building?.buname}</TableCell>
             <TableCell align="center">{item.edges?.room?.rname}</TableCell>
             <TableCell align="center">
             <Button
                 onClick={() => {
                   deleteUsers(item.id);
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