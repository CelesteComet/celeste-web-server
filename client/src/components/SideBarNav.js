import React, { Component } from 'react';
import styles from '../scss/sideBarNav';
import { Link } from 'react-router-dom';
import ReactCSSTransitionGroup from 'react-addons-css-transition-group'; 

class SideBarNav extends Component {
  constructor(props) {
    super(props);
    this.state = {
      visible: false 
    }
    this.handleMenuItemClick = this.handleMenuItemClick.bind(this);
  }

  componentDidUpdate(prevProps) {
    // Typical usage (don't forget to compare props):
    if (this.props.visible !== prevProps.visible) {
      this.setState({
        visible: this.props.visible
      })
    }    
  }

  handleMenuItemClick(e) {
    this.setState({
      visible: false
    })
  }

  render() {
    return (
      <div>
        <div className={`sidebar-nav ${this.state.visible ? 'on' : ''}`}>
          <nav>
            <ul onClick={this.props.handleMenuItemClick}>
              <li><Link disabled to="/">HOME</Link></li>
              <li><Link to="">DATABASE</Link></li>
              <li><Link to="">ADD BAG</Link></li>
              <li><Link to="">STORIES</Link></li>
              <li><Link to="">TAGS</Link></li>
            </ul>
          </nav>
        </div>
      </div>
    );
  }
}

export default SideBarNav;