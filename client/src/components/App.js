import React, { Component, Fragment } from 'react';
import resetStyle from '../scss/reset.scss';
import Home from './Home';
import { BrowserRouter as Router, Route, Link, Switch } from "react-router-dom";
import { connect } from 'react-redux';
import { fetchUser } from '../actions/userActions';

import LoginPage from './LoginPage';
import NotFound from './NotFound';
import BagsIndexPage from './BagsIndexPage';
import AuthRoute from './HOC/AuthRoute';

import { toggleSideBarNav } from '../actions/uiActions';
import styles from '../scss/app.scss';
import HeaderContainer from './HeaderContainer';
import SideBarNav from './SideBarNav';
import BagDetailPage from './BagDetailPage';
import TypographyPage from './TypographyPage';

class App extends Component {
  constructor(props) {
    super(props); 
  }

  componentDidMount() {
    // Check if user is logged in
    console.log("Checking if user is logged in");
    const { dispatch } = this.props;
    dispatch(fetchUser())
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
          <HeaderContainer />
          <SideBarNav 
            visible={vSideBarNav} 
            handleMenuItemClick={this.handleMenuItemClick} />
          <div className="main">
            <Switch>
              <Route path="/" exact component={ Home } />
              <Route path="/login" exact component={ LoginPage } />
              <Route path="/bags" exact component={ BagsIndexPage } />
              <Route path="/bags/:id" component={ BagDetailPage } />
              <Route path="/typo" component={ TypographyPage } />
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


