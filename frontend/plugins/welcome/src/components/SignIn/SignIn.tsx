import React, { FC,useState,useEffect,Component } from 'react';
import Avatar from '@material-ui/core/Avatar';
import Button from '@material-ui/core/Button';
import TextField from '@material-ui/core/TextField';
import LockOutlinedIcon from '@material-ui/icons/LockOutlined';
import Typography from '@material-ui/core/Typography';
import { makeStyles } from '@material-ui/core/styles';
import { Alert, AlertTitle } from '@material-ui/lab';
import {
  Header,
  Page,
  pageTheme,
  Content,
 } from '@backstage/core';
import { DefaultApi } from '../../api/apis';

function Copyright() {
  return (
    <Typography variant="body2" color="textSecondary" align="center">
      {'Copyright © '}
      Group 18 SA-63{' '}
      {new Date().getFullYear()}
      {'.'}
    </Typography>
  );
}

const useStyles = makeStyles(theme => ({
  paper: {
    marginTop: theme.spacing(1),
    marginBottom: theme.spacing(1),
    marginLeft: theme.spacing(1),
    //display: 'flex',
   // flexDirection: 'column',
    //alignItems: 'center',
    //align: 'center',
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
  },
  textField: {
    width: 350 ,
    marginLeft:7,
    marginRight:-7,
   },
   margin: {
    margin: theme.spacing(2),
 },

}));


const SignIn: FC<{}> = () => {
  const classes = useStyles();
  const http = new DefaultApi();

  const [alert, setAlert] = useState(true);

  const [employeeemail, setEmployeeEmail] = useState(String);
  const [password, setPassword] = useState(String);

  const handleEmployeeEmailChange = (event: any) => {
    setEmployeeEmail(event.target.value as string);
  };

  const handlePasswordChange = (event: any) => {
    setPassword(event.target.value as string);
  };

  const EmployeeLogin = async () => {
  if (employeeemail === 'admin' && password === 'system') {
    setAlert(true);
  } else {
    setAlert(false);
  }
}

  return (
    <Page theme={pageTheme.tool}>

<Header
       title="Login" type="Repairing computer systems"> 
     </Header>

     <Content align="center">
  <div className={classes.paper}> <Avatar className={classes.avatar}>
          <LockOutlinedIcon />
        </Avatar></div>
     <div className={classes.paper}><strong>พนักงานศูนย์บริการแจ้งซ่อมคอมพิวเตอร์</strong></div>
     <TextField className={classes.textField}
    //          style={{ width: 500 ,marginLeft:7,marginRight:-7}}
                variant="outlined"
                margin="normal"
                required
                fullWidth
                id="employeeemail"
                label="Email Address"
                type="email"
                name="email"
                autoFocus
                onChange={handleEmployeeEmailChange}
               // value={personalid}
                //onChange={handlePersonalIDChange}
              />
 <div></div>
      <TextField className={classes.textField}
    //          style={{ width: 500 ,marginLeft:7,marginRight:-7}}
                variant="outlined"
                margin="normal"
                required
                fullWidth
                name="password"
                label="Password"
                type="password"
                id="password"
                onChange={handlePasswordChange}
               // value={personalid}
                //onChange={handlePersonalIDChange}  
              />
            <div> 
            <Button
              onClick={() => {EmployeeLogin();}}
              href="/home"
              type="submit"
              variant="contained"
              color="primary"
              className={classes.submit}
            >
              Sign In
            </Button></div>

            {status ? ( 
                      <div className={classes.margin} style={{ width: 500 ,marginLeft:30,marginRight:-7,marginTop:16}}>
              {alert ? ( 
                      <Alert severity="success"> <AlertTitle>Success</AlertTitle> Complete data — check it out! </Alert>) 
              : (     <Alert severity="warning"> <AlertTitle>Warining</AlertTitle> Incomplete data — please try again!</Alert>)} </div>
            ) : null}
     </Content>
    </Page>
  );
};

export default SignIn;