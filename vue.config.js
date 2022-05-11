const { defineConfig } = require('@vue/cli-service')
// module.exports = defineConfig({
//   transpileDependencies: true
// },
// devServer: {
//   proxy:'https://localhost:8088'

// })

module.exports = {
  devServer: {
    proxy: {
      "^/api": {
        target: "http://127.0.0.1:8088",
        changeOrigin: true,
        logLevel: "debug",
        pathRewrite: { "^/api": "/" }
      }
    }
  }
};