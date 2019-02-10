import React, { Component } from 'react';
import { connect } from 'react-redux';
import { Link } from 'react-router-dom';
import LogoutLink from './LogoutLink';
import { openLoginForm, toggleSideBarNav } from '../actions/uiActions';
import Hamburger from './Hamburger';

import styles from '../scss/header.scss';

function Header({login, logout, user}) {
  return (
    <header>
      <nav>
        <ul>
          <li><Link to="/">Pursey</Link></li>
          <li><Link to="/bags">Bags</Link></li>
          <li><Link to="/bags">Add Bag</Link></li>
          { !user && 
            <li>
              <Link to="/login">Login</Link>
            </li> 
          }
          {
            user &&
            <li>
              <a onClick={logout}>Logout</a>
            </li>
          }
        </ul>
      </nav>
      {/* MOBILE NAV */}
      <nav className="mobile">
        <ul>
          <li>
            <Link to="/">
              <span onClick={(e) => {e.preventDefault(); dispatch(toggleSideBarNav()) }}>
                <Hamburger />
              </span>
            </Link>
          </li>
        </ul>
      </nav>
    </header>  
  );
}

export default Header;

