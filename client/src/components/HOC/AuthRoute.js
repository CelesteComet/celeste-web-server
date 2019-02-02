import React, { Component } from 'react';
import { connect } from 'react-redux';
import { BrowserRouter as Router, Route, Link, Switch } from "react-router-dom";

class AuthRoute extends Component {

  constructor(props) {
    super(props)
  }

  render() {
    const {path, component} = this.props;

    // implement if not logged in
    


    return (
      <Route path={path} component={component} />
    );
  }
}

const mapStateToProps = state => {
  return { state }
}

export default connect(mapStateToProps, null)(AuthRoute);