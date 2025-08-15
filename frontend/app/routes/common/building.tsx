import React from "react";
import type { Route } from "./+types/building";
import { Breadcrumb } from "antd";
import { Link } from "react-router";
import { Spin } from "antd";

export function meta({}: Route.MetaArgs) {
  return [
    { title: "Building Page" },
    { name: "description", content: "Building component page" },
  ];
}

export default function Building() {
  const [SpinCheck, setSpinCheck] = React.useState(false);

  // React.useEffect(() => {
  //   // Simulate a loading delay
  //   const timer = setTimeout(() => {
  //     setSpinCheck(false);
  //   }, 1000); // 2 seconds delay

  //   return () => clearTimeout(timer); // Cleanup the timer on unmount
  // }, []);

  return (
    <Spin tip="載入中..." spinning={SpinCheck} size="large"  >
      <div style={{ padding: "2rem" }}>
        <Breadcrumb
          items={[{ title: <Link to="/">Home</Link> }, { title: 'Building' }]}
          style={{ margin: 1 }}
        />

        <h1>Building Component</h1>
        <p>This is the Building component in the common routes.</p>
        <p>You have successfully navigated to the building page!</p>
      </div>
    </Spin>
  );
}