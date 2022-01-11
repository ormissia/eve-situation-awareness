import React, {Component} from "react";
import {Layout} from 'antd';

import HomepageContent from "../components/HomepageContent";

const {Header, Footer, Content} = Layout;

class Homepage extends Component {

    render() {
        return (
            <Layout style={{background: 'white'}}>
                <Header style={{position: 'fixed', zIndex: 1, width: '100%', background: 'whitesmoke'}}>

                </Header>
                <Content className="site-layout" style={{padding: '32px', marginTop: 64}}>
                    <HomepageContent />
                </Content>
                <Footer style={{background: 'whitesmoke'}}>
                    foooooot
                </Footer>
            </Layout>
        )
    }
}


export default Homepage
