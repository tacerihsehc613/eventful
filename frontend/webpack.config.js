module.exports = {
    mode: 'development',
    entry: "./src/index.tsx",
    output: {
        publicPath: '/',
        filename: "bundle.js",
        path: __dirname + "/dist"
    },
    resolve: {
        // modules: ["src", "node_modules"],
        extensions: [".ts", ".tsx", ".js"]
    },
    module: {
        rules: [
            {
                test: /\.tsx?$/,
                loader: "ts-loader"
            }
        ]
    },
    externals: {
        "react": "React",
        "react-dom": "ReactDOM"
    },
    // experiments: {
	// 	lazyCompilation: {
	// 		entries: false,
	// 		imports: true,
	// 		backend: {
	// 			listen: {
	// 				host: "0.0.0.0"
	// 			}
	// 		}
	// 	}
	// },
    devServer: {
        //devMiddleware: { publicPath: '/' },
        client: {
            webSocketURL: {
              hostname: 'www.myevents.example',
              port: 8080,
            },
        },
        static: __dirname + "/",
        // devMiddleware: { publicPath: '/dist' },
        // static: { directory: path.resolve(__dirname) },
        // hot: true,
        historyApiFallback: true,
        host: '0.0.0.0',
        port: 8080,
        allowedHosts: 'all'
        // Other devServer configuration options can be added here
    }
}