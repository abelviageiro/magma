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
import Popout from '../components/Popout.react';
import React from 'react';
import Typography from '@material-ui/core/Typography';
import {storiesOf} from '@storybook/react';

storiesOf('Popout', module).add('default', () => (
  <div style={{padding: 100}}>
    <Popout
      content={
        <div style={{padding: 20}}>
          <Typography>Content</Typography>
        </div>
      }>
      <Button variant="contained" color="primary">
        Click me!
      </Button>
    </Popout>
  </div>
));
