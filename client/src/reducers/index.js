import bagsReducer from './bagsReducer';
import usersReducer from './usersReducer';
import uiReducer from './uiReducer';
import errorsReducer from './errorsReducer';

import { combineReducers } from 'redux'

export default combineReducers({
  bags: bagsReducer,
  user: usersReducer,
  ui: uiReducer,
  errors: errorsReducer 
})


