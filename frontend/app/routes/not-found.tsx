import React from "react";
import { Breadcrumb, Button, Result } from "antd";
import { Link } from "react-router";

export function meta() {
  return [
    { title: "404 - Page Not Found" },
    { name: "description", content: "The requested page could not be found." },
  ];
}

export default function NotFound() {
  return (
    <div>
      <Breadcrumb
        items={[{ title: <Link to="/">Home</Link> }, { title: '404' }]}
        style={{ margin: 1 }}
      />
      
      <Result
        status="404"
        title="404"
        subTitle="抱歉，您訪問的頁面不存在。"
        extra={
          <Button type="primary">
            <Link to="/">返回首頁</Link>
          </Button>
        }
      />
    </div>
  );
}
