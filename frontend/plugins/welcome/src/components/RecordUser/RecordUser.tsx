import React, { useState, useEffect } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import { ContentHeader, Content, Header, Page, pageTheme } from '@backstage/core';
import {
  Select,
  MenuItem,
} from '@material-ui/core';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import FormControl from '@material-ui/core/FormControl';
import { Alert, AlertTitle } from '@material-ui/lab';
import InputLabel from '@material-ui/core/InputLabel';
import AddCircleOutlinedIcon from '@material-ui/icons/AddCircleOutlined';
import CancelRoundedIcon from '@material-ui/icons/CancelRounded';
import InputAdornment from '@material-ui/core/InputAdornment';
import PersonIcon from '@material-ui/icons/Person';
import { DefaultApi } from '../../api/apis';

import { EntFaculty } from '../../api/models/EntFaculty'; // import interface Faculty
import { EntBranch } from '../../api/models/EntBranch'; // import interface Branch
import { EntBuilding } from '../../api/models/EntBuilding'; // import interface Building
import { EntRoom } from '../../api/models/EntRoom'; // import interface Room
import { EntUser } from '../../api';

// css style 
const useStyles = makeStyles((theme: Theme) =>
 createStyles({
   root: {
     display: 'flex',
     flexWrap: 'wrap',
     justifyContent: 'center',
   },
   margin: {
      margin: theme.spacing(2),
   },
   insideLabel: {
    margin: theme.spacing(1),
  },
   button: {
    marginLeft: '40px',
  },
   textField: {
    width: 500 ,
    marginLeft:7,
    marginRight:-7,
   },
    select: {
      width: 500 ,
      marginLeft:7,
      marginRight:-7,
    },
    paper: {
      marginTop: theme.spacing(1),
      marginBottom: theme.spacing(1),
      marginLeft: theme.spacing(1),
    },
  }),
);

/* interface recordUser {
  personalID: string;
  personalName: string;
  faculty: number;
  branch: number;
  building: number;
  room: number;
} */

export default function RecordUser() {
 const classes = useStyles();
 const http = new DefaultApi();
 
 const [users, setUser] = React.useState<EntUser[]>([]);

 const [facultys, setFacultys] = React.useState<EntFaculty[]>([]);
 const [branchs, setBranchs] = React.useState<EntBranch[]>([]);
 const [buildings, setBuildings] = React.useState<EntBuilding[]>([]);
 const [rooms, setRooms] = React.useState<EntRoom[]>([]);

 const [status, setStatus] = useState(false);
 const [alert, setAlert] = useState(true);
 const [loading, setLoading] = useState(true);

 const [personalid, setPersonalID] = useState(String);
 const [personalname, setPersonalName] = useState(String);
 const [faculty, setFaculty] = useState(Number);
 const [branch, setBranch] = useState(Number);
 const [building, setBuilding] = useState(Number);
 const [room, setRoom] = useState(Number);

 useEffect(() => {
  const getFacultys = async () => {
    const res = await http.listFaculty({ limit: 10, offset: 0 });
    setLoading(false);
    setFacultys(res);
    console.log(res);
  };
  getFacultys();

  const getBranchs = async () => {
    const res = await http.listBranch({ limit: 10, offset: 0 });
    setLoading(false);
    setBranchs(res);
    console.log(res);
  };
  getBranchs();

  const getBuildings = async () => {
    const res = await http.listBuilding({ limit: 10, offset: 0 });
    setLoading(false);
    setBuildings(res);
    console.log(res);
  };
  getBuildings();

  const getRooms = async () => {
    const res = await http.listRoom({ limit: 10, offset: 0 });
    setLoading(false);
    setRooms(res);
    console.log(res);
  };
  getRooms();

}, [loading]);


const getUser = async () => {
  const res = await http.listUser({ limit: 10, offset: 0 });
  setUser(res);
};

const handlePersonalIDChange = (event: any) => {
  setPersonalID(event.target.value as string);
};

const handlePersonalNameChange = (event: any) => {
  setPersonalName(event.target.value as string);
};

const FacultyhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
  setFaculty(event.target.value as number);
};

const BranchhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
  setBranch(event.target.value as number);
};

const BuildinghandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
  setBuilding(event.target.value as number);
};

const RoomhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
  setRoom(event.target.value as number);
};

