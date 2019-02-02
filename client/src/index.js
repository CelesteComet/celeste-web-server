import React from 'react'
import { render } from 'react-dom'
import { Provider } from 'react-redux'

// REDUX STUFF
import logger from 'redux-logger';
import thunk from 'redux-thunk';
import { createStore, applyMiddleware } from 'redux'
import { composeWithDevTools } from 'redux-devtools-extension';
import rootReducer from './reducers'

import App from './components/App'

const store = createStore(rootReducer, composeWithDevTools(
  applyMiddleware(thunk, logger)
));

render(
  <Provider store={store}>
    <App />
  </Provider>,
  document.getElementById('app')
)