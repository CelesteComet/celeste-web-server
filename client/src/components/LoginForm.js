import React from 'react';
import ReactDOM from 'react-dom'; // you used 'react-dom' as 'ReactDOM'
import { connect } from 'react-redux';

import ReactCSSTransitionGroup from 'react-addons-css-transition-group'; 

import styles from '../scss/loginForm';
import buttonStyles from '../scss/buttons.scss';

class LoginForm extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      user: {email: '', password: ''},
      errors: []
    }
    this.bindRefs = this.bindRefs.bind(this);
    this.handleOnChange = this.handleOnChange.bind(this);
  }

  componentDidUpdate(prevProps, prevState) {
    if (prevProps.errors !== this.props.errors) {
      this.setState({
        errors: this.props.errors
      })
    }
  }

  handleOnChange(e) {
    const user = {};
    user[e.target.name] = e.target.value;
    this.setState({
      user,
      errors: []
    })
  }

  bindRefs(input) {
    this.emailInput = input;
  }

  render() {
    let errors;
    if (this.state.errors) {
      errors = (
        this.state.errors.map(err => {
          return <li className="errors">{err}</li>
        })
      );
    }
    const {user} = this.state;
    return (
      <form className="login-form" onSubmit={this.props.handleLogin.bind(null, user)}> 
        <a className="login-form__close" href="#" onClick={this.props.handleClose}>
          {/*<span className="icon-cross" />*/}
        </a>

        <input ref={this.bindRefs} type='email' name='email' placeholder="EMAIL" onChange={this.handleOnChange} />
        <input type='password' name='password' placeholder="PASSWORD" onChange={this.handleOnChange} />
        <a href="#">forgot your password?</a>
        <input className="button__submit" type='submit' value='login' />
        <a href="/auth/google" >LOGIN WITH GOOGLE</a>
        <ul> 
          { errors }
        </ul>
      </form>          
    );
  }
}

export default LoginForm;



