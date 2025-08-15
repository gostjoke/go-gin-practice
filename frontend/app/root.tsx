import React from 'react';
import {
  isRouteErrorResponse,
  Links,
  Meta,
  Outlet,
  Scripts,
  ScrollRestoration,
} from "react-router";

import type { Route } from "./+types/root";
import "./app.css";
import "antd/dist/reset.css"; // Antd v5 樣式

import { 
  LaptopOutlined, 
  NotificationOutlined, 
  UserOutlined, 
  MenuUnfoldOutlined, 
  MenuFoldOutlined  
} from '@ant-design/icons';
import type { MenuProps } from 'antd';
import {Layout, Menu, theme, Button } from 'antd';
import { Link } from "react-router";

const { Header, Content, Sider, Footer  } = Layout;


const items2: MenuProps['items'] = [
  {
    key: 'user',
    icon: React.createElement(UserOutlined),
    label: 'User Management',
    children: [
      { key: '1', label: 'User List' },
      { key: '2', label: 'User Profile' },
      { key: '3', label: 'User Settings' },
      { key: '4', label: 'User Permissions' },
    ],
  },
  {
    key: 'system',
    icon: React.createElement(LaptopOutlined),
    label: 'System',
    children: [
      { key: '5', label: 'Dashboard' },
      { key: '6', label: 'Monitoring' },
      { key: '7', label: 'Logs' },
      { key: '8', label: 'Configuration' },
    ],
  },
  {
    key: 'notifications',
    icon: React.createElement(NotificationOutlined),
    label: 'Notifications',
    children: [
      { key: '9', label: 'Inbox' },
      { key: '10', label: 'Sent' },
      { key: '11', label: 'Draft' },
      { key: '12', label: 'Settings' },
    ],
  },
];

export const links: Route.LinksFunction = () => [
  { rel: "preconnect", href: "https://fonts.googleapis.com" },
  {
    rel: "preconnect",
    href: "https://fonts.gstatic.com",
    crossOrigin: "anonymous",
  },
  {
    rel: "stylesheet",
    href: "https://fonts.googleapis.com/css2?family=Inter:ital,opsz,wght@0,14..32,100..900;1,14..32,100..900&display=swap",
  },
];


export default function App() {
  const [collapsed, setCollapsed] = React.useState(false);
  const {
    token: { colorBgContainer, borderRadiusLG },
  } = theme.useToken();

  const items1: MenuProps['items'] = [
    {
      key: '0',
      label: <>
        <Button
          type="text"
          icon={collapsed ? <MenuUnfoldOutlined /> : <MenuFoldOutlined />}
          onClick={() => setCollapsed(!collapsed)}
          style={{
            fontSize: '16px',
            width: 20,
            height: 20,
            background: '#4F4F4F',
          }}
        />
      </>
    },
    {
      key: '1',
      label: <Link to="/">Home</Link>,
    },
    {
      key: '2',
      label: <Link to="/building">Building</Link>,
    },
    {
      key: '3',
      label: 'nav 3',
    }
  ];

  return (
    <html lang="en">
      <head>
        <meta charSet="utf-8" />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <Meta />
        <Links />
      </head>
      <body>
        <Layout>
          <Layout>
            <Sider 
              trigger={null} collapsible collapsed={collapsed}
              style={{
                left: 0,
                top: 0,
                bottom: 0,
                height: '100vh',      // 撐滿整個視窗高度
                overflow: 'auto',     // 如果內容太多就內部滾動
              }}
              >
                <img
                  src="./img/FedExLogo.webp"
                  alt="Not find"  
                  style={{
                    width: '100%',      // 讓圖片寬度跟 Sider 一樣寬
                    height: '8%',     // 高度等比例縮放
                    padding: '10px',    // 留點內距
                    objectFit: 'contain', // 保持比例不裁切
                    background: 'white', // 背景色
                  }}
                />
              <Menu
                className="custom-menu"
                mode="inline"
                theme="light"  
                defaultSelectedKeys={['1']}
                defaultOpenKeys={['sub1']}
                    style={{
                      height: '100%',
                    }}
                items={items2}
              />
            </Sider>
            <Layout style={{ height: '100vh' }}>
                <Header
                  style={{
                    display: 'flex',
                    alignItems: 'center',
                    position: 'sticky',  // 固定
                    top: 0,               // 距離頂部 0
                    zIndex: 1,            // 保證在上方
                    width: '100%',
                  }}
                >
                <Menu
                  theme="dark"
                  mode="horizontal"
                  defaultSelectedKeys={['1']}
                  items={items1}
                  style={{ flex: 1, minWidth: 0, paddingInline: 0,}}
                />
              </Header>
              <Content
                style={{
                  padding: 24,
                  margin: 0,
                  // minHeight: 280,
                  overflow: 'auto',     // 只有內容可滾動
                  height: '100vh',
                  background: "#E0E0E0",
                  borderRadius: borderRadiusLG,
                }}
              >
                <Outlet />
              </Content>
              <Footer
                style={{
                  textAlign: 'center',
                  background: '#000000',
                  // padding: '10px 50px',
                  color: 'white',
                }}
              >
                © {new Date().getFullYear()} Your Company Name
              </Footer>
            </Layout>
          </Layout>
        </Layout>
        <ScrollRestoration />
        <Scripts />
      </body>
    </html>
  );
}

export function ErrorBoundary({ error }: Route.ErrorBoundaryProps) {
  let message = "Oops!";
  let details = "An unexpected error occurred.";
  let stack: string | undefined;

  if (isRouteErrorResponse(error)) {
    message = error.status === 404 ? "404" : "Error";
    details =
      error.status === 404
        ? "The requested page could not be found."
        : error.statusText || details;
  } else if (import.meta.env.DEV && error && error instanceof Error) {
    details = error.message;
    stack = error.stack;
  }

  return (
    <main style={{ padding: "20px" }}>
      <h1>{message}</h1>
      <p>{details}</p>
      {stack && (
        <pre style={{ overflowX: "auto", background: "#f5f5f5", padding: "10px" }}>
          <code>{stack}</code>
        </pre>
      )}
    </main>
  );
}
