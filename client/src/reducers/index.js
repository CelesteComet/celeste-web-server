import bagsReducer from './bagsReducer';

import { combineReducers } from 'redux'

export default combineReducers({
  bags: bagsReducer
})


