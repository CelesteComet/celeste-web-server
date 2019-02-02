const merge = require('webpack-merge');
const common = require('./webpack.common.js');

module.exports = merge(common, {
  mode: 'development',
  devtool: 'inline-source-map',
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