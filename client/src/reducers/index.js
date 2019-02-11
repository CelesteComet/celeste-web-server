import { combineReducers } from 'redux'
import bagsReducer from './bagsReducer';
import usersReducer from './usersReducer';
import uiReducer from './uiReducer';
import errorsReducer from './errorsReducer';
import commentsReducer from './commentsReducer';

export default combineReducers({
  bags: bagsReducer,
  user: usersReducer,
  comments: commentsReducer,
  ui: uiReducer,
  errors: errorsReducer 
})


