const path = require('path');

module.exports = {
	mode: 'development',
  entry: './src/index.js',
  output: {
    filename: 'bundle.js',
    path: path.resolve(__dirname, 'dist')
  },
	devServer: {	
		contentBase: './dist',
    historyApiFallback: true,// Fall back to index.html in case file in server can't be found
    proxy: {
      '/api': {
        target: 'http://localhost:8080',
        secure: false
      }
    }    
	},
	module: {
		rules: [ 
			{ test: /\.js$/, exclude: /node_modules/, loader: "babel-loader" }
		]
	}
}
