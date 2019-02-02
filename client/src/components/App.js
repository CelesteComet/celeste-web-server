import React, { Component, Fragment } from 'react';
import Home from './Home';
import { HashRouter as Router, Route, Link, Switch } from "react-router-dom";
import LoginForm from './LoginForm';
import NotFound from './NotFound';
import BagsIndexPage from './BagsIndexPage';
import Header from './Header';

class App extends Component {
  constructor(props) {
    super(props); 
  }

  render() {
    return (
      <Router>
        <Fragment>
          <Header />
          <Switch>
            <Route path="/" exact component={ Home } />
            <Route path="/bags" component={ BagsIndexPage } />
            {/*<Redirect from="/old-match" to="/will-match" />*/}
            <Route component={ NotFound } />
          </Switch>
        </Fragment>
      </Router>
    );
  }
}

export default App;