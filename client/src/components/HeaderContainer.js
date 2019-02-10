import React, { Component } from 'react';
import { connect }  from 'react-redux';
import Header from './Header';
import { loginUser, logoutUser } from '../actions/userActions';

class HeaderContainer extends Component {
  constructor(props) {
    super(props)
  }

  render() {
    const { user, login, logout } = this.props;
    return (
      <Header 
        login={ login }
        logout={ logout }
        user={ user } />
    )
  }
}

const mSTP = state => {
  return { 
    user: state.user
  }
}

const mDTP = dispatch => {
  return { 
    login: (user) => {
      dispatch(loginUser(user));
    },
    logout: () => {
      dispatch(logoutUser());
    }
  }
}

export default connect(mSTP, mDTP)(HeaderContainer)