// create user
const CreateUser = async () => {
  const user = {
    personalID: personalid,
    personalName: personalname,
    faculty: faculty,
    branch: branch,
    building: building,
    room: room,
  };
  console.log(user);
  const res: any = await http.createUser({ user: user });
    setStatus(true);
  if (res.id != '') {
    setAlert(true);
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
        title={`User information record`}
        type="Repairing computer systems"> 
      </Header>

      <Content>
        <ContentHeader title="User information"> 
              <Button onClick={() => {CreateUser();}} variant="contained"  color="primary" startIcon={<AddCircleOutlinedIcon/>}> Create new user </Button>
              <Button style={{ marginLeft: 20 }} component={RouterLink} to="/recordusertable" variant="contained" startIcon={<CancelRoundedIcon/>}> Dismiss </Button>
        </ContentHeader>  
        <div className={classes.root}>
          <form noValidate autoComplete="off">
            <FormControl
              variant="outlined"
            >
               <div className={classes.paper}><strong>รหัสประจำตัวผู้ใช้บริการ</strong></div>
              <TextField className={classes.textField}
              InputProps={{
                startAdornment: (
                  <InputAdornment position="start">
                    <PersonIcon />
                  </InputAdornment>
                ),
              }}
                id="personalID"
                label=""
                variant="standard"
                color="secondary"
                type="string"
                size="medium"
                value={personalid}
                onChange={handlePersonalIDChange}
              />

            <div className={classes.paper}><strong>ชื่อผู้ใช้บริการ</strong></div>
              <TextField className={classes.textField}
              InputProps={{
                startAdornment: (
                  <InputAdornment position="start">
                    <PersonIcon />
                  </InputAdornment>
                ),
              }}
                id="personalName"
                label=""
                variant="standard"
                color="secondary"
                type="string"
                size="medium"
                value={personalname}
                onChange={handlePersonalNameChange}
              />

              <div className={classes.paper}><strong>คณะ/สำนักวิชา</strong></div>
              <Select className={classes.select}
                color="secondary"
                labelId="faculty-label"
                id="faculty"
                value={faculty}
                onChange={FacultyhandleChange}
              >
                <InputLabel className={classes.insideLabel} id="faculty-label">เลือกคณะ/สำนักวิชา(Faculty)</InputLabel>

                {facultys.map((item: EntFaculty) => (
                  <MenuItem value={item.id}>{item.fname}</MenuItem>
                ))}
              </Select>

              <div className={classes.paper}><strong>สาขาวิชา</strong></div>
              <Select className={classes.select}
                color="secondary"
                id="branch"
                value={branch}
                onChange={BranchhandleChange}
              >
                <InputLabel className={classes.insideLabel}>เลือกสาขาวิชา(Branch)</InputLabel>

                {branchs.map((item: EntBranch) => (
                  <MenuItem value={item.id}>{item.brname}</MenuItem>
                ))}
              </Select>

              <div className={classes.paper}><strong>อาคาร</strong></div>
              <Select className={classes.select}
                color="secondary"
                id="building"
                value={building}
                onChange={BuildinghandleChange}
              >
                <InputLabel className={classes.insideLabel}>เลือกอาคาร(Building)</InputLabel>

                {buildings.map((item: EntBuilding) => (
                  <MenuItem value={item.id}>{item.buname}</MenuItem>
                ))}
              </Select>

              <div className={classes.paper}><strong>ห้อง</strong></div>
              <Select className={classes.select}
                color="secondary"
                id="room"
                value={room}
                onChange={RoomhandleChange}
              >
                <InputLabel className={classes.insideLabel}>เลือกห้อง(Room)</InputLabel>

                {rooms.map((item: EntRoom) => (
                  <MenuItem value={item.id}>{item.rname}</MenuItem>
                ))}
              </Select>

              {status ? ( 
                      <div className={classes.margin} style={{ width: 500 ,marginLeft:30,marginRight:-7,marginTop:16}}>
              {alert ? 
                (<Alert severity="success"> <AlertTitle>Success</AlertTitle> Complete data — check it out! </Alert>) 
              : (<Alert severity="warning"> <AlertTitle>Warning</AlertTitle> Incomplete data — please try again!</Alert>)} </div>
            ) : null}
            
            </FormControl>
          </form>
        </div>
      </Content>
    </Page>
  );
 }

