import React, { useState } from 'react';
import {
    AppstoreOutlined,
    ContainerOutlined,
    DesktopOutlined,
    MailOutlined,
    MenuFoldOutlined,
    MenuUnfoldOutlined,
    PieChartOutlined,
} from '@ant-design/icons';
import type { MenuProps } from 'antd';
import { Button, Menu } from 'antd';

type MenuItem = Required<MenuProps>['items'][number];

const items: MenuItem[] = [
    {
        key: 'sub1',
        label: 'Navigation One',
        icon: <MailOutlined />,
        children: [
            { key: '5', label: 'Option 5' },
            { key: '6', label: 'Option 6' },
            { key: '7', label: 'Option 7' },
            { key: '8', label: 'Option 8' },
        ],
    },
    {
        key: 'sub2',
        label: 'Navigation Two',
        icon: <AppstoreOutlined />,
        children: [
            { key: '9', label: 'Option 9' },
            { key: '10', label: 'Option 10' },
        ],
    },
];

const Sidebar: React.FC = () => {
    const [collapsed, setCollapsed] = useState(true); // 初始为收缩状态

    return (
        <div
            className="h-[100vh]"
            style={{ width: collapsed ? 80 : 200, }} // 根据状态设置宽度
            onMouseEnter={() => setCollapsed(false)} // 鼠标悬停时展开
            onMouseLeave={() => setCollapsed(true)} // 鼠标移开时收缩
        >
            <Menu
                className="h-[100vh]"
                style={{ backgroundColor: "rgb(245,245,245)", }}
                mode="inline" inlineCollapsed={collapsed} items={items} />
        </div>
    );
};

export default Sidebar;