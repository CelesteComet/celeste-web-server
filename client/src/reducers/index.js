import bagsReducer from './bagsReducer';
import usersReducer from './usersReducer';

import { combineReducers } from 'redux'

export default combineReducers({
  bags: bagsReducer,
  users: usersReducer
})


