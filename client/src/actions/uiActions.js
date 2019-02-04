export const OPEN_LOGIN_FORM = 'OPEN_LOGIN_FORM';
export const CLOSE_LOGIN_FORM = 'CLOSE_LOGIN_FORM';
export const TOGGLE_SIDEBAR_NAV = 'TOGGLE_SIDEBAR_NAV';

export const openLoginForm = () => {
  return {
    type: OPEN_LOGIN_FORM
  }
}

export const closeLoginForm = () => {
  return {
    type: CLOSE_LOGIN_FORM 
  }
}

export const toggleSideBarNav = () => {
  return {
    type: TOGGLE_SIDEBAR_NAV 
  }
}

