import React, { useState, useEffect } from 'react';
import { Link } from "react-router-dom";

import { Typography, Grid, Button } from '@material-ui/core';
import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableContainer from '@mui/material/TableContainer';
import TableHead from '@mui/material/TableHead';
import TableRow from '@mui/material/TableRow';
import Paper from '@mui/material/Paper';

function createData(name, calories, fat, carbs, protein) {
  return { name, calories, fat, carbs, protein };
}

const rows = [
  createData('Frozen yoghurt', 159, 6.0, 24, 4.0),
  createData('Ice cream sandwich', 237, 9.0, 37, 4.3),
  createData('Eclair', 262, 16.0, 24, 6.0),
  createData('Cupcake', 305, 3.7, 67, 4.3),
  createData('Gingerbread', 356, 16.0, 49, 3.9),
];

export default function ALLVms() {
  const [allVms, setVMs] = useState([])
  useEffect(() => {
    fetch('http://localhost:8030/vm_config', {
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json',
                'Authorization': `Bearer ${localStorage.getItem('token')}`
            },
        })
        .then(async (response) => {
          const res = await response.json()
          setVMs(res)
        })
        .catch((error) => {
            console.log(error)
        })
  }, [])

  console.log(allVms)

  return (
    <Grid item container>
      <Grid item xs={2} />
      <Grid item xs={7}>
        <Grid container style={{ padding: '20px' }}>
          <Grid item xs={9}>
            <Typography gutterBottom>My Virtual Machines</Typography>
          </Grid>
          <Grid item xs={3} >
            <Link
              to={'/vms/add'}
            >
              <Button
                type="submit"
                variant="contained"
                color="primary"
                style={{ float: 'right' }}
              >
                Create VM
              </Button>
            </Link>
          </Grid>
        </Grid>
        <TableContainer component={Paper}>
          <Table sx={{ minWidth: 650 }} aria-label="simple table">
            <TableHead>
              <TableRow>
                <TableCell align="left">VM name</TableCell>
                <TableCell align="left">CPUs</TableCell>
                <TableCell align="left">Disk&nbsp;(mb)</TableCell>
                <TableCell align="left">Memory&nbsp;(mb)</TableCell>
                <TableCell align="left">Template</TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {allVms.map((row) => (
                <TableRow
                  key={row.VM_name}
                  sx={{ '&:last-child td, &:last-child th': { border: 0 } }}
                >
                  <TableCell >{row.VM_name}</TableCell>
                  <TableCell >{row.CPUs}</TableCell>
                  <TableCell >{row.Disk}</TableCell>
                  <TableCell >{row.Memory}</TableCell>
                  <TableCell >{row.Template}</TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </TableContainer>
      </Grid>
      <Grid item xs={3} />
    </Grid>
  );
}