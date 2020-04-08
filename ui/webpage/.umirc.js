// ref: https://umijs.org/config/
export default {
  treeShaking: true,
  routes: [
    {
      path: '/',
      component: '../layouts/index',
      routes: [
        {
          path: '/config',
          component: './config',
        },
        {
          path: '/help',
          component: './help',
        },
        {
          path: '/login',
          component: './login',
        },
        {
          path: '/register',
          component: './register',
        },
        {
          path: '/about',
          component: './about',
        },
        {
          path: '/qqlogin',
          component: './qqlogin',
        },
        {
          path: '/',
          component: '../pages/index',
        },
      ],
    },
  ],
  plugins: [
    // ref: https://umijs.org/plugin/umi-plugin-react.html
    [
      'umi-plugin-react',
      {
        antd: false,
        dva: false,
        dynamicImport: false,
        title: 'webpage',
        dll: false,
        routes: {
          exclude: [/components\//],
        },
      },
    ],
  ],
};
