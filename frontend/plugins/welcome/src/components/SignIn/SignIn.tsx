import React, { FC,useState } from 'react';
import { Link } from "react-router-dom";
import Avatar from '@material-ui/core/Avatar';
import Button from '@material-ui/core/Button';
import TextField from '@material-ui/core/TextField';
import LockOutlinedIcon from '@material-ui/icons/LockOutlined';
import { makeStyles } from '@material-ui/core/styles';
import { Alert } from '@material-ui/lab';
import {
  Header,
  Page,
  pageTheme,
  Content,
 } from '@backstage/core';
import { DefaultApi } from '../../api/apis';

const useStyles = makeStyles(theme => ({
  paper: {
    marginTop: theme.spacing(1),
    marginBottom: theme.spacing(1),
    marginLeft: theme.spacing(1),
  },
  head: {
    marginLeft: theme.spacing(70),
    fontSize: '18px',
  },
  avatar: {
    marginTop: theme.spacing(1),
    marginBottom: theme.spacing(1),
    marginLeft: theme.spacing(84),
    backgroundColor: theme.palette.secondary.main,
  },
  form: {
    width: '100%', // Fix IE 11 issue.
    marginTop: theme.spacing(1),
  },
  submit: {
    margin: theme.spacing(2, 0, 2),
    marginLeft: theme.spacing(83),
  },
  textField: {
    width: 350 ,
    marginLeft: theme.spacing(67),
   },
   margin: {
    margin: theme.spacing(2),
 },
 signin: {
   margin: theme.spacing(2, 0, 2),
   width: 350 ,
   marginLeft: theme.spacing(67),
 }

}));


const SignIn: FC<{}> = () => {
  const classes = useStyles();

  const [alert, setAlert] = useState(true);
  const [status, setStatus] = useState(false);

  const [username, setUsername] = useState(String);
  const [password, setPassword] = useState(String);

  const handleUsernameChange = (event: any) => {
    setUsername(event.target.value as string);
  };
  
  const handlePasswordChange = (event: any) => {
    setPassword(event.target.value as string);
  };

  /*function EmployeeLogin() {
  if (username === 'admin' && password === 'system') {
    <Link to="/home" /> ;
 } else {
   setAlert(true);
 }
}*/

const EmployeeLogin = async () => {
  const employee = {
    username: 'admin',
    password: 'system',
  };
  console.log(employee);
    setStatus(true);
    console.log('Enter this Username and Password');
  if (username == 'admin' && password == 'system') {
    window.location.href = "http://localhost:3000/Home";
    setAlert(true);
    console.log('kluay');
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
       title="Signin" type="Repairing computer systems"> 
     </Header>

     <Content>
  <div className={classes.paper}> <Avatar className={classes.avatar}>
          <LockOutlinedIcon />
        </Avatar></div>
     <div className={classes.head}><strong>พนักงานศูนย์บริการแจ้งซ่อมคอมพิวเตอร์</strong></div>

     <form noValidate autoComplete="off">
     <div><TextField className={classes.textField}
                variant="outlined"
                margin="normal"
                required
                fullWidth
                id="username"
                label="Username"
                type="username"
                name="username"
                onChange={handleUsernameChange}
              /></div>
      <div><TextField className={classes.textField}
                variant="outlined"
                margin="normal"
                required
                fullWidth
                name="password"
                label="Password"
                type="password"
                id="password"
                onChange={handlePasswordChange}
              /></div></form>

            <div> 
            <Button
              onClick={() => {EmployeeLogin();}}
              type="submit"
              variant="contained"
              color="primary"
              className={classes.submit}
            >
              Sign In
            </Button></div>

            {status ? ( 
                      <div className={classes.signin}>
              { alert ? ( 
                    <Link to="/home" />   ) 
              : ( <Alert variant="outlined" severity="info"> Incorrect Username or Password </Alert> )} </div>
            ) : null}
     </Content>
    </Page>
  );
};

export default SignIn;