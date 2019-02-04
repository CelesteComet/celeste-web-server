import { 
  OPEN_LOGIN_FORM,
  CLOSE_LOGIN_FORM
} from '../actions/uiActions'

const initialState = {
  vLoginForm: false 
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

