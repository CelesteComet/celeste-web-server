import React, { Component } from 'react';
import { connect } from 'react-redux';

function Header() {
  return (
    <header>
      I AM THE HEADER
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

