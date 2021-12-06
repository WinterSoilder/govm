import React from 'react';
import { Typography, Grid } from '@material-ui/core';
import useStyles from '../../shared/Styles/styles';
import { useFormikContext } from 'formik';

function ReviewVMDetails(props) {
  const classes = useStyles();
  const { values: formValues } = useFormikContext();

  const {
      VM_Name,
      CPU,
      Disk,
      Memory,
      Template,
      SSHUsername,
      SSHPassword
    } = formValues;
  return (
    <Grid item container direction="column" xs={12} sm={6}>
      <Typography variant="h6" gutterBottom className={classes.title}>
        Virtual Machine details
      </Typography>
      <Grid container>
        <React.Fragment>
          <Grid item xs={6}>
            <Typography gutterBottom>Virtaul Machine Name</Typography>
          </Grid>
          <Grid item xs={6}>
            <Typography gutterBottom>{VM_Name}</Typography>
          </Grid>
        </React.Fragment>
        <React.Fragment>
          <Grid item xs={6}>
            <Typography gutterBottom>CPU Cores</Typography>
          </Grid>
          <Grid item xs={6}>
            <Typography gutterBottom>{CPU}</Typography>
          </Grid>
        </React.Fragment>
        <React.Fragment>
          <Grid item xs={6}>
            <Typography gutterBottom>Hard Disk</Typography>
          </Grid>
          <Grid item xs={6}>
            <Typography gutterBottom>{Disk}</Typography>
          </Grid>
        </React.Fragment>
        <React.Fragment>
          <Grid item xs={6}>
            <Typography gutterBottom>RAM</Typography>
          </Grid>
          <Grid item xs={6}>
            <Typography gutterBottom>{Memory}</Typography>
          </Grid>
        </React.Fragment>
        <React.Fragment>
          <Grid item xs={6}>
            <Typography gutterBottom>OS Template</Typography>
          </Grid>
          <Grid item xs={6}>
            <Typography gutterBottom>
              {Template}
            </Typography>
          </Grid>
        </React.Fragment>
        <React.Fragment>
          <Grid item xs={6}>
            <Typography gutterBottom>{SSHUsername}</Typography>
          </Grid>
          <Grid item xs={6}>
            <Typography gutterBottom>
              {SSHUsername}
            </Typography>
          </Grid>
        </React.Fragment>
        <React.Fragment>
          <Grid item xs={6}>
            <Typography gutterBottom>{SSHPassword}</Typography>
          </Grid>
          <Grid item xs={6}>
            <Typography gutterBottom>
              {SSHPassword}
            </Typography>
          </Grid>
        </React.Fragment>
      </Grid>
    </Grid>
  );
}

export default ReviewVMDetails;
