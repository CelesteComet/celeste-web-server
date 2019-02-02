import React, { Component } from 'react';
import axios from 'axios';

class APICall extends Component {

  constructor(props) {
    super(props);
  }

  componentDidMount() {
    const { query } = this.props;
    axios.get(query)
      .then(data => {
        console.log(data);
      })
  }

  render() {
    return (
      <div>awdaw</div>
    );
  }
}

export default APICall;
