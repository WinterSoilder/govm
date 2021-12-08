import React from 'react';
import { Grid, Typography } from '@material-ui/core';
import { InputField, SelectField } from '../../shared/FormFields';

export default function AddressForm(props) {
  const {
    formField: {
      VM_Name,
      CPU,
      Disk,
      Memory,
      Template
    }
  } = props;
  return (
    <React.Fragment>
      <Typography variant="h6" gutterBottom>
        VM Configuration Details  
      </Typography>
      <Grid container spacing={3}>
        <Grid item xs={12} sm={6}>
          <InputField name={VM_Name.name} label={VM_Name.label} fullWidth />
        </Grid>
        <Grid item xs={12} sm={6}>
          <InputField  name={CPU.name} label={CPU.label} fullWidth />
        </Grid>
        <Grid item xs={12} sm={6}>
          <InputField  name={Disk.name} label={Disk.label} fullWidth />
        </Grid>
        <Grid item xs={12} sm={6}>
          <InputField  name={Memory.name} label={Memory.label} fullWidth />
        </Grid>
        <Grid item xs={12} sm={6}>
          <SelectField
            name={Template.name}
            label={Template.label}
            data={Template.data}
            fullWidth
          />
        </Grid>
      </Grid>
    </React.Fragment>
  );
}
