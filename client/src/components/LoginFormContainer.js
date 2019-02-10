import React, { Component } from 'react';
import LoginForm from './LoginForm';
import ReactCSSTransitionGroup from 'react-addons-css-transition-group'; 
import { connect } from 'react-redux';
import styles from '../scss/loginForm';
import { createUser, loginUser, loginWithGoogle } from '../actions/userActions';
import { closeLoginForm } from '../actions/uiActions';

class LoginFormContainer extends Component {
  constructor(props) {
    super(props)
    this.handleLogin = this.handleLogin.bind(this);
    this.handleClose = this.handleClose.bind(this);
    this.handleGoogle = this.handleGoogle.bind(this);
  }

  handleLogin(user, e) {
    e.preventDefault();
    const { dispatch } = this.props;
    dispatch(loginUser(user))
  }

  handleGoogle(e) {
    e.preventDefault();
    const { dispatch } = this.props;
  }

  handleClose(e) {
    e.preventDefault();
    const { dispatch } = this.props;
    dispatch(closeLoginForm());
  }

  render() {
    const { ui, errors } = this.props.state;
    const { vLoginForm } = ui;

    return (
      <LoginForm 
        visible={true} 
        handleLogin={this.handleLogin} 
        handleClose={this.handleClose} 
        handleGoogle={this.handleGoogle}
        errors={ errors } />              
    );
  }
}

const mapStateToProps = state => {
  return { state }
}
const mapDispatchToProps = dispatch => {
  return { dispatch }
}
export default connect(mapStateToProps, mapDispatchToProps)(LoginFormContainer);
