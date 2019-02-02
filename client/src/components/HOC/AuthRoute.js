import React, { Component } from 'react';
import { connect } from 'react-redux';
import { BrowserRouter as Router, Route, Link, Switch } from "react-router-dom";
import { fetchUser } from '../../actions/userActions';

class AuthRoute extends Component {

  constructor(props) {
    super(props)
  }

  componentDidMount() {
    const { dispatch } = this.props;
    dispatch(fetchUser())
  }

  render() {
    const {path, component} = this.props;
    const { users } = this.props.state;
    console.log(users);

    if (users && !users.email) {
      return <p>Please Log In</p>
    }

    return (
      <Route path={path} component={component} />
    );
  }
}

const mapStateToProps = state => {
  return { state }
}

const mapDispatchToProps = dispatch => {
  return { dispatch }
}

export default connect(mapStateToProps, mapDispatchToProps)(AuthRoute);