import React from 'react'
import { BrowserRouter as Router, Route } from 'react-router-dom'
import { Layout, Menu, Icon } from 'antd'

import '../scss/App.scss'
import 'antd/dist/antd.css'

import Navbar from '../components/Navbar'
import Login from './Login'

const { SubMenu } = Menu
const { Header, Sider, Content } = Layout

const App = () => {
  return (
    <Router>
      <Layout style={{ height: '100%', width: '100%' }}>
        <Header style={{ color: '#FFF' }}>
          <Navbar />
        </Header>
        <Layout>
          <Sider>
            <Menu
              mode='inline'
              defaultSelectedKeys={['1']}
              defaultOpenKeys={['sub1']}
              style={{ height: '100%', borderRight: 0 }}
            >
              <SubMenu
                key='sub1'
                title={
                  <span>
                    <Icon type='user' />
                    subnav 1
                  </span>
                }
              >
                <Menu.Item key='1'>option1</Menu.Item>
                <Menu.Item key='2'>option2</Menu.Item>
              </SubMenu>
              <SubMenu
                key='sub2'
                title={
                  <span>
                    <Icon type='laptop' />
                    subnav 2
                  </span>
                }
              >
                <Menu.Item key='3'>option5</Menu.Item>
                <Menu.Item key='4'>option6</Menu.Item>
              </SubMenu>
            </Menu>
          </Sider>
          <Layout>
            <Content>
              <Route path='/' component={Login} />
            </Content>
          </Layout>
        </Layout>
      </Layout>
    </Router>
  )
}

export default App
