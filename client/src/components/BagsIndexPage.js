import React, { Component } from 'react';
import { fetchBags } from '../actions/bagActions';
import { connect } from 'react-redux';
import BagsIndexPageItem from './BagsIndexPageItem';

class BagsIndexPage extends Component {

  constructor(props) {
    super(props);
  }

  componentDidMount() {
    const { dispatch } = this.props;
    dispatch(fetchBags());
  }

  render() {
    const { state } = this.props;
    const { bags } = state;

    if (!bags) {
      return <div>LOADING</div>
    }

    return (
      <div className="bags-index">
        {bags.map((bag, i) => {
          return <BagsIndexPageItem key={i} bag={bag} />
        })}
      </div>
    );      
  }
}

const mapDispatchToProps = dispatch => {
  return { dispatch }
}

const mapStateToProps = state => {
  return { state }
}

export default connect(mapStateToProps, mapDispatchToProps)(BagsIndexPage);