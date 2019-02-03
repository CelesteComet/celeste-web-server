export const CLOSE_LOGIN_FORM = 'CLOSE_LOGIN_FORM';
export const OPEN_LOGIN_FORM = 'OPEN_LOGIN_FORM';
export const CLOSE_MODAL = 'CLOSE_MODAL';
export const OPEN_MODAL = 'OPEN_MODAL';

export const closeLoginForm = () => {
  return {
    type: CLOSE_LOGIN_FORM 
  }
}

export const openLoginForm = () => {
  return {
    type: OPEN_LOGIN_FORM
  }
}

export const closeModal = () => {
  return {
    type: CLOSE_MODAL
  }
}

export const openModal = () => {
  return {
    type: OPEN_MODAL
  }
}
