import React, { useState } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import {
  Content,
  Header,
  Page,
  pageTheme,
  ContentHeader,
} from '@backstage/core';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import FormControl from '@material-ui/core/FormControl';
import { Alert } from '@material-ui/lab';
import { DefaultApi } from '../../api/apis';
import Select from '@material-ui/core/Select';
import InputLabel from '@material-ui/core/InputLabel';
import MenuItem from '@material-ui/core/MenuItem';

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      display: 'flex',
      flexWrap: 'wrap',
      justifyContent: 'center',
    },
    margin: {
      margin: theme.spacing(3),
    },
    withoutLabel: {
      marginTop: theme.spacing(3),
    },
    textField: {
      width: '25ch',
    },
  }),
);

const initialUserState = {
  name: 'System Analysis and Design',
  age: 20,
};

export default function Create() {
  const classes = useStyles();
  const profile = { givenName: 'ระบบฝากครรภ์' };
  const api = new DefaultApi();

  const [user, setUser] = useState(initialUserState);
  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(true);

  const handleInputChange = (event: any) => {
    const { id, value } = event.target;
    setUser({ ...user, [id]: value });
  };

  const CreateUser = async () => {
    const res: any = await api.createUser({ user });
    setStatus(true);
    if (res.id != '') {
      setAlert(true);
    } else {
      setAlert(false);
    }
    const timer = setTimeout(() => {
      setStatus(false);
    }, 1000);
  };

  return (
    <Page theme={pageTheme.home}>
      <Header
        title={`Welcome to ${profile.givenName || 'to Backstage'}`}
        subtitle="นพ.AAA BBBB"
      ></Header>
      <Content>
        <table border='0' width='90%' align='center'>
          <tr>
            <td><ContentHeader title="เพิ่มข้อมูลฝากครรภ์"></ContentHeader></td>
            <td></td>
            <td></td>
            <td><Button variant="outlined" color="secondary" href="#outlined-buttons">
              LogOut
                </Button>
            </td>
          </tr>
          <tr>
            <td></td>
            <td><font size='3'>ผู้ฝากครรภ์</font></td>
            <td><FormControl
              ///fullWidth
              className={classes.margin}
              variant="outlined"
            >
              <InputLabel id="label"></InputLabel>
              <Select labelId="label" id="select" value="1" style={{ width: 400 }}>
                <MenuItem value="1">30001 นาง สมหญิง นอนน้อย</MenuItem>
                <MenuItem value="2">30002 นาง มาดี จริงใจ</MenuItem>
              </Select>
            </FormControl></td>
            <td></td>
          </tr>
          <tr>
            <td></td>
            <td><font size='3'>แพทย์ที่ดูแล</font></td>
            <td>
              <FormControl
                ///fullWidth
                className={classes.margin}
                variant="outlined"
              >
                <InputLabel id="label"></InputLabel>
                <Select labelId="label" id="select" value="2" style={{ width: 400 }}>
                  <MenuItem value="2">20001 นพ.AAA BBBB</MenuItem>
                </Select>
              </FormControl>
            </td>
            <td></td>
          </tr>
          <tr>
            <td></td>
            <td><font size='3'>สถานะเด็ก</font></td>
            <td>
              <FormControl
                ///fullWidth
                className={classes.margin}
                variant="outlined"
              >
                <InputLabel id="label"></InputLabel>
                <Select labelId="label" id="select" value="ปกติ" style={{ width: 400 }}>
                  <MenuItem value="ปกติ">7001 ปกติ</MenuItem>
                  <MenuItem value="ไม่ปกติ">7002 ไม่ปกติ</MenuItem>
                </Select>
              </FormControl>
            </td>
            <td></td>
          </tr>
          <tr>
            <td></td>
            <td><font size='3'>เวลา</font></td>
            <td>
              <FormControl className={classes.margin} >
                <TextField
                  id="datetime-local"
                  label="DATE-TIME"
                  type="datetime-local"
                  defaultValue="2017-05-24T10:30"
                  className={classes.textField}
                  InputLabelProps={{
                    shrink: true,
                  }}
                />
              </FormControl>
            </td>
            <td></td>
          </tr>
        </table>
        <div className={classes.root}>
          <form noValidate autoComplete="off">
            <div className={classes.margin}>
              <Button
                onClick={() => {
                  CreateUser();
                }}
                variant="contained"
                color="primary"
              >
                Submit
             </Button>
              <Button
                style={{ marginLeft: 20 }}
                component={RouterLink}
                to="/"
                variant="contained"
              >
                Back
             </Button>
            </div>
          </form>
        </div>
      </Content>
    </Page>
  );
}
