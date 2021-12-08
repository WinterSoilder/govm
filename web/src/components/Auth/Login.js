import React from 'react';
import { Button, Grid, Typography } from '@material-ui/core';
import { InputField } from '../../shared/FormFields';
import userModel from '../../shared/FormModel/userModel';
import initialValues from '../../shared/FormModel/initialValues';
import { useNavigate } from "react-router-dom";

import { Form, Formik } from 'formik';

export default function Login() {
    const { formId, formField } = userModel;
    const { Email, Password } = formField;
    const navigate = useNavigate();

    function _handleSubmit(values) {
        const loginValues = { "email": values.email, "password": values.password }
        fetch('http://localhost:8030/userlogin', {
            method: 'POST',
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(loginValues)
        })
            .then(async (response) => {
                let user = await response.json()
                localStorage.setItem('token', user.Response.AuthToken)
                navigate('/');
            })
            .catch((error) => {
                console.log(error)
            })
    }

    return (
        <Grid item container>
            <Grid item xs={4} />
            <Grid item xs={2}>
                <Formik
                    initialValues={initialValues}
                    onSubmit={_handleSubmit}
                >
                    <Form id={formId}>
                        <React.Fragment>
                            <Typography variant="h6" gutterBottom>
                                Sign In:
                            </Typography>
                            <Grid container spacing={3}>
                                <Grid item xs={12} md={12}>
                                    <InputField
                                        name={Email.name}
                                        label={Email.label}
                                        fullWidth
                                    />
                                </Grid>
                                <Grid item xs={12} md={12}>
                                    <InputField
                                        type="password"
                                        name={Password.name}
                                        label={Password.label}
                                        fullWidth
                                    />
                                </Grid>
                                <Grid item xs={12} md={12}>
                                    <Button
                                        type="submit"
                                        variant="contained"
                                        color="primary"
                                    >
                                        {'Sign In'}
                                    </Button>
                                </Grid>
                            </Grid>
                        </React.Fragment>
                    </Form>
                </Formik>
            </Grid>
            <Grid item xs={4} />
        </Grid>
    );
};