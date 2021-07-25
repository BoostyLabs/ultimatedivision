const path = require('path');
const { CleanWebpackPlugin } = require('clean-webpack-plugin');
const HtmlWebpackPlugin = require('html-webpack-plugin');

module.exports = {
    mode: 'development',
    entry: './src/index.tsx',
    target: 'web',
    devtool: 'inline-source-map',
    output: {
        path: path.resolve(__dirname, 'dist/'),
        filename: '[name].[hash].js',
        publicPath: '/'
    },
    plugins: [
        new HtmlWebpackPlugin({
            title: 'Ultimate Division',
            template: './public/index.html'
        }),
        new CleanWebpackPlugin()
    ],
    devServer: {
        port: 3000,
        open: true,
        historyApiFallback: true
    },
    resolve: {
        alias: {
            "@FootballerCard": path.resolve(__dirname, './src/app/components/FootballerCardPage/'),
            "@FootballField": path.resolve(__dirname, './src/app/components/FootballFieldPage/'),
            "@MarketPlace": path.resolve(__dirname, './src/app/components/MarketPlacePage/'),
            "@Navbar": path.resolve(__dirname, './src/app/components/Navbar/'),
            "@Paginator": path.resolve(__dirname, './src/app/components/Paginator/'),
            "@PlayerCard": path.resolve(__dirname, './src/app/components/PlayerCard/'),
            "@Img": path.resolve(__dirname, './src/app/static/img/'),
            "@Store": path.resolve(__dirname, './src/app/store/'),
            "@Types": path.resolve(__dirname, './src/app/types/'),
            "@Routes": path.resolve(__dirname, './src/app/routes/'),
        },
        extensions: [
            '.ts',
            '.tsx',
            '.js',
            '.jsx',
        ],
        modules: ['node_modules']
    },
    module: {
        rules: [
            {
                test: /\.m?(tsx|ts)$/i,
                exclude: /(node_modules)/,
                use: [
                    {
                        loader: 'ts-loader'
                    }
                ],
            },
            {
                test: /\.(s[c]ss|css)$/,
                use: [
                    'style-loader',
                    {
                        loader: 'css-loader',
                        options: {

                            url: false
                        }
                    },
                    'sass-loader'
                ]
            },
            {
                test: /\.(woff|woff2|eot|ttf|otf)$/i,
                use: [
                    {
                        loader: 'url-loader',
                        options: {
                            name: './fonts/[name].[ext]',
                        }
                    },
                ]
            },
            {
                test: /\.(jpe|jpg|png|svg)(\?.*$|$)/,
                exclude: /(node_modules)/,
                use: [
                    {
                        loader: 'file-loader',
                        options: {
                            name: './images/[name].[ext]',
                        }
                    },
                    {
                        loader: 'image-webpack-loader',
                        options: {
                            mozjpeg: {
                                progressive: true,
                            },
                            // optipng.enabled: false will disable optipng
                            optipng: {
                                enabled: false,
                            },
                            pngquant: {
                                quality: [0.8, 0.90],
                                speed: 2
                            },
                            gifsicle: {
                                interlaced: false,
                            },
                            // the webp option will enable WEBP
                            webp: {
                                quality: 75
                            }
                        }
                    },
                ],
            }
        ]
    },
};