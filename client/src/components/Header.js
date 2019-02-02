import React, { Component } from 'react';
import { connect } from 'react-redux';
import { Link } from 'react-router-dom';

//import styles from '../scss/header.module.scss';

function Header() {
  return (
    <header>
      <nav>
        <ul>
          <li><Link to="/">Pursey</Link></li>
          <li><Link to="/bags">Bags</Link></li>
          <li><Link to="/bags">Add Bag</Link></li>
          <li>Login</li>
        </ul>
      </nav>
    </header>
  );
}

const mapStateToProps = state => {
  return { state }
}

const mapDispatchToProps = dispatch => {
  return { dispatch }
}

export default connect(mapStateToProps, mapDispatchToProps)(Header)

