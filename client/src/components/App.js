import React, { Component, Fragment } from 'react';
import LoginForm from './LoginForm';

class App extends Component {
  constructor(props) {
    super(props); 
  }

  render() {
    return (
      <div>
        <LoginForm />
      </div>
    );
  }
}

export default App;