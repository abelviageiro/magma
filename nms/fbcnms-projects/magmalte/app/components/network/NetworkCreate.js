/**
 * Copyright 2004-present Facebook. All Rights Reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * @flow
 * @format
 */

import Button from '@material-ui/core/Button';
import DialogActions from '@material-ui/core/DialogActions';
import DialogContent from '@material-ui/core/DialogContent';
import DialogTitle from '@material-ui/core/DialogTitle';
import FormGroup from '@material-ui/core/FormGroup';
import Grid from '@material-ui/core/Grid';
import MagmaTopBar from '../MagmaTopBar';
import Paper from '@material-ui/core/Paper';
import React, {useState} from 'react';
import TextField from '@material-ui/core/TextField';

import axios from 'axios';
import {makeStyles} from '@material-ui/styles';
import {useEnqueueSnackbar} from '@fbcnms/ui/hooks/useSnackbar';

const useStyles = makeStyles(theme => ({
  container: {
    marginTop: theme.spacing.unit * 2,
  },
  content: {
    padding: theme.spacing.unit * 2,
  },
}));

export default function NetworkCreate() {
  const [name, setName] = useState('');
  const classes = useStyles();

  const enqueueSnackbar = useEnqueueSnackbar();

  const onCreate = async () => {
    let createResponse;
    try {
      createResponse = await axios.post('/nms/network/create', {name});
      if (!createResponse?.data.success) {
        enqueueSnackbar(
          `Unable to create network: ${createResponse?.data.message}`,
          {
            variant: 'error',
            autoHideDuration: 10000,
          },
        );
        return;
      }
    } catch (error) {
      const message = error.response?.data?.error || error;
      enqueueSnackbar(`Unable to create network: ${message}`, {
        variant: 'error',
        autoHideDuration: 10000,
      });
      return;
    }
    enqueueSnackbar('Created network ' + name, {
      variant: 'success',
      autoHideDuration: 10000,
    });
    window.location.href = `/nms/${name}/map`;
  };

  return (
    <>
      <MagmaTopBar />
      <div className={classes.container}>
        <Grid container spacing={24}>
          <Grid item xs />
          <Grid item xs={6}>
            <Paper className={classes.content}>
              <DialogTitle>Create Network</DialogTitle>
              <DialogContent>
                <FormGroup row>
                  <TextField
                    required
                    label="Network Name"
                    margin="normal"
                    value={name}
                    onChange={evt => setName(evt.target.value)}
                  />
                </FormGroup>
              </DialogContent>
              <DialogActions>
                <Button
                  disabled={name === ''}
                  onClick={onCreate}
                  color="primary"
                  variant="contained">
                  Create
                </Button>
              </DialogActions>
            </Paper>
          </Grid>
          <Grid item xs />
        </Grid>
      </div>
    </>
  );
}
