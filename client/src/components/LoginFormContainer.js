import React, { Component } from 'react';
import LoginForm from './LoginForm';
import ReactCSSTransitionGroup from 'react-addons-css-transition-group'; 
import { connect } from 'react-redux';
import styles from '../scss/loginForm';
import { createUser, loginUser } from '../actions/userActions';
import { closeLoginForm } from '../actions/uiActions';

class LoginFormContainer extends Component {
  constructor(props) {
    super(props)
    this.handleLogin = this.handleLogin.bind(this);
    this.handleClose = this.handleClose.bind(this);
  }

  handleLogin(user, e) {
    e.preventDefault();
    const { dispatch } = this.props;
    dispatch(loginUser(user))
  }

  handleClose(e) {
    e.preventDefault();
    const { dispatch } = this.props;
    dispatch(closeLoginForm());
  }

  render() {
    const { ui } = this.props.state;
    const { vLoginForm } = ui;
    return (
      <ReactCSSTransitionGroup
        transitionName="swipe-down"
        transitionEnterTimeout={300}
        transitionLeaveTimeout={300}>   
        {vLoginForm && 
          <LoginForm 
            visible={true} 
            handleLogin={this.handleLogin} 
            handleClose={this.handleClose} />
        }
      </ReactCSSTransitionGroup>     
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
