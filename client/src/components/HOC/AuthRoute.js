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
    const { user } = this.props.state;

    if (!user) { return <p>PLEASE LOGIN</p>; }

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