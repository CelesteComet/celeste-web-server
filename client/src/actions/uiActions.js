export const OPEN_LOGIN_FORM = 'OPEN_LOGIN_FORM';
export const CLOSE_LOGIN_FORM = 'CLOSE_LOGIN_FORM';

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
