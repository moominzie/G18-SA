import React, { FC } from 'react';
import Button from '@material-ui/core/Button';
import AddBoxRoundedIcon from '@material-ui/icons/AddBoxRounded';
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
              startIcon={<AddBoxRoundedIcon />}> 
              User information record
            </Button>
           <div></div>

            <Button 
              variant="outlined" 
              color="secondary" 
              href="/recordreturninvoicetable" 
              startIcon={<AddBoxRoundedIcon />}> 
              Record Returninvoice Table
            </Button>
        </Content>
      </div>
   </Page>
 );
};

