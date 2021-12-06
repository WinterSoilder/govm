import React from 'react';
import { Grid, Typography } from '@material-ui/core';
import { InputField } from '../../shared/FormFields';

export default function PaymentForm(props) {
  const {
    formField: { SSHName, SSHPassword }
  } = props;

  return (
    <React.Fragment>
      <Typography variant="h6" gutterBottom>
        User Credentials
      </Typography>
      <Grid container spacing={3}>
        <Grid item xs={12} md={6}>
          <InputField
            name={SSHName.name}
            label={SSHName.label}
            fullWidth
          />
        </Grid>
        <Grid item xs={12} md={6}>
          <InputField
            name={SSHPassword.name}
            label={SSHPassword.label}
            fullWidth
          />
        </Grid>
      </Grid>
    </React.Fragment>
  );
}
