import React from 'react';
import { connect } from 'react-redux';
import { createUser, loginUser } from '../actions/userActions';
import { closeLoginForm } from '../actions/uiActions';
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
    this.handleSubmit = this.handleSubmit.bind(this);
    this.handleOnChange = this.handleOnChange.bind(this);
    this.handleLogout = this.handleLogout.bind(this);
    this.handleClose = this.handleClose.bind(this);
  }

  componentDidUpdate(prevProps, prevState) {
    if (prevProps.state.ui.vLoginForm !== this.props.state.ui.vLoginForm) {
      this.emailInput.focus();
    }
  }

  handleSubmit(e) {
    e.preventDefault();
    const { dispatch } = this.props;
    const user = this.state;
    console.log(user);
    dispatch(loginUser(user))
  }

  handleOnChange(e) {
    this.setState({
      [e.target.name]: e.target.value
    })
  }

  handleLogout() {
    const { dispatch } = this.props;
  }

  handleClose(e) {
    e.preventDefault();
    const { dispatch } = this.props;
    dispatch(closeLoginForm());
  }

  bindRefs(input) {
    this.emailInput = input;
  }

  render() {
    const { ui } = this.props.state;

    return (
      <div>
        <ReactCSSTransitionGroup
          transitionName="swipe-down"
          transitionEnterTimeout={300}
          transitionLeaveTimeout={300}>
        {ui.vLoginForm && 
          <form className={styles['login-form']} onSubmit={this.handleSubmit}> 
            <img src=""></img>
            <a className={styles['close-button']} href="#" onClick={this.handleClose}>CLOSE</a>
            <label htmlFor='email'>EMAIL</label>
            <input ref={this.bindRefs} type='email' name='email' onChange={this.handleOnChange} />
            <label htmlFor='password'>PASSWORD</label>
              <input type='password' name='password' onChange={this.handleOnChange} />
            <input className={buttonStyles.submit} type='submit' value='submit' /> 
          </form>          
        }
        </ReactCSSTransitionGroup>
        {ui.vLoginForm && 
          <div className={styles.backDrop}>BACKDROP</div>
        }   
      </div>
    );
  }
}

const mapStateToProps = state => {
  return { state }
}

const mapDispatchToProps = dispatch => {
  return { dispatch}
}

export default connect(mapStateToProps, mapDispatchToProps)(LoginForm);



