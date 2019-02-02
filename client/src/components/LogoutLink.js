import React, { Component } from 'react';
import { Link } from 'react-router-dom';
import { logoutUser } from '../actions/userActions';
import { connect } from 'react-redux';

class LogoutLink extends Component {
  constructor(props) {
    super(props);
    this.handleLogout = this.handleLogout.bind(this);
  }

  handleLogout(e) {
    e.preventDefault();
    const { dispatch } = this.props
    dispatch(logoutUser());
  }

  render() {
    return (
      <Link to="/logout" onClick={this.handleLogout}>Logout</Link>
    );
  }
}

const mapDispatchToProps = dispatch => {
  return { dispatch }
}

export default connect(null, mapDispatchToProps)(LogoutLink);