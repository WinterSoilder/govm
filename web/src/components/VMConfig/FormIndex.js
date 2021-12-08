import React, { useState } from 'react';
import {
  Stepper,
  Step,
  StepLabel,
  Button,
  Grid,
  CircularProgress
} from '@material-ui/core';
import { Formik, Form } from 'formik';

import CreateVM from './CreateVM';
import SSHUserForm from './SSHUserForm';
import ReviewVMDetails from './ReviewVMDetails';

import createVMModel from '../../shared/FormModel/createVMModel';

import useStyles from '../../shared/Styles/styles';
import initialValues from '../../shared/FormModel/initialValues';

const steps = ['VM Configuration Details', 'SSH User', 'Review'];
const { formId, formField } = createVMModel;

function _renderStepContent(step) {
  switch (step) {
    case 0:
      return <CreateVM formField={formField} />;
    case 1:
      return <SSHUserForm formField={formField} />;
    case 2:
      return <ReviewVMDetails />;
    default:
      return <div>Not Found</div>;
  }
}

export default function FormIndex() {
  const classes = useStyles();
  const [activeStep, setActiveStep] = useState(0);
  const isLastStep = activeStep === steps.length - 1;

  function _handleSubmit(values, actions) {
    if (isLastStep) {
      fetch('http://localhost:8030/vm_config', {
            method: 'POST',
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json',
                'Authorization': `Bearer ${localStorage.getItem('token')}`
            },
            body: JSON.stringify(values)
        })
        .then(async (response) => {
          actions.setSubmitting(false);
        })
        .catch((error) => {
            console.log(error)
        })
    } else {
      setActiveStep(activeStep + 1);
      actions.setTouched({});
      actions.setSubmitting(false);
    }
  }

  function _handleBack() {
    setActiveStep(activeStep - 1);
  }

  return (
    <Grid item container>
      <Grid item xs={3}/>
      <Grid item xs={6}>
    <React.Fragment>
      <Stepper activeStep={activeStep} className={classes.stepper}>
        {steps.map(label => (
          <Step key={label}>
            <StepLabel>{label}</StepLabel>
          </Step>
        ))}
      </Stepper>
      <React.Fragment>
        {activeStep === steps.length ? (
          <div>Success</div>
        ) : (
          <Formik
            initialValues={initialValues}
            onSubmit={_handleSubmit}
          >
            {({ isSubmitting }) => (
              <Form id={formId}>
                {_renderStepContent(activeStep)}

                <div className={classes.buttons}>
                  {activeStep !== 0 && (
                    <Button onClick={_handleBack} className={classes.button}>
                      Back
                    </Button>
                  )}
                  <div className={classes.wrapper}>
                    <Button
                      disabled={isSubmitting}
                      type="submit"
                      variant="contained"
                      color="primary"
                      className={classes.button}
                    >
                      {isLastStep ? 'Create VM' : 'Next'}
                    </Button>
                    {isSubmitting && (
                      <CircularProgress
                        size={24}
                        className={classes.buttonProgress}
                      />
                    )}
                  </div>
                </div>
              </Form>
            )}
          </Formik>
        )}
      </React.Fragment>
    </React.Fragment>
    <Grid item xs={3}/>
    </Grid>
    </Grid>
  );
}