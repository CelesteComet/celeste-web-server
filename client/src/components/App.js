import React, { Component, Fragment } from 'react';
import CssBaseline from '@material-ui/core/CssBaseline';
import { MuiThemeProvider, createMuiTheme } from '@material-ui/core/styles';
import blue from '@material-ui/core/colors/blue';
import orange from '@material-ui/core/colors/orange';
import PrimarySearchAppBar from './AppBar/AppBar';

const theme = createMuiTheme({
  palette: {
    primary: blue,
    secondary: {
      main: '#a7ffeb',
    },
  },
  status: {
    danger: 'orange',
  },
});

class App extends Component {
  constructor(props) {
    super(props); 
  }

  render() {
    return (
      <MuiThemeProvider theme={theme}>
        <CssBaseline />
        <div>
          <PrimarySearchAppBar />
        </div>
      </MuiThemeProvider>
    );
  }
}

export default App;