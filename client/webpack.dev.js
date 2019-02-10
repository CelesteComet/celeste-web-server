const merge = require('webpack-merge');
const common = require('./webpack.common.js');
const webpack = require('webpack');
const BundleAnalyzerPlugin = require('webpack-bundle-analyzer').BundleAnalyzerPlugin;
const OptimizeCSSAssetsPlugin = require("optimize-css-assets-webpack-plugin");

module.exports = merge(common, {
  mode: 'development',
  devtool: 'inline-source-map',
  plugins: [
    new webpack.EnvironmentPlugin({
      NODE_ENV: 'development', // use 'development' unless process.env.NODE_ENV is defined
      AUTH_URL: 'http://localhost:1337'
    })
  ],    
  module: {
    rules: [ 
      { test: /\.js$/, exclude: /node_modules/, loader: "babel-loader" },
      {
        test: /\.scss$/,
        use: [
          { loader: "style-loader" },
          { loader: 'css-loader' },          
          { loader: "sass-loader"} // compiles Sass to CSS, using Node Sass by default
        ],
      },
      {
        test: /\.(woff(2)?|ttf|eot|svg)(\?v=\d+\.\d+\.\d+)?$/,
        use: [{
          loader: 'file-loader',
          options: { name: '[name].[ext]' }
        }]
      }      
    ]
  },  
  devServer: {  
    contentBase: './dist',
    historyApiFallback: true, // Fall back to index.html in case file in server can't be found
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