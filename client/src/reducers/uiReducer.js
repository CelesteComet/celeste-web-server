import { 
  OPEN_LOGIN_FORM,
  CLOSE_LOGIN_FORM,
  TOGGLE_SIDEBAR_NAV,
} from '../actions/uiActions'

const initialState = {
  vLoginForm: false,
  vSideBarNav: false 
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
    case TOGGLE_SIDEBAR_NAV:
      newState.vSideBarNav = !state.vSideBarNav;
      console.log(newState.vSideBarNav)
      return newState 
    default:
      return state
  }
}

export default uiReducer;

