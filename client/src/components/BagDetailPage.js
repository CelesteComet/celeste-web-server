import React, { Component } from 'react';
import { connect }  from 'react-redux';
import { fetchBag } from '../actions/bagActions';
import PageLoadSpinner from './PageLoadSpinner';

class BagDetailPage extends Component {
  constructor(props) {
    super(props)

    this.state = {
      loading: true 
    }
  }

  componentDidMount() {
    const { fetchBag } = this.props;
    const { id } = this.props.match.params;
    fetchBag(id);
  }

  componentDidUpdate(prevState, prevProps) {
    if (prevState.bag != this.props.bag) {
      this.setState({
        loading: false
      })
    }
  }

  render() {
    if (!this.props.bag || this.state.loading) {
      return <PageLoadSpinner />
    }

    const { name, brand, image_url, created_by_member } = this.props.bag;

    return (
      <div className="wrapper">
        <h1>{ name }</h1>
        <h1>{ brand }</h1>
        <h1>created by {created_by_member}</h1>
        <img src={image_url} alt={name} style={{width: '400px'}}></img>
      </div>
    );
  }
}

const mSTP = state => {
  return { 
    bag: state.bags[0] ? state.bags[0] : null
  }
}

const mDTP = dispatch => {
  return { 
    fetchBag: (i) => {
      dispatch(fetchBag(i))
    }
  }
}

export default connect(mSTP, mDTP)(BagDetailPage);