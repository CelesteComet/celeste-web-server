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
      email: "",
      password: ""
    }
    this.bindRefs = this.bindRefs.bind(this);
    this.handleOnChange = this.handleOnChange.bind(this);
  }

  componentDidUpdate(prevProps, prevState) {
    // if (prevProps.state.ui.vLoginForm !== this.props.state.ui.vLoginForm) {
    //   this.emailInput.focus();
    // }
  }

  handleOnChange(e) {
    this.setState({
      [e.target.name]: e.target.value
    })
  }

  bindRefs(input) {
    this.emailInput = input;
  }

  render() {
    const user = this.state;
    return (
      <form className="login-form" onSubmit={this.props.handleLogin.bind(null, user)}> 
        <a className="login-form__close" href="#" onClick={this.props.handleClose}>
          <span className="icon-cross" />
        </a>
        <input ref={this.bindRefs} type='email' name='email' placeholder="EMAIL" onChange={this.handleOnChange} />
        <input type='password' name='password' placeholder="PASSWORD" onChange={this.handleOnChange} />
        <a href="#">forgot your password?</a>
        <input className="button__submit" type='submit' value='login' /> 
        {ReactDOM.createPortal(
          <div className="backdrop">BACKDROP</div>,
          document.body
        )}        
      </form>          
    );
  }
}

export default LoginForm;



