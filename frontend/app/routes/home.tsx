import React from "react";

import type { Route } from "./+types/home";
import { Welcome } from "../welcome/welcome";
import { Breadcrumb, Spin } from "antd";
import { Divider, Typography } from 'antd';

export function meta({}: Route.MetaArgs) {
  return [
    { title: "New React Router App" },
    { name: "description", content: "Welcome to React Router!" },
  ];
}

export default function Home() {
  const [SpinCheck, setSpinCheck] = React.useState(false);

  // React.useEffect(() => {
  //   // Simulate a loading delay
  //   const timer = setTimeout(() => {
  //     setSpinCheck(false);
  //   }, 1000); // 2 seconds delay

  //   return () => clearTimeout(timer); // Cleanup the timer on unmount
  // }, []);

  return(
    <Spin tip="載入中..." spinning={SpinCheck} size="large">
      <Breadcrumb
        items={[{ title: 'Home' }]}
        style={{ margin: 1 }}
      />
      <Divider style={{ borderColor: "black", borderWidth: 1 }} />
      
      <Typography.Title level={3}>Welcome to the Home Page</Typography.Title>
      <Welcome />
    </Spin>
  );
}
