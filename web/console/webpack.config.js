const path = require('path');
const { CleanWebpackPlugin } = require('clean-webpack-plugin');
const HtmlWebpackPlugin = require('html-webpack-plugin');

module.exports = {
    mode: 'development',
    experiments: {
        asset: true
    },
    entry: './src/index.tsx',
    target: 'web',
    devtool: 'inline-source-map',
    output: {
        path: path.resolve(__dirname, 'dist/'),
        filename: '[name].[hash].js',
        publicPath: 'https://ultimatedivision.com/test/'
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
                    'css-loader',
                    'sass-loader',
                ]
            },
            {
                test: /\.(woff|woff2|eot|ttf|otf)$/i,
                exclude: /(node_modules)/,
                type: 'asset/resource',
                generator: {
                    filename: 'fonts/[name][hash:5][ext]'
                },
            },
            {
                test: /\.(jpe|jpg|png|svg)(\?.*$|$)/,
                exclude: /(node_modules)/,
                type: 'asset/resource',
                generator: {
                    filename: 'images/[name][hash:5][ext]'
                },
                use: [
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