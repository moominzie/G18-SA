import React, { FC } from 'react';
import Button from '@material-ui/core/Button';
import LocalFloristRoundedIcon from '@material-ui/icons/LocalFloristRounded';
import ExitToAppRoundedIcon from '@material-ui/icons/ExitToAppRounded';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import {
 Content,
 Header,
 Page,
 pageTheme,
} from '@backstage/core';

const useStyles = makeStyles((theme: Theme) =>
 createStyles({
    paper: {
      marginTop: theme.spacing(1),
      marginBottom: theme.spacing(1),
      marginLeft: theme.spacing(70),
    },
    margin: {
      marginBottom: theme.spacing(2),

    },
  }),
);

export default function WelcomePage() {
  const classes = useStyles();
 return (
   
   <Page theme={pageTheme.tool}>
     <Header
       title="Repairing computer systems" type="group 18">
          <Button variant="contained" color="primary" href="/" startIcon={<ExitToAppRoundedIcon />}> Logout
           </Button>
     </Header>

      <div className={classes.paper}>
        <Content>
            <Button className={classes.margin}
              variant="outlined" 
              color="secondary" 
              href="/recordusertable" 
              startIcon={<LocalFloristRoundedIcon />}> 
              User information
            </Button>
           <div></div>

            <Button className={classes.margin}
               variant="outlined" 
              color="secondary" 
             href="/recordrepairinvoicetable" 
              startIcon={<LocalFloristRoundedIcon />}> 
             Repair invoice
               </Button>
               <div></div>

               <Button className={classes.margin}
              variant="outlined" 
              color="secondary" 
              href="/recordusertable" 
              startIcon={<LocalFloristRoundedIcon />}> 
              Part Order
            </Button>
           <div></div>

            <Button className={classes.margin}
              variant="outlined" 
              color="secondary" 
              href="/recordreturninvoicetable" 
              startIcon={<LocalFloristRoundedIcon />}> 
              Return invoice
            </Button>
            <div></div>

            <Button className={classes.margin}
              variant="outlined" 
              color="secondary" 
              href="/recordusertable" 
              startIcon={<LocalFloristRoundedIcon />}> 
              Payment Bill
            </Button>
           <div></div>

            

        </Content>
      </div>
   </Page>
 );
};