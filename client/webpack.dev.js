const merge = require('webpack-merge');
const common = require('./webpack.common.js');
const webpack = require('webpack');

module.exports = merge(common, {
  mode: 'development',
  devtool: 'inline-source-map',
  plugins: [
    new webpack.EnvironmentPlugin({
      NODE_ENV: 'development', // use 'development' unless process.env.NODE_ENV is defined
      AUTH_URL: 'http://localhost:1337'
    })
  ],    
  devServer: {  
    contentBase: './dist',
    historyApiFallback: true,// Fall back to index.html in case file in server can't be found
    proxy: {
      '/api': {
        target: 'http://localhost:8080',
        secure: false
      },
      '/auth': {
        target: 'http://localhost:8080',
        secure: false
      }
    }    
  }
});