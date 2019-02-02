import React, { Component, Fragment } from 'react';
import styles from '../scss/reset.scss';
import Home from './Home';
import { BrowserRouter as Router, Route, Link, Switch } from "react-router-dom";
import { connect } from 'react-redux';
import { fetchUser } from '../actions/userActions';

import LoginForm from './LoginForm';
import NotFound from './NotFound';
import BagsIndexPage from './BagsIndexPage';
import Header from './Header';
import AuthRoute from './HOC/AuthRoute';

class App extends Component {
  constructor(props) {
    super(props); 
  }

  componentDidMount() {
    const { dispatch } = this.props;
    dispatch(fetchUser())    
  }

  render() {
    return (
      <Router>
        <Fragment>
          <Header />
          <LoginForm />
          <Switch>
            <Route path="/" exact component={ Home } />
            <Route path="/bags" component={ BagsIndexPage } />
            {/*<Redirect from="/old-match" to="/will-match" />*/}
            <AuthRoute path="/p" component={ BagsIndexPage } />
            <Route component={ NotFound } />
          </Switch>
        </Fragment>
      </Router>
    );
  }
}

const mapStateToProps = state => {
  return { state }
}

const mapDispatchToProps = dispatch => {
  return { dispatch }
}

export default connect(mapStateToProps, mapDispatchToProps)(App);


