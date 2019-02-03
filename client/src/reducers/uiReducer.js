import { 
  CLOSE_LOGIN_FORM,
  OPEN_LOGIN_FORM,
  START_CLOSE_ANIMATION,
  END_CLOSE_ANIMATION
} 
from '../actions/uiActions'

const initialState = {
  vLoginForm: true 
};

const uiReducer = (state = initialState, action) => {
  let newState = Object.assign({}, state);
  switch (action.type) {
    case CLOSE_LOGIN_FORM:
      newState.vLoginForm = false;
      return newState
    case OPEN_LOGIN_FORM:
      newState.vLoginForm = true;
      return newState      
    default:
      return state
  }
}

export default uiReducer;
