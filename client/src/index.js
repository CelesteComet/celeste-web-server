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

// Import Font
require('./scss/main.scss')

let store;
if (process.env.NODE_ENV === 'development') {
  store = createStore(rootReducer, composeWithDevTools(
    applyMiddleware(thunk, logger)
  ));
} else {
  store = createStore(rootReducer, applyMiddleware(thunk));
}

render(
  <Provider store={store}>
    <App />
  </Provider>,
  document.getElementById('app')
)