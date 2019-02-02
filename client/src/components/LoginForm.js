import React from 'react';

class LoginForm extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      email: "",
      password: ""
    }
    this.handleSubmit = this.handleSubmit.bind(this);
    this.handleOnChange = this.handleOnChange.bind(this);
  }

  handleSubmit(e) {
    e.preventDefault();
    let url = "http://localhost:1337/users";
    fetch(url, {
      method: "POST", // *GET, POST, PUT, DELETE, etc.
      mode: "cors",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify(this.state)
    })  
    .then((res) => {
      document.cookie = `JWT=${res.headers.get('JWT')}`;
    })  
  }

  handleOnChange(e) {
    this.setState({
      [e.target.name]: e.target.value
    })
  }

  render() {
    return (
      <form action='http://ec2-3-82-107-155.compute-1.amazonaws.com:1337/user' method='POST'>
        <label htmlFor='email'>EMAIL 
          <input type='email' name='email' onChange={this.handleOnChange} />
        </label>
        <label htmlFor='password'>PASSWORD
          <input type='password' name='password' onChange={this.handleOnChange} />
        </label>
        <input type='submit' value='submit' onClick={this.handleSubmit} />
      </form>
    );
  }
}

export default LoginForm;