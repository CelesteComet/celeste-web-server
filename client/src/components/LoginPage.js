import React from 'react';
import LoginFormContainer from './LoginFormContainer';

function LoginPage() {
  return (
    <div className='wrapper'>
      <div>
        <h1>Sign In</h1>
        <LoginFormContainer />
      </div>
    </div>
  );
}

export default LoginPage;