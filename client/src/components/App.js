import React, { Component, Fragment } from 'react';
import resetStyle from '../scss/reset.scss';
import Home from './Home';
import { BrowserRouter as Router, Route, Link, Switch } from "react-router-dom";
import { connect } from 'react-redux';
import { fetchUser } from '../actions/userActions';

import LoginFormContainer from './LoginFormContainer';
import NotFound from './NotFound';
import BagsIndexPage from './BagsIndexPage';
import Header from './Header';
import SideBarNav from './SideBarNav';
import AuthRoute from './HOC/AuthRoute';

import { toggleSideBarNav } from '../actions/uiActions';
import styles from '../scss/app.scss';

class App extends Component {
  constructor(props) {
    super(props); 
  }

  componentDidMount() {
    const { dispatch } = this.props;
setTimeout(function(){
    document.body.className="";
},500);    
  }

  handleMenuItemClick(e) {
    e.preventDefault();
    const dispatch = this.props;
    dispatch(toggleSideBarNav());
  }

  render() {
    const { state } = this.props;
    const { ui } = state;
    const vSideBarNav = ui.vSideBarNav;
    const { dispatch } = this.props;
    return (
      <Router>
        <Fragment>
          <Header />
          <SideBarNav 
            visible={vSideBarNav} 
            handleMenuItemClick={this.handleMenuItemClick} />
          <LoginFormContainer />
          <div className={styles.app}>
            <Switch>
              <Route path="/" exact component={ Home } />
              <Route path="/bags" component={ BagsIndexPage } />
              {/*<Redirect from="/old-match" to="/will-match" />*/}
              <AuthRoute path="/p" component={ BagsIndexPage } />
              <Route component={ NotFound } />
            </Switch>
          </div>
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


