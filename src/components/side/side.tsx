import { AppstoreOutlined, MailOutlined } from "@ant-design/icons";
import type { MenuProps } from "antd";
import { Menu } from "antd";
type MenuItem = Required<MenuProps>["items"][number];

const items: MenuItem[] = [
  {
    key: "sub1",
    label: "Navigation One",
    icon: <MailOutlined />,
    children: [
      { key: "1", label: "Option 5" },
      { key: "2", label: "Option 6" },
    ],
  },
  {
    key: "sub2",
    label: "Navigation Two",
    icon: <AppstoreOutlined />,
    children: [
      { key: "3", label: "Option 5" },
      { key: "4", label: "Option 6" },
    ],
  },
];

export const Side = ({ open }: { open: boolean }) => {
  return (
    <Menu
      style={{ backgroundColor: "rgb(245,245,245)" }}
      mode="inline"
      inlineCollapsed={!open}
      items={items}
    />
  );
};